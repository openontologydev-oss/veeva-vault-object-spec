import json
import os

SCHEMA_PATH = "../experiment_07_postman_collection/mdl_components_schema.json"
OUTPUT_DIR = "."

# Core component categories for separate diagrams
CATEGORIES = {
    "core_objects": ["Object", "Objectlifecycle", "Objecttype", "Objectaction", "Objectvalidation", "Objectworkflow"],
    "documents": ["Docfield", "Doctype", "Doclifecycle", "Docatomicsecurity", "Docfielddependency", "Docfieldlayout"],
    "security": ["Securityprofile", "Permissionset", "Atomicsecurity", "Appsecurityrule", "Sharingrule"],
    "workflows": ["Workflow", "Recordtrigger", "Actiontrigger", "Recordaction", "Recordworkflowaction"],
    "sdk": ["Customwebapi", "Userdefinedclass", "Userdefinedservice", "Sdkjob", "Queue", "Job"],
    "ui": ["Tab", "Dashboard", "Report", "Page", "Pagelayout", "Searchcollection"]
}

def generate_erd_for_category(category_name, component_names, all_components):
    lines = []
    lines.append("classDiagram")
    lines.append(f"    %% {category_name.upper()} Components")

    for comp_name in component_names:
        if comp_name not in all_components:
            continue
        comp_data = all_components[comp_name]
        process_component(comp_name, comp_data, lines, all_components)

    return "\n".join(lines)

def process_component(comp_name, comp_data, lines, all_components, parent_name=None):
    # Add class definition
    lines.append(f"    class {comp_name} {{")
    attrs = comp_data.get("attributes", {})
    # Limit attributes to avoid size issues
    for i, (attr_name, attr_def) in enumerate(attrs.items()):
        if i >= 10:  # Max 10 attrs per class
            lines.append("        ...")
            break
        attr_type = attr_def.get("type", "String")[:15]  # Truncate type
        lines.append(f"        {attr_type} {attr_name}")
    lines.append("    }")

    # Parent composition
    if parent_name:
        lines.append(f"    {parent_name} *-- {comp_name}")

    # Process subcomponents
    for sub_name, sub_data in comp_data.get("subcomponents", {}).items():
        sub_full_name = f"{comp_name}_{sub_name}"
        process_component(sub_full_name, sub_data, lines, all_components, parent_name=comp_name)

def main():
    if not os.path.exists(SCHEMA_PATH):
        print(f"Error: {SCHEMA_PATH} not found")
        return

    with open(SCHEMA_PATH, 'r') as f:
        data = json.load(f)

    components = data.get("components", {})

    for cat_name, comp_names in CATEGORIES.items():
        mmd_content = generate_erd_for_category(cat_name, comp_names, components)
        output_file = f"erd_{cat_name}.mmd"
        with open(output_file, 'w') as f:
            f.write(mmd_content)
        print(f"Generated {output_file}")

if __name__ == "__main__":
    main()
