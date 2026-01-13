import json
import os
import re

SCHEMA_PATH = "../experiment_07_postman_collection/mdl_components_schema.json"
OUTPUT_PATH = "vault/enums_gen.go"

def clean_name(s):
    # Remove non-alphanumeric, TitleCase
    s = re.sub(r'[^a-zA-Z0-9]', '_', s)
    parts = s.split('_')
    return ''.join(p.capitalize() for p in parts if p)

def clean_val(s):
    # Handle values like "Dark Orange" -> "DarkOrange"
    # Handle "hide__v" -> "HideV"
    # Handle "Read-Only" -> "ReadOnly"
    # Handle numbers or symbols
    if not s:
        return "Empty"
    
    clean = re.sub(r'[^a-zA-Z0-9]', ' ', s).title().replace(' ', '')
    if clean and clean[0].isdigit():
        clean = "Val" + clean
    return clean

def generate_go():
    if not os.path.exists(SCHEMA_PATH):
        print(f"Error: {SCHEMA_PATH} not found")
        return

    with open(SCHEMA_PATH, 'r') as f:
        data = json.load(f)

    lines = []
    lines.append("// Package vault provides Go ORM models for Veeva Vault objects and components.")
    lines.append("// This file is auto-generated from the official Veeva MDL Components documentation.")
    lines.append("// It contains Enum definitions derived from validation rules.")
    lines.append("package vault")
    lines.append("")

    components = data.get("components", {})
    
    # Track generated types to avoid duplicates if multiple fields share the same Enum logic (rare in this schema structure)
    # Actually in MDL schema, enums are defined per attribute usually.
    
    generated_types = set()

    def process_comp(comp_name, comp_data):
        attributes = comp_data.get("attributes", {})
        for attr_name, attr_def in attributes.items():
            values = attr_def.get("enum_values") or attr_def.get("allowedValues")
            
            if values and len(values) > 0:
                # Construct Type Name: CompNameAttrName
                # e.g. Appsecurityfield + Type -> AppsecurityfieldType
                
                # Check for existing types in types.go? 
                # We'll prefix with "Enum" or just use the Component prefix to avoid collision with standard types?
                # Actually, let's use ComponentName + AttributeName (TitleCase)
                
                type_name = clean_name(comp_name) + clean_name(attr_name)
                
                if type_name in generated_types:
                    continue
                
                generated_types.add(type_name)
                
                lines.append(f"// {type_name} represents values for {comp_name}.{attr_name}")
                lines.append(f"type {type_name} string")
                lines.append("")
                lines.append("const (")
                
                for v in values:
                    # Const Name: TypeName + Value
                    # e.g. AppsecurityfieldTypeOBJECT
                    val_name = clean_val(v)
                    if not val_name: 
                        val_name = "Empty"
                        
                    const_name = f"{type_name}{val_name}"
                    lines.append(f'\t{const_name} {type_name} = "{v}"')
                
                lines.append(")")
                lines.append("")

        for sub_name, sub_data in comp_data.get("subcomponents", {}).items():
            # Subcomponent name usually already capitalized in dict?
            sub_full_name = f"{comp_name}{sub_name}"
            process_comp(sub_full_name, sub_data)

    for name, data in components.items():
        process_comp(name, data)

    with open(OUTPUT_PATH, 'w') as f:
        f.write("\n".join(lines))
        
    print(f"Generated {OUTPUT_PATH} with {len(lines)} lines")

if __name__ == "__main__":
    generate_go()
