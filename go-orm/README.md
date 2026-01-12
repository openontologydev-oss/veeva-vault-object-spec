# Veeva Vault Go ORM

A comprehensive Go ORM representation of Veeva Vault objects and components, generated from the official MDL Components documentation.

## Installation

```bash
go get github.com/veeva/vault-orm
```

## Quick Start

```go
package main

import (
    "log"
    
    "github.com/veeva/vault-orm/vault"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    // Connect to database
    db, err := gorm.Open(sqlite.Open("vault.db"), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // Run migrations
    if err := vault.Migrate(db); err != nil {
        log.Fatal(err)
    }

    // Create indexes for performance
    if err := vault.CreateIndexes(db); err != nil {
        log.Fatal(err)
    }

    // Create a product object
    product := vault.ExampleProduct()
    if err := db.Create(product).Error; err != nil {
        log.Fatal(err)
    }

    // Create a picklist
    picklist := vault.ExamplePicklist()
    if err := db.Create(picklist).Error; err != nil {
        log.Fatal(err)
    }

    // Create a lifecycle
    lifecycle := vault.ExampleLifecycle()
    if err := db.Create(lifecycle).Error; err != nil {
        log.Fatal(err)
    }

    log.Println("Vault schema created successfully!")
}
```

## Models

### Core Objects

| Model | MDL Component | Description |
|-------|---------------|-------------|
| `VaultObject` | Object | Vault object definition |
| `ObjectField` | Object.Field | Field on an object |
| `ObjectIndex` | Object.Index | Index on an object |
| `ObjectType` | Objecttype | Object type variation |
| `ObjectLifecycle` | Objectlifecycle | Object lifecycle |
| `ObjectLifecycleState` | Objectlifecyclestate | Lifecycle state |

### Documents

| Model | MDL Component | Description |
|-------|---------------|-------------|
| `Doctype` | Doctype | Document type |
| `Docfield` | Docfield | Document field |
| `Doclifecycle` | Doclifecycle | Document lifecycle |
| `Docrelationshiptype` | Docrelationshiptype | Document relationship |

### Workflows & Triggers

| Model | MDL Component | Description |
|-------|---------------|-------------|
| `Workflow` | Workflow | Workflow definition |
| `ObjectWorkflow` | Objectworkflow | Object-specific workflow |
| `Recordtrigger` | Recordtrigger | SDK record trigger |
| `Actiontrigger` | Actiontrigger | Low-code action trigger |
| `Recordaction` | Recordaction | Custom record action |

### Security

| Model | MDL Component | Description |
|-------|---------------|-------------|
| `Securityprofile` | Securityprofile | Security profile |
| `Permissionset` | Permissionset | Permission set |
| `Sharingrule` | Sharingrule | Sharing rule |
| `Atomicsecurity` | Atomicsecurity | Atomic security config |

### SDK Components

| Model | MDL Component | Description |
|-------|---------------|-------------|
| `Customwebapi` | Customwebapi | Custom REST API |
| `Userdefinedclass` | Userdefinedclass | Custom Java class |
| `Sdkjob` | Sdkjob | SDK job handler |
| `Queue` | Queue | Message queue |

### Configuration

| Model | MDL Component | Description |
|-------|---------------|-------------|
| `Picklist` | Picklist | Picklist definition |
| `PicklistEntry` | Picklistentry | Picklist value |
| `Tab` | Tab | Navigation tab |
| `Dashboard` | Dashboard | Dashboard |
| `Report` | Report | Report |

## Field Types

```go
const (
    FieldTypeString          FieldType = "String"
    FieldTypeNumber          FieldType = "Number"
    FieldTypeBoolean         FieldType = "Boolean"
    FieldTypeDate            FieldType = "Date"
    FieldTypeDateTime        FieldType = "DateTime"
    FieldTypeObject          FieldType = "Object"
    FieldTypePicklist        FieldType = "Picklist"
    FieldTypeID              FieldType = "ID"
    FieldTypeUsers           FieldType = "users"
    FieldTypeComponent       FieldType = "Component"
    FieldTypeDocuments       FieldType = "documents"
    FieldTypeLongText        FieldType = "LongText"
    FieldTypeRichText        FieldType = "RichText"
    FieldTypeBinary          FieldType = "Binary"
    FieldTypeObjectReference FieldType = "ObjectReference"
    FieldTypeObjectParent    FieldType = "ObjectParent"
)
```

## Relationship Types

```go
const (
    RelationshipTypeParent           RelationshipType = "parent"
    RelationshipTypeReference        RelationshipType = "reference"
    RelationshipTypeChild            RelationshipType = "child"
    RelationshipTypeReferenceInbound RelationshipType = "reference_inbound"
    RelationshipTypeReferenceOutbound RelationshipType = "reference_outbound"
)

const (
    RelationshipDeletionBlock   RelationshipDeletion = "block"
    RelationshipDeletionCascade RelationshipDeletion = "cascade"
    RelationshipDeletionSetNull RelationshipDeletion = "setnull"
)
```

## Trigger Events

```go
const (
    TriggerEventBeforeInsert TriggerEvent = "BEFORE_INSERT"
    TriggerEventBeforeUpdate TriggerEvent = "BEFORE_UPDATE"
    TriggerEventBeforeDelete TriggerEvent = "BEFORE_DELETE"
    TriggerEventAfterInsert  TriggerEvent = "AFTER_INSERT"
    TriggerEventAfterUpdate  TriggerEvent = "AFTER_UPDATE"
    TriggerEventAfterDelete  TriggerEvent = "AFTER_DELETE"
)
```

## Example: Creating a Product Object

```go
product := &vault.VaultObject{
    ComponentBase: vault.ComponentBase{
        Name:   "product__c",
        Label:  "Product",
        Active: true,
    },
    LabelPlural:      "Products",
    Source:           vault.SourceCustom,
    ObjectClass:      vault.ObjectClassBase,
    InMenu:           true,
    AllowAttachments: true,
    Audit:            true,
    Fields: []vault.ObjectField{
        {
            Name:     "name__v",
            Label:    "Name",
            Type:     vault.FieldTypeString,
            Required: true,
            MaxLength: func() *int { v := 128; return &v }(),
        },
        {
            Name:              "category__c",
            Label:             "Category",
            Type:              vault.FieldTypePicklist,
            PicklistName:      func() *string { v := "product_category__c"; return &v }(),
        },
        {
            Name:              "manufacturer__c",
            Label:             "Manufacturer",
            Type:              vault.FieldTypeObject,
            ReferencedObject:  func() *string { v := "company__c"; return &v }(),
            RelationshipType:  func() *vault.RelationshipType { v := vault.RelationshipTypeReference; return &v }(),
        },
    },
}

db.Create(product)
```

## Example: Creating a Lifecycle

```go
lifecycle := &vault.ObjectLifecycle{
    ComponentBase: vault.ComponentBase{
        Name:   "product_lifecycle__c",
        Label:  "Product Lifecycle",
        Active: true,
    },
    States: []vault.ObjectLifecycleState{
        {
            ComponentBase: vault.ComponentBase{
                Name:   "draft__c",
                Label:  "Draft",
                Active: true,
            },
            RecordStatus: vault.RecordStatusActive,
        },
        {
            ComponentBase: vault.ComponentBase{
                Name:   "approved__c",
                Label:  "Approved",
                Active: true,
            },
            RecordStatus: vault.RecordStatusActive,
        },
    },
    Roles: []vault.ObjectLifecycleRole{
        {
            Name:            "owner__v",
            Active:          true,
            ApplicationRole: "owner__v",
            Permissions:     vault.StringSlice{"Read", "Edit", "Delete"},
        },
    },
}

db.Create(lifecycle)
```

## Database Support

The ORM is built on GORM and supports:

- PostgreSQL (recommended)
- SQLite
- MySQL
- SQL Server

## Source

Generated from the official Veeva MDL Components documentation:
- https://developer.veevavault.com/mdl/components
- 117 top-level components
- 88 subcomponents
- Generated: 2026-01-12

## License

MIT License
