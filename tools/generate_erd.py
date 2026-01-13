import json
import os

SCHEMA_PATH = "../experiment_07_postman_collection/mdl_components_schema.json"
OUTPUT_PATH = "veeva_mdl_erd.mmd"

def generate_erd():
    if not os.path.exists(SCHEMA_PATH):
        print(f"Error: {SCHEMA_PATH} not found")
        return

    with open(SCHEMA_PATH, 'r') as f:
        data = json.load(f)

    lines = []
    lines.append("classDiagram")

    components = data.get("components", {})
    
    # We want to identify relationships.
    # Relationships can be:
    # 1. Attribute Type == "ComponentReference" or "SubcomponentReference"
    # 2. Attribute Type == "Object" (Special reference to Object component)
    # 3. Attribute Type == "Doclifecycle" (Special reference)
    # 4. Subcomponents (Composition)

    # First, declare all classes to ensure they exist
    # But Mermaid handles implicit declarations fine.

    known_components = set(components.keys())

    # We might want to filter to avoid a 1000-node graph if it's too big. 
    # But let's try generating the full one first.

    def process_comp_relationships(comp_name, comp_data, parent_name=None):
        # Add class definition with attributes
        lines.append(f"    class {comp_name} {{")
        for attr_name, attr_def in comp_data.get("attributes", {}).items():
            attr_type = attr_def.get("type", "String")
            # Cleanup type for display
            display_type = attr_type
            if attr_def.get("required"):
                display_type += "*"
            lines.append(f"        {display_type} {attr_name}")
        lines.append("    }")

        # Parent composition relationship
        if parent_name:
            lines.append(f"    {parent_name} *-- {comp_name} : contains")

        # Attribute relationships
        for attr_name, attr_def in comp_data.get("attributes", {}).items():
            raw_type = attr_def.get("type")
            target = None
            
            # Explicit references
            if raw_type in ["ComponentReference", "SubcomponentReference"]:
                # The schema doesn't always strictly say WHICH component it references in the 'type' field
                # It often says it in the description or it's implied by name.
                # However, sometimes 'type' is literally the component name (e.g. "Object", "Doclifecycle")
                pass

            # Implicit references where Type matches a Component Name
            if raw_type in known_components:
                target = raw_type
            
            # Special Known Types that are components
            if raw_type == "Object": target = "Object"
            if raw_type == "Doclifecycle": target = "Doclifecycle"
            if raw_type == "Doctype": target = "Doctype"
            
            if target and target != comp_name:
                # Add relationship
                # Use dashed line for reference
                lines.append(f"    {comp_name} ..> {target} : {attr_name}")


        # Process Subcomponents recursively
        for sub_name, sub_data in comp_data.get("subcomponents", {}).items():
            # Subcomponent names are often TitleCase in keys but let's standardize
            sub_full_name = f"{comp_name}_{sub_name}"
            # Actually, standardizing the name for unique class ID
            clean_sub_name = f"{comp_name}_{sub_name}"
            process_comp_relationships(clean_sub_name, sub_data, parent_name=comp_name)

    for name, data in components.items():
        process_comp_relationships(name, data)

    with open(OUTPUT_PATH, 'w') as f:
        f.write("\n".join(lines))
    
    print(f"Generated {OUTPUT_PATH} with {len(lines)} lines")

if __name__ == "__main__":
    generate_erd()
