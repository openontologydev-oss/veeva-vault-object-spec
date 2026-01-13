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
    lines.append(f"// Generated date: {data.get('_metadata', {}).get('extracted_date', 'Unknown')}\n")

    components = data.get("components", {})
    
    # Forward declarations aren't strictly needed in TS interfaces as long as they are defined in the file
    
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

    # Process subcomponents (explicitly defined subcomponents usually map to array attributes)
    # However, sometimes they are not listed in 'attributes' but exist in 'subcomponents'
    # In MDL, usually subcomponents are contained in the parent.
    # We should check if we missed any subcomponents that weren't in attributes.
    
    subcomponents = data.get("subcomponents", {})
    
    # Check if there are subcomponents that act as attributes but weren't in the attributes list?
    # Usually in the schema I extracted, attributes cover everything. 
    # But let's handle the definition of the subcomponent types themselves.
    
    lines.append("}\n")
    
    # Generate types for subcomponents
    for sub_name, sub_data in subcomponents.items():
        # Subcomponent type name: Parent_Child
        sub_type_name = f"{name}_{sub_name}"
        # Some subcomponent names might be generic like 'field', so namespacing is good.
        # But MDL capitalization is usually TitleCase. 'field' might be 'Field'.
        # The schema keys for subcomponents seem to be lowercase or mixed.
        # Let's Capitalize for the type name.
        sub_type_name_cap = f"{name}_{sub_name[0].upper() + sub_name[1:]}"
        
        process_component(sub_type_name_cap, sub_data, lines)


def map_type(parent_name, attr_name, attr_def, parent_data):
    raw_type = attr_def.get("type")
    
    if raw_type == "Subcomponent":
        # It references a subcomponent.
        # Check if there is a matching subcomponent key
        # Try exact match, or case insensitive
        subcomponents = parent_data.get("subcomponents", {})
        target_sub = None
        if attr_name in subcomponents:
            target_sub = attr_name
        else:
            # Maybe singular/plural match? Or capitalization?
            # We'll guess the type name based on our naming convention
            pass
        
        # If we can't find it, we assume it's constructed as Parent_AttributeName(Capitalized)
        type_name = f"{parent_name}_{attr_name[0].upper() + attr_name[1:]}"
        return f"{type_name}[]" # Subcomponents are usually arrays
        
    return TYPE_MAPPING.get(raw_type, "any")

if __name__ == "__main__":
    generate_ts()
