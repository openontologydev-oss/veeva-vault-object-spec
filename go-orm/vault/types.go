// Package vault provides Go ORM models for Veeva Vault objects and components.
// This package is auto-generated from the official Veeva MDL Components documentation.
//
// Source: https://developer.veevavault.com/mdl/components
// Generated: 2026-01-12
// Components: 117 top-level, 88 subcomponents
package vault

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// =============================================================================
// BASE TYPES
// =============================================================================

// BaseModel contains common fields for all Vault components
type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id" example:"1"`
	CreatedAt time.Time      `json:"created_at" example:"2024-01-15T10:30:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2024-01-15T10:30:00Z"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// ComponentBase contains common fields for all MDL components
type ComponentBase struct {
	Name   string `gorm:"size:255;uniqueIndex;not null" json:"name" example:"product__v"`
	Label  string `gorm:"size:60" json:"label" example:"Product"`
	Active bool   `gorm:"default:true" json:"active" example:"true"`
}

// AuditFields contains audit tracking fields
type AuditFields struct {
	CreatedBy  uint       `json:"created_by" example:"1"`
	ModifiedBy uint       `json:"modified_by" example:"1"`
	CreatedDate  time.Time `json:"created_date" example:"2024-01-15T10:30:00Z"`
	ModifiedDate time.Time `json:"modified_date" example:"2024-01-15T10:30:00Z"`
}

// =============================================================================
// CUSTOM TYPES FOR SPECIAL MDL TYPES
// =============================================================================

// LongString represents a long text field (up to 32000 chars)
type LongString string

// Expression represents a VQL or formula expression
type Expression string

// ActionScript represents action trigger script code
type ActionScript string

// SdkCode represents Vault Java SDK source code
type SdkCode string

// XMLString represents XML configuration content
type XMLString string

// StringSlice is a custom type for storing string arrays in a single column
type StringSlice []string

// Value implements driver.Valuer for StringSlice
func (s StringSlice) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}
	return json.Marshal(s)
}

// Scan implements sql.Scanner for StringSlice
func (s *StringSlice) Scan(value interface{}) error {
	if value == nil {
		*s = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, s)
}

// LocalizedData stores translations for labels
type LocalizedData map[string]map[string]string

// Value implements driver.Valuer for LocalizedData
func (l LocalizedData) Value() (driver.Value, error) {
	if l == nil {
		return nil, nil
	}
	return json.Marshal(l)
}

// Scan implements sql.Scanner for LocalizedData
func (l *LocalizedData) Scan(value interface{}) error {
	if value == nil {
		*l = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, l)
}

// =============================================================================
// ENUMS - Field Types
// =============================================================================

// FieldType represents the type of an object field
type FieldType string

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

// AllFieldTypes returns all valid field types
func AllFieldTypes() []FieldType {
	return []FieldType{
		FieldTypeString, FieldTypeNumber, FieldTypeBoolean, FieldTypeDate,
		FieldTypeDateTime, FieldTypeObject, FieldTypePicklist, FieldTypeID,
		FieldTypeUsers, FieldTypeComponent, FieldTypeDocuments, FieldTypeLongText,
		FieldTypeRichText, FieldTypeBinary, FieldTypeObjectReference, FieldTypeObjectParent,
	}
}

// =============================================================================
// ENUMS - Relationship Types
// =============================================================================

// RelationshipType represents the type of object relationship
type RelationshipType string

const (
	RelationshipTypeParent           RelationshipType = "parent"
	RelationshipTypeReference        RelationshipType = "reference"
	RelationshipTypeChild            RelationshipType = "child"
	RelationshipTypeReferenceInbound RelationshipType = "reference_inbound"
	RelationshipTypeReferenceOutbound RelationshipType = "reference_outbound"
)

// RelationshipDeletion represents behavior when related record is deleted
type RelationshipDeletion string

const (
	RelationshipDeletionBlock   RelationshipDeletion = "block"
	RelationshipDeletionCascade RelationshipDeletion = "cascade"
	RelationshipDeletionSetNull RelationshipDeletion = "setnull"
)

// =============================================================================
// ENUMS - Object Class
// =============================================================================

// ObjectClass represents the class of a Vault object
type ObjectClass string

const (
	ObjectClassBase                ObjectClass = "base"
	ObjectClassComponent           ObjectClass = "component"
	ObjectClassUserRoleSetup       ObjectClass = "userrolesetup"
	ObjectClassLegalHold           ObjectClass = "legalhold"
	ObjectClassUserTask            ObjectClass = "usertask"
	ObjectClassESignature          ObjectClass = "esignature"
	ObjectClassChecklist           ObjectClass = "checklist"
	ObjectClassSection             ObjectClass = "section"
	ObjectClassResponse            ObjectClass = "response"
	ObjectClassLifecycleStagesObject ObjectClass = "lifecyclestagesobject"
)

// =============================================================================
// ENUMS - Data Store
// =============================================================================

// DataStore represents the storage type for an object
type DataStore string

const (
	DataStoreStandard DataStore = "standard"
	DataStoreRaw      DataStore = "raw"
)

// =============================================================================
// ENUMS - Record Status
// =============================================================================

// RecordStatus represents the status of an object record
type RecordStatus string

const (
	RecordStatusActive      RecordStatus = "active__v"
	RecordStatusInactive    RecordStatus = "inactive__v"
	RecordStatusInMigration RecordStatus = "in_migration__v"
	RecordStatusArchived    RecordStatus = "archived__v"
)

// =============================================================================
// ENUMS - Trigger Events
// =============================================================================

// TriggerEvent represents when a trigger fires
type TriggerEvent string

const (
	TriggerEventBeforeInsert TriggerEvent = "BEFORE_INSERT"
	TriggerEventBeforeUpdate TriggerEvent = "BEFORE_UPDATE"
	TriggerEventBeforeDelete TriggerEvent = "BEFORE_DELETE"
	TriggerEventAfterInsert  TriggerEvent = "AFTER_INSERT"
	TriggerEventAfterUpdate  TriggerEvent = "AFTER_UPDATE"
	TriggerEventAfterDelete  TriggerEvent = "AFTER_DELETE"
)

// AllTriggerEvents returns all valid trigger events
func AllTriggerEvents() []TriggerEvent {
	return []TriggerEvent{
		TriggerEventBeforeInsert, TriggerEventBeforeUpdate, TriggerEventBeforeDelete,
		TriggerEventAfterInsert, TriggerEventAfterUpdate, TriggerEventAfterDelete,
	}
}

// =============================================================================
// ENUMS - Lifecycle Permissions
// =============================================================================

// LifecyclePermission represents a permission level
type LifecyclePermission string

const (
	LifecyclePermissionRead   LifecyclePermission = "Read"
	LifecyclePermissionEdit   LifecyclePermission = "Edit"
	LifecyclePermissionDelete LifecyclePermission = "Delete"
)

// =============================================================================
// ENUMS - Security Types
// =============================================================================

// SecurityType represents field/action security level
type SecurityType string

const (
	SecurityTypeHide    SecurityType = "hide__v"
	SecurityTypeRead    SecurityType = "read__v"
	SecurityTypeEdit    SecurityType = "edit__v"
	SecurityTypeView    SecurityType = "view__v"
	SecurityTypeExecute SecurityType = "execute__v"
)

// =============================================================================
// ENUMS - App Security Field Types
// =============================================================================

// AppSecurityFieldType represents the type of app security field
type AppSecurityFieldType string

const (
	AppSecurityFieldTypeObject   AppSecurityFieldType = "OBJECT"
	AppSecurityFieldTypePicklist AppSecurityFieldType = "PICKLIST"
	AppSecurityFieldTypeBoolean  AppSecurityFieldType = "BOOLEAN"
)

// =============================================================================
// SOURCE TYPE
// =============================================================================

// Source represents whether a component is standard or custom
type Source string

const (
	SourceStandard Source = "standard"
	SourceCustom   Source = "custom"
)
