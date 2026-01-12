package vault

// =============================================================================
// PICKLIST
// =============================================================================

// Picklist represents a picklist definition (MDL: Picklist)
//
// Example MDL:
//   CREATE Picklist status__c (
//     label('Status'),
//     active(true)
//   );
type Picklist struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:255" json:"description,omitempty"`
	HelpContent *string `gorm:"size:255" json:"help_content,omitempty"`

	// Child relationships
	Entries []PicklistEntry `gorm:"foreignKey:PicklistID" json:"entries,omitempty"`
}

// TableName returns the table name for Picklist
func (Picklist) TableName() string {
	return "vault_picklists"
}

// PicklistEntry represents a picklist value (MDL: Picklistentry)
//
// Example MDL:
//   Picklistentry draft__c (
//     label('Draft'),
//     active(true),
//     order(1)
//   );
type PicklistEntry struct {
	BaseModel
	ComponentBase

	// Parent picklist
	PicklistID uint      `gorm:"not null;index" json:"picklist_id"`
	Picklist   *Picklist `gorm:"foreignKey:PicklistID" json:"-"`

	// Properties
	Order int `gorm:"default:0" json:"order" example:"1"`

	// Parent entry (for hierarchical picklists)
	ParentID *uint          `json:"parent_id,omitempty"`
	Parent   *PicklistEntry `gorm:"foreignKey:ParentID" json:"parent,omitempty"`

	// Children
	Children []PicklistEntry `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}

// TableName returns the table name for PicklistEntry
func (PicklistEntry) TableName() string {
	return "vault_picklist_entries"
}

// =============================================================================
// SECURITY PROFILE
// =============================================================================

// Securityprofile represents a security profile (MDL: Securityprofile)
type Securityprofile struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:255" json:"description,omitempty"`
	HelpContent *string `gorm:"size:255" json:"help_content,omitempty"`
}

// TableName returns the table name for Securityprofile
func (Securityprofile) TableName() string {
	return "vault_securityprofiles"
}

// =============================================================================
// PERMISSION SET
// =============================================================================

// Permissionset represents a permission set (MDL: Permissionset)
type Permissionset struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:255" json:"description,omitempty"`
	HelpContent *string `gorm:"size:255" json:"help_content,omitempty"`

	// Permissions
	Permissions StringSlice `gorm:"type:json" json:"permissions,omitempty"`
}

// TableName returns the table name for Permissionset
func (Permissionset) TableName() string {
	return "vault_permissionsets"
}

// =============================================================================
// SHARING RULE
// =============================================================================

// Sharingrule represents a sharing rule (MDL: Sharingrule)
type Sharingrule struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:255" json:"description,omitempty"`

	// Target object
	ObjectName string `gorm:"size:100" json:"object,omitempty"`

	// Child relationships
	Roles []Sharingrole `gorm:"foreignKey:RuleID" json:"roles,omitempty"`
}

// TableName returns the table name for Sharingrule
func (Sharingrule) TableName() string {
	return "vault_sharingrules"
}

// Sharingrole represents a role in a sharing rule (MDL: Sharingrole)
type Sharingrole struct {
	BaseModel

	// Parent rule
	RuleID uint         `gorm:"not null;index" json:"rule_id"`
	Rule   *Sharingrule `gorm:"foreignKey:RuleID" json:"-"`

	// Properties
	Name        string      `gorm:"size:100;not null" json:"name"`
	Permissions StringSlice `gorm:"type:json" json:"permissions,omitempty"`
}

// TableName returns the table name for Sharingrole
func (Sharingrole) TableName() string {
	return "vault_sharingroles"
}

// =============================================================================
// ATOMIC SECURITY
// =============================================================================

// Atomicsecurity represents atomic security configuration (MDL: Atomicsecurity)
type Atomicsecurity struct {
	BaseModel
	ComponentBase

	// Object reference
	ObjectName string `gorm:"size:100;not null" json:"object" example:"product__c"`

	// Lifecycle reference
	LifecycleName string `gorm:"size:100;not null" json:"object_lifecycle" example:"product_lifecycle__c"`

	// State reference
	StateName string `gorm:"size:100;not null" json:"state" example:"active_state__c"`

	// Child relationships
	FieldSecurities        []FieldSecurity        `gorm:"foreignKey:AtomicSecurityID" json:"field_securities,omitempty"`
	ObjectControlSecurities []ObjectControlSecurity `gorm:"foreignKey:AtomicSecurityID" json:"object_control_securities,omitempty"`
	ActionSecurities       []ActionSecurity       `gorm:"foreignKey:AtomicSecurityID" json:"action_securities,omitempty"`
	WorkflowActionSecurities []WorkflowActionSecurity `gorm:"foreignKey:AtomicSecurityID" json:"workflow_action_securities,omitempty"`
	RelationshipSecurities []RelationshipSecurity `gorm:"foreignKey:AtomicSecurityID" json:"relationship_securities,omitempty"`
}

// TableName returns the table name for Atomicsecurity
func (Atomicsecurity) TableName() string {
	return "vault_atomicsecurities"
}

// FieldSecurity represents field-level security (MDL: Fieldsecurity)
type FieldSecurity struct {
	BaseModel
	ComponentBase

	// Parent
	AtomicSecurityID uint            `gorm:"not null;index" json:"atomic_security_id"`
	AtomicSecurity   *Atomicsecurity `gorm:"foreignKey:AtomicSecurityID" json:"-"`

	// Properties
	Role   string       `gorm:"size:100;not null" json:"role" example:"editor__v"`
	Type   SecurityType `gorm:"size:20;not null" json:"type" example:"edit__v"`
	Fields StringSlice  `gorm:"type:json;not null" json:"fields"`
}

// TableName returns the table name for FieldSecurity
func (FieldSecurity) TableName() string {
	return "vault_field_securities"
}

// ObjectControlSecurity represents object control security (MDL: Objectcontrolsecurity)
type ObjectControlSecurity struct {
	BaseModel
	ComponentBase

	// Parent
	AtomicSecurityID uint            `gorm:"not null;index" json:"atomic_security_id"`
	AtomicSecurity   *Atomicsecurity `gorm:"foreignKey:AtomicSecurityID" json:"-"`

	// Properties
	Role           string       `gorm:"size:100;not null" json:"role" example:"editor__v"`
	Type           SecurityType `gorm:"size:20;not null" json:"type" example:"view__v"`
	ObjectControls StringSlice  `gorm:"type:json;not null" json:"object_controls"`
}

// TableName returns the table name for ObjectControlSecurity
func (ObjectControlSecurity) TableName() string {
	return "vault_object_control_securities"
}

// ActionSecurity represents action-level security (MDL: Actionsecurity)
type ActionSecurity struct {
	BaseModel
	ComponentBase

	// Parent
	AtomicSecurityID uint            `gorm:"not null;index" json:"atomic_security_id"`
	AtomicSecurity   *Atomicsecurity `gorm:"foreignKey:AtomicSecurityID" json:"-"`

	// Properties
	Role             string       `gorm:"size:100;not null" json:"role" example:"editor__v"`
	Type             SecurityType `gorm:"size:20;not null" json:"type" example:"execute__v"`
	ObjectActions    StringSlice  `gorm:"type:json" json:"object_actions,omitempty"`
	LifecycleActions StringSlice  `gorm:"type:json" json:"lifecycle_actions,omitempty"`
}

// TableName returns the table name for ActionSecurity
func (ActionSecurity) TableName() string {
	return "vault_action_securities"
}

// WorkflowActionSecurity represents workflow action security (MDL: Workflowactionsecurity)
type WorkflowActionSecurity struct {
	BaseModel
	ComponentBase

	// Parent
	AtomicSecurityID uint            `gorm:"not null;index" json:"atomic_security_id"`
	AtomicSecurity   *Atomicsecurity `gorm:"foreignKey:AtomicSecurityID" json:"-"`

	// Properties
	Role                string       `gorm:"size:100;not null" json:"role" example:"editor__v"`
	Type                SecurityType `gorm:"size:20;not null" json:"type" example:"execute__v"`
	WorkflowActions     StringSlice  `gorm:"type:json" json:"workflow_actions,omitempty"`
	WorkflowTaskActions StringSlice  `gorm:"type:json" json:"workflow_task_actions,omitempty"`
}

// TableName returns the table name for WorkflowActionSecurity
func (WorkflowActionSecurity) TableName() string {
	return "vault_workflow_action_securities"
}

// RelationshipSecurity represents relationship security (MDL: Relationshipsecurity)
type RelationshipSecurity struct {
	BaseModel
	ComponentBase

	// Parent
	AtomicSecurityID uint            `gorm:"not null;index" json:"atomic_security_id"`
	AtomicSecurity   *Atomicsecurity `gorm:"foreignKey:AtomicSecurityID" json:"-"`

	// Properties
	Role           string       `gorm:"size:100;not null" json:"role" example:"editor__v"`
	Type           SecurityType `gorm:"size:20;not null" json:"type" example:"edit__v"`
	DocumentFields StringSlice  `gorm:"type:json" json:"document_fields,omitempty"`
	ObjectFields   StringSlice  `gorm:"type:json" json:"object_fields,omitempty"`
}

// TableName returns the table name for RelationshipSecurity
func (RelationshipSecurity) TableName() string {
	return "vault_relationship_securities"
}

// =============================================================================
// APP SECURITY RULE
// =============================================================================

// Appsecurityrule represents an application security rule (MDL: Appsecurityrule)
type Appsecurityrule struct {
	BaseModel
	ComponentBase

	// Properties
	Description        *string `gorm:"size:255" json:"description,omitempty"`
	HelpContent        *string `gorm:"size:255" json:"help_content,omitempty"`
	AllowCustomObjects bool    `gorm:"not null" json:"allow_custom_objects" example:"true"`

	// Child relationships
	Fields      []Appsecurityfield      `gorm:"foreignKey:RuleID" json:"fields,omitempty"`
	Assignments []Appsecurityassignment `gorm:"foreignKey:RuleID" json:"assignments,omitempty"`
}

// TableName returns the table name for Appsecurityrule
func (Appsecurityrule) TableName() string {
	return "vault_appsecurityrules"
}

// Appsecurityfield represents a field in an app security rule (MDL: Appsecurityfield)
type Appsecurityfield struct {
	BaseModel
	ComponentBase

	// Parent rule
	RuleID uint             `gorm:"not null;index" json:"rule_id"`
	Rule   *Appsecurityrule `gorm:"foreignKey:RuleID" json:"-"`

	// Properties
	Type       AppSecurityFieldType `gorm:"size:20;not null" json:"type" example:"OBJECT"`
	MultiValue bool                 `gorm:"default:false" json:"multi_value" example:"false"`

	// References
	PicklistName *string `gorm:"size:100" json:"picklist,omitempty"`
	ObjectName   *string `gorm:"size:100" json:"object,omitempty"`
}

// TableName returns the table name for Appsecurityfield
func (Appsecurityfield) TableName() string {
	return "vault_appsecurityfields"
}

// Appsecurityassignment represents an assignment in an app security rule (MDL: Appsecurityassignment)
type Appsecurityassignment struct {
	BaseModel
	ComponentBase

	// Parent rule
	RuleID uint             `gorm:"not null;index" json:"rule_id"`
	Rule   *Appsecurityrule `gorm:"foreignKey:RuleID" json:"-"`

	// Properties
	Field         string `gorm:"size:100;not null" json:"field" example:"study_migration__v"`
	ObjectName    string `gorm:"size:100;not null" json:"object" example:"procedure__v"`
	SecurityField *string `gorm:"size:100" json:"securityfield,omitempty"`
}

// TableName returns the table name for Appsecurityassignment
func (Appsecurityassignment) TableName() string {
	return "vault_appsecurityassignments"
}
