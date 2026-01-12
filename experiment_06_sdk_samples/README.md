# Experiment 06: SDK Sample MDL Extraction

## Overview

**Source:** Official Veeva SDK sample repositories  
**Type:** Real-world MDL files from VPK packages  
**Viability:** HIGH  
**Value:** HIGH - Contains 55 production-quality MDL examples

## What Was Found

### Repositories Cloned

| Repository | MDL Files | Key Components |
|------------|-----------|----------------|
| `vsdk-object-sample` | 15 | Objects, Picklists, Lifecycles, Workflows |
| `vsdk-document-sample` | 40 | Doctypes, Docfields, Doclifecycles |
| `vsdk-spark-integration-rules-sample` | - | Integration rules (VPK-based) |
| `vsdk-common-services-sample` | - | VPK with Object definitions |

### MDL Files Discovered (55 total)

**Object Definitions:**
- `Object.vsdk_product__c.mdl` - Product with 20+ fields
- `Object.vsdk_country__c.mdl` - Country reference object
- `Object.vsdk_region__c.mdl` - Region reference object
- `Object.vsdk_country_brand__c.mdl` - Country Brand with relationships
- `Object.vsdk_document_task__c.mdl` - Document task object

**Picklists:**
- `Picklist.vsdk_product_family__c.mdl`
- `Picklist.vsdk_market_size__c.mdl`
- `Picklist.vsdk_therapeutic_area__c.mdl`

**Lifecycles:**
- `Objectlifecycle.vsdk_product_trade_lifecycle__c.mdl`
- `Objectlifecycle.vsdk_country_brand_penetration_lifecycle__c.mdl`
- `Doclifecycle.vsdk_document_lifecycle__c.mdl`

**Workflows:**
- `Objectworkflow.vsdk_country_brand_penetration__c.mdl`

**Document Components:**
- `Doctype.vsdk_document__c.mdl`
- `Docfield.vsdk_external_id__c.mdl`
- `Docfieldlayout.vsdk_sample_section__c.mdl`

**Actions & Layouts:**
- `Objectaction.vsdk_product__c.run_vsdk_change_owner_user_action__c.mdl`
- `Pagelayout.vsdk_product_detail_page_layout__c.mdl`
- `Pagelayout.vsdk_country_brand_detail_page_layout__c.mdl`
- `Notificationtemplate.vsdk_object_notification_template__c.mdl`

### Sample Object Definition

From `Object.vsdk_product__c.mdl`:

```mdl
RECREATE Object vsdk_product__c (
   label('vSDK Product'),
   label_plural('vSDK Products'),
   active(true),
   in_menu(true),
   allow_attachments(false),
   enable_esignatures(false),
   audit(true),
   dynamic_security(false),
   system_managed(false),
   available_lifecycles('Objectlifecycle.vsdk_product_trade_lifecycle__c'),
   object_class('base'),
   allow_types(false),
   
   Field name__v(
      label('Product Name'),
      type('String'),
      required(true),
      unique(true),
      max_length(128),
      list_column(true)
   ),
   
   Field abbreviation__c(
      label('Product Abbreviation'),
      type('String'),
      max_length(10),
      help_content('Internal abbreviation or code')
   ),
   
   Field compound__c(
      label('Compound'),
      type('Boolean'),
      checkbox(false)
   ),
   
   Field expected_date__c(
      label('Expected Date'),
      type('Date')
   ),
   
   Field region__c(
      label('Region'),
      type('Object'),
      required(true),
      object('vsdk_region__c'),
      relationship_type('reference'),
      relationship_outbound_name('vsdk_region__cr'),
      relationship_inbound_name('vsdk_products__cr'),
      relationship_deletion('block')
   ),
   
   Field product_family__c(
      label('Product Family'),
      type('Picklist'),
      multi_value(true),
      picklist('Picklist.vsdk_product_family__c')
   ),
   
   Field lifecycle__v(
      label('Lifecycle'),
      type('Component'),
      component('Objectlifecycle')
   ),
   
   Field status__v(
      label('Status'),
      type('Picklist'),
      picklist('Picklist.default_status__v')
   )
);
```

### Field Types Demonstrated

| Type | Example | Attributes |
|------|---------|------------|
| `String` | `name__v` | `max_length`, `unique` |
| `Boolean` | `compound__c` | `checkbox` |
| `Date` | `expected_date__c` | - |
| `DateTime` | `created_date__v` | - |
| `Object` | `region__c` | `object`, `relationship_type`, `relationship_deletion` |
| `Picklist` | `status__v` | `picklist`, `multi_value` |
| `Component` | `lifecycle__v` | `component` |
| `ID` | `id` | System-generated |

### Relationship Patterns

**Object Reference:**
```mdl
Field region__c(
   type('Object'),
   object('vsdk_region__c'),
   relationship_type('reference'),
   relationship_outbound_name('vsdk_region__cr'),
   relationship_inbound_name('vsdk_products__cr'),
   relationship_deletion('block')
)
```

**User Reference:**
```mdl
Field owner__c(
   type('Object'),
   object('user__sys'),
   relationship_type('reference'),
   relationship_deletion('block'),
   secure_relationship(false)
)
```

### VPK Files Available

```
vsdk-object-sample/deploy-vpk/
├── vsdk-object-sample-components/
│   └── vsdk-object-sample-components.vpk
└── code/
    └── vsdk-object-sample-action-code.vpk

vsdk-document-sample/deploy-vpk/
├── code/vsdk-document-sample-code.vpk
└── components/
    ├── Base_vsdk-document-sample-components.vpk
    ├── Clinical_vsdk-document-sample-components.vpk
    ├── Multichannel_vsdk-document-sample-components.vpk
    ├── Quality_vsdk-document-sample-components.vpk
    └── RIM_vsdk-document-sample-components.vpk
```

## Parsing with Meddle

```python
from meddle import Command

# Read MDL file
with open('Object.vsdk_product__c.mdl') as f:
    mdl = f.read()

# Parse
cmd = Command.loads(mdl)

print(f"Component: {cmd.component_type_name}")  # "Object"
print(f"Name: {cmd.component_name}")            # "vsdk_product__c"

# Object-level attributes
for attr in cmd.attributes:
    print(f"  {attr.name}: {attr.value}")

# Fields (subcomponents)
for field in cmd.components:
    if field.component_type_name == 'Field':
        print(f"\nField: {field.component_name}")
        for attr in field.attributes:
            print(f"    {attr.name}: {attr.value}")
```

## Conclusion

The SDK sample repositories provide **real-world MDL examples** that demonstrate:

1. **Object structure** - Full object definitions with all attributes
2. **Field patterns** - All field types with proper attribute combinations
3. **Relationships** - Object references, user references, picklists
4. **Lifecycles** - Object and document lifecycle configurations
5. **Workflows** - Object workflow definitions
6. **Actions** - Custom actions and notifications

These examples, combined with the `meddle` parser, allow extracting and understanding production-quality Vault configurations.
