// Package vault provides Go ORM models for Veeva Vault objects and components.
// This file is auto-generated from the official Veeva MDL Components documentation.
// It contains Enum definitions derived from validation rules.
package vault

// ActiontriggerEvent represents values for Actiontrigger.event
type ActiontriggerEvent string

const (
	ActiontriggerEventBeforeInsert ActiontriggerEvent = "BEFORE_INSERT"
	ActiontriggerEventBeforeUpdate ActiontriggerEvent = "BEFORE_UPDATE"
	ActiontriggerEventBeforeDelete ActiontriggerEvent = "BEFORE_DELETE"
	ActiontriggerEventAfterInsert  ActiontriggerEvent = "AFTER_INSERT"
	ActiontriggerEventAfterUpdate  ActiontriggerEvent = "AFTER_UPDATE"
	ActiontriggerEventAfterDelete  ActiontriggerEvent = "AFTER_DELETE"
)

// AppsecurityruleappsecurityfieldType represents values for AppsecurityruleAppsecurityfield.type
type AppsecurityruleappsecurityfieldType string

const (
	AppsecurityruleappsecurityfieldTypeObject   AppsecurityruleappsecurityfieldType = "OBJECT"
	AppsecurityruleappsecurityfieldTypePicklist AppsecurityruleappsecurityfieldType = "PICKLIST"
	AppsecurityruleappsecurityfieldTypeBoolean  AppsecurityruleappsecurityfieldType = "BOOLEAN"
)

// AtomicsecurityfieldsecurityType represents values for AtomicsecurityFieldsecurity.type
type AtomicsecurityfieldsecurityType string

const (
	AtomicsecurityfieldsecurityTypeHideV AtomicsecurityfieldsecurityType = "hide__v"
	AtomicsecurityfieldsecurityTypeReadV AtomicsecurityfieldsecurityType = "read__v"
	AtomicsecurityfieldsecurityTypeEditV AtomicsecurityfieldsecurityType = "edit__v"
)

// AtomicsecurityobjectcontrolsecurityType represents values for AtomicsecurityObjectcontrolsecurity.type
type AtomicsecurityobjectcontrolsecurityType string

const (
	AtomicsecurityobjectcontrolsecurityTypeHideV AtomicsecurityobjectcontrolsecurityType = "hide__v"
	AtomicsecurityobjectcontrolsecurityTypeViewV AtomicsecurityobjectcontrolsecurityType = "view__v"
)

// AtomicsecurityactionsecurityType represents values for AtomicsecurityActionsecurity.type
type AtomicsecurityactionsecurityType string

const (
	AtomicsecurityactionsecurityTypeHideV    AtomicsecurityactionsecurityType = "hide__v"
	AtomicsecurityactionsecurityTypeViewV    AtomicsecurityactionsecurityType = "view__v"
	AtomicsecurityactionsecurityTypeExecuteV AtomicsecurityactionsecurityType = "execute__v"
)

// AtomicsecurityworkflowactionsecurityType represents values for AtomicsecurityWorkflowactionsecurity.type
type AtomicsecurityworkflowactionsecurityType string

const (
	AtomicsecurityworkflowactionsecurityTypeHideV    AtomicsecurityworkflowactionsecurityType = "hide__v"
	AtomicsecurityworkflowactionsecurityTypeExecuteV AtomicsecurityworkflowactionsecurityType = "execute__v"
)

// AtomicsecurityworkflowactionsecurityWorkflowActions represents values for AtomicsecurityWorkflowactionsecurity.workflow_actions
type AtomicsecurityworkflowactionsecurityWorkflowActions string

const (
	AtomicsecurityworkflowactionsecurityWorkflowActionsAddparticipants       AtomicsecurityworkflowactionsecurityWorkflowActions = "addparticipants"
	AtomicsecurityworkflowactionsecurityWorkflowActionsCancel                AtomicsecurityworkflowactionsecurityWorkflowActions = "cancel"
	AtomicsecurityworkflowactionsecurityWorkflowActionsEmailparticipants     AtomicsecurityworkflowactionsecurityWorkflowActions = "emailparticipants"
	AtomicsecurityworkflowactionsecurityWorkflowActionsRemovecontent         AtomicsecurityworkflowactionsecurityWorkflowActions = "removecontent"
	AtomicsecurityworkflowactionsecurityWorkflowActionsUpdateworkflowduedate AtomicsecurityworkflowactionsecurityWorkflowActions = "updateworkflowduedate"
)

// AtomicsecurityworkflowactionsecurityWorkflowTaskActions represents values for AtomicsecurityWorkflowactionsecurity.workflow_task_actions
type AtomicsecurityworkflowactionsecurityWorkflowTaskActions string

const (
	AtomicsecurityworkflowactionsecurityWorkflowTaskActionsCancel        AtomicsecurityworkflowactionsecurityWorkflowTaskActions = "cancel"
	AtomicsecurityworkflowactionsecurityWorkflowTaskActionsReassign      AtomicsecurityworkflowactionsecurityWorkflowTaskActions = "reassign"
	AtomicsecurityworkflowactionsecurityWorkflowTaskActionsUpdateduedate AtomicsecurityworkflowactionsecurityWorkflowTaskActions = "updateduedate"
)

// AtomicsecurityrelationshipsecurityType represents values for AtomicsecurityRelationshipsecurity.type
type AtomicsecurityrelationshipsecurityType string

const (
	AtomicsecurityrelationshipsecurityTypeReadV AtomicsecurityrelationshipsecurityType = "read__v"
	AtomicsecurityrelationshipsecurityTypeEditV AtomicsecurityrelationshipsecurityType = "edit__v"
)

// DisclosureruleComplianceDateShiftUnits represents values for Disclosurerule.compliance_date_shift_units
type DisclosureruleComplianceDateShiftUnits string

const (
	DisclosureruleComplianceDateShiftUnitsDays   DisclosureruleComplianceDateShiftUnits = "days"
	DisclosureruleComplianceDateShiftUnitsMonths DisclosureruleComplianceDateShiftUnits = "months"
	DisclosureruleComplianceDateShiftUnitsYears  DisclosureruleComplianceDateShiftUnits = "years"
)

// DisclosureruleDueDateShiftUnits represents values for Disclosurerule.due_date_shift_units
type DisclosureruleDueDateShiftUnits string

const (
	DisclosureruleDueDateShiftUnitsDays   DisclosureruleDueDateShiftUnits = "days"
	DisclosureruleDueDateShiftUnitsMonths DisclosureruleDueDateShiftUnits = "months"
)

// DisclosurexmldoctypemappingClinicalTrialsDocumentType represents values for Disclosurexmldoctypemapping.clinical_trials_document_type
type DisclosurexmldoctypemappingClinicalTrialsDocumentType string

const (
	DisclosurexmldoctypemappingClinicalTrialsDocumentTypeProtocol             DisclosurexmldoctypemappingClinicalTrialsDocumentType = "Protocol"
	DisclosurexmldoctypemappingClinicalTrialsDocumentTypeSap                  DisclosurexmldoctypemappingClinicalTrialsDocumentType = "SAP"
	DisclosurexmldoctypemappingClinicalTrialsDocumentTypeIcf                  DisclosurexmldoctypemappingClinicalTrialsDocumentType = "ICF"
	DisclosurexmldoctypemappingClinicalTrialsDocumentTypeProtocolAndSap       DisclosurexmldoctypemappingClinicalTrialsDocumentType = "Protocol_and_SAP"
	DisclosurexmldoctypemappingClinicalTrialsDocumentTypeProtocolAndIcf       DisclosurexmldoctypemappingClinicalTrialsDocumentType = "Protocol_and_ICF"
	DisclosurexmldoctypemappingClinicalTrialsDocumentTypeProtocolAndSapAndIcf DisclosurexmldoctypemappingClinicalTrialsDocumentType = "Protocol_and_SAP_and_ICF"
)

// DocatomicsecuritydocworkflowactionsecurityType represents values for DocatomicsecurityDocworkflowactionsecurity.type
type DocatomicsecuritydocworkflowactionsecurityType string

const (
	DocatomicsecuritydocworkflowactionsecurityTypeHideV    DocatomicsecuritydocworkflowactionsecurityType = "hide__v"
	DocatomicsecuritydocworkflowactionsecurityTypeViewV    DocatomicsecuritydocworkflowactionsecurityType = "view__v"
	DocatomicsecuritydocworkflowactionsecurityTypeExecuteV DocatomicsecuritydocworkflowactionsecurityType = "execute__v"
)

// DocatomicsecuritydocworkflowactionsecurityWorkflowActions represents values for DocatomicsecurityDocworkflowactionsecurity.workflow_actions
type DocatomicsecuritydocworkflowactionsecurityWorkflowActions string

const (
	DocatomicsecuritydocworkflowactionsecurityWorkflowActionsCancel                DocatomicsecuritydocworkflowactionsecurityWorkflowActions = "cancel"
	DocatomicsecuritydocworkflowactionsecurityWorkflowActionsAddparticipants       DocatomicsecuritydocworkflowactionsecurityWorkflowActions = "addparticipants"
	DocatomicsecuritydocworkflowactionsecurityWorkflowActionsRemovecontent         DocatomicsecuritydocworkflowactionsecurityWorkflowActions = "removecontent"
	DocatomicsecuritydocworkflowactionsecurityWorkflowActionsEmailparticipants     DocatomicsecuritydocworkflowactionsecurityWorkflowActions = "emailparticipants"
	DocatomicsecuritydocworkflowactionsecurityWorkflowActionsUpdateworkflowduedate DocatomicsecuritydocworkflowactionsecurityWorkflowActions = "updateworkflowduedate"
)

// DocatomicsecuritydocworkflowactionsecurityWorkflowTaskActions represents values for DocatomicsecurityDocworkflowactionsecurity.workflow_task_actions
type DocatomicsecuritydocworkflowactionsecurityWorkflowTaskActions string

const (
	DocatomicsecuritydocworkflowactionsecurityWorkflowTaskActionsCancel        DocatomicsecuritydocworkflowactionsecurityWorkflowTaskActions = "cancel"
	DocatomicsecuritydocworkflowactionsecurityWorkflowTaskActionsReassign      DocatomicsecuritydocworkflowactionsecurityWorkflowTaskActions = "reassign"
	DocatomicsecuritydocworkflowactionsecurityWorkflowTaskActionsUpdateduedate DocatomicsecuritydocworkflowactionsecurityWorkflowTaskActions = "updateduedate"
)

// DocatomicsecuritydocactionsecurityType represents values for DocatomicsecurityDocactionsecurity.type
type DocatomicsecuritydocactionsecurityType string

const (
	DocatomicsecuritydocactionsecurityTypeHideV    DocatomicsecuritydocactionsecurityType = "hide__v"
	DocatomicsecuritydocactionsecurityTypeViewV    DocatomicsecuritydocactionsecurityType = "view__v"
	DocatomicsecuritydocactionsecurityTypeExecuteV DocatomicsecuritydocactionsecurityType = "execute__v"
)

// DocfieldType represents values for Docfield.type
type DocfieldType string

const (
	DocfieldTypeString          DocfieldType = "String"
	DocfieldTypeNumber          DocfieldType = "Number"
	DocfieldTypeBoolean         DocfieldType = "Boolean"
	DocfieldTypePicklist        DocfieldType = "Picklist"
	DocfieldTypeDate            DocfieldType = "Date"
	DocfieldTypeUrl             DocfieldType = "URL"
	DocfieldTypeFormula         DocfieldType = "Formula"
	DocfieldTypeObjectreference DocfieldType = "ObjectReference"
	DocfieldTypeLookup          DocfieldType = "Lookup"
)

// DocfieldScope represents values for Docfield.scope
type DocfieldScope string

const (
	DocfieldScopeDocument        DocfieldScope = "Document"
	DocfieldScopeDocumentversion DocfieldScope = "DocumentVersion"
)

// DocfielddependencyDepruleDocfieldAction represents values for Docfielddependency.deprule_docfield_action
type DocfielddependencyDepruleDocfieldAction string

const (
	DocfielddependencyDepruleDocfieldActionIshidden   DocfielddependencyDepruleDocfieldAction = "isHidden"
	DocfielddependencyDepruleDocfieldActionIsreadonly DocfielddependencyDepruleDocfieldAction = "isReadOnly"
	DocfielddependencyDepruleDocfieldActionIsvisible  DocfielddependencyDepruleDocfieldAction = "isVisible"
	DocfielddependencyDepruleDocfieldActionIsrequired DocfielddependencyDepruleDocfieldAction = "isRequired"
	DocfielddependencyDepruleDocfieldActionPicklist   DocfielddependencyDepruleDocfieldAction = "picklist"
)

// DocfieldlayoutIcon represents values for Docfieldlayout.icon
type DocfieldlayoutIcon string

const (
	DocfieldlayoutIconUndefined     DocfieldlayoutIcon = "UNDEFINED"
	DocfieldlayoutIconAustralia     DocfieldlayoutIcon = "AUSTRALIA"
	DocfieldlayoutIconAustria       DocfieldlayoutIcon = "AUSTRIA"
	DocfieldlayoutIconBelgium       DocfieldlayoutIcon = "BELGIUM"
	DocfieldlayoutIconCanada        DocfieldlayoutIcon = "CANADA"
	DocfieldlayoutIconChina         DocfieldlayoutIcon = "CHINA"
	DocfieldlayoutIconEuropeanUnion DocfieldlayoutIcon = "EUROPEAN_UNION"
	DocfieldlayoutIconFrance        DocfieldlayoutIcon = "FRANCE"
	DocfieldlayoutIconGermany       DocfieldlayoutIcon = "GERMANY"
	DocfieldlayoutIconGlobal        DocfieldlayoutIcon = "GLOBAL"
	DocfieldlayoutIconIreland       DocfieldlayoutIcon = "IRELAND"
	DocfieldlayoutIconItaly         DocfieldlayoutIcon = "ITALY"
	DocfieldlayoutIconJapan         DocfieldlayoutIcon = "JAPAN"
	DocfieldlayoutIconNetherlands   DocfieldlayoutIcon = "NETHERLANDS"
	DocfieldlayoutIconSpain         DocfieldlayoutIcon = "SPAIN"
	DocfieldlayoutIconSweden        DocfieldlayoutIcon = "SWEDEN"
	DocfieldlayoutIconSwitzerland   DocfieldlayoutIcon = "SWITZERLAND"
	DocfieldlayoutIconUnitedKingdom DocfieldlayoutIcon = "UNITED_KINGDOM"
	DocfieldlayoutIconUnitedStates  DocfieldlayoutIcon = "UNITED_STATES"
	DocfieldlayoutIconVeeva         DocfieldlayoutIcon = "VEEVA"
)

// DoclifecycledoclifecycleroleNoteColor represents values for DoclifecycleDoclifecyclerole.note_color
type DoclifecycledoclifecycleroleNoteColor string

const (
	DoclifecycledoclifecycleroleNoteColorDarkOrange DoclifecycledoclifecycleroleNoteColor = "Dark Orange"
	DoclifecycledoclifecycleroleNoteColorOrange     DoclifecycledoclifecycleroleNoteColor = "Orange"
	DoclifecycledoclifecycleroleNoteColorDarkYellow DoclifecycledoclifecycleroleNoteColor = "Dark Yellow"
	DoclifecycledoclifecycleroleNoteColorYellow     DoclifecycledoclifecycleroleNoteColor = "Yellow"
	DoclifecycledoclifecycleroleNoteColorLime       DoclifecycledoclifecycleroleNoteColor = "Lime"
	DoclifecycledoclifecycleroleNoteColorGreen      DoclifecycledoclifecycleroleNoteColor = "Green"
	DoclifecycledoclifecycleroleNoteColorAqua       DoclifecycledoclifecycleroleNoteColor = "Aqua"
	DoclifecycledoclifecycleroleNoteColorRoyal      DoclifecycledoclifecycleroleNoteColor = "Royal"
	DoclifecycledoclifecycleroleNoteColorPurple     DoclifecycledoclifecycleroleNoteColor = "Purple"
	DoclifecycledoclifecycleroleNoteColorLilac      DoclifecycledoclifecycleroleNoteColor = "Lilac"
	DoclifecycledoclifecycleroleNoteColorPink       DoclifecycledoclifecycleroleNoteColor = "Pink"
	DoclifecycledoclifecycleroleNoteColorSalmon     DoclifecycledoclifecycleroleNoteColor = "Salmon"
)

// EmailprocessorAllowedSenders represents values for Emailprocessor.allowed_senders
type EmailprocessorAllowedSenders string

const (
	EmailprocessorAllowedSendersVaultUsersAndPersons EmailprocessorAllowedSenders = "VAULT_USERS_AND_PERSONS"
	EmailprocessorAllowedSendersVaultUsers           EmailprocessorAllowedSenders = "VAULT_USERS"
	EmailprocessorAllowedSendersVaultGroups          EmailprocessorAllowedSenders = "VAULT_GROUPS"
)

// FormattedoutputInputFileType represents values for Formattedoutput.input_file_type
type FormattedoutputInputFileType string

const (
	FormattedoutputInputFileTypeXfaPdf FormattedoutputInputFileType = "XFA_PDF"
	FormattedoutputInputFileTypeMsWord FormattedoutputInputFileType = "MS_WORD"
)

// FormattedoutputOutputFormat represents values for Formattedoutput.output_format
type FormattedoutputOutputFormat string

const (
	FormattedoutputOutputFormatPdfFlat FormattedoutputOutputFormat = "PDF_FLAT"
	FormattedoutputOutputFormatMsWord  FormattedoutputOutputFormat = "MS_WORD"
)

// InboundemailaddressAllowedSenders represents values for Inboundemailaddress.allowed_senders
type InboundemailaddressAllowedSenders string

const (
	InboundemailaddressAllowedSendersVaultUsersAndPersons InboundemailaddressAllowedSenders = "VAULT_USERS_AND_PERSONS"
	InboundemailaddressAllowedSendersVaultUsers           InboundemailaddressAllowedSenders = "VAULT_USERS"
	InboundemailaddressAllowedSendersVaultGroups          InboundemailaddressAllowedSenders = "VAULT_GROUPS"
	InboundemailaddressAllowedSendersUnspecified          InboundemailaddressAllowedSenders = "UNSPECIFIED"
)

// IntegrationrulefieldruleQueryFieldType represents values for IntegrationruleFieldrule.query_field_type
type IntegrationrulefieldruleQueryFieldType string

const (
	IntegrationrulefieldruleQueryFieldTypeDocumentBooleanSys              IntegrationrulefieldruleQueryFieldType = "document_boolean__sys"
	IntegrationrulefieldruleQueryFieldTypeDocumentComponentSys            IntegrationrulefieldruleQueryFieldType = "document_component__sys"
	IntegrationrulefieldruleQueryFieldTypeDocumentDateSys                 IntegrationrulefieldruleQueryFieldType = "document_date__sys"
	IntegrationrulefieldruleQueryFieldTypeDocumentDatetimeSys             IntegrationrulefieldruleQueryFieldType = "document_datetime__sys"
	IntegrationrulefieldruleQueryFieldTypeDocumentFormulaSys              IntegrationrulefieldruleQueryFieldType = "document_formula__sys"
	IntegrationrulefieldruleQueryFieldTypeDocumentLookupSys               IntegrationrulefieldruleQueryFieldType = "document_lookup__sys"
	IntegrationrulefieldruleQueryFieldTypeDocumentNumberSys               IntegrationrulefieldruleQueryFieldType = "document_number__sys"
	IntegrationrulefieldruleQueryFieldTypeDocumentObjectReferenceSys      IntegrationrulefieldruleQueryFieldType = "document_object_reference__sys"
	IntegrationrulefieldruleQueryFieldTypeDocumentPicklistSys             IntegrationrulefieldruleQueryFieldType = "document_picklist__sys"
	IntegrationrulefieldruleQueryFieldTypeDocumentTextSys                 IntegrationrulefieldruleQueryFieldType = "document_text__sys"
	IntegrationrulefieldruleQueryFieldTypeDocumentUrlSys                  IntegrationrulefieldruleQueryFieldType = "document_url__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectBooleanSys               IntegrationrulefieldruleQueryFieldType = "vobject_boolean__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectComponentSys             IntegrationrulefieldruleQueryFieldType = "vobject_component__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectCurrencySys              IntegrationrulefieldruleQueryFieldType = "vobject_currency__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectDateSys                  IntegrationrulefieldruleQueryFieldType = "vobject_date__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectDatetimeSys              IntegrationrulefieldruleQueryFieldType = "vobject_datetime__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectDocumentReferenceSys     IntegrationrulefieldruleQueryFieldType = "vobject_document_reference__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectFormulaSys               IntegrationrulefieldruleQueryFieldType = "vobject_formula__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectIdSys                    IntegrationrulefieldruleQueryFieldType = "vobject_id__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectLongtextSys              IntegrationrulefieldruleQueryFieldType = "vobject_longtext__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectLookupSys                IntegrationrulefieldruleQueryFieldType = "vobject_lookup__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectNumberSys                IntegrationrulefieldruleQueryFieldType = "vobject_number__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectObjectReferenceSys       IntegrationrulefieldruleQueryFieldType = "vobject_object_reference__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectParentObjectReferenceSys IntegrationrulefieldruleQueryFieldType = "vobject_parent_object_reference__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectPicklistSys              IntegrationrulefieldruleQueryFieldType = "vobject_picklist__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectRichtextSys              IntegrationrulefieldruleQueryFieldType = "vobject_richtext__sys"
	IntegrationrulefieldruleQueryFieldTypeVobjectTextSys                  IntegrationrulefieldruleQueryFieldType = "vobject_text__sys"
)

// JobRepeatOnDays represents values for Job.repeat_on_days
type JobRepeatOnDays string

const (
	JobRepeatOnDaysSun JobRepeatOnDays = "Sun"
	JobRepeatOnDaysMon JobRepeatOnDays = "Mon"
	JobRepeatOnDaysTue JobRepeatOnDays = "Tue"
	JobRepeatOnDaysWed JobRepeatOnDays = "Wed"
	JobRepeatOnDaysThu JobRepeatOnDays = "Thu"
	JobRepeatOnDaysFri JobRepeatOnDays = "Fri"
	JobRepeatOnDaysSat JobRepeatOnDays = "Sat"
)

// JobWeekDay represents values for Job.week_day
type JobWeekDay string

const (
	JobWeekDaySun JobWeekDay = "Sun"
	JobWeekDayMon JobWeekDay = "Mon"
	JobWeekDayTue JobWeekDay = "Tue"
	JobWeekDayWed JobWeekDay = "Wed"
	JobWeekDayThu JobWeekDay = "Thu"
	JobWeekDayFri JobWeekDay = "Fri"
	JobWeekDaySat JobWeekDay = "Sat"
)

// JobMonthRepeatType represents values for Job.month_repeat_type
type JobMonthRepeatType string

const (
	JobMonthRepeatTypeDayofthemonth JobMonthRepeatType = "dayOfTheMonth"
	JobMonthRepeatTypeDayoftheweek  JobMonthRepeatType = "dayOfTheWeek"
)

// JobUsages represents values for Job.usages
type JobUsages string

const (
	JobUsagesWorkflow JobUsages = "WORKFLOW"
)

// JobscheduleddataexportjobactionDataStorageOption represents values for JobScheduleddataexportjobaction.data_storage_option
type JobscheduleddataexportjobactionDataStorageOption string

const (
	JobscheduleddataexportjobactionDataStorageOptionVaultFileStaging JobscheduleddataexportjobactionDataStorageOption = "VAULT_FILE_STAGING"
	JobscheduleddataexportjobactionDataStorageOptionCustomS3Bucket   JobscheduleddataexportjobactionDataStorageOption = "CUSTOM_S3_BUCKET"
)

// JobmetadataSingleInstanceStates represents values for Jobmetadata.single_instance_states
type JobmetadataSingleInstanceStates string

const (
	JobmetadataSingleInstanceStatesScheduled JobmetadataSingleInstanceStates = "scheduled"
	JobmetadataSingleInstanceStatesQueued    JobmetadataSingleInstanceStates = "queued"
	JobmetadataSingleInstanceStatesRunning   JobmetadataSingleInstanceStates = "running"
)

// JobmetadataTimeoutDuration represents values for Jobmetadata.timeout_duration
type JobmetadataTimeoutDuration string

const (
	JobmetadataTimeoutDurationVal60   JobmetadataTimeoutDuration = "60"
	JobmetadataTimeoutDurationVal240  JobmetadataTimeoutDuration = "240"
	JobmetadataTimeoutDurationVal720  JobmetadataTimeoutDuration = "720"
	JobmetadataTimeoutDurationVal1380 JobmetadataTimeoutDuration = "1380"
)

// LifecyclestatetypeLifecycleScope represents values for Lifecyclestatetype.lifecycle_scope
type LifecyclestatetypeLifecycleScope string

const (
	LifecyclestatetypeLifecycleScopeAppliestoall   LifecyclestatetypeLifecycleScope = "AppliesToAll"
	LifecyclestatetypeLifecycleScopeAvailabletoall LifecyclestatetypeLifecycleScope = "AvailableToAll"
	LifecyclestatetypeLifecycleScopeSpecific       LifecyclestatetypeLifecycleScope = "Specific"
)

// LifecyclestatetypeLifecycleType represents values for Lifecyclestatetype.lifecycle_type
type LifecyclestatetypeLifecycleType string

const (
	LifecyclestatetypeLifecycleTypeDocument LifecyclestatetypeLifecycleType = "Document"
	LifecyclestatetypeLifecycleTypeObject   LifecyclestatetypeLifecycleType = "Object"
)

// LinkComponent represents values for Link.component
type LinkComponent string

const (
	LinkComponentDocument LinkComponent = "Document"
	LinkComponentObject   LinkComponent = "Object"
)

// LinkTarget represents values for Link.target
type LinkTarget string

const (
	LinkTargetSelf  LinkTarget = "self"
	LinkTargetBlank LinkTarget = "blank"
)

// MobileshareactionconfigAllowedFor represents values for Mobileshareactionconfig.allowed_for
type MobileshareactionconfigAllowedFor string

const (
	MobileshareactionconfigAllowedForAnyone         MobileshareactionconfigAllowedFor = "ANYONE"
	MobileshareactionconfigAllowedForSpecificGroups MobileshareactionconfigAllowedFor = "SPECIFIC_GROUPS"
)

// NotificationtemplateEntityType represents values for Notificationtemplate.entity_type
type NotificationtemplateEntityType string

const (
	NotificationtemplateEntityTypeDocument NotificationtemplateEntityType = "document"
	NotificationtemplateEntityTypeObject   NotificationtemplateEntityType = "object"
	NotificationtemplateEntityTypeCustom   NotificationtemplateEntityType = "custom"
)

// ObjectObjectClass represents values for Object.object_class
type ObjectObjectClass string

const (
	ObjectObjectClassBase                  ObjectObjectClass = "base"
	ObjectObjectClassComponent             ObjectObjectClass = "component"
	ObjectObjectClassUserrolesetup         ObjectObjectClass = "userrolesetup"
	ObjectObjectClassLegalhold             ObjectObjectClass = "legalhold"
	ObjectObjectClassUsertask              ObjectObjectClass = "usertask"
	ObjectObjectClassEsignature            ObjectObjectClass = "esignature"
	ObjectObjectClassChecklist             ObjectObjectClass = "checklist"
	ObjectObjectClassSection               ObjectObjectClass = "section"
	ObjectObjectClassResponse              ObjectObjectClass = "response"
	ObjectObjectClassResponserefdoc        ObjectObjectClass = "responserefdoc"
	ObjectObjectClassResponsedoc           ObjectObjectClass = "responsedoc"
	ObjectObjectClassAvailableanswer       ObjectObjectClass = "availableanswer"
	ObjectObjectClassDependency            ObjectObjectClass = "dependency"
	ObjectObjectClassSubchecklist          ObjectObjectClass = "subchecklist"
	ObjectObjectClassQualityteammember     ObjectObjectClass = "qualityteammember"
	ObjectObjectClassLifecyclestagesobject ObjectObjectClass = "lifecyclestagesobject"
	ObjectObjectClassPublicaccesslink      ObjectObjectClass = "publicaccesslink"
	ObjectObjectClassQualityoneteammember  ObjectObjectClass = "qualityoneteammember"
)

// ObjectDataStore represents values for Object.data_store
type ObjectDataStore string

const (
	ObjectDataStoreStandard ObjectDataStore = "standard"
	ObjectDataStoreRaw      ObjectDataStore = "raw"
)

// ObjectfieldType represents values for ObjectField.type
type ObjectfieldType string

const (
	ObjectfieldTypeString          ObjectfieldType = "String"
	ObjectfieldTypeNumber          ObjectfieldType = "Number"
	ObjectfieldTypeBoolean         ObjectfieldType = "Boolean"
	ObjectfieldTypeDate            ObjectfieldType = "Date"
	ObjectfieldTypeDatetime        ObjectfieldType = "DateTime"
	ObjectfieldTypeObject          ObjectfieldType = "Object"
	ObjectfieldTypePicklist        ObjectfieldType = "Picklist"
	ObjectfieldTypeId              ObjectfieldType = "ID"
	ObjectfieldTypeUsers           ObjectfieldType = "users"
	ObjectfieldTypeComponent       ObjectfieldType = "Component"
	ObjectfieldTypeDocuments       ObjectfieldType = "documents"
	ObjectfieldTypeLongtext        ObjectfieldType = "LongText"
	ObjectfieldTypeRichtext        ObjectfieldType = "RichText"
	ObjectfieldTypeBinary          ObjectfieldType = "Binary"
	ObjectfieldTypeObjectreference ObjectfieldType = "ObjectReference"
	ObjectfieldTypeObjectparent    ObjectfieldType = "ObjectParent"
)

// ObjectfieldRelationshipType represents values for ObjectField.relationship_type
type ObjectfieldRelationshipType string

const (
	ObjectfieldRelationshipTypeParent            ObjectfieldRelationshipType = "parent"
	ObjectfieldRelationshipTypeReference         ObjectfieldRelationshipType = "reference"
	ObjectfieldRelationshipTypeChild             ObjectfieldRelationshipType = "child"
	ObjectfieldRelationshipTypeReferenceInbound  ObjectfieldRelationshipType = "reference_inbound"
	ObjectfieldRelationshipTypeReferenceOutbound ObjectfieldRelationshipType = "reference_outbound"
)

// ObjectfieldRelationshipDeletion represents values for ObjectField.relationship_deletion
type ObjectfieldRelationshipDeletion string

const (
	ObjectfieldRelationshipDeletionBlock   ObjectfieldRelationshipDeletion = "block"
	ObjectfieldRelationshipDeletionCascade ObjectfieldRelationshipDeletion = "cascade"
	ObjectfieldRelationshipDeletionSetnull ObjectfieldRelationshipDeletion = "setnull"
)

// ObjectfieldSubtype represents values for ObjectField.subtype
type ObjectfieldSubtype string

const (
	ObjectfieldSubtypeCurrency ObjectfieldSubtype = "Currency"
	ObjectfieldSubtypeLink     ObjectfieldSubtype = "Link"
	ObjectfieldSubtypeTime     ObjectfieldSubtype = "Time"
	ObjectfieldSubtypePercent  ObjectfieldSubtype = "Percent"
	ObjectfieldSubtypePhone    ObjectfieldSubtype = "Phone"
	ObjectfieldSubtypeEmail    ObjectfieldSubtype = "Email"
)

// ObjectfieldRollupFunction represents values for ObjectField.rollup_function
type ObjectfieldRollupFunction string

const (
	ObjectfieldRollupFunctionCount ObjectfieldRollupFunction = "count"
	ObjectfieldRollupFunctionMin   ObjectfieldRollupFunction = "min"
	ObjectfieldRollupFunctionMax   ObjectfieldRollupFunction = "max"
	ObjectfieldRollupFunctionSum   ObjectfieldRollupFunction = "sum"
)

// ObjectlifecycleobjectlifecyclestateRecordStatus represents values for ObjectlifecycleObjectlifecyclestate.record_status
type ObjectlifecycleobjectlifecyclestateRecordStatus string

const (
	ObjectlifecycleobjectlifecyclestateRecordStatusActiveV      ObjectlifecycleobjectlifecyclestateRecordStatus = "active__v"
	ObjectlifecycleobjectlifecyclestateRecordStatusInactiveV    ObjectlifecycleobjectlifecyclestateRecordStatus = "inactive__v"
	ObjectlifecycleobjectlifecyclestateRecordStatusInMigrationV ObjectlifecycleobjectlifecyclestateRecordStatus = "in_migration__v"
	ObjectlifecycleobjectlifecyclestateRecordStatusArchivedV    ObjectlifecycleobjectlifecyclestateRecordStatus = "archived__v"
)

// ObjectlifecycleobjectlifecyclerolePermissions represents values for ObjectlifecycleObjectlifecyclerole.permissions
type ObjectlifecycleobjectlifecyclerolePermissions string

const (
	ObjectlifecycleobjectlifecyclerolePermissionsRead   ObjectlifecycleobjectlifecyclerolePermissions = "Read"
	ObjectlifecycleobjectlifecyclerolePermissionsEdit   ObjectlifecycleobjectlifecyclerolePermissions = "Edit"
	ObjectlifecycleobjectlifecyclerolePermissionsDelete ObjectlifecycleobjectlifecyclerolePermissions = "Delete"
)

// ObjectlifecycleobjectlifecyclepermissionPermission represents values for ObjectlifecycleObjectlifecyclepermission.permission
type ObjectlifecycleobjectlifecyclepermissionPermission string

const (
	ObjectlifecycleobjectlifecyclepermissionPermissionRead   ObjectlifecycleobjectlifecyclepermissionPermission = "Read"
	ObjectlifecycleobjectlifecyclepermissionPermissionEdit   ObjectlifecycleobjectlifecyclepermissionPermission = "Edit"
	ObjectlifecycleobjectlifecyclepermissionPermissionDelete ObjectlifecycleobjectlifecyclepermissionPermission = "Delete"
)

// ObjecttypetypefieldSource represents values for ObjecttypeTypefield.source
type ObjecttypetypefieldSource string

const (
	ObjecttypetypefieldSourceStandard    ObjecttypetypefieldSource = "standard"
	ObjecttypetypefieldSourceSample      ObjecttypetypefieldSource = "sample"
	ObjecttypetypefieldSourceCustom      ObjecttypetypefieldSource = "custom"
	ObjecttypetypefieldSourceApplication ObjecttypetypefieldSource = "application"
	ObjecttypetypefieldSourceSystem      ObjecttypetypefieldSource = "system"
)

// ObjectworkflowWorkflowContentType represents values for Objectworkflow.workflow_content_type
type ObjectworkflowWorkflowContentType string

const (
	ObjectworkflowWorkflowContentTypeDocument ObjectworkflowWorkflowContentType = "Document"
	ObjectworkflowWorkflowContentTypeObject   ObjectworkflowWorkflowContentType = "Object"
)

// ObjectworkflowClass represents values for Objectworkflow.class
type ObjectworkflowClass string

const (
	ObjectworkflowClassReadandunderstood ObjectworkflowClass = "ReadAndUnderstood"
)

// ObjectworkflowCardinality represents values for Objectworkflow.cardinality
type ObjectworkflowCardinality string

const (
	ObjectworkflowCardinalityOne       ObjectworkflowCardinality = "One"
	ObjectworkflowCardinalityOneormany ObjectworkflowCardinality = "OneOrMany"
)

// ObjectworkflowobjectworkflowstepType represents values for ObjectworkflowObjectworkflowstep.type
type ObjectworkflowobjectworkflowstepType string

const (
	ObjectworkflowobjectworkflowstepTypeStart                         ObjectworkflowobjectworkflowstepType = "start"
	ObjectworkflowobjectworkflowstepTypeEnd                           ObjectworkflowobjectworkflowstepType = "end"
	ObjectworkflowobjectworkflowstepTypeUsertask                      ObjectworkflowobjectworkflowstepType = "usertask"
	ObjectworkflowobjectworkflowstepTypeContenttask                   ObjectworkflowobjectworkflowstepType = "contenttask"
	ObjectworkflowobjectworkflowstepTypeNotification                  ObjectworkflowobjectworkflowstepType = "notification"
	ObjectworkflowobjectworkflowstepTypePlaceholder                   ObjectworkflowobjectworkflowstepType = "placeholder"
	ObjectworkflowobjectworkflowstepTypeChangestate                   ObjectworkflowobjectworkflowstepType = "changestate"
	ObjectworkflowobjectworkflowstepTypeDecision                      ObjectworkflowobjectworkflowstepType = "decision"
	ObjectworkflowobjectworkflowstepTypeJoin                          ObjectworkflowobjectworkflowstepType = "join"
	ObjectworkflowobjectworkflowstepTypeAction                        ObjectworkflowobjectworkflowstepType = "action"
	ObjectworkflowobjectworkflowstepTypeUpdaterecordfield             ObjectworkflowobjectworkflowstepType = "updaterecordfield"
	ObjectworkflowobjectworkflowstepTypeUpdatesharingsettings         ObjectworkflowobjectworkflowstepType = "updatesharingsettings"
	ObjectworkflowobjectworkflowstepTypeUpdatedocumentsharingsettings ObjectworkflowobjectworkflowstepType = "updatedocumentsharingsettings"
	ObjectworkflowobjectworkflowstepTypeUpdaterelatedrecordfield      ObjectworkflowobjectworkflowstepType = "updaterelatedrecordfield"
	ObjectworkflowobjectworkflowstepTypeContentaction                 ObjectworkflowobjectworkflowstepType = "contentaction"
)

// OutboundemaildomainVerificationStatus represents values for Outboundemaildomain.verification_status
type OutboundemaildomainVerificationStatus string

const (
	OutboundemaildomainVerificationStatusUnverified OutboundemaildomainVerificationStatus = "UNVERIFIED"
	OutboundemaildomainVerificationStatusVerified   OutboundemaildomainVerificationStatus = "VERIFIED"
	OutboundemaildomainVerificationStatusError      OutboundemaildomainVerificationStatus = "ERROR"
)

// OverlaytemplateoverlayoverrideHeaderFooterPlacement represents values for OverlaytemplateOverlayoverride.header_footer_placement
type OverlaytemplateoverlayoverrideHeaderFooterPlacement string

const (
	OverlaytemplateoverlayoverrideHeaderFooterPlacementHeaderFooterV    OverlaytemplateoverlayoverrideHeaderFooterPlacement = "header_footer__v"
	OverlaytemplateoverlayoverrideHeaderFooterPlacementLeftRightMarginV OverlaytemplateoverlayoverrideHeaderFooterPlacement = "left_right_margin__v"
)

// OverlaytemplateoverlayoverrideOrientation represents values for OverlaytemplateOverlayoverride.orientation
type OverlaytemplateoverlayoverrideOrientation string

const (
	OverlaytemplateoverlayoverrideOrientationPortraitV  OverlaytemplateoverlayoverrideOrientation = "portrait__v"
	OverlaytemplateoverlayoverrideOrientationLandscapeV OverlaytemplateoverlayoverrideOrientation = "landscape__v"
)

// OverlaytemplateoverlayoverrideOverlayFormatType represents values for OverlaytemplateOverlayoverride.overlay_format_type
type OverlaytemplateoverlayoverrideOverlayFormatType string

const (
	OverlaytemplateoverlayoverrideOverlayFormatTypeXfaOverrideV        OverlaytemplateoverlayoverrideOverlayFormatType = "xfa_override__v"
	OverlaytemplateoverlayoverrideOverlayFormatTypeBasicOverlayFormatV OverlaytemplateoverlayoverrideOverlayFormatType = "basic_overlay_format__v"
)

// OverlaytemplateoverlayoverridePaperSize represents values for OverlaytemplateOverlayoverride.paper_size
type OverlaytemplateoverlayoverridePaperSize string

const (
	OverlaytemplateoverlayoverridePaperSizeLetterV  OverlaytemplateoverlayoverridePaperSize = "letter__v"
	OverlaytemplateoverlayoverridePaperSizeA4V      OverlaytemplateoverlayoverridePaperSize = "a4__v"
	OverlaytemplateoverlayoverridePaperSizeA3V      OverlaytemplateoverlayoverridePaperSize = "a3__v"
	OverlaytemplateoverlayoverridePaperSizeLegalV   OverlaytemplateoverlayoverridePaperSize = "legal__v"
	OverlaytemplateoverlayoverridePaperSizeTabloidV OverlaytemplateoverlayoverridePaperSize = "tabloid__v"
)

// PagelinkMode represents values for Pagelink.mode
type PagelinkMode string

const (
	PagelinkModeCreate PagelinkMode = "Create"
	PagelinkModeEdit   PagelinkMode = "Edit"
	PagelinkModeCopy   PagelinkMode = "Copy"
	PagelinkModeView   PagelinkMode = "View"
)

// PagelinkPageType represents values for Pagelink.page_type
type PagelinkPageType string

const (
	PagelinkPageTypeObject PagelinkPageType = "Object"
)

// PicklistOrderType represents values for Picklist.order_type
type PicklistOrderType string

const (
	PicklistOrderTypeOrderAscSys       PicklistOrderType = "order_asc__sys"
	PicklistOrderTypeValueLabelAscSys  PicklistOrderType = "value_label_asc__sys"
	PicklistOrderTypeValueLabelDescSys PicklistOrderType = "value_label_desc__sys"
)

// QualitycurriculumsmartmatchruleMatchingType represents values for Qualitycurriculumsmartmatchrule.matching_type
type QualitycurriculumsmartmatchruleMatchingType string

const (
	QualitycurriculumsmartmatchruleMatchingTypeExact QualitycurriculumsmartmatchruleMatchingType = "EXACT"
)

// QualitydynamicenrollmentruleMatchingType represents values for Qualitydynamicenrollmentrule.matching_type
type QualitydynamicenrollmentruleMatchingType string

const (
	QualitydynamicenrollmentruleMatchingTypeExact QualitydynamicenrollmentruleMatchingType = "EXACT"
)

// QualityexternalnotificationtemplateDownloadOption represents values for Qualityexternalnotificationtemplate.download_option
type QualityexternalnotificationtemplateDownloadOption string

const (
	QualityexternalnotificationtemplateDownloadOptionSource QualityexternalnotificationtemplateDownloadOption = "SOURCE"
	QualityexternalnotificationtemplateDownloadOptionPdf    QualityexternalnotificationtemplateDownloadOption = "PDF"
	QualityexternalnotificationtemplateDownloadOptionBoth   QualityexternalnotificationtemplateDownloadOption = "BOTH"
	QualityexternalnotificationtemplateDownloadOptionNone   QualityexternalnotificationtemplateDownloadOption = "NONE"
)

// QualityteamqualityteamroleCascadeBehavior represents values for QualityteamQualityteamrole.cascade_behavior
type QualityteamqualityteamroleCascadeBehavior string

const (
	QualityteamqualityteamroleCascadeBehaviorNotApplicable            QualityteamqualityteamroleCascadeBehavior = "NOT_APPLICABLE"
	QualityteamqualityteamroleCascadeBehaviorInheritAllowOverride     QualityteamqualityteamroleCascadeBehavior = "INHERIT_ALLOW_OVERRIDE"
	QualityteamqualityteamroleCascadeBehaviorMultiSourceSingleInherit QualityteamqualityteamroleCascadeBehavior = "MULTI_SOURCE_SINGLE_INHERIT"
)

// QueueType represents values for Queue.type
type QueueType string

const (
	QueueTypeInbound  QueueType = "inbound"
	QueueTypeOutbound QueueType = "outbound"
	QueueTypeJob      QueueType = "job"
)

// RecordroletriggerEventSegment represents values for Recordroletrigger.event_segment
type RecordroletriggerEventSegment string

const (
	RecordroletriggerEventSegmentPreCustom   RecordroletriggerEventSegment = "PRE_CUSTOM"
	RecordroletriggerEventSegmentPostCustom  RecordroletriggerEventSegment = "POST_CUSTOM"
	RecordroletriggerEventSegmentUnspecified RecordroletriggerEventSegment = "UNSPECIFIED"
)

// RecordroletriggerEvents represents values for Recordroletrigger.events
type RecordroletriggerEvents string

const (
	RecordroletriggerEventsBefore RecordroletriggerEvents = "BEFORE"
	RecordroletriggerEventsAfter  RecordroletriggerEvents = "AFTER"
)

// RecordroletriggerOrder represents values for Recordroletrigger.order
type RecordroletriggerOrder string

const (
	RecordroletriggerOrderNumber1     RecordroletriggerOrder = "NUMBER_1"
	RecordroletriggerOrderNumber2     RecordroletriggerOrder = "NUMBER_2"
	RecordroletriggerOrderNumber3     RecordroletriggerOrder = "NUMBER_3"
	RecordroletriggerOrderNumber4     RecordroletriggerOrder = "NUMBER_4"
	RecordroletriggerOrderNumber5     RecordroletriggerOrder = "NUMBER_5"
	RecordroletriggerOrderNumber6     RecordroletriggerOrder = "NUMBER_6"
	RecordroletriggerOrderNumber7     RecordroletriggerOrder = "NUMBER_7"
	RecordroletriggerOrderNumber8     RecordroletriggerOrder = "NUMBER_8"
	RecordroletriggerOrderNumber9     RecordroletriggerOrder = "NUMBER_9"
	RecordroletriggerOrderNumber10    RecordroletriggerOrder = "NUMBER_10"
	RecordroletriggerOrderUnspecified RecordroletriggerOrder = "UNSPECIFIED"
)

// RecordtriggerEventSegment represents values for Recordtrigger.event_segment
type RecordtriggerEventSegment string

const (
	RecordtriggerEventSegmentPreCustom   RecordtriggerEventSegment = "PRE_CUSTOM"
	RecordtriggerEventSegmentPostCustom  RecordtriggerEventSegment = "POST_CUSTOM"
	RecordtriggerEventSegmentUnspecified RecordtriggerEventSegment = "UNSPECIFIED"
)

// RecordtriggerEvents represents values for Recordtrigger.events
type RecordtriggerEvents string

const (
	RecordtriggerEventsBeforeInsert RecordtriggerEvents = "BEFORE_INSERT"
	RecordtriggerEventsAfterInsert  RecordtriggerEvents = "AFTER_INSERT"
	RecordtriggerEventsBeforeUpdate RecordtriggerEvents = "BEFORE_UPDATE"
	RecordtriggerEventsAfterUpdate  RecordtriggerEvents = "AFTER_UPDATE"
	RecordtriggerEventsBeforeDelete RecordtriggerEvents = "BEFORE_DELETE"
	RecordtriggerEventsAfterDelete  RecordtriggerEvents = "AFTER_DELETE"
)

// RecordtriggerOrder represents values for Recordtrigger.order
type RecordtriggerOrder string

const (
	RecordtriggerOrderNumber1     RecordtriggerOrder = "NUMBER_1"
	RecordtriggerOrderNumber2     RecordtriggerOrder = "NUMBER_2"
	RecordtriggerOrderNumber3     RecordtriggerOrder = "NUMBER_3"
	RecordtriggerOrderNumber4     RecordtriggerOrder = "NUMBER_4"
	RecordtriggerOrderNumber5     RecordtriggerOrder = "NUMBER_5"
	RecordtriggerOrderNumber6     RecordtriggerOrder = "NUMBER_6"
	RecordtriggerOrderNumber7     RecordtriggerOrder = "NUMBER_7"
	RecordtriggerOrderNumber8     RecordtriggerOrder = "NUMBER_8"
	RecordtriggerOrderNumber9     RecordtriggerOrder = "NUMBER_9"
	RecordtriggerOrderNumber10    RecordtriggerOrder = "NUMBER_10"
	RecordtriggerOrderUnspecified RecordtriggerOrder = "UNSPECIFIED"
)

// RecordworkflowactionStepTypes represents values for Recordworkflowaction.step_types
type RecordworkflowactionStepTypes string

const (
	RecordworkflowactionStepTypesStart RecordworkflowactionStepTypes = "START"
	RecordworkflowactionStepTypesTask  RecordworkflowactionStepTypes = "TASK"
)

// RenditionprofileBookmarkExpansionLevel represents values for Renditionprofile.bookmark_expansion_level
type RenditionprofileBookmarkExpansionLevel string

const (
	RenditionprofileBookmarkExpansionLevelAllLevels  RenditionprofileBookmarkExpansionLevel = "ALL_LEVELS"
	RenditionprofileBookmarkExpansionLevelLevelOne   RenditionprofileBookmarkExpansionLevel = "LEVEL_ONE"
	RenditionprofileBookmarkExpansionLevelLevelTwo   RenditionprofileBookmarkExpansionLevel = "LEVEL_TWO"
	RenditionprofileBookmarkExpansionLevelLevelThree RenditionprofileBookmarkExpansionLevel = "LEVEL_THREE"
	RenditionprofileBookmarkExpansionLevelLevelFour  RenditionprofileBookmarkExpansionLevel = "LEVEL_FOUR"
	RenditionprofileBookmarkExpansionLevelLevelFive  RenditionprofileBookmarkExpansionLevel = "LEVEL_FIVE"
)

// RenditionprofileEmbedFonts represents values for Renditionprofile.embed_fonts
type RenditionprofileEmbedFonts string

const (
	RenditionprofileEmbedFontsDefault  RenditionprofileEmbedFonts = "DEFAULT"
	RenditionprofileEmbedFontsFullFont RenditionprofileEmbedFonts = "FULL_FONT"
	RenditionprofileEmbedFontsSubFont  RenditionprofileEmbedFonts = "SUB_FONT"
)

// RenditionprofilePdfFormat represents values for Renditionprofile.pdf_format
type RenditionprofilePdfFormat string

const (
	RenditionprofilePdfFormatDefault RenditionprofilePdfFormat = "DEFAULT"
	RenditionprofilePdfFormatPdfa1B  RenditionprofilePdfFormat = "PDFA_1B"
)

// ReportFormat represents values for Report.format
type ReportFormat string

const (
	ReportFormatTabular ReportFormat = "tabular"
	ReportFormatMatrix  ReportFormat = "matrix"
)

// ReportContentType represents values for Report.content_type
type ReportContentType string

const (
	ReportContentTypeDocument          ReportContentType = "Document"
	ReportContentTypeWorkflow          ReportContentType = "Workflow"
	ReportContentTypeReadandunderstand ReportContentType = "ReadAndUnderstand"
	ReportContentTypeDistribution      ReportContentType = "Distribution"
	ReportContentTypeBindernode        ReportContentType = "Bindernode"
	ReportContentTypeObject            ReportContentType = "Object"
	ReportContentTypeRelationship      ReportContentType = "Relationship"
	ReportContentTypeMatcheddocument   ReportContentType = "MatchedDocument"
	ReportContentTypeBinder            ReportContentType = "Binder"
	ReportContentTypeBindersection     ReportContentType = "BinderSection"
	ReportContentTypeMultipass         ReportContentType = "MultiPass"
	ReportContentTypeUnion             ReportContentType = "Union"
)

// ReportClass represents values for Report.class
type ReportClass string

const (
	ReportClassStandard ReportClass = "Standard"
	ReportClassView     ReportClass = "View"
)

// ReporttypeClass represents values for Reporttype.class
type ReporttypeClass string

const (
	ReporttypeClassStandard  ReporttypeClass = "Standard"
	ReporttypeClassMultipass ReporttypeClass = "MultiPass"
	ReporttypeClassUnionall  ReporttypeClass = "UnionAll"
)

// SavedviewViewLayoutType represents values for Savedview.view_layout_type
type SavedviewViewLayoutType string

const (
	SavedviewViewLayoutTypeDetail    SavedviewViewLayoutType = "DETAIL"
	SavedviewViewLayoutTypeTile      SavedviewViewLayoutType = "TILE"
	SavedviewViewLayoutTypeCompact   SavedviewViewLayoutType = "COMPACT"
	SavedviewViewLayoutTypeReporting SavedviewViewLayoutType = "REPORTING"
)

// SignaturepageLocation represents values for Signaturepage.location
type SignaturepageLocation string

const (
	SignaturepageLocationStartV SignaturepageLocation = "start__v"
	SignaturepageLocationEndV   SignaturepageLocation = "end__v"
)

// TabcollectiontabcollectionitemType represents values for TabcollectionTabcollectionitem.type
type TabcollectiontabcollectionitemType string

const (
	TabcollectiontabcollectionitemTypeTab  TabcollectiontabcollectionitemType = "Tab"
	TabcollectiontabcollectionitemTypeMenu TabcollectiontabcollectionitemType = "Menu"
)

// UserdefinedmodelSerializationInclude represents values for Userdefinedmodel.serialization_include
type UserdefinedmodelSerializationInclude string

const (
	UserdefinedmodelSerializationIncludeAlways      UserdefinedmodelSerializationInclude = "ALWAYS"
	UserdefinedmodelSerializationIncludeIgnore      UserdefinedmodelSerializationInclude = "IGNORE"
	UserdefinedmodelSerializationIncludeNonNull     UserdefinedmodelSerializationInclude = "NON_NULL"
	UserdefinedmodelSerializationIncludeUnspecified UserdefinedmodelSerializationInclude = "UNSPECIFIED"
)

// VaulttokenCloneBehavior represents values for Vaulttoken.clone_behavior
type VaulttokenCloneBehavior string

const (
	VaulttokenCloneBehaviorClearSys   VaulttokenCloneBehavior = "clear__sys"
	VaulttokenCloneBehaviorPersistSys VaulttokenCloneBehavior = "persist__sys"
)

// VaulttokenType represents values for Vaulttoken.type
type VaulttokenType string

const (
	VaulttokenTypeStringSys     VaulttokenType = "string__sys"
	VaulttokenTypeBooleanSys    VaulttokenType = "boolean__sys"
	VaulttokenTypeNumberSys     VaulttokenType = "number__sys"
	VaulttokenTypeDateSys       VaulttokenType = "date__sys"
	VaulttokenTypeDatetimeSys   VaulttokenType = "datetime__sys"
	VaulttokenTypeListStringSys VaulttokenType = "list_string__sys"
)

// WorkflowworkflowstepFlowType represents values for WorkflowWorkflowstep.flow_type
type WorkflowworkflowstepFlowType string

const (
	WorkflowworkflowstepFlowTypeStart                 WorkflowworkflowstepFlowType = "start"
	WorkflowworkflowstepFlowTypeEnd                   WorkflowworkflowstepFlowType = "end"
	WorkflowworkflowstepFlowTypeUsertask              WorkflowworkflowstepFlowType = "usertask"
	WorkflowworkflowstepFlowTypeNotification          WorkflowworkflowstepFlowType = "notification"
	WorkflowworkflowstepFlowTypePlaceholder           WorkflowworkflowstepFlowType = "placeholder"
	WorkflowworkflowstepFlowTypeChangestate           WorkflowworkflowstepFlowType = "changestate"
	WorkflowworkflowstepFlowTypeDecision              WorkflowworkflowstepFlowType = "decision"
	WorkflowworkflowstepFlowTypeJoin                  WorkflowworkflowstepFlowType = "join"
	WorkflowworkflowstepFlowTypeAction                WorkflowworkflowstepFlowType = "action"
	WorkflowworkflowstepFlowTypeUpdaterecordfield     WorkflowworkflowstepFlowType = "updaterecordfield"
	WorkflowworkflowstepFlowTypeUpdatesharingsettings WorkflowworkflowstepFlowType = "updatesharingsettings"
)
