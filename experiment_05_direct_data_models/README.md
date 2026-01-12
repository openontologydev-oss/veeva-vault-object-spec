# Experiment 05: Direct Data API Python Models

## Overview

**Repository:** [veeva/Vault-Direct-Data-API-Accelerators](https://github.com/veeva/Vault-Direct-Data-API-Accelerators)  
**Type:** Official Veeva Python accelerator for data extraction  
**Viability:** MEDIUM  
**Value:** MEDIUM - Contains useful Python dataclasses for field metadata

## What Was Found

### Purpose

This accelerator is designed for **high-volume, read-only data extraction** from Veeva Vault to external systems (Snowflake, Databricks, Redshift, Azure SQL, SQLite).

### Key Python Models

Located in `common/api/model/`:

#### VaultObjectField (`metadata/vault_object_field.py`)

```python
@dataclass
class VaultObjectField(VaultModel):
    checkbox: bool = None
    component: str = None
    created_by: int = None
    created_date: str = None
    create_object_inline: bool = None
    editable: bool = None
    encrypted: bool = None
    facetable: bool = None
    help_content: str = None
    label: str = None
    list_column: bool = None
    lookup_relationship_name: str = None
    lookup_source_field: str = None
    max_length: int = None
    max_value: int = None
    min_value: int = None
    multi_part_field: str = None
    multi_part_readonly: str = None
    multi_value: bool = None
    name: str = None
    no_copy: bool = None
    object: ObjectReference = None
    order: int = None
    picklist: str = None
    relationship_criteria: str = None
    relationship_deletion: str = None
    relationship_inbound_label: str = None
    relationship_inbound_name: str = None
    relationship_outbound_name: str = None
    relationship_type: str = None
    required: bool = None
    scale: int = None
    secure_relationship: bool = None
    sequential_naming: bool = None
    source: str = None
    start_number: int = None
    status: List[str] = None
    system_managed_name: bool = None
    type: str = None
    value_format: str = None
    unique: bool = None
```

#### DocumentField (`component/document_field.py`)

```python
@dataclass
class DocumentField(VaultModel):
    definedIn: str = None
    definedInType: str = None
    disabled: bool = None
    editable: bool = None
    facetable: bool = None
    helpContent: str = None
    hidden: bool = None
    label: str = None
    maxLength: int = None
    maxValue: int = None
    minValue: int = None
    name: str = None
    noCopy: bool = None
    queryable: bool = None
    repeating: bool = None
    required: bool = None
    scope: str = None
    section: str = None
    sectionPosition: int = None
    setOnCreateOnly: bool = None
    shared: bool = None
    systemAttribute: bool = None
    type: str = None
```

#### Document (`component/document.py`)

Standard Vault document fields:

```python
@dataclass
class Document(VaultModel):
    id: int
    version_id: str
    major_version_number__v: int
    minor_version_number__v: int
    name__v: str
    title__v: str
    type__v: str
    subtype__v: str
    classification__v: str
    lifecycle__v: str
    status__v: str
    filename__v: str
    format__v: str
    size__v: int
    pages__v: int
    md5checksum__v: str
    # ... 50+ more fields
```

### API Endpoints Implemented

| Endpoint | Purpose |
|----------|---------|
| `POST /api/{version}/auth` | Authentication |
| `GET /api/{version}/services/directdata/files` | List Direct Data files |
| `GET /api/{version}/services/directdata/files/{name}` | Download extract |
| `POST /api/{version}/query` | Execute VQL |
| `GET /api/{version}/metadata/objects/documents/properties` | Document fields |
| `GET /api/{version}/metadata/objects/documents/types` | Document types |

### Direct Data Extract Structure

When extracting data, the Direct Data API returns:

```
extract.tar.gz
├── manifest.csv           # Index of all files
├── Metadata/
│   └── metadata.csv       # Schema for all tables
├── Object/
│   └── {object_name}.csv  # Object records
└── Document/
    └── {type}.csv         # Document metadata
```

The `metadata.csv` contains:
```csv
extract,column_name,type,length
study__v,id,id,
study__v,name__v,string,128
study__v,created_date__v,datetime,
```

### Field Types from Metadata

| Type | Description |
|------|-------------|
| `id` | Object/record identifier |
| `string` | Text field |
| `datetime` | Timestamp |
| `boolean` | True/false |
| `number` / `numeric` | Numeric value |
| `date` | Date only |

## Limitations

This accelerator is focused on **data extraction**, not metadata discovery:

- No endpoints for `/metadata/vobjects` (object definitions)
- No endpoints for `/objects/picklists` (picklist values)
- No lifecycle/workflow metadata endpoints
- Limited to QueryDescribe for field metadata

## Usefulness

| Use Case | Rating | Notes |
|----------|--------|-------|
| **Field Model Reference** | Good | Python dataclass for VaultObjectField |
| **Document Model** | Good | Complete Document dataclass |
| **Object Definitions** | Poor | No object metadata endpoints |
| **Integration Patterns** | Good | Shows Vault API usage |

## Conclusion

The Direct Data API Accelerators provide **useful Python models** for field metadata (`VaultObjectField`, `DocumentField`), which complement the Java models in VAPIL. However, for complete object specifications, you would need to extend this with metadata API endpoints.

The `VaultObjectField` dataclass with 35+ properties is particularly useful as a reference for field attributes.
