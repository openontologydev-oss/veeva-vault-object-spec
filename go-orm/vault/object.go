package vault

// =============================================================================
// VAULT OBJECT - Core object definition
// =============================================================================

// VaultObject represents a Vault Object component (MDL: Object)
// Objects define the data model for the Vault.
//
// Example MDL:
//   CREATE Object product__c (
//     label('Product'),
//     label_plural('Products'),
//     active(true),
//     in_menu(true),
//     audit(true)
//   );
type VaultObject struct {
	BaseModel
	ComponentBase
	AuditFields

	// Core Properties
	LabelPlural   string      `gorm:"size:40;not null" json:"label_plural" example:"Products"`
	Description   *string     `gorm:"size:255" json:"description,omitempty" example:"Custom product object"`
	HelpContent   *string     `gorm:"size:255" json:"help_content,omitempty"`
	Prefix        string      `gorm:"size:10" json:"prefix" example:"00P"`
	Source        Source      `gorm:"size:20;default:'custom'" json:"source" example:"custom"`
	ObjectClass   ObjectClass `gorm:"size:30;default:'base'" json:"object_class" example:"base"`
	DataStore     DataStore   `gorm:"size:20;default:'standard'" json:"data_store" example:"standard"`
	Order         int         `gorm:"default:0" json:"order" example:"12"`

	// Menu & UI
	InMenu bool `gorm:"not null;default:true" json:"in_menu" example:"true"`

	// Features
	AllowAttachments     bool `gorm:"default:false" json:"allow_attachments" example:"false"`
	AllowTypes           bool `gorm:"default:false" json:"allow_types" example:"false"`
	EnableESignatures    bool `gorm:"default:false" json:"enable_esignatures" example:"false"`
	EnableMerges         bool `gorm:"default:false" json:"enable_merges" example:"false"`
	DynamicSecurity      bool `gorm:"default:false" json:"dynamic_security" example:"false"`
	SystemManaged        bool `gorm:"default:false" json:"system_managed" example:"false"`
	MultiSelect          bool `gorm:"default:false" json:"multi_select" example:"false"`
	RelateRecordsSelectAll bool `gorm:"default:false" json:"relate_records_select_all" example:"false"`

	// Audit & Security
	Audit                 bool `gorm:"not null;default:true" json:"audit" example:"true"`
	SecureAuditTrail      bool `gorm:"default:false" json:"secure_audit_trail" example:"false"`
	SecureSharingSettings bool `gorm:"default:false" json:"secure_sharing_settings" example:"false"`
	SecureAttachments     bool `gorm:"default:false" json:"secure_attachments" example:"false"`
	SecureCopyRecord      bool `gorm:"default:false" json:"secure_copy_record" example:"false"`
	PreventRecordOverwrite bool `gorm:"default:false" json:"prevent_record_overwrite" example:"false"`
	RoleOverrides         bool `gorm:"default:false" json:"role_overrides" example:"false"`
	Facetable             bool `gorm:"default:false" json:"facetable" example:"false"`

	// Relationships to other objects
	UserRoleSetupObjectID *uint        `json:"user_role_setup_object_id,omitempty"`
	UserRoleSetupObject   *VaultObject `gorm:"foreignKey:UserRoleSetupObjectID" json:"user_role_setup_object,omitempty"`

	SecurityTreeObjectID *uint        `json:"security_tree_object_id,omitempty"`
	SecurityTreeObject   *VaultObject `gorm:"foreignKey:SecurityTreeObjectID" json:"security_tree_object,omitempty"`

	// Status
	Status       StringSlice `gorm:"type:json" json:"status" example:"[\"active__v\"]"`
	DefaultObjType string    `gorm:"size:100" json:"default_obj_type" example:"base__v"`

	// Localization
	LocalizedData LocalizedData `gorm:"type:json" json:"localized_data,omitempty"`

	// Summary fields (multi-value)
	SummaryFields StringSlice `gorm:"type:json" json:"summary_fields,omitempty"`

	// Child relationships
	Fields              []ObjectField         `gorm:"foreignKey:ObjectID" json:"fields,omitempty"`
	Indexes             []ObjectIndex         `gorm:"foreignKey:ObjectID" json:"indexes,omitempty"`
	Types               []ObjectType          `gorm:"foreignKey:ObjectID" json:"types,omitempty"`
	AvailableLifecycles []ObjectLifecycle     `gorm:"many2many:object_available_lifecycles" json:"available_lifecycles,omitempty"`
	Relationships       []ObjectRelationship  `gorm:"foreignKey:ObjectID" json:"relationships,omitempty"`
}

// TableName returns the table name for VaultObject
func (VaultObject) TableName() string {
	return "vault_objects"
}

// =============================================================================
// OBJECT FIELD - Field definition within an object
// =============================================================================

// ObjectField represents a field on a Vault Object (MDL: Object.Field)
//
// Example MDL:
//   Field name__v (
//     label('Name'),
//     type('String'),
//     required(true),
//     max_length(128)
//   );
type ObjectField struct {
	BaseModel
	AuditFields

	// Parent relationship
	ObjectID uint        `gorm:"not null;index" json:"object_id"`
	Object   *VaultObject `gorm:"foreignKey:ObjectID" json:"-"`

	// Core Properties
	Name        string    `gorm:"size:100;not null;index" json:"name" example:"name__v"`
	Label       string    `gorm:"size:40;not null" json:"label" example:"Name"`
	Type        FieldType `gorm:"size:30;not null" json:"type" example:"String"`
	Active      bool      `gorm:"not null;default:true" json:"active" example:"true"`
	Required    bool      `gorm:"not null;default:false" json:"required" example:"false"`
	Editable    bool      `gorm:"default:true" json:"editable" example:"true"`
	Source      Source    `gorm:"size:20;default:'custom'" json:"source" example:"custom"`
	Description *string   `gorm:"size:1000" json:"description,omitempty"`
	HelpContent *string   `gorm:"size:255" json:"help_content,omitempty"`

	// Display
	ListColumn bool `gorm:"not null;default:false" json:"list_column" example:"false"`
	Order      int  `gorm:"default:0" json:"order" example:"1"`
	NoCopy     bool `gorm:"default:false" json:"no_copy" example:"false"`

	// Type-specific constraints
	MaxLength  *int    `json:"max_length,omitempty" example:"128"`
	MaxValue   *int    `json:"max_value,omitempty"`
	MinValue   *int    `json:"min_value,omitempty"`
	Scale      *int    `json:"scale,omitempty" example:"2"`
	FormatMask *string `gorm:"size:3900" json:"format_mask,omitempty"`

	// Unique & System
	Unique            bool `gorm:"default:false" json:"unique" example:"false"`
	SystemManagedName bool `gorm:"default:false" json:"system_managed_name" example:"false"`

	// Auto-numbering
	StartNumber *int    `json:"start_number,omitempty"`
	ValueFormat *string `gorm:"size:255" json:"value_format,omitempty"`

	// Multi-value
	MultiValue bool `gorm:"default:false" json:"multi_value" example:"false"`

	// Search & Filter
	Queryable bool `gorm:"default:true" json:"queryable" example:"true"`
	Facetable bool `gorm:"default:false" json:"facetable" example:"false"`

	// Object Reference (when type is Object, ObjectReference, ObjectParent)
	ReferencedObject *string `gorm:"size:100" json:"object,omitempty" example:"product_family__v"`

	// Relationship properties
	RelationshipType          *RelationshipType    `gorm:"size:30" json:"relationship_type,omitempty" example:"reference"`
	RelationshipDeletion      *RelationshipDeletion `gorm:"size:20" json:"relationship_deletion,omitempty" example:"block"`
	RelationshipOutboundName  *string              `gorm:"size:50" json:"relationship_outbound_name,omitempty"`
	RelationshipInboundName   *string              `gorm:"size:50" json:"relationship_inbound_name,omitempty"`
	RelationshipInboundLabel  *string              `gorm:"size:40" json:"relationship_inbound_label,omitempty"`
	RelationshipCriteria      *Expression          `gorm:"size:4000" json:"relationship_criteria,omitempty"`

	// Picklist reference (when type is Picklist)
	PicklistName      *string `gorm:"size:100" json:"picklist,omitempty" example:"status__v"`
	ControllingField  *string `gorm:"size:50" json:"controlling_field,omitempty"`

	// Component reference (when type is Component)
	ComponentType *string `gorm:"size:100" json:"component,omitempty"`

	// Lookup (for lookup fields)
	LookupObjectID           *string `gorm:"size:100" json:"lookup_object_id,omitempty"`
	LookupFieldID            *string `gorm:"size:100" json:"lookup_field_id,omitempty"`
	LookupSourceField        *string `gorm:"size:100" json:"lookup_source_field,omitempty"`
	LookupRelationshipName   *string `gorm:"size:100" json:"lookup_relationship_name,omitempty"`

	// Localization
	LocalizedData LocalizedData `gorm:"type:json" json:"localized_data,omitempty"`

	// Status
	Status StringSlice `gorm:"type:json" json:"status" example:"[\"active__v\"]"`
}

// TableName returns the table name for ObjectField
func (ObjectField) TableName() string {
	return "vault_object_fields"
}

// =============================================================================
// OBJECT INDEX - Index definition within an object
// =============================================================================

// ObjectIndex represents an index on a Vault Object (MDL: Object.Index)
//
// Example MDL:
//   Index idx_product_name__c (
//     unique(true),
//     fields('name__v')
//   );
type ObjectIndex struct {
	BaseModel

	// Parent relationship
	ObjectID uint        `gorm:"not null;index" json:"object_id"`
	Object   *VaultObject `gorm:"foreignKey:ObjectID" json:"-"`

	// Properties
	Name   string `gorm:"size:100;not null" json:"name" example:"idx_product_name__c"`
	Unique bool   `gorm:"default:false" json:"unique" example:"true"`

	// Fields in the index (stored as JSON array)
	Fields StringSlice `gorm:"type:json;not null" json:"fields" example:"[\"name__v\"]"`
}

// TableName returns the table name for ObjectIndex
func (ObjectIndex) TableName() string {
	return "vault_object_indexes"
}

// =============================================================================
// OBJECT RELATIONSHIP - Relationship metadata
// =============================================================================

// ObjectRelationship represents a relationship between objects
// This is derived from Object Field metadata for reference/parent fields
type ObjectRelationship struct {
	BaseModel

	// Parent object
	ObjectID uint        `gorm:"not null;index" json:"object_id"`
	Object   *VaultObject `gorm:"foreignKey:ObjectID" json:"-"`

	// Relationship details
	RelationshipName  string           `gorm:"size:100;not null" json:"relationship_name" example:"product_family__vr"`
	RelationshipLabel string           `gorm:"size:60" json:"relationship_label" example:"Product Family"`
	RelationshipType  RelationshipType `gorm:"size:30;not null" json:"relationship_type" example:"reference_outbound"`
	FieldName         string           `gorm:"size:100;not null" json:"field" example:"product_family__v"`

	// Related object
	TargetObjectName string `gorm:"size:100;not null" json:"target_object_name" example:"product_family__v"`

	// Deletion behavior
	RelationshipDeletion *RelationshipDeletion `gorm:"size:20" json:"relationship_deletion,omitempty" example:"block"`

	// Localization
	LocalizedData LocalizedData `gorm:"type:json" json:"localized_data,omitempty"`
}

// TableName returns the table name for ObjectRelationship
func (ObjectRelationship) TableName() string {
	return "vault_object_relationships"
}

// =============================================================================
// URL METADATA - API URLs for objects
// =============================================================================

// ObjectURLs contains API endpoint URLs for an object
type ObjectURLs struct {
	Field    string `json:"field" example:"/api/v21.2/metadata/vobjects/product__v/fields/{name}"`
	Record   string `json:"record" example:"/api/v21.2/vobjects/product__v/{id}"`
	List     string `json:"list" example:"/api/v21.2/vobjects/product__v"`
	Metadata string `json:"metadata" example:"/api/v21.2/metadata/vobjects/product__v"`
}
