package vault

// =============================================================================
// CUSTOM WEB API
// =============================================================================

// Customwebapi represents a custom web API (MDL: Customwebapi)
//
// Example MDL:
//   CREATE Customwebapi my_api__c (
//     label('My Custom API'),
//     active(true),
//     api_group(Webapigroup.custom_apis__c),
//     endpoint_name('myapi'),
//     minimum_version('v21.1')
//   );
type Customwebapi struct {
	BaseModel
	ComponentBase

	// Properties
	APIGroup       string `gorm:"size:100;not null" json:"api_group" example:"custom_apis__c"`
	EndpointName   string `gorm:"size:1500;not null" json:"endpoint_name" example:"myapi"`
	MinimumVersion string `gorm:"size:1500;not null" json:"minimum_version" example:"v21.1"`

	// SDK Code
	Checksum   *string `gorm:"size:1500" json:"checksum,omitempty"`
	SourceCode SdkCode `gorm:"type:text" json:"source_code,omitempty"`
}

// TableName returns the table name for Customwebapi
func (Customwebapi) TableName() string {
	return "vault_customwebapis"
}

// =============================================================================
// WEB API GROUP
// =============================================================================

// Webapigroup represents a web API group (MDL: Webapigroup)
type Webapigroup struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:255" json:"description,omitempty"`
}

// TableName returns the table name for Webapigroup
func (Webapigroup) TableName() string {
	return "vault_webapigroups"
}

// =============================================================================
// USER DEFINED CLASS
// =============================================================================

// Userdefinedclass represents a user-defined Java class (MDL: Userdefinedclass)
type Userdefinedclass struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// SDK Code
	Checksum   *string `gorm:"size:1500" json:"checksum,omitempty"`
	SourceCode SdkCode `gorm:"type:text" json:"source_code,omitempty"`
}

// TableName returns the table name for Userdefinedclass
func (Userdefinedclass) TableName() string {
	return "vault_userdefinedclasses"
}

// =============================================================================
// USER DEFINED MODEL
// =============================================================================

// Userdefinedmodel represents a user-defined model class (MDL: Userdefinedmodel)
type Userdefinedmodel struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// SDK Code
	Checksum   *string `gorm:"size:1500" json:"checksum,omitempty"`
	SourceCode SdkCode `gorm:"type:text" json:"source_code,omitempty"`
}

// TableName returns the table name for Userdefinedmodel
func (Userdefinedmodel) TableName() string {
	return "vault_userdefinedmodels"
}

// =============================================================================
// USER DEFINED SERVICE
// =============================================================================

// Userdefinedservice represents a user-defined service class (MDL: Userdefinedservice)
type Userdefinedservice struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// SDK Code
	Checksum   *string `gorm:"size:1500" json:"checksum,omitempty"`
	SourceCode SdkCode `gorm:"type:text" json:"source_code,omitempty"`
}

// TableName returns the table name for Userdefinedservice
func (Userdefinedservice) TableName() string {
	return "vault_userdefinedservices"
}

// =============================================================================
// SDK JOB
// =============================================================================

// Sdkjob represents an SDK job handler (MDL: Sdkjob)
type Sdkjob struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// SDK Code
	Checksum   *string `gorm:"size:1500" json:"checksum,omitempty"`
	SourceCode SdkCode `gorm:"type:text" json:"source_code,omitempty"`
}

// TableName returns the table name for Sdkjob
func (Sdkjob) TableName() string {
	return "vault_sdkjobs"
}

// =============================================================================
// JOB - Scheduled Jobs
// =============================================================================

// JobActionType represents the type of job action
type JobActionType string

const (
	JobActionTypeSendNotification     JobActionType = "sendnotification"
	JobActionTypeChangeObjectState    JobActionType = "changeobjectstate"
	JobActionTypeNoop                 JobActionType = "noop"
	JobActionTypeHTTPPost             JobActionType = "httppost"
	JobActionTypeSdkBatch             JobActionType = "sdkbatch"
	JobActionTypeEdlMatch             JobActionType = "edlmatch"
	JobActionTypeFlashReport          JobActionType = "flashreport"
	JobActionTypeScheduledDataExport  JobActionType = "scheduleddataexport"
	JobActionTypeDocumentPeriodicReview JobActionType = "documentperiodicreview"
)

// Job represents a scheduled job (MDL: Job)
type Job struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Schedule
	CronExpression *string `gorm:"size:255" json:"cron_expression,omitempty"`
	Timezone       *string `gorm:"size:100" json:"timezone,omitempty"`

	// Child relationships (job actions)
	SendNotificationActions []SendNotificationJobAction `gorm:"foreignKey:JobID" json:"send_notification_actions,omitempty"`
	ChangeStateActions      []ChangeObjectStateJobAction `gorm:"foreignKey:JobID" json:"change_state_actions,omitempty"`
	SdkBatchActions         []SdkBatchJobAction          `gorm:"foreignKey:JobID" json:"sdk_batch_actions,omitempty"`
}

// TableName returns the table name for Job
func (Job) TableName() string {
	return "vault_jobs"
}

// SendNotificationJobAction represents a send notification job action
type SendNotificationJobAction struct {
	BaseModel
	ComponentBase

	// Parent job
	JobID uint `gorm:"not null;index" json:"job_id"`
	Job   *Job `gorm:"foreignKey:JobID" json:"-"`

	// Properties
	Order int `gorm:"default:0" json:"order"`
}

// TableName returns the table name for SendNotificationJobAction
func (SendNotificationJobAction) TableName() string {
	return "vault_job_send_notifications"
}

// ChangeObjectStateJobAction represents a change state job action
type ChangeObjectStateJobAction struct {
	BaseModel
	ComponentBase

	// Parent job
	JobID uint `gorm:"not null;index" json:"job_id"`
	Job   *Job `gorm:"foreignKey:JobID" json:"-"`

	// Properties
	Order int `gorm:"default:0" json:"order"`
}

// TableName returns the table name for ChangeObjectStateJobAction
func (ChangeObjectStateJobAction) TableName() string {
	return "vault_job_change_states"
}

// SdkBatchJobAction represents an SDK batch job action
type SdkBatchJobAction struct {
	BaseModel
	ComponentBase

	// Parent job
	JobID uint `gorm:"not null;index" json:"job_id"`
	Job   *Job `gorm:"foreignKey:JobID" json:"-"`

	// Properties
	Order      int     `gorm:"default:0" json:"order"`
	Checksum   *string `gorm:"size:1500" json:"checksum,omitempty"`
	SourceCode SdkCode `gorm:"type:text" json:"source_code,omitempty"`
}

// TableName returns the table name for SdkBatchJobAction
func (SdkBatchJobAction) TableName() string {
	return "vault_job_sdk_batches"
}

// =============================================================================
// QUEUE - Message Queues
// =============================================================================

// Queue represents a message queue (MDL: Queue)
type Queue struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// SDK Code
	Checksum   *string `gorm:"size:1500" json:"checksum,omitempty"`
	SourceCode SdkCode `gorm:"type:text" json:"source_code,omitempty"`

	// Child relationships
	Connections []QueueConnection `gorm:"foreignKey:QueueID" json:"connections,omitempty"`
}

// TableName returns the table name for Queue
func (Queue) TableName() string {
	return "vault_queues"
}

// QueueConnection represents a queue connection (MDL: Queueconnection)
type QueueConnection struct {
	BaseModel
	ComponentBase

	// Parent queue
	QueueID uint   `gorm:"not null;index" json:"queue_id"`
	Queue   *Queue `gorm:"foreignKey:QueueID" json:"-"`

	// Properties
	ConnectionType *string `gorm:"size:100" json:"connection_type,omitempty"`
}

// TableName returns the table name for QueueConnection
func (QueueConnection) TableName() string {
	return "vault_queue_connections"
}

// =============================================================================
// MESSAGE GROUP
// =============================================================================

// Messagegroup represents a message group (MDL: Messagegroup)
type Messagegroup struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Child relationships
	Messages []Message `gorm:"foreignKey:GroupID" json:"messages,omitempty"`
}

// TableName returns the table name for Messagegroup
func (Messagegroup) TableName() string {
	return "vault_messagegroups"
}

// Message represents a message in a message group (MDL: Message)
type Message struct {
	BaseModel
	ComponentBase

	// Parent group
	GroupID uint          `gorm:"not null;index" json:"group_id"`
	Group   *Messagegroup `gorm:"foreignKey:GroupID" json:"-"`

	// Properties
	MessageText string `gorm:"type:text" json:"message_text"`
}

// TableName returns the table name for Message
func (Message) TableName() string {
	return "vault_messages"
}

// =============================================================================
// NOTIFICATION TEMPLATE
// =============================================================================

// Notificationtemplate represents a notification template (MDL: Notificationtemplate)
type Notificationtemplate struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string    `gorm:"size:500" json:"description,omitempty"`
	Subject     string     `gorm:"size:255" json:"subject" example:"Record Updated: {name}"`
	Body        LongString `gorm:"type:text" json:"body"`
}

// TableName returns the table name for Notificationtemplate
func (Notificationtemplate) TableName() string {
	return "vault_notificationtemplates"
}

// =============================================================================
// ACCOUNT MESSAGE
// =============================================================================

// Accountmessage represents an account message (MDL: Accountmessage)
type Accountmessage struct {
	BaseModel
	ComponentBase

	// Properties
	DefaultLang *string    `gorm:"size:255" json:"default_lang,omitempty"`
	Subject     string     `gorm:"size:255;not null" json:"subject" example:"Welcome to Vault"`
	EmailBody   LongString `gorm:"type:text;not null" json:"email_body"`
}

// TableName returns the table name for Accountmessage
func (Accountmessage) TableName() string {
	return "vault_accountmessages"
}
