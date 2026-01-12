# Experiment 07: Official Documentation Extraction

## Overview

This experiment extracts the complete Veeva Vault MDL component specification directly from official Veeva documentation sources.

## Sources

| Source | URL | Status |
|--------|-----|--------|
| MDL Components Reference | https://developer.veevavault.com/mdl/components | Scraped |
| Vault API Reference (25.3) | https://developer.veevavault.com/api/25.3/ | Scraped |
| Postman Collection | https://god.gw.postman.com/run-collection/25881067-751d9171-bee8-40ff-82f2-c47cab21893c | Link captured |
| Veeva Postman Team | https://www.postman.com/veevavault | Link captured |

## Outputs

### 1. mdl_components_schema.json (216 KB)

Complete JSON schema extracted from the official MDL Components documentation.

**Statistics:**
- **117 top-level components**
- **88 subcomponents**
- **205 total component types**
- **8,227 lines of JSON**

**Structure:**
```json
{
  "source": "https://developer.veevavault.com/mdl/components",
  "extracted_date": "2026-01-12",
  "components": {
    "ComponentName": {
      "description": "Component description",
      "attributes": {
        "attribute_name": {
          "type": "String|Boolean|Number|Enum|ComponentReference|...",
          "required": true|false,
          "maxLength": 60,
          "allowedValues": ["value1", "value2"],
          "multiValue": true|false,
          "description": "Attribute description"
        }
      },
      "subcomponents": {
        "SubcomponentName": {
          "attributes": {...}
        }
      }
    }
  }
}
```

### 2. api_vault_objects_examples.json (38 KB)

Vault Objects API endpoint definitions and response examples extracted from the API Reference.

**Statistics:**
- **15+ endpoints** documented
- **Full cURL request examples**
- **JSON response examples** with real field data
- **1,018 lines of JSON**

**Endpoints included:**
- Retrieve Object Metadata
- Retrieve Object Field Metadata
- Retrieve Object Collection
- Create/Update/Delete Object Records
- Cascade Delete
- Object Types
- Object Record Roles

## Component Categories

### Core Platform (24 components)
- `Object`, `Objecttype`, `Objectlifecycle`, `Objectworkflow`, `Objectaction`, `Objectvalidation`
- `Picklist`, `Tab`, `Tabcollection`, `Dashboard`, `Report`, `Reporttype`
- `Securityprofile`, `Permissionset`, `Atomicsecurity`, `Sharingrule`
- `Job`, `Queue`, `Notificationtemplate`, `Messagegroup`

### Document Management (18 components)
- `Doctype`, `Docfield`, `Docfieldlayout`, `Docfielddependency`
- `Doclifecycle`, `Docrelationshiptype`, `Docmatchingrule`, `Docparticipantrule`
- `Documentaction`, `Documentcheck`, `Documentchecksection`, `Documentstagegroup`
- `Docatomicsecurity`, `Formattedoutput`, `Renditionprofile`, `Renditiontype`

### SDK & Development (12 components)
- `Recordtrigger`, `Recordaction`, `Recordroletrigger`, `Recordworkflowaction`
- `Customwebapi`, `Webapigroup`, `Sdkjob`
- `Userdefinedclass`, `Userdefinedmodel`, `Userdefinedservice`
- `Actiontrigger` (with `Actionblock`, `Draftactionblock`)

### Workflow & Lifecycle (8 components)
- `Workflow`, `Objectworkflow`, `Integrationrule`
- `Lifecyclestatetype`, `Lifecyclestatetypeassociation`
- `Objectlifecyclestagegroup`, `Matchingrule`

### Quality (QMS) Specific (20 components)
- `Qualityteam`, `Qualityteamlifecycleassociation`
- `Qualityrecordcheck`, `Qualityrecordcheckinsight`, `Qualityrecordchecklifecycleassociation`
- `Qualityreasonforchange`, `Qualitycurriculumsmartmatchrule`
- `Qualitydynamicenrollmentrule`, `Qualityexternalnotification`
- `Qmsactionpathconfiguration`, `Qmsautomationusertemplate`
- ... and 9 more

### RIM (Regulatory) Specific (5 components)
- `Rimdoctypeconfig`
- `Rimeventchangedetail`, `Rimeventchangetype`
- `Rimobjectconfiguration`, `Rimobjectmapping`

### Clinical Specific (4 components)
- `Clinicalstandardmapping`
- `Disclosurerule`, `Disclosurexmldoctypemapping`, `Disclosurexmlfieldmapping`

### Mobile & UI (6 components)
- `Page`, `Pagelayout`, `Pagelink`
- `Mobiletab`, `Mobileshareactionconfig`
- `Layoutprofile`, `Layoutrule`

### Other (20 components)
- `Accountmessage`, `Appsecurityrule`, `Checklisttype`
- `Link`, `Tag`, `Tagsecurityrule`, `Vaulttoken`
- `Searchcollection`, `Savedview`, `Signaturepage`
- `Overlaytemplate`, `Labelset`, `Selfevidentcorrection`
- ... and more

## Comparison with Meddle

| Metric | Meddle (Exp 01) | Official Docs (Exp 07) |
|--------|-----------------|------------------------|
| Top-level components | 95 | 117 |
| Subcomponents | 79 | 88 |
| Total types | 174 | 205 |
| Attribute parsing | Merged type_data | Separate type, maxLength, etc. |
| Required field | In type_data | Separate boolean |
| Enum values | Not extracted | Fully extracted |

### Components in Official Docs but NOT in Meddle (22)

1. `Actiontrigger` - Action Triggers (LOW-CODE automation)
2. `Casechildconfig` - MedInquiry case child config
3. `Clientdistribution` - Custom Pages code distributions
4. `Clinicalstandardmapping` - CTMS Transfer mappings
5. `Disclosurerule` - Clinical disclosure rules
6. `Disclosurexmldoctypemapping` - Disclosure doc type mapping
7. `Disclosurexmlfieldmapping` - Disclosure field mapping
8. `Docatomicsecurity` - Document atomic security
9. `Docfielddependency` - Document field dependencies
10. `Docfieldlayout` - Document field layout sections
11. `Docparticipantrule` - Document participant rules
12. `Documentcheck` - Document readiness checks
13. `Documentchecksection` - Document check sections
14. `Documentstagegroup` - Document lifecycle stages
15. `Lifecyclestatetypeassociation` - State type associations
16. `Link` - Web Actions
17. `Mobileshareactionconfig` - Mobile sharing actions
18. `Mobiletab` - Vault Mobile tabs
19. `Objectlifecyclestagegroup` - Object lifecycle stages
20. `Outboundemaildomain` - Outbound email domain config
21. `Overlaytemplate` - Document overlays
22. `Pagelink` - Page links to custom URLs

## Usage

### Load the MDL Components Schema

```python
import json

with open('mdl_components_schema.json') as f:
    schema = json.load(f)

# List all components
for name in sorted(schema['components'].keys()):
    comp = schema['components'][name]
    sub_count = len(comp.get('subcomponents', {}))
    print(f"{name}: {len(comp['attributes'])} attrs, {sub_count} subs")

# Get Object component details
obj = schema['components']['Object']
for attr_name, attr_info in obj['attributes'].items():
    print(f"  {attr_name}: {attr_info['type']} (required={attr_info.get('required', False)})")
```

### Load API Examples

```python
import json

with open('api_vault_objects_examples.json') as f:
    api = json.load(f)

# Get object metadata response example
endpoint = api['endpoints']['retrieve_object_metadata']
print(f"Path: {endpoint['path']}")
print(f"Response: {json.dumps(endpoint['response_example'], indent=2)}")
```

## Postman Collection

The official Veeva Postman collection can be accessed via:

1. **Run in Postman button**: [Import Collection](https://god.gw.postman.com/run-collection/25881067-751d9171-bee8-40ff-82f2-c47cab21893c)
2. **Veeva Vault Postman Team**: [https://www.postman.com/veevavault](https://www.postman.com/veevavault)

Note: The Postman collection requires a Postman account to import.

## Extraction Date

2026-01-12

## Next Steps

1. **Create TypeScript/Python type definitions** from the MDL schema
2. **Compare field-by-field** with meddle's validation.json
3. **Download Postman collection** via Postman account
4. **Parse API response examples** for additional field metadata
