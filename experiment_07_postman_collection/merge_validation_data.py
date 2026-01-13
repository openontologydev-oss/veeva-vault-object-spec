import json
import re
import os

VALIDATION_PATH = "../experiment_01_meddle_schema/meddle/src/meddle/validation.json"
SCHEMA_PATH = "mdl_components_schema.json"
OUTPUT_PATH = "mdl_components_schema_enriched.json"

def clean_enum_value(val):
    return val.strip()

def extract_validation_rules(type_data_str, desc_str):
    rules = {}
    
    # Extract Allowed values: A|B|C|
    allowed_match = re.search(r"Allowed values\s*:\s*([^:\n]+)", type_data_str)
    if allowed_match:
        raw_vals = allowed_match.group(1).split('|')
        rules['enum_values'] = [v.strip() for v in raw_vals if v.strip()]
        
    # Extract options: A| B|
    # Sometimes in description, sometimes in type_data? 
    # The grep showed "options: editable| readonly| hidden|" in description
    if desc_str and "options:" in desc_str:
        options_match = re.search(r"options:\s*([^:\n]+)", desc_str)
        if options_match:
            raw_vals = options_match.group(1).split('|')
            # If we didn't find enums yet, or if this looks like a valid set
            if 'enum_values' not in rules:
                rules['enum_values'] = [v.strip() for v in raw_vals if v.strip()]

    # Extract Max/Min length/value
    # "Maximum length : 60"
    max_len_match = re.search(r"Maximum length\s*:\s*(\d+)", type_data_str)
    if max_len_match:
        rules['maxLength'] = int(max_len_match.group(1))

    # "Maximum value : 9"
    max_val_match = re.search(r"Maximum value\s*:\s*(\d+)", type_data_str)
    if max_val_match:
        rules['maxValue'] = int(max_val_match.group(1))

    # "Minimum value : 1"
    min_val_match = re.search(r"Minimum value\s*:\s*(\d+)", type_data_str)
    if min_val_match:
        rules['minValue'] = int(min_val_match.group(1))

    return rules

def merge_validation():
    if not os.path.exists(VALIDATION_PATH):
        print(f"Error: {VALIDATION_PATH} not found")
        return

    with open(VALIDATION_PATH, 'r') as f:
        validation_data = json.load(f)

    with open(SCHEMA_PATH, 'r') as f:
        schema_data = json.load(f)

    stats = {
        "enums_added": 0,
        "constraints_added": 0
    }

    # Recursive function to handle nested subcomponents
    def process_component_merge(target_comp, source_comp_data):
        target_attrs = target_comp.get("attributes", {})
        source_attrs = source_comp_data.get("attributes", {})
        
        for attr_name, source_attr_def in source_attrs.items():
            if attr_name in target_attrs:
                target_attr = target_attrs[attr_name]
                type_data = source_attr_def.get("type_data", "")
                description = source_attr_def.get("description", "")
                
                rules = extract_validation_rules(type_data, description)
                
                if 'enum_values' in rules:
                    target_attr['enum_values'] = rules['enum_values']
                    stats['enums_added'] += 1
                    
                if 'maxValue' in rules:
                    target_attr['maxValue'] = rules['maxValue']
                    stats['constraints_added'] += 1
                
                if 'minValue' in rules:
                    target_attr['minValue'] = rules['minValue']
                    stats['constraints_added'] += 1
                    
                # We already have maxLength from the scrape potentially, but this might be more accurate
                if 'maxLength' in rules and 'maxLength' not in target_attr:
                    target_attr['maxLength'] = rules['maxLength']

        # Recursively handle subcomponents
        # In validation.json: "subcomponents": { "SubName": { ... } }
        # In schema.json: "subcomponents": { "SubName": { ... } } (Usually lowercase keys in schema)
        
        source_subs = source_comp_data.get("subcomponents", {})
        target_subs = target_comp.get("subcomponents", {})
        
        for sub_name, sub_data in source_subs.items():
            # Schema keys might be case sensitive or not. 
            # In extraction, keys were often TitleCase or lowercase.
            # Let's try direct match first, then Case insensitive
            target_sub_key = sub_name
            if sub_name not in target_subs:
                # search case insensitive
                for k in target_subs.keys():
                    if k.lower() == sub_name.lower():
                        target_sub_key = k
                        break
            
            if target_sub_key in target_subs:
                process_component_merge(target_subs[target_sub_key], sub_data)


    for comp_name, comp_data in validation_data.items():
        if comp_name in schema_data['components']:
            process_component_merge(schema_data['components'][comp_name], comp_data)

    schema_data['_metadata']['enrichment'] = "Enriched with validation data from meddle experiment"
    
    with open(OUTPUT_PATH, 'w') as f:
        json.dump(schema_data, f, indent=2)
    
    print(f"Merged validation data. Added {stats['enums_added']} enums and {stats['constraints_added']} constraints.")
    
    # Overwrite original
    os.rename(OUTPUT_PATH, SCHEMA_PATH)

if __name__ == "__main__":
    merge_validation()
