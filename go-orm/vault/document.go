package vault

// =============================================================================
// DOCUMENT TYPE
// =============================================================================

// Doctype represents a document type (MDL: Doctype)
//
// Example MDL:
//   CREATE Doctype promotional_piece__c (
//     label('Promotional Piece'),
//     active(true)
//   );
type Doctype struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:255" json:"description,omitempty"`
	HelpContent *string `gorm:"size:500" json:"help_content,omitempty"`

	// Hierarchy (Doctype can have subtypes/classifications)
	ParentID *uint    `json:"parent_id,omitempty"`
	Parent   *Doctype `gorm:"foreignKey:ParentID" json:"parent,omitempty"`

	// Binder properties
	BinderTemplateFile *string `gorm:"size:1500" json:"binder_template_file,omitempty"`

	// Rendition properties
	Renditions StringSlice `gorm:"type:json" json:"renditions,omitempty"`

	// Child relationships
	Children     []Doctype    `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	SharedFields []Docfield   `gorm:"many2many:doctype_shared_fields" json:"shared_fields,omitempty"`
	LocalFields  []Docfield   `gorm:"foreignKey:DoctypeID" json:"local_fields,omitempty"`
	Lifecycles   []Doclifecycle `gorm:"many2many:doctype_lifecycles" json:"lifecycles,omitempty"`
}

// TableName returns the table name for Doctype
func (Doctype) TableName() string {
	return "vault_doctypes"
}

// =============================================================================
// DOCUMENT FIELD
// =============================================================================

// DocfieldDisplaySection represents the display section for a document field
type DocfieldDisplaySection string

const (
	DocfieldDisplaySectionGeneralProperties DocfieldDisplaySection = "general_properties__v"
	DocfieldDisplaySectionVersionProperties DocfieldDisplaySection = "version_properties__v"
	DocfieldDisplaySectionCustomProperties  DocfieldDisplaySection = "custom_properties__v"
)

// Docfield represents a document field (MDL: Docfield)
//
// Example MDL:
//   CREATE Docfield product__c (
//     label('Product'),
//     type('Object'),
//     object('product__v'),
//     shared(true),
//     display_section(Docfieldlayout.custom_properties__v)
//   );
type Docfield struct {
	BaseModel
	ComponentBase

	// Doctype (for local fields)
	DoctypeID *uint    `json:"doctype_id,omitempty"`
	Doctype   *Doctype `gorm:"foreignKey:DoctypeID" json:"-"`

	// Core Properties
	Type           FieldType              `gorm:"size:30;not null" json:"type" example:"String"`
	Required       bool                   `gorm:"default:false" json:"required" example:"false"`
	Shared         bool                   `gorm:"not null" json:"shared" example:"true"`
	DisplaySection DocfieldDisplaySection `gorm:"size:100;not null" json:"display_section" example:"general_properties__v"`
	Description    *string                `gorm:"size:1000" json:"description,omitempty"`
	HelpContent    *string                `gorm:"size:500" json:"help_content,omitempty"`

	// Constraints
	MaxLength *int `json:"max_length,omitempty" example:"255"`
	Scale     *int `json:"scale,omitempty"`

	// Multi-value
	MultiValue bool `gorm:"default:false" json:"multi_value" example:"false"`

	// Queryable
	Queryable bool `gorm:"default:true" json:"queryable" example:"true"`

	// Relationship (when type is Object)
	RelationshipName    *string               `gorm:"size:100" json:"relationship_name,omitempty"`
	SecureRelationship  bool                  `gorm:"default:false" json:"secure_relationship" example:"false"`
	RelationshipCriteria *string              `gorm:"size:4000" json:"relationship_criteria,omitempty"`

	// Picklist reference
	PicklistName     *string `gorm:"size:100" json:"picklist,omitempty"`
	ControllingField *string `gorm:"size:100" json:"controlling_field,omitempty"`

	// Lookup
	LookupObjectFieldKey *string `gorm:"size:100" json:"lookupObjectFieldKey,omitempty"`
	LookupObjectID       *string `gorm:"size:100" json:"lookupObjectId,omitempty"`
	LookupFieldID        *string `gorm:"size:100" json:"lookupFieldId,omitempty"`

	// Formula
	Formula     *Expression `gorm:"size:4000" json:"formula,omitempty"`
	BlankFields *string     `gorm:"size:20" json:"blank_fields,omitempty"` // "zeros" or "blanks"

	// Security
	DefaultSecurity           *string     `gorm:"size:20" json:"default_security,omitempty"` // editable, readonly, hidden
	Editable                  bool        `gorm:"default:true" json:"editable" example:"true"`
	Hidden                    bool        `gorm:"default:false" json:"hidden" example:"false"`
	SecurityOverrideEditable  StringSlice `gorm:"type:json" json:"security_override_editable,omitempty"`
	SecurityOverrideReadonly  StringSlice `gorm:"type:json" json:"security_override_readonly,omitempty"`
	SecurityOverrideHidden    StringSlice `gorm:"type:json" json:"security_override_hidden,omitempty"`

	// Copy behavior
	NoCopy         bool `gorm:"default:false" json:"no_copy" example:"false"`
	SetOnCreateOnly bool `gorm:"default:false" json:"set_on_create_only" example:"false"`

	// Permissions (can change)
	CanChangeSection   bool `gorm:"default:false" json:"can_change_section" example:"false"`
	CanChangeRequired  bool `gorm:"default:false" json:"can_change_required" example:"false"`
	CanChangeDisabled  bool `gorm:"default:false" json:"can_change_disabled" example:"false"`
	CanChangeMultiple  bool `gorm:"default:false" json:"can_change_multiple" example:"false"`
	CanChangeNoCopy    bool `gorm:"default:false" json:"can_change_no_copy" example:"false"`

	// Used in document types (for shared fields)
	UsedInDoctypes []Doctype `gorm:"many2many:doctype_shared_fields" json:"used_in,omitempty"`
}

// TableName returns the table name for Docfield
func (Docfield) TableName() string {
	return "vault_docfields"
}

// =============================================================================
// DOCUMENT FIELD LAYOUT
// =============================================================================

// Docfieldlayout represents a document field layout section (MDL: Docfieldlayout)
type Docfieldlayout struct {
	BaseModel
	ComponentBase

	// Properties
	Order int `gorm:"default:0" json:"order" example:"1"`
}

// TableName returns the table name for Docfieldlayout
func (Docfieldlayout) TableName() string {
	return "vault_docfieldlayouts"
}

// =============================================================================
// DOCUMENT LIFECYCLE
// =============================================================================

// Doclifecycle represents a document lifecycle (MDL: Doclifecycle)
type Doclifecycle struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:255" json:"description,omitempty"`

	// Child relationships
	States []DoclifecycleState `gorm:"foreignKey:LifecycleID" json:"states,omitempty"`
	Roles  []DoclifecycleRole  `gorm:"foreignKey:LifecycleID" json:"roles,omitempty"`
	Events []DoclifecycleEvent `gorm:"foreignKey:LifecycleID" json:"events,omitempty"`

	// Document types using this lifecycle
	Doctypes []Doctype `gorm:"many2many:doctype_lifecycles" json:"doctypes,omitempty"`
}

// TableName returns the table name for Doclifecycle
func (Doclifecycle) TableName() string {
	return "vault_doclifecycles"
}

// DoclifecycleState represents a state in a document lifecycle (MDL: Doclifecyclestate)
type DoclifecycleState struct {
	BaseModel
	ComponentBase

	// Parent lifecycle
	LifecycleID uint          `gorm:"not null;index" json:"lifecycle_id"`
	Lifecycle   *Doclifecycle `gorm:"foreignKey:LifecycleID" json:"-"`

	// Properties
	Description *string `gorm:"size:255" json:"description,omitempty"`
}

// TableName returns the table name for DoclifecycleState
func (DoclifecycleState) TableName() string {
	return "vault_doclifecycle_states"
}

// DoclifecycleRole represents a role in a document lifecycle (MDL: Doclifecyclerole)
type DoclifecycleRole struct {
	BaseModel

	// Parent lifecycle
	LifecycleID uint          `gorm:"not null;index" json:"lifecycle_id"`
	Lifecycle   *Doclifecycle `gorm:"foreignKey:LifecycleID" json:"-"`

	// Properties
	Name            string      `gorm:"size:100;not null" json:"name" example:"owner__v"`
	Active          bool        `gorm:"not null;default:true" json:"active" example:"true"`
	ApplicationRole string      `gorm:"size:60;not null" json:"application_role" example:"owner__v"`
	Permissions     StringSlice `gorm:"type:json" json:"permissions"`
}

// TableName returns the table name for DoclifecycleRole
func (DoclifecycleRole) TableName() string {
	return "vault_doclifecycle_roles"
}

// DoclifecycleEvent represents an event in a document lifecycle (MDL: Doclifecycleevent)
type DoclifecycleEvent struct {
	BaseModel

	// Parent lifecycle
	LifecycleID uint          `gorm:"not null;index" json:"lifecycle_id"`
	Lifecycle   *Doclifecycle `gorm:"foreignKey:LifecycleID" json:"-"`

	// Properties
	Event string    `gorm:"size:1500;not null" json:"event"`
	Rule  XMLString `gorm:"type:text" json:"rule"`
	Order int       `gorm:"not null" json:"order" example:"1"`
}

// TableName returns the table name for DoclifecycleEvent
func (DoclifecycleEvent) TableName() string {
	return "vault_doclifecycle_events"
}

// =============================================================================
// DOCUMENT RELATIONSHIP TYPE
// =============================================================================

// Docrelationshiptype represents a document relationship type (MDL: Docrelationshiptype)
type Docrelationshiptype struct {
	BaseModel
	ComponentBase

	// Properties
	SourceDocType *string `gorm:"size:100" json:"source_doc_type,omitempty"`
	TargetDocType *string `gorm:"size:100" json:"target_doc_type,omitempty"`
	SingleUse     bool    `gorm:"default:false" json:"single_use" example:"false"`
}

// TableName returns the table name for Docrelationshiptype
func (Docrelationshiptype) TableName() string {
	return "vault_docrelationshiptypes"
}

// =============================================================================
// DOCUMENT ACTION
// =============================================================================

// Documentaction represents a document action (MDL: Documentaction)
type Documentaction struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string   `gorm:"size:500" json:"description,omitempty"`
	Rule        XMLString `gorm:"type:text" json:"rule,omitempty"`
}

// TableName returns the table name for Documentaction
func (Documentaction) TableName() string {
	return "vault_documentactions"
}
