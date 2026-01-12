package vault

// =============================================================================
// OBJECT TYPE - Object type variations
// =============================================================================

// ObjectType represents an object type (MDL: Objecttype)
// Object types allow variations of an object with different field configurations.
//
// Example MDL:
//   CREATE Objecttype product__c.software__c (
//     label('Software Product'),
//     active(true)
//   );
type ObjectType struct {
	BaseModel
	ComponentBase

	// Parent object
	ObjectID uint         `gorm:"not null;index" json:"object_id"`
	Object   *VaultObject `gorm:"foreignKey:ObjectID" json:"-"`

	// Properties
	Description *string `gorm:"size:255" json:"description,omitempty"`
	HelpContent *string `gorm:"size:255" json:"help_content,omitempty"`

	// Child relationships
	TypeFields  []TypeField  `gorm:"foreignKey:ObjectTypeID" json:"type_fields,omitempty"`
	TypeActions []TypeAction `gorm:"foreignKey:ObjectTypeID" json:"type_actions,omitempty"`
}

// TableName returns the table name for ObjectType
func (ObjectType) TableName() string {
	return "vault_object_types"
}

// TypeField represents a field override for an object type (MDL: Objecttype.Typefield)
type TypeField struct {
	BaseModel

	// Parent
	ObjectTypeID uint        `gorm:"not null;index" json:"object_type_id"`
	ObjectType   *ObjectType `gorm:"foreignKey:ObjectTypeID" json:"-"`

	// Field reference
	FieldName string `gorm:"size:100;not null" json:"field" example:"product_code__c"`

	// Overrides
	Required *bool   `json:"required,omitempty"`
	Hidden   *bool   `json:"hidden,omitempty"`
	Editable *bool   `json:"editable,omitempty"`
	Label    *string `gorm:"size:40" json:"label,omitempty"`
}

// TableName returns the table name for TypeField
func (TypeField) TableName() string {
	return "vault_type_fields"
}

// TypeAction represents an action override for an object type (MDL: Objecttype.Typeaction)
type TypeAction struct {
	BaseModel

	// Parent
	ObjectTypeID uint        `gorm:"not null;index" json:"object_type_id"`
	ObjectType   *ObjectType `gorm:"foreignKey:ObjectTypeID" json:"-"`

	// Action reference
	ActionName string `gorm:"size:100;not null" json:"action" example:"approve__c"`

	// Overrides
	Active *bool   `json:"active,omitempty"`
	Label  *string `gorm:"size:60" json:"label,omitempty"`
}

// TableName returns the table name for TypeAction
func (TypeAction) TableName() string {
	return "vault_type_actions"
}

// =============================================================================
// OBJECT LIFECYCLE - Lifecycle definition
// =============================================================================

// ObjectLifecycle represents an object lifecycle (MDL: Objectlifecycle)
// Lifecycles define the states and transitions for object records.
//
// Example MDL:
//   CREATE Objectlifecycle product_lifecycle__c (
//     label('Product Lifecycle'),
//     active(true),
//     starting_state(Objectlifecyclestate.draft__c)
//   );
type ObjectLifecycle struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:255" json:"description,omitempty"`

	// Starting state
	StartingStateID *uint                 `json:"starting_state_id,omitempty"`
	StartingState   *ObjectLifecycleState `gorm:"foreignKey:StartingStateID" json:"starting_state,omitempty"`

	// State labels (ordered list)
	StateLabels   StringSlice `gorm:"type:json" json:"state_labels,omitempty"`
	LabeledStates StringSlice `gorm:"type:json" json:"labeled_states,omitempty"`

	// Child relationships
	States      []ObjectLifecycleState      `gorm:"foreignKey:LifecycleID" json:"states,omitempty"`
	Roles       []ObjectLifecycleRole       `gorm:"foreignKey:LifecycleID" json:"roles,omitempty"`
	Permissions []ObjectLifecyclePermission `gorm:"foreignKey:LifecycleID" json:"permissions,omitempty"`
	Events      []ObjectLifecycleEvent      `gorm:"foreignKey:LifecycleID" json:"events,omitempty"`

	// Objects using this lifecycle
	Objects []VaultObject `gorm:"many2many:object_available_lifecycles" json:"objects,omitempty"`
}

// TableName returns the table name for ObjectLifecycle
func (ObjectLifecycle) TableName() string {
	return "vault_object_lifecycles"
}

// =============================================================================
// OBJECT LIFECYCLE STATE
// =============================================================================

// ObjectLifecycleState represents a state in an object lifecycle (MDL: Objectlifecyclestate)
//
// Example MDL:
//   Objectlifecyclestate draft__c (
//     label('Draft'),
//     active(true),
//     record_status(active__v)
//   );
type ObjectLifecycleState struct {
	BaseModel
	ComponentBase

	// Parent lifecycle
	LifecycleID uint             `gorm:"not null;index" json:"lifecycle_id"`
	Lifecycle   *ObjectLifecycle `gorm:"foreignKey:LifecycleID" json:"-"`

	// Properties
	Description  *string      `gorm:"size:255" json:"description,omitempty"`
	RecordStatus RecordStatus `gorm:"size:30;default:'active__v'" json:"record_status" example:"active__v"`

	// Cancel state configuration
	CancelStateID              *uint                 `json:"cancel_state_id,omitempty"`
	CancelState                *ObjectLifecycleState `gorm:"foreignKey:CancelStateID" json:"cancel_state,omitempty"`
	SkipCancelState            bool                  `gorm:"default:false" json:"skip_cancel_state" example:"false"`
	SkipEntryActionsCancelState bool                 `gorm:"default:false" json:"skip_entry_actions_cancel_state" example:"false"`

	// Entry criteria and actions
	EntryCriteria []ObjectLifecycleStateEntryCriterion `gorm:"foreignKey:StateID" json:"entry_criteria,omitempty"`
	EntryActions  []ObjectLifecycleStateEntryAction    `gorm:"foreignKey:StateID" json:"entry_actions,omitempty"`
	UserActions   []ObjectLifecycleStateUserAction     `gorm:"foreignKey:StateID" json:"user_actions,omitempty"`
}

// TableName returns the table name for ObjectLifecycleState
func (ObjectLifecycleState) TableName() string {
	return "vault_object_lifecycle_states"
}

// =============================================================================
// ENTRY CRITERION
// =============================================================================

// ObjectLifecycleStateEntryCriterion represents an entry criterion (MDL: Objectlifecyclestateentrycriterion)
type ObjectLifecycleStateEntryCriterion struct {
	BaseModel

	// Parent state
	StateID uint                  `gorm:"not null;index" json:"state_id"`
	State   *ObjectLifecycleState `gorm:"foreignKey:StateID" json:"-"`

	// Properties
	Rule  XMLString `gorm:"type:text" json:"rule"`
	Order int       `gorm:"default:0" json:"order" example:"1"`
}

// TableName returns the table name for ObjectLifecycleStateEntryCriterion
func (ObjectLifecycleStateEntryCriterion) TableName() string {
	return "vault_object_lifecycle_state_entry_criteria"
}

// =============================================================================
// ENTRY ACTION
// =============================================================================

// ObjectLifecycleStateEntryAction represents an entry action (MDL: Objectlifecyclestateentryaction)
type ObjectLifecycleStateEntryAction struct {
	BaseModel

	// Parent state
	StateID uint                  `gorm:"not null;index" json:"state_id"`
	State   *ObjectLifecycleState `gorm:"foreignKey:StateID" json:"-"`

	// Properties
	Rule  XMLString `gorm:"type:text" json:"rule"`
	Order int       `gorm:"default:0" json:"order" example:"1"`
}

// TableName returns the table name for ObjectLifecycleStateEntryAction
func (ObjectLifecycleStateEntryAction) TableName() string {
	return "vault_object_lifecycle_state_entry_actions"
}

// =============================================================================
// USER ACTION
// =============================================================================

// ObjectLifecycleStateUserAction represents a user action (MDL: Objectlifecyclestateuseraction)
//
// Example MDL:
//   Objectlifecyclestateuseraction approve__c (
//     label('Approve'),
//     order(1),
//     rule(<action-rule>...</action-rule>)
//   );
type ObjectLifecycleStateUserAction struct {
	BaseModel

	// Parent state
	StateID uint                  `gorm:"not null;index" json:"state_id"`
	State   *ObjectLifecycleState `gorm:"foreignKey:StateID" json:"-"`

	// Properties
	Name             string      `gorm:"size:100;not null" json:"name" example:"approve__c"`
	Label            string      `gorm:"size:60;not null" json:"label" example:"Approve"`
	Rule             XMLString   `gorm:"type:text" json:"rule"`
	Order            int         `gorm:"not null" json:"order" example:"1"`
	KeyboardShortcut StringSlice `gorm:"type:json" json:"keyboard_shortcut,omitempty"`
}

// TableName returns the table name for ObjectLifecycleStateUserAction
func (ObjectLifecycleStateUserAction) TableName() string {
	return "vault_object_lifecycle_state_user_actions"
}

// =============================================================================
// LIFECYCLE EVENT
// =============================================================================

// ObjectLifecycleEvent represents a lifecycle event handler (MDL: Objectlifecycleevent)
type ObjectLifecycleEvent struct {
	BaseModel

	// Parent lifecycle
	LifecycleID uint             `gorm:"not null;index" json:"lifecycle_id"`
	Lifecycle   *ObjectLifecycle `gorm:"foreignKey:LifecycleID" json:"-"`

	// Properties
	Event string    `gorm:"size:1500;not null" json:"event" example:"RecordStateChange"`
	Rule  XMLString `gorm:"type:text" json:"rule"`
	Order int       `gorm:"not null" json:"order" example:"1"`
}

// TableName returns the table name for ObjectLifecycleEvent
func (ObjectLifecycleEvent) TableName() string {
	return "vault_object_lifecycle_events"
}

// =============================================================================
// LIFECYCLE ROLE
// =============================================================================

// ObjectLifecycleRole represents a role in a lifecycle (MDL: Objectlifecyclerole)
//
// Example MDL:
//   Objectlifecyclerole editor__v (
//     active(true),
//     application_role('editor__v'),
//     permissions(Read, Edit)
//   );
type ObjectLifecycleRole struct {
	BaseModel

	// Parent lifecycle
	LifecycleID uint             `gorm:"not null;index" json:"lifecycle_id"`
	Lifecycle   *ObjectLifecycle `gorm:"foreignKey:LifecycleID" json:"-"`

	// Properties
	Name            string      `gorm:"size:100;not null" json:"name" example:"editor__v"`
	Active          bool        `gorm:"not null;default:true" json:"active" example:"true"`
	ApplicationRole string      `gorm:"size:60;not null" json:"application_role" example:"editor__v"`
	Permissions     StringSlice `gorm:"type:json" json:"permissions" example:"[\"Read\",\"Edit\"]"`
}

// TableName returns the table name for ObjectLifecycleRole
func (ObjectLifecycleRole) TableName() string {
	return "vault_object_lifecycle_roles"
}

// =============================================================================
// LIFECYCLE PERMISSION
// =============================================================================

// ObjectLifecyclePermission represents state-specific permissions (MDL: Objectlifecyclepermission)
type ObjectLifecyclePermission struct {
	BaseModel

	// Parent lifecycle
	LifecycleID uint             `gorm:"not null;index" json:"lifecycle_id"`
	Lifecycle   *ObjectLifecycle `gorm:"foreignKey:LifecycleID" json:"-"`

	// Properties
	States      StringSlice `gorm:"type:json" json:"states"`
	Permissions StringSlice `gorm:"type:json" json:"permissions"`
}

// TableName returns the table name for ObjectLifecyclePermission
func (ObjectLifecyclePermission) TableName() string {
	return "vault_object_lifecycle_permissions"
}
