# Veeva Vault Object Specification Research

This repository contains a comprehensive research effort to extract and document the Veeva Vault object specification from publicly available sources. Since Veeva doesn't publish a central object specification, this project aggregates information from multiple sources.

## Executive Summary

**Key Finding:** The most complete Vault object specification is available through the `meddle` library, which contains a JSON schema with **95 component types**, **79 subcomponents**, and **1,285 attribute definitions** scraped from official Veeva documentation.

## Experiment Results

| Experiment | Source | Viability | Value | Key Output |
|------------|--------|-----------|-------|------------|
| **01** | [meddle](https://github.com/gorkaerana/meddle) | HIGH | CRITICAL | `validation.json` with complete schema |
| **02** | [overpack](https://github.com/gorkaerana/overpack) | HIGH | HIGH | VPK parser with meddle integration |
| **03** | [VAPIL](https://github.com/veeva/vault-api-library) | HIGH | HIGH | Java models for API responses |
| **04** | [OpenAPI Spec](https://github.com/fderuiter/veeva-vault-openapi-spec) | MEDIUM | MEDIUM | 310 endpoints (no schemas) |
| **05** | [Direct Data API](https://github.com/veeva/Vault-Direct-Data-API-Accelerators) | MEDIUM | MEDIUM | Python dataclasses |
| **06** | Veeva SDK Samples | HIGH | HIGH | 55 real MDL files |

## Resources Discovered

### Primary Sources (Recommended)

1. **`gorkaerana/meddle`** - Python MDL parser with validation schema
   - Contains `validation.json` with all component/attribute definitions
   - Lark-based parser for MDL syntax
   - 175+ test fixtures from real Vault packages

2. **`veeva/vault-api-library` (VAPIL)** - Official Java SDK
   - `VaultObject` - 30+ properties for object metadata
   - `VaultObjectField` - 40+ properties for field metadata
   - 40+ request classes for all API endpoints

3. **Veeva SDK Samples** - Real-world MDL examples
   - `vsdk-object-sample` - Objects, relationships, lifecycles
   - `vsdk-document-sample` - Document types, fields, lifecycles
   - 55 MDL files demonstrating all patterns

### Secondary Sources

4. **`gorkaerana/overpack`** - VPK file parser
   - Parses Vault Package files (ZIP with MDL)
   - Integrates with meddle for MDL parsing

5. **`fderuiter/veeva-vault-openapi-spec`** - OpenAPI 3.0 spec
   - 310 endpoint definitions
   - Missing response schemas (limited value)

6. **`veeva/Vault-Direct-Data-API-Accelerators`** - Python models
   - `VaultObjectField` dataclass (35+ properties)
   - `DocumentField` dataclass

## Component Types (95 total)

From the meddle `validation.json`:

**Core Object Components:**
- `Object` - Custom objects with fields
- `Objecttype` - Object type variations
- `Objectlifecycle` - Lifecycle states/transitions
- `Objectworkflow` - Workflow definitions
- `Objectaction` - Custom actions
- `Objectvalidation` - Validation rules

**Document Components:**
- `Doctype` - Document types
- `Docfield` - Document fields
- `Doclifecycle` - Document lifecycles
- `Docrelationshiptype` - Document relationships

**Configuration:**
- `Picklist` - Picklist definitions
- `Pagelayout` - Page layouts
- `Tab`, `Tabcollection` - Navigation
- `Dashboard` - Dashboards
- `Securityprofile`, `Permissionset` - Security

**SDK Components:**
- `Recordtrigger`, `Recordaction` - SDK triggers/actions
- `Customwebapi` - Custom REST APIs
- `Userdefinedclass` - Custom Java classes
- `Sdkjob` - SDK job handlers

[Full list of 95 components in experiment_01_meddle_schema/README.md]

## Field Types

| Type | Description | MDL Example |
|------|-------------|-------------|
| `String` | Text field | `type('String'), max_length(128)` |
| `Number` | Numeric | `type('Number'), scale(2)` |
| `Boolean` | True/false | `type('Boolean'), checkbox(false)` |
| `Date` | Date only | `type('Date')` |
| `DateTime` | Date and time | `type('DateTime')` |
| `Object` | Object reference | `type('Object'), object('product__v')` |
| `Picklist` | Picklist reference | `type('Picklist'), picklist('status__v')` |
| `Component` | Component reference | `type('Component'), component('Objectlifecycle')` |
| `ID` | System identifier | `type('ID')` |
| `LongText` | Long text | `type('LongText')` |

## Relationship Types

| Type | Description |
|------|-------------|
| `reference` | Standard lookup relationship |
| `parent` | Parent-child (hierarchical) |
| `child` | Child side of parent relationship |

**Deletion Behaviors:**
- `block` - Prevent deletion if related records exist
- `cascade` - Delete related records
- `setnull` - Set reference to null

## Quick Start

### Extract Schema from Meddle

```python
import json

# Load the complete schema
with open('experiment_01_meddle_schema/meddle/src/meddle/validation.json') as f:
    schema = json.load(f)

# List all component types
print(f"Components: {len(schema)}")
for component in sorted(schema.keys()):
    print(f"  - {component}")

# Get Field attributes for Object component
object_schema = schema['Object']
field_attrs = object_schema['subcomponents']['Field']['attributes']
for name, info in field_attrs.items():
    print(f"{name}: {info['type_data']}")
```

### Parse MDL Files

```python
from meddle import Command

mdl = open('experiment_06_sdk_samples/vsdk-object-sample/.../Object.vsdk_product__c.mdl').read()
cmd = Command.loads(mdl)

print(f"Object: {cmd.component_name}")
for field in cmd.components:
    if field.component_type_name == 'Field':
        field_attrs = {a.name: a.value for a in field.attributes}
        print(f"  {field.component_name}: {field_attrs.get('type')}")
```

### Use VAPIL Java Models

```java
import com.veeva.vault.vapil.api.model.metadata.*;

VaultObject obj = response.getObject();
for (VaultObjectField field : obj.getFields()) {
    System.out.printf("%s (%s)%n", field.getName(), field.getType());
}
```

## Directory Structure

```
veeva-vault-object-spec/
├── README.md                          # This file
├── experiment_01_meddle_schema/       # MDL schema extraction
│   ├── README.md
│   └── meddle/                        # Cloned repo
├── experiment_02_vpk_extraction/      # VPK parsing
│   ├── README.md
│   └── overpack/                      # Cloned repo
├── experiment_03_vapil_models/        # Java API models
│   ├── README.md
│   └── vault-api-library/             # Cloned repo
├── experiment_04_openapi_spec/        # OpenAPI specification
│   ├── README.md
│   └── veeva-vault-openapi-spec/      # Cloned repo
├── experiment_05_direct_data_models/  # Python dataclasses
│   ├── README.md
│   └── Vault-Direct-Data-API-Accelerators/  # Cloned repo
└── experiment_06_sdk_samples/         # Real MDL examples
    ├── README.md
    ├── vsdk-object-sample/
    ├── vsdk-document-sample/
    ├── vsdk-spark-integration-rules-sample/
    └── vsdk-common-services-sample/
```

## Recommendations

### For Building a Complete Object Specification

1. **Start with `meddle/validation.json`** - Most comprehensive schema
2. **Supplement with VAPIL models** - API response perspective
3. **Reference SDK sample MDL files** - Real-world patterns
4. **Use `overpack` for VPK parsing** - Extract from actual packages

### For Runtime Metadata Discovery

1. Use VAPIL's `MetaDataRequest` class
2. Call `/api/{version}/metadata/vobjects/{object_name}`
3. Parse response using `VaultObject` / `VaultObjectField` models

### For Generating Types/SDK

1. Parse `validation.json` to extract types and constraints
2. Generate code from the schema (TypeScript, Python, etc.)
3. Validate against SDK sample MDL files

## Limitations

1. **No Central Published Spec** - Veeva doesn't publish a unified schema
2. **Schema Lag** - Community schemas may lag behind Veeva releases
3. **SDK Code Not Parsed** - `SdkCode` and `Expression` types are raw strings
4. **Vault-Specific Objects** - Some objects vary by Vault application type

## License

This research repository aggregates open-source projects under their respective licenses:
- meddle/overpack: Check repository for license
- VAPIL: Apache 2.0
- veeva-vault-openapi-spec: MIT
- Veeva SDK Samples: Check Veeva's terms

## Contributing

To extend this research:
1. Add new sources to appropriate experiment directories
2. Update experiment READMEs with findings
3. Update this main README with summary changes
