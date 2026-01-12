package vault

// =============================================================================
// WORKFLOW
// =============================================================================

// Workflow represents a workflow definition (MDL: Workflow)
//
// Example MDL:
//   CREATE Workflow approval_workflow__c (
//     label('Approval Workflow'),
//     active(true)
//   );
type Workflow struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Child relationships
	Steps []WorkflowStep `gorm:"foreignKey:WorkflowID" json:"steps,omitempty"`
}

// TableName returns the table name for Workflow
func (Workflow) TableName() string {
	return "vault_workflows"
}

// WorkflowStep represents a step in a workflow (MDL: Workflowstep)
type WorkflowStep struct {
	BaseModel
	ComponentBase

	// Parent workflow
	WorkflowID uint      `gorm:"not null;index" json:"workflow_id"`
	Workflow   *Workflow `gorm:"foreignKey:WorkflowID" json:"-"`

	// Properties
	Order int       `gorm:"default:0" json:"order" example:"1"`
	Rule  XMLString `gorm:"type:text" json:"rule,omitempty"`
}

// TableName returns the table name for WorkflowStep
func (WorkflowStep) TableName() string {
	return "vault_workflow_steps"
}

// =============================================================================
// OBJECT WORKFLOW
// =============================================================================

// ObjectWorkflow represents an object-specific workflow (MDL: Objectworkflow)
type ObjectWorkflow struct {
	BaseModel
	ComponentBase

	// Parent object
	ObjectID uint        `gorm:"not null;index" json:"object_id"`
	Object   *VaultObject `gorm:"foreignKey:ObjectID" json:"-"`

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Child relationships
	Steps             []ObjectWorkflowStep          `gorm:"foreignKey:WorkflowID" json:"steps,omitempty"`
	CancelationActions []WorkflowCancelationAction  `gorm:"foreignKey:WorkflowID" json:"cancelation_actions,omitempty"`
}

// TableName returns the table name for ObjectWorkflow
func (ObjectWorkflow) TableName() string {
	return "vault_object_workflows"
}

// ObjectWorkflowStep represents a step in an object workflow (MDL: Objectworkflowstep)
type ObjectWorkflowStep struct {
	BaseModel
	ComponentBase

	// Parent workflow
	WorkflowID uint            `gorm:"not null;index" json:"workflow_id"`
	Workflow   *ObjectWorkflow `gorm:"foreignKey:WorkflowID" json:"-"`

	// Properties
	Order int       `gorm:"default:0" json:"order" example:"1"`
	Rule  XMLString `gorm:"type:text" json:"rule,omitempty"`
}

// TableName returns the table name for ObjectWorkflowStep
func (ObjectWorkflowStep) TableName() string {
	return "vault_object_workflow_steps"
}

// WorkflowCancelationAction represents a cancelation action (MDL: WorkflowCancelationAction)
type WorkflowCancelationAction struct {
	BaseModel

	// Parent workflow
	WorkflowID uint            `gorm:"not null;index" json:"workflow_id"`
	Workflow   *ObjectWorkflow `gorm:"foreignKey:WorkflowID" json:"-"`

	// Properties
	Name  string    `gorm:"size:100;not null" json:"name"`
	Rule  XMLString `gorm:"type:text" json:"rule,omitempty"`
	Order int       `gorm:"default:0" json:"order" example:"1"`
}

// TableName returns the table name for WorkflowCancelationAction
func (WorkflowCancelationAction) TableName() string {
	return "vault_workflow_cancelation_actions"
}

// =============================================================================
// RECORD TRIGGER - SDK Triggers
// =============================================================================

// Recordtrigger represents a record trigger (MDL: Recordtrigger)
// Record triggers execute Vault Java SDK code when record events occur.
//
// Example MDL:
//   CREATE Recordtrigger product_validation__c (
//     label('Product Validation'),
//     active(true),
//     object('product__c'),
//     event('BEFORE_INSERT')
//   );
type Recordtrigger struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Target object
	ObjectName string `gorm:"size:100;not null" json:"object" example:"product__c"`

	// Event configuration
	Event TriggerEvent `gorm:"size:20;not null" json:"event" example:"BEFORE_INSERT"`
	Order int          `gorm:"default:0" json:"order" example:"1"`

	// SDK Code
	Checksum   *string `gorm:"size:1500" json:"checksum,omitempty"`
	SourceCode SdkCode `gorm:"type:text" json:"source_code,omitempty"`
}

// TableName returns the table name for Recordtrigger
func (Recordtrigger) TableName() string {
	return "vault_recordtriggers"
}

// =============================================================================
// ACTION TRIGGER - Low-code Triggers
// =============================================================================

// Actiontrigger represents an action trigger (MDL: Actiontrigger)
// Action triggers execute low-code actions when record events occur.
//
// Example MDL:
//   CREATE Actiontrigger update_status__c (
//     label('Update Status'),
//     active(true),
//     object(Object.product__c),
//     event(AFTER_INSERT)
//   );
type Actiontrigger struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Target object
	ObjectName string `gorm:"size:100;not null" json:"object" example:"product__c"`

	// Event configuration
	Event TriggerEvent `gorm:"size:20;not null" json:"event" example:"AFTER_INSERT"`
	Order int          `gorm:"default:0" json:"order" example:"1"`

	// Child relationships
	ActionBlocks      []Actionblock      `gorm:"foreignKey:TriggerID" json:"action_blocks,omitempty"`
	DraftActionBlocks []Draftactionblock `gorm:"foreignKey:TriggerID" json:"draft_action_blocks,omitempty"`
}

// TableName returns the table name for Actiontrigger
func (Actiontrigger) TableName() string {
	return "vault_actiontriggers"
}

// Actionblock represents an action block within an action trigger (MDL: Actionblock)
type Actionblock struct {
	BaseModel
	ComponentBase

	// Parent trigger
	TriggerID uint           `gorm:"not null;index" json:"trigger_id"`
	Trigger   *Actiontrigger `gorm:"foreignKey:TriggerID" json:"-"`

	// Properties
	Description *string      `gorm:"size:500" json:"description,omitempty"`
	Order       int          `gorm:"not null" json:"order" example:"1"`
	Code        ActionScript `gorm:"type:text" json:"code,omitempty"`
}

// TableName returns the table name for Actionblock
func (Actionblock) TableName() string {
	return "vault_actionblocks"
}

// Draftactionblock represents a draft action block (MDL: Draftactionblock)
type Draftactionblock struct {
	BaseModel
	ComponentBase

	// Parent trigger
	TriggerID uint           `gorm:"not null;index" json:"trigger_id"`
	Trigger   *Actiontrigger `gorm:"foreignKey:TriggerID" json:"-"`

	// Properties
	Description *string      `gorm:"size:500" json:"description,omitempty"`
	Order       int          `gorm:"not null" json:"order" example:"1"`
	Code        ActionScript `gorm:"type:text" json:"code,omitempty"`
}

// TableName returns the table name for Draftactionblock
func (Draftactionblock) TableName() string {
	return "vault_draftactionblocks"
}

// =============================================================================
// RECORD ACTION
// =============================================================================

// Recordaction represents a custom record action (MDL: Recordaction)
type Recordaction struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Target object
	ObjectName string `gorm:"size:100;not null" json:"object" example:"product__c"`

	// SDK Code
	Checksum   *string `gorm:"size:1500" json:"checksum,omitempty"`
	SourceCode SdkCode `gorm:"type:text" json:"source_code,omitempty"`
}

// TableName returns the table name for Recordaction
func (Recordaction) TableName() string {
	return "vault_recordactions"
}

// =============================================================================
// OBJECT ACTION
// =============================================================================

// Objectaction represents a custom object action (MDL: Objectaction)
type Objectaction struct {
	BaseModel
	ComponentBase

	// Parent object
	ObjectID uint        `gorm:"not null;index" json:"object_id"`
	Object   *VaultObject `gorm:"foreignKey:ObjectID" json:"-"`

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`
	HelpContent *string `gorm:"size:255" json:"help_content,omitempty"`

	// Action rule
	Rule XMLString `gorm:"type:text" json:"rule,omitempty"`
}

// TableName returns the table name for Objectaction
func (Objectaction) TableName() string {
	return "vault_objectactions"
}

// =============================================================================
// RECORD ROLE TRIGGER
// =============================================================================

// Recordroletrigger represents a record role trigger (MDL: Recordroletrigger)
type Recordroletrigger struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Target object
	ObjectName string `gorm:"size:100;not null" json:"object" example:"product__c"`

	// SDK Code
	Checksum   *string `gorm:"size:1500" json:"checksum,omitempty"`
	SourceCode SdkCode `gorm:"type:text" json:"source_code,omitempty"`
}

// TableName returns the table name for Recordroletrigger
func (Recordroletrigger) TableName() string {
	return "vault_recordroletriggers"
}

// =============================================================================
// RECORD WORKFLOW ACTION
// =============================================================================

// Recordworkflowaction represents a workflow action (MDL: Recordworkflowaction)
type Recordworkflowaction struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// SDK Code
	Checksum   *string `gorm:"size:1500" json:"checksum,omitempty"`
	SourceCode SdkCode `gorm:"type:text" json:"source_code,omitempty"`
}

// TableName returns the table name for Recordworkflowaction
func (Recordworkflowaction) TableName() string {
	return "vault_recordworkflowactions"
}
