# Veeva Vault Object Specification Research

This repository contains a comprehensive research effort to extract and document the Veeva Vault object specification from publicly available sources. Since Veeva doesn't publish a central object specification, this project aggregates information from multiple sources.

## Executive Summary

**Key Finding:** We have successfully extracted the **complete Veeva Vault MDL specification** with:
- **117 top-level component types**
- **88 subcomponent types**
- **205 total component definitions**
- **15+ API endpoint examples** with response schemas

The authoritative source is now the **official Veeva MDL Components documentation** (Experiment 07), which we have fully scraped and converted to JSON.

## Experiment Results

| Experiment | Source | Viability | Value | Key Output |
|------------|--------|-----------|-------|------------|
| **01** | [meddle](https://github.com/gorkaerana/meddle) | HIGH | HIGH | `validation.json` (95 components) |
| **02** | [overpack](https://github.com/gorkaerana/overpack) | HIGH | HIGH | VPK parser with meddle integration |
| **03** | [VAPIL](https://github.com/veeva/vault-api-library) | HIGH | HIGH | Java models for API responses |
| **04** | [OpenAPI Spec](https://github.com/fderuiter/veeva-vault-openapi-spec) | MEDIUM | MEDIUM | 310 endpoints (no schemas) |
| **05** | [Direct Data API](https://github.com/veeva/Vault-Direct-Data-API-Accelerators) | MEDIUM | MEDIUM | Python dataclasses |
| **06** | Veeva SDK Samples | HIGH | HIGH | 55 real MDL files |
| **07** | **Official Veeva Docs** | **CRITICAL** | **CRITICAL** | **117 components + API examples** |

## Best Sources by Use Case

| Use Case | Recommended Source | Location |
|----------|-------------------|----------|
| **Complete MDL schema** | Experiment 07 | `experiment_07_postman_collection/mdl_components_schema.json` |
| **Python MDL parsing** | Experiment 01 | `experiment_01_meddle_schema/meddle/` |
| **Java API integration** | Experiment 03 | `experiment_03_vapil_models/vault-api-library/` |
| **Real MDL examples** | Experiment 06 | `experiment_06_sdk_samples/vsdk-*/` |
| **API endpoint reference** | Experiment 07 | `experiment_07_postman_collection/api_vault_objects_examples.json` |
| **VPK file parsing** | Experiment 02 | `experiment_02_vpk_extraction/overpack/` |

## Component Types (117 Total)

### Core Platform (24)
- `Object`, `Objecttype`, `Objectlifecycle`, `Objectworkflow`, `Objectaction`, `Objectvalidation`
- `Picklist`, `Tab`, `Tabcollection`, `Dashboard`, `Report`, `Reporttype`
- `Securityprofile`, `Permissionset`, `Atomicsecurity`, `Sharingrule`
- `Job`, `Queue`, `Notificationtemplate`, `Messagegroup`, `Messageprocessor`

### Document Management (18)
- `Doctype`, `Docfield`, `Docfieldlayout`, `Docfielddependency`
- `Doclifecycle`, `Docrelationshiptype`, `Docmatchingrule`, `Docparticipantrule`
- `Documentaction`, `Documentcheck`, `Documentchecksection`, `Documentstagegroup`
- `Docatomicsecurity`, `Formattedoutput`, `Renditionprofile`, `Renditiontype`

### SDK & Development (12)
- `Recordtrigger`, `Recordaction`, `Recordroletrigger`, `Recordworkflowaction`
- `Customwebapi`, `Webapigroup`, `Sdkjob`
- `Userdefinedclass`, `Userdefinedmodel`, `Userdefinedservice`
- `Actiontrigger` (with `Actionblock`, `Draftactionblock`)

### Quality (QMS) Specific (20)
- `Qualityteam`, `Qualityrecordcheck`, `Qualityreasonforchange`
- `Qmsactionpathconfiguration`, `Qmsautomationusertemplate`
- And 15 more quality-specific components

### RIM (Regulatory) Specific (5)
- `Rimdoctypeconfig`, `Rimeventchangedetail`, `Rimeventchangetype`
- `Rimobjectconfiguration`, `Rimobjectmapping`

### Clinical Specific (4)
- `Clinicalstandardmapping`, `Disclosurerule`
- `Disclosurexmldoctypemapping`, `Disclosurexmlfieldmapping`

### Other (34)
- See `experiment_07_postman_collection/README.md` for full categorization

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
| `Enum` | Enumerated values | See schema for `allowedValues` |

## Quick Start

### Load Complete MDL Schema (Recommended)

```python
import json

# Load the complete schema from official docs
with open('experiment_07_postman_collection/mdl_components_schema.json') as f:
    schema = json.load(f)

# List all 117 component types
print(f"Components: {len(schema['components'])}")
for name in sorted(schema['components'].keys()):
    comp = schema['components'][name]
    attr_count = len(comp.get('attributes', {}))
    sub_count = len(comp.get('subcomponents', {}))
    print(f"  {name}: {attr_count} attrs, {sub_count} subs")

# Get Object field details
obj = schema['components']['Object']
for attr_name, attr_info in obj['attributes'].items():
    type_str = attr_info.get('type', 'unknown')
    required = attr_info.get('required', False)
    print(f"  {attr_name}: {type_str} (required={required})")
```

### Load API Examples

```python
import json

with open('experiment_07_postman_collection/api_vault_objects_examples.json') as f:
    api = json.load(f)

# List all documented endpoints
for name, endpoint in api['endpoints'].items():
    print(f"{endpoint['method']} {endpoint['path']}")
    print(f"  {endpoint['description']}")
```

### Parse MDL Files (using meddle)

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

## Official Veeva Resources

| Resource | URL |
|----------|-----|
| MDL Components Reference | https://developer.veevavault.com/mdl/components |
| API Reference (25.3 GA) | https://developer.veevavault.com/api/25.3/ |
| Postman Collection | https://god.gw.postman.com/run-collection/25881067-751d9171-bee8-40ff-82f2-c47cab21893c |
| Veeva Postman Team | https://www.postman.com/veevavault |
| Vault Java SDK Javadocs | http://repo.veevavault.com/javadoc/vault-sdk-api/25.3.2/docs/api/index.html |
| Vault Help (Platform) | https://platform.veevavault.help/ |

## Directory Structure

```
veeva-vault-object-spec/
├── README.md                              # This file
├── experiment_01_meddle_schema/           # Python MDL parser
│   ├── README.md
│   └── meddle/                            # 95 components, Lark parser
├── experiment_02_vpk_extraction/          # VPK file parser
│   ├── README.md
│   └── overpack/
├── experiment_03_vapil_models/            # Official Java SDK
│   ├── README.md
│   └── vault-api-library/                 # VaultObject, VaultObjectField
├── experiment_04_openapi_spec/            # OpenAPI 3.0 spec
│   ├── README.md
│   └── veeva-vault-openapi-spec/          # 310 endpoints
├── experiment_05_direct_data_models/      # Python dataclasses
│   ├── README.md
│   └── Vault-Direct-Data-API-Accelerators/
├── experiment_06_sdk_samples/             # Real MDL examples
│   ├── README.md
│   ├── vsdk-object-sample/                # 55 MDL files total
│   ├── vsdk-document-sample/
│   ├── vsdk-spark-integration-rules-sample/
│   └── vsdk-common-services-sample/
└── experiment_07_postman_collection/      # OFFICIAL DOCS EXTRACTION
    ├── README.md
    ├── mdl_components_schema.json         # 117 components, 88 subcomponents
    └── api_vault_objects_examples.json    # 15+ API endpoints with examples
```

## Comparison: Meddle vs Official Docs

| Metric | Meddle (Exp 01) | Official Docs (Exp 07) |
|--------|-----------------|------------------------|
| Top-level components | 95 | **117** |
| Subcomponents | 79 | **88** |
| Total types | 174 | **205** |
| Attribute types | Merged in `type_data` | Separate fields |
| Required flag | In `type_data` string | Boolean field |
| Enum values | Not extracted | **Fully extracted** |
| Data freshness | Community-maintained | 2026-01-12 |

### Components in Official Docs but NOT in Meddle (22)

`Actiontrigger`, `Casechildconfig`, `Clientdistribution`, `Clinicalstandardmapping`, 
`Disclosurerule`, `Disclosurexmldoctypemapping`, `Disclosurexmlfieldmapping`, 
`Docatomicsecurity`, `Docfielddependency`, `Docfieldlayout`, `Docparticipantrule`, 
`Documentcheck`, `Documentchecksection`, `Documentstagegroup`, `Lifecyclestatetypeassociation`, 
`Link`, `Mobileshareactionconfig`, `Mobiletab`, `Objectlifecyclestagegroup`, 
`Outboundemaildomain`, `Overlaytemplate`, `Pagelink`

## Recommendations

### For Building a Type-Safe SDK

1. **Use `experiment_07_postman_collection/mdl_components_schema.json`** as the source of truth
2. Generate types from the JSON schema (TypeScript, Python dataclasses, etc.)
3. Validate against real MDL files from Experiment 06

### For MDL Parsing

1. **Use `meddle` library** from Experiment 01
2. Extend its `validation.json` with missing components from Experiment 07

### For API Integration

1. **Use VAPIL (Java)** from Experiment 03 - official Veeva library
2. Reference API examples from `api_vault_objects_examples.json`

### For Understanding Vault Patterns

1. Study the 55 real MDL files in Experiment 06
2. See how objects, lifecycles, workflows, and triggers are structured

## Limitations

1. **No Central Published Spec** - This research aggregates multiple sources
2. **Schema Lag** - Community schemas may lag behind Veeva releases
3. **SDK Code Not Parsed** - `SdkCode` and `Expression` types are raw strings
4. **Vault-Specific Objects** - Some objects vary by Vault application type (QMS, RIM, Clinical, etc.)
5. **Postman Collection** - Requires Postman account to import

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

## Research Date

Last updated: 2026-01-12
