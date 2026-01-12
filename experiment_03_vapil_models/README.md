# Experiment 03: VAPIL Java Model Analysis

## Overview

**Repository:** [veeva/vault-api-library](https://github.com/veeva/vault-api-library)  
**Type:** Official Veeva Java REST API client library  
**Viability:** HIGH  
**Value:** HIGH - Complete metadata models with typed fields

## What Was Found

### Key Model Classes

Located in `src/main/java/com/veeva/vault/vapil/api/model/`:

| Class | Location | Purpose |
|-------|----------|---------|
| `VaultObject` | `metadata/` | Object metadata (30+ properties) |
| `VaultObjectField` | `metadata/` | Field metadata (40+ properties) |
| `DocumentField` | `metadata/` | Document field metadata |
| `VaultObjectPageLayout` | `metadata/` | Page layout metadata |
| `Document` | `common/` | Document data model |
| `ObjectRecord` | `common/` | Object record data model |

### VaultObject Properties (30+ fields)

```java
// Core properties
String name              // Object API name (e.g., "product__v")
String label             // Display label
String label_plural      // Plural label
String description       // Description
String object_class      // Object classification
String source            // SYSTEM, STANDARD, CUSTOM, APPLICATION
String prefix            // Object prefix (e.g., "OT")
Integer order            // Display order

// Feature flags
Boolean allow_attachments     // Attachments enabled
Boolean allow_types           // Object types enabled
Boolean auditable             // Auditing enabled
Boolean dynamic_security      // Dynamic security enabled
Boolean enable_esignatures    // E-signatures enabled
Boolean in_menu               // Shown in menu
Boolean system_managed        // System managed object
Boolean configuration_data    // Configuration data flag
Boolean role_overrides        // Role overrides enabled
Boolean secure_attachments    // Secure attachments
Boolean secure_sharing_settings // Secure sharing
Boolean prevent_record_overwrite // Overwrite prevention

// Relationships and types
List<VaultObjectField> fields       // Field definitions
List<Relationship> relationships    // Object relationships
List<ObjectType> object_types       // Object type definitions
List<String> available_lifecycles   // Available lifecycles
String default_obj_type             // Default object type

// Audit fields
Integer created_by, modified_by
String created_date, modified_date
```

### VaultObjectField Properties (40+ fields)

```java
// Core properties
String name              // Field API name
String label             // Display label
String type              // Field type (String, Object, Picklist, etc.)
String subtype           // Field subtype
Integer order            // Display order
String source            // Field source
List<String> status      // Status values

// Constraints
Boolean required         // Required field
Boolean editable         // Editable field
Boolean unique           // Unique constraint
Boolean encrypted        // Encrypted field
Integer max_length       // Maximum length
Long min_value, max_value // Numeric constraints
Integer scale            // Decimal scale
String value_format      // Value format pattern
String format_mask       // Format mask

// UI properties
Boolean checkbox         // Checkbox display
Boolean list_column      // Show in list view
Boolean facetable        // Facetable for search
Boolean no_copy          // Exclude from copy
Boolean multi_value      // Multi-value field
String help_content      // Help text

// Relationship properties
String picklist                    // Picklist reference
String component                   // Component reference
ObjectReference object             // Object reference
String relationship_type           // reference, parent, child
String relationship_deletion       // block, cascade, set_null
String relationship_criteria       // VQL criteria
String relationship_inbound_name   // Inbound relationship name
String relationship_inbound_label  // Inbound relationship label
String relationship_outbound_name  // Outbound relationship name
Boolean secure_relationship        // Secure relationship
Boolean create_object_inline       // Inline object creation

// Lookup properties
String lookup_relationship_name    // Lookup relationship
String lookup_source_field         // Lookup source field

// Auto-naming
Boolean sequential_naming          // Sequential naming
Boolean system_managed_name        // System managed name
Integer start_number               // Auto-number start
```

### Relationship Class

```java
class Relationship {
    String field                   // Field on this object
    String relationship_name       // Relationship name
    String relationship_label      // Display label
    String relationship_type       // reference_inbound, child, reference_outbound, parent
    String relationship_deletion   // block, cascade, set_null
    ObjectReference object         // Related object reference
}
```

### API Request Classes (40+)

Key metadata-related requests:

| Request Class | Endpoint | Purpose |
|---------------|----------|---------|
| `MetaDataRequest` | `/metadata/vobjects/*` | Object/field metadata |
| `PicklistRequest` | `/objects/picklists/*` | Picklist operations |
| `DocumentRequest` | `/objects/documents/*` | Document operations |
| `QueryRequest` | `/query` | VQL queries |

### Field Types

Based on the models and MDL samples:

| Type | Description |
|------|-------------|
| `String` | Text field with max_length |
| `Number` | Numeric field with scale, min/max |
| `Boolean` | True/false |
| `Date` | Date only |
| `DateTime` | Date and time |
| `Object` | Reference to another object |
| `Picklist` | Reference to picklist |
| `Component` | Reference to component (lifecycle, etc.) |
| `ID` | System identifier |
| `LongText` | Long text/rich text |
| `Currency` | Currency field |
| `Formula` | Calculated formula |
| `Rollup` | Rollup summary |

## Usage

```java
import com.veeva.vault.vapil.api.model.metadata.*;
import com.veeva.vault.vapil.api.request.*;

// Get object metadata
MetaDataRequest request = vaultClient.newRequest(MetaDataRequest.class);
MetaDataObjectResponse response = request.retrieveObjectMetadata("product__v");

VaultObject obj = response.getObject();
System.out.println(obj.getName());       // "product__v"
System.out.println(obj.getLabel());      // "Product"

// Iterate fields
for (VaultObjectField field : obj.getFields()) {
    System.out.printf("%s (%s): %s%n", 
        field.getName(), 
        field.getType(), 
        field.getLabel());
}

// Check relationships
for (VaultObject.Relationship rel : obj.getRelationships()) {
    System.out.printf("Relationship: %s -> %s (%s)%n",
        rel.getField(),
        rel.getObjectReference().getName(),
        rel.getRelationshipType());
}
```

## Conclusion

VAPIL provides **official, typed Java models** for Vault metadata:

1. **Complete VaultObject model** - All object-level properties
2. **Comprehensive VaultObjectField** - 40+ field properties
3. **Relationship modeling** - Full relationship structure
4. **Request classes** - All metadata API endpoints

This complements the MDL schema from meddle by providing the **runtime API perspective** - how metadata is returned from the Vault REST API.
