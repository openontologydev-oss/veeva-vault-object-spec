# Experiment 01: Meddle Schema Extraction

## Overview

**Repository:** [gorkaerana/meddle](https://github.com/gorkaerana/meddle)  
**Type:** Python library for Veeva Vault MDL (Metadata Definition Language)  
**Viability:** HIGH  
**Value:** CRITICAL - Contains the most comprehensive schema available

## What Was Found

### Schema Statistics

| Metric | Count |
|--------|-------|
| **Top-level Components** | 95 |
| **Subcomponents** | 79 |
| **Total Attributes** | 1,285 |

### Key Files

- `src/meddle/validation.json` - Complete schema with all component types, attributes, types, and constraints
- `src/meddle/mdl_grammar.lark` - Lark parser grammar for MDL syntax
- `tests/mdl_examples/` - Sample MDL files and test fixtures

### Component Types Found (95 total)

**Core Object Components:**
- `Object` - Custom objects with fields and relationships
- `Objecttype` - Object type variations
- `Objectlifecycle` - Lifecycle states and transitions
- `Objectworkflow` - Workflow definitions
- `Objectaction` - Custom actions
- `Objectvalidation` - Validation rules

**Document Components:**
- `Doctype` - Document types
- `Docfield` - Document fields
- `Doclifecycle` - Document lifecycles
- `Docrelationshiptype` - Document relationships
- `Documentaction` - Document actions

**Configuration Components:**
- `Picklist` - Picklist definitions with entries
- `Pagelayout` - Page layouts
- `Tab`, `Tabcollection` - UI navigation
- `Dashboard` - Dashboard configurations
- `Securityprofile`, `Permissionset` - Security

**SDK Components:**
- `Recordtrigger` - SDK record triggers
- `Recordaction` - SDK record actions
- `Customwebapi` - Custom REST APIs
- `Userdefinedclass` - Custom Java classes
- `Sdkjob` - SDK job handlers

### Attribute Types Discovered

From the `validation.json`, attributes have these types:

| Type | Description | Example |
|------|-------------|---------|
| `String` | Text with max_length | `label`, `description` |
| `Boolean` | True/false | `active`, `required` |
| `Number` | Numeric values | `order`, `max_value` |
| `Enum` | Allowed values | `type`, `relationship_type` |
| `Object` | Object reference | `object('vsdk_product__c')` |
| `Picklist` | Picklist reference | `picklist('Picklist.status__v')` |
| `Component` | Component reference | `component('Objectlifecycle')` |
| `Subcomponent` | Subcomponent reference | `Field.field_name__c` |
| `SdkCode` | Java source code | `source_code(...)` |
| `Expression` | Formula expressions | `formula(...)` |
| `XMLString` | Embedded XML | `dashboard_markup(...)` |
| `LongString` | Long text (1500+ chars) | `email_body` |

### MDL Grammar Structure

```
Commands:
  CREATE   - Create new component
  RECREATE - Create or replace
  ALTER    - Modify existing (with ADD, MODIFY, DROP subcommands)
  DROP     - Delete component
  RENAME   - Rename component

Logical Operators:
  IF EXISTS
  IF NOT EXISTS

Value Types:
  string   - 'quoted text'
  number   - 123, 45.67
  boolean  - true, false
  xml      - {xml content}
```

## Usage

```python
from meddle import Command
import json

# Load validation schema
with open('src/meddle/validation.json') as f:
    schema = json.load(f)

# List all component types
for component_name in schema.keys():
    print(component_name)

# Get Field subcomponent attributes for Object
object_schema = schema['Object']
field_attrs = object_schema['subcomponents']['Field']['attributes']
for attr_name, attr_info in field_attrs.items():
    print(f"{attr_name}: {attr_info['type_data']}")

# Parse MDL file
mdl = """
RECREATE Object my_object__c (
   label('My Object'),
   active(true)
);
"""
cmd = Command.loads(mdl)
print(cmd.component_type_name)  # "Object"
print(cmd.component_name)       # "my_object__c"
```

## Conclusion

**This is the most valuable resource found.** The `validation.json` file contains a comprehensive schema of all Veeva Vault component types, derived from official Veeva documentation. Combined with the Lark parser, this provides:

1. Complete type definitions for all 95 component types
2. All attribute names with type constraints
3. Enum values for restricted attributes
4. Required vs optional flags
5. Max length constraints for string fields

This can serve as the foundation for generating a complete Veeva Vault object specification.
