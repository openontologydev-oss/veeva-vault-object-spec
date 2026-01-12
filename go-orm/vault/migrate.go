package vault

import (
	"gorm.io/gorm"
)

// AllModels returns all Vault ORM models for migration
// Use this with db.AutoMigrate(AllModels()...)
func AllModels() []interface{} {
	return []interface{}{
		// Core Objects
		&VaultObject{},
		&ObjectField{},
		&ObjectIndex{},
		&ObjectRelationship{},

		// Object Types
		&ObjectType{},
		&TypeField{},
		&TypeAction{},

		// Object Lifecycles
		&ObjectLifecycle{},
		&ObjectLifecycleState{},
		&ObjectLifecycleStateEntryCriterion{},
		&ObjectLifecycleStateEntryAction{},
		&ObjectLifecycleStateUserAction{},
		&ObjectLifecycleEvent{},
		&ObjectLifecycleRole{},
		&ObjectLifecyclePermission{},

		// Documents
		&Doctype{},
		&Docfield{},
		&Docfieldlayout{},
		&Doclifecycle{},
		&DoclifecycleState{},
		&DoclifecycleRole{},
		&DoclifecycleEvent{},
		&Docrelationshiptype{},
		&Documentaction{},

		// Workflows
		&Workflow{},
		&WorkflowStep{},
		&ObjectWorkflow{},
		&ObjectWorkflowStep{},
		&WorkflowCancelationAction{},

		// Triggers & Actions
		&Recordtrigger{},
		&Actiontrigger{},
		&Actionblock{},
		&Draftactionblock{},
		&Recordaction{},
		&Objectaction{},
		&Recordroletrigger{},
		&Recordworkflowaction{},

		// Picklists
		&Picklist{},
		&PicklistEntry{},

		// Security
		&Securityprofile{},
		&Permissionset{},
		&Sharingrule{},
		&Sharingrole{},
		&Atomicsecurity{},
		&FieldSecurity{},
		&ObjectControlSecurity{},
		&ActionSecurity{},
		&WorkflowActionSecurity{},
		&RelationshipSecurity{},
		&Appsecurityrule{},
		&Appsecurityfield{},
		&Appsecurityassignment{},

		// SDK Components
		&Customwebapi{},
		&Webapigroup{},
		&Userdefinedclass{},
		&Userdefinedmodel{},
		&Userdefinedservice{},
		&Sdkjob{},

		// Jobs & Queues
		&Job{},
		&SendNotificationJobAction{},
		&ChangeObjectStateJobAction{},
		&SdkBatchJobAction{},
		&Queue{},
		&QueueConnection{},

		// Messages
		&Messagegroup{},
		&Message{},
		&Notificationtemplate{},
		&Accountmessage{},

		// UI Components
		&Tab{},
		&Subtab{},
		&Tabcollection{},
		&TabcollectionItem{},
		&Dashboard{},
		&Report{},
		&Reporttype{},
		&Page{},
		&Pagelayout{},
		&Pagelink{},
		&Link{},
		&Savedview{},
		&Searchcollection{},
		&SearchcollectionSection{},
		&Tag{},
		&Layoutprofile{},
		&Objectlayout{},
		&Labelset{},
		&Renditionprofile{},
		&Renditiontype{},
	}
}

// Migrate runs auto-migration for all Vault models
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(AllModels()...)
}

// CreateIndexes creates additional indexes for performance
func CreateIndexes(db *gorm.DB) error {
	// Object fields compound index
	if err := db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_object_fields_object_name 
		ON vault_object_fields (object_id, name)
	`).Error; err != nil {
		return err
	}

	// Lifecycle states compound index
	if err := db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_lifecycle_states_lifecycle_name 
		ON vault_object_lifecycle_states (lifecycle_id, name)
	`).Error; err != nil {
		return err
	}

	// Picklist entries compound index
	if err := db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_picklist_entries_picklist_name 
		ON vault_picklist_entries (picklist_id, name)
	`).Error; err != nil {
		return err
	}

	return nil
}

// DropAllTables drops all Vault tables (use with caution!)
func DropAllTables(db *gorm.DB) error {
	models := AllModels()
	for i := len(models) - 1; i >= 0; i-- {
		if err := db.Migrator().DropTable(models[i]); err != nil {
			return err
		}
	}
	return nil
}

// =============================================================================
// EXAMPLE USAGE
// =============================================================================

// ExampleProduct returns an example VaultObject representing a Product
func ExampleProduct() *VaultObject {
	description := "Custom product object for tracking products"
	return &VaultObject{
		ComponentBase: ComponentBase{
			Name:   "product__c",
			Label:  "Product",
			Active: true,
		},
		LabelPlural:      "Products",
		Description:      &description,
		Source:           SourceCustom,
		ObjectClass:      ObjectClassBase,
		DataStore:        DataStoreStandard,
		InMenu:           true,
		AllowAttachments: true,
		Audit:            true,
		Prefix:           "PRD",
		Fields: []ObjectField{
			{
				Name:     "name__v",
				Label:    "Name",
				Type:     FieldTypeString,
				Active:   true,
				Required: true,
				MaxLength: func() *int { v := 128; return &v }(),
			},
			{
				Name:     "product_code__c",
				Label:    "Product Code",
				Type:     FieldTypeString,
				Active:   true,
				Required: true,
				Unique:   true,
				MaxLength: func() *int { v := 50; return &v }(),
			},
			{
				Name:              "product_family__c",
				Label:             "Product Family",
				Type:              FieldTypeObject,
				Active:            true,
				ReferencedObject:  func() *string { v := "product_family__c"; return &v }(),
				RelationshipType:  func() *RelationshipType { v := RelationshipTypeReference; return &v }(),
				RelationshipDeletion: func() *RelationshipDeletion { v := RelationshipDeletionBlock; return &v }(),
			},
			{
				Name:        "status__c",
				Label:       "Status",
				Type:        FieldTypePicklist,
				Active:      true,
				Required:    true,
				PicklistName: func() *string { v := "product_status__c"; return &v }(),
			},
			{
				Name:   "launch_date__c",
				Label:  "Launch Date",
				Type:   FieldTypeDate,
				Active: true,
			},
			{
				Name:   "price__c",
				Label:  "Price",
				Type:   FieldTypeNumber,
				Active: true,
				Scale:  func() *int { v := 2; return &v }(),
			},
			{
				Name:       "is_active__c",
				Label:      "Is Active",
				Type:       FieldTypeBoolean,
				Active:     true,
				ListColumn: true,
			},
		},
	}
}

// ExamplePicklist returns an example Picklist
func ExamplePicklist() *Picklist {
	return &Picklist{
		ComponentBase: ComponentBase{
			Name:   "product_status__c",
			Label:  "Product Status",
			Active: true,
		},
		Entries: []PicklistEntry{
			{
				ComponentBase: ComponentBase{
					Name:   "draft__c",
					Label:  "Draft",
					Active: true,
				},
				Order: 1,
			},
			{
				ComponentBase: ComponentBase{
					Name:   "active__c",
					Label:  "Active",
					Active: true,
				},
				Order: 2,
			},
			{
				ComponentBase: ComponentBase{
					Name:   "discontinued__c",
					Label:  "Discontinued",
					Active: true,
				},
				Order: 3,
			},
		},
	}
}

// ExampleLifecycle returns an example ObjectLifecycle
func ExampleLifecycle() *ObjectLifecycle {
	return &ObjectLifecycle{
		ComponentBase: ComponentBase{
			Name:   "product_lifecycle__c",
			Label:  "Product Lifecycle",
			Active: true,
		},
		States: []ObjectLifecycleState{
			{
				ComponentBase: ComponentBase{
					Name:   "draft__c",
					Label:  "Draft",
					Active: true,
				},
				RecordStatus: RecordStatusActive,
			},
			{
				ComponentBase: ComponentBase{
					Name:   "in_review__c",
					Label:  "In Review",
					Active: true,
				},
				RecordStatus: RecordStatusActive,
			},
			{
				ComponentBase: ComponentBase{
					Name:   "approved__c",
					Label:  "Approved",
					Active: true,
				},
				RecordStatus: RecordStatusActive,
			},
			{
				ComponentBase: ComponentBase{
					Name:   "retired__c",
					Label:  "Retired",
					Active: true,
				},
				RecordStatus: RecordStatusInactive,
			},
		},
		Roles: []ObjectLifecycleRole{
			{
				Name:            "owner__v",
				Active:          true,
				ApplicationRole: "owner__v",
				Permissions:     StringSlice{"Read", "Edit", "Delete"},
			},
			{
				Name:            "editor__v",
				Active:          true,
				ApplicationRole: "editor__v",
				Permissions:     StringSlice{"Read", "Edit"},
			},
			{
				Name:            "viewer__v",
				Active:          true,
				ApplicationRole: "viewer__v",
				Permissions:     StringSlice{"Read"},
			},
		},
	}
}
