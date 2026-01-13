import json
import os

SCHEMA_PATH = "../experiment_07_postman_collection/mdl_components_schema.json"
OUTPUT_PATH = "index.ts"

TYPE_MAPPING = {
    "Boolean": "boolean",
    "Number": "number",
    "String": "string",
    "LongString": "string",
    "XMLString": "string",
    "Expression": "string",
    "File": "string",
    "Enum": "string",
    "SdkCode": "string",
    "ComponentReference": "string",
    "SubcomponentReference": "string",
    "ActionScript": "string"
}

def generate_ts():
    if not os.path.exists(SCHEMA_PATH):
        print(f"Error: {SCHEMA_PATH} not found")
        return

    with open(SCHEMA_PATH, 'r') as f:
        data = json.load(f)

    lines = []
    lines.append("// Generated TypeScript definitions for Veeva Vault MDL Components")
    lines.append("// Source: https://developer.veevavault.com/mdl/components")
    lines.append(f"// Generated date: {data.get('_metadata', {}).get('extracted_date', 'Unknown')}")
    lines.append(f"// {data.get('_metadata', {}).get('enrichment', '')}\n")

    components = data.get("components", {})
    
    for comp_name, comp_data in components.items():
        process_component(comp_name, comp_data, lines)

    with open(OUTPUT_PATH, 'w') as f:
        f.write("\n".join(lines))
    
    print(f"Generated {OUTPUT_PATH} with {len(lines)} lines")

def process_component(name, data, lines):
    description = data.get("description", "").replace("\n", " ")
    if description:
        lines.append(f"/** {description} */")
    
    lines.append(f"export interface {name} {{")
    
    # Process attributes
    attributes = data.get("attributes", {})
    for attr_name, attr_def in attributes.items():
        ts_type = map_type(name, attr_name, attr_def, data)
        required = "" if attr_def.get("required") else "?"
        lines.append(f"  {attr_name}{required}: {ts_type};")

    lines.append("}\n")
    
    # Generate types for subcomponents
    subcomponents = data.get("subcomponents", {})
    for sub_name, sub_data in subcomponents.items():
        # Subcomponent type name: Parent_Child
        # Capitalize for the type name.
        sub_type_name_cap = f"{name}_{sub_name[0].upper() + sub_name[1:]}"
        process_component(sub_type_name_cap, sub_data, lines)


def map_type(parent_name, attr_name, attr_def, parent_data):
    raw_type = attr_def.get("type")
    
    # 1. Handle Enums (Implicit or Explicit)
    # Check for 'enum_values' or 'allowedValues'
    # These might appear even if type is 'String' (e.g. implicit enums from 'options:' scraping)
    
    values = attr_def.get("enum_values") or attr_def.get("allowedValues")
    if values and isinstance(values, list) and len(values) > 0:
        # Create a union type: "Val1" | "Val2"
        # Escape quotes in values just in case
        quoted_values = []
        for v in values:
            clean_v = v.replace("'", "\\'")
            quoted_values.append(f"'{clean_v}'")
        return " | ".join(quoted_values)

    # 2. Handle Subcomponents
    if raw_type == "Subcomponent":
        # It references a subcomponent.
        # Check if there is a matching subcomponent key in parent
        subcomponents = parent_data.get("subcomponents", {})
        
        # We assume it's constructed as Parent_AttributeName(Capitalized)
        # Note: In MDL, the attribute name often matches the subcomponent name (e.g. 'field' -> 'Field')
        # But sometimes not exactly. 
        type_name = f"{parent_name}_{attr_name[0].upper() + attr_name[1:]}"
        return f"{type_name}[]" # Subcomponents are usually arrays
        
    return TYPE_MAPPING.get(raw_type, "any")

if __name__ == "__main__":
    generate_ts()
