// Generated TypeScript definitions for Veeva Vault MDL Components
// Source: https://developer.veevavault.com/mdl/components
// Generated date: 2026-01-12

/** View and edit messages that Vault sends to users as both emails and in-application notifications. */
export interface Accountmessage {
  label?: string;
  active?: boolean;
  default_lang?: string;
  email_body: string;
  subject: string;
}

/** An Action Trigger allows you to create and execute custom business logic when a create, update, or delete record operation occurs on an object record. */
export interface Actiontrigger {
  label?: string;
  active?: boolean;
  description?: string;
  object: string;
  event: string;
  order?: number;
}

export interface Actiontrigger_Actionblock {
  label?: string;
  active?: boolean;
  description?: string;
  order: number;
  code?: string;
}

export interface Actiontrigger_Draftactionblock {
  label?: string;
  active?: boolean;
  description?: string;
  order: number;
  code?: string;
}

/** Internal Vault developers use Application Security Rules to control access to records. */
export interface Appsecurityrule {
  label?: string;
  active?: boolean;
  help_content?: string;
  allow_custom_objects: boolean;
  description?: string;
}

export interface Appsecurityrule_Appsecurityfield {
  label?: string;
  active?: boolean;
  multi_value?: boolean;
  picklist?: string;
  type: string;
  object?: string;
}

export interface Appsecurityrule_Appsecurityassignment {
  label?: string;
  active?: boolean;
  field: Appsecurityrule_Appsecurityassignment_Field[];
  securityrule?: string;
  securityfield?: string;
  object: string;
}

/** Atomic security provides an additional layer of control by allowing you to configure security at the lifecycle state and/or assigned role level. */
export interface Atomicsecurity {
  label: string;
  active?: boolean;
  object: string;
  object_lifecycle: string;
  state: string;
}

export interface Atomicsecurity_Fieldsecurity {
  label: string;
  active?: boolean;
  role: string;
  type: string;
  fields: string;
}

export interface Atomicsecurity_Objectcontrolsecurity {
  label: string;
  active?: boolean;
  role: string;
  type: string;
  object_controls: string;
}

export interface Atomicsecurity_Actionsecurity {
  label: string;
  active?: boolean;
  role: string;
  type: string;
  object_actions?: string;
  lifecycle_actions?: string;
}

export interface Atomicsecurity_Workflowactionsecurity {
  label: string;
  active?: boolean;
  role: string;
  type: string;
  workflow_actions?: string;
  workflow_task_actions?: string;
}

export interface Atomicsecurity_Relationshipsecurity {
  label: string;
  active?: boolean;
  role: string;
  type: string;
  document_fields?: string;
  object_fields?: string;
}

/** Configuration details of a Case Child section in the Medical Inquiry user interface, available only in MedInquiry Vaults. */
export interface Casechildconfig {
  label?: string;
  active?: boolean;
  object_key: string;
  field_keys: string;
  parent_field_key: string;
  save_all_on_create?: boolean;
  auto_suggest_faqs?: boolean;
  request_details_field_key?: string;
  grandparent_field_key?: string;
  children?: string;
  config_key: string;
  order?: number;
}

/** Checklists are a data entry method that allows a user to enter data in question and answer format. */
export interface Checklisttype {
  label?: string;
  active?: boolean;
  target_object: string;
  target_object_type: string;
  version_controlled?: boolean;
}

export interface Checklisttype_Checklistmatchingfield {
  label?: string;
  active?: boolean;
  target_matching_field?: Checklisttype_Checklistmatchingfield_Target_matching_field[];
  target_object: string;
  checklist_design_matching_field: Checklisttype_Checklistmatchingfield_Checklist_design_matching_field[];
  checklist_design_object: string;
}

/** Represents a Client Code Distribution containing Custom Pages code uploaded via Vault API. This component type is read-only. */
export interface Clientdistribution {
  label?: string;
  active?: boolean;
  checksum: string;
}

export interface Clientdistribution_Pageclientcode {
  label?: string;
  active?: boolean;
  export?: string;
  file: string;
}

/** Specific to Clinical Operations Vaults, supports CTMS Transfer feature. Maps custom fields and values to standard fields for Study data transfer between CRO and Sponsor Vaults. */
export interface Clinicalstandardmapping {
  label?: string;
  active?: boolean;
  custom_component: string;
  custom_subcomponent?: string;
  standard_component?: string;
  standard_subcomponent?: string;
}

/** Developers can create custom web APIs using the Vault Java SDK. */
export interface Customwebapi {
  label?: string;
  active?: boolean;
  api_group: string;
  checksum?: string;
  endpoint_name: string;
  minimum_version: string;
  source_code?: string;
}

/** Dashboards provide an at-a-glance understanding of key metrics. */
export interface Dashboard {
  active?: boolean;
  label: string;
  description?: string;
  editors?: string;
  viewers?: string;
  owners?: string;
  dashboard_markup: string;
  tags?: string;
}

/** Specific to Clinical Operations Vaults, supports Veeva Disclosures feature. Controls user actions that trigger Vault to create Disclosure records for supported registries. */
export interface Disclosurerule {
  label?: string;
  active?: boolean;
  assessment_value?: string;
  authority: string;
  compliance_date_shift: number;
  compliance_date_shift_units: string;
  country: string;
  disclosure_subtype: string;
  disclosure_type: string;
  due_date_shift: number;
  due_date_shift_units: string;
  milestone_level?: string;
  milestone_type?: string;
  object_key: string;
  rules_version: string;
  trigger_or_sys_action: string;
  workflow_initiated?: string;
}

/** Specific to Clinical Operations Vaults, supports Veeva Disclosures feature. Maps document types for inclusion when posting Related Documents to ClinicalTrials.gov or Custom XML. */
export interface Disclosurexmldoctypemapping {
  label?: string;
  active?: boolean;
  clinical_trials_document_type?: string;
  custom_xml_document_type?: string;
  document_type: string;
}

/** Specific to Clinical Operations Vaults, supports Veeva Disclosures feature. Maps custom Disclosure object fields for inclusion in Disclosure XML files. */
export interface Disclosurexmlfieldmapping {
  label?: string;
  active?: boolean;
  field: Disclosurexmlfieldmapping_Field[];
  mapping_label: string;
  object: string;
}

/** Atomic security for documents provides granular control over who can perform specific lifecycle user actions based on lifecycle state and role. */
export interface Docatomicsecurity {
  label?: string;
  active?: boolean;
  document_lifecycle: string;
  state: string;
}

export interface Docatomicsecurity_Docworkflowactionsecurity {
  label?: string;
  active?: boolean;
  role: string;
  type: string;
  workflow_actions?: string;
  workflow_task_actions?: string;
}

export interface Docatomicsecurity_Docactionsecurity {
  label?: string;
  active?: boolean;
  role: string;
  type: string;
  lifecycle_actions?: string;
}

/** Attribute of a document associated to one or more document types. */
export interface Docfield {
  label: string;
  active: boolean;
  required: boolean;
  shared: boolean;
  type: string;
  picklist?: string;
  multi_value?: boolean;
  display_section: string;
  scope?: string;
  max_length?: number;
  min_value?: number;
  max_value?: number;
  scale?: number;
  formula?: string;
  editable?: boolean;
  hidden?: boolean;
  no_copy?: boolean;
  object?: string;
  default_value?: string;
  help_content?: string;
  queryable?: boolean;
}

/** Dependent fields configuration that defines relationships between a controlling field and one or more dependent fields. */
export interface Docfielddependency {
  doctype?: string;
  docfield?: string;
  docfield_multi_value?: string;
  docfield_value?: string;
  deprule_docfield: string;
  deprule_docfield_action: string;
  deprule_docfield_multi_values?: string;
  deprule_docfield_values?: string;
  lifecycle?: string;
  state?: string;
  role?: string;
}

/** Defines how a field layout section appears on the Doc Info page in the Vault UI. */
export interface Docfieldlayout {
  active: boolean;
  allow_field_assignment?: boolean;
  show_in_edit_mode?: boolean;
  panel?: string;
  label: string;
  icon?: string;
  order?: number;
  fields?: string;
}

/** Document lifecycles are the sequences of states a document goes through during its life. */
export interface Doclifecycle {
  label: string;
  active: boolean;
  description?: string;
  state_labels: string;
  labeled_states: string;
  expiration?: string;
  signature_pages?: string;
  overlays?: string;
}

export interface Doclifecycle_Doclifecyclerole {
  label: string;
  active: boolean;
  description?: string;
  application_role: string;
  note_color?: string;
  security_options?: string;
}

export interface Doclifecycle_Doclifecyclestate {
  label: string;
  active: boolean;
  description?: string;
  order: number;
  cancel_state?: string;
  skip_cancel_state?: boolean;
  skip_entry_actions_cancel_state?: boolean;
  entry_criteria?: string;
  entry_action?: string;
  user_action?: string;
  security_settings?: string;
}

export interface Doclifecycle_Doclifecycleevent {
  active: boolean;
  event: string;
  rule: string;
  order: number;
}

/** Dynamic Access Control (DAC) is an access control model for documents, which automates the assignment of users to document roles. */
export interface Docmatchingrule {
  lifecycle: string;
  role: string;
  active: boolean;
  description?: string;
  label: string;
}

export interface Docmatchingrule_Docmatchingcriterion {
  object_field: Docmatchingrule_Docmatchingcriterion_Object_field[];
  data_field: string;
  object: string;
}

/** Workflow participant rules control which document fields force the selected role to be required or hidden in the workflow start dialog. */
export interface Docparticipantrule {
  lifecycle: string;
  role: string;
  active: boolean;
  description?: string;
  label: string;
}

export interface Docparticipantrule_Docparticipantcriterion {
  object_field: Docparticipantrule_Docparticipantcriterion_Object_field[];
  data_field: string;
  object: string;
}

/** Document relationship types define the type of relationship between two documents. */
export interface Docrelationshiptype {
  label: string;
  active: boolean;
  single_use?: boolean;
  target_doctype?: string;
  source_version_specific?: boolean;
  carry_forward_source_specific?: boolean;
  target_version_specific?: boolean;
  carry_forward_target_specific?: boolean;
  relationship_filter?: string;
}

/** Document types refer to a hierarchical taxonomy to differentiate the various types of documents stored in Vault. */
export interface Doctype {
  label: string;
  active: boolean;
  description?: string;
  fields?: string;
  rendition_types?: string;
  available_lifecycles?: string;
  relationship_types?: string;
  allow_attachments?: boolean;
  document_number_format?: string;
  document_name_format?: string;
  filters?: string;
  create_document_permissions?: string;
  create_binder_permissions?: string;
}

/** Documents can have custom actions that perform specific instructions. */
export interface Documentaction {
  label?: string;
  active?: boolean;
  user_input_object_type?: string;
  user_input_object?: string;
  icon?: string;
  lifecycle?: string;
  checksum?: string;
  usages?: string;
  source_code?: string;
}

/** Document checks are document type-specific checklist items displayed in the Document Readiness panel to help guide users in documentation reviews. */
export interface Documentcheck {
  label?: string;
  active?: boolean;
  document_check_section: string;
  help_text?: string;
  order: number;
  pending_delete?: boolean;
}

/** Document check sections contain document checks within the Document Readiness panel in the Vault UI. */
export interface Documentchecksection {
  label?: string;
  active?: boolean;
  document_check_count?: number;
  document_type?: string;
  help_text?: string;
  pending_delete?: boolean;
}

/** Lifecycle Stages allow organizations to group document lifecycle states into ordered collections of stages. */
export interface Documentstagegroup {
  label?: string;
  active?: boolean;
  document_lifecycle: string;
  description?: string;
}

export interface Documentstagegroup_Documentstage {
  label?: string;
  active?: boolean;
  document_lifecycle_states_ref?: string;
  description?: string;
  order: number;
}

/** When creating Inbound Email Addresses in the Vault Admin UI, Admins select an Email Processor. */
export interface Emailprocessor {
  label?: string;
  active?: boolean;
  checksum?: string;
  source_code?: string;
  allowed_senders: string;
}

/** Formatted output templates allow quick transfer of complete object record data into a customized template file. */
export interface Formattedoutput {
  label?: string;
  active?: boolean;
  input_file_type?: string;
  output_format?: string;
  root_object: string;
  root_object_type: string;
  template_file: string;
}

/** When configuring Email to Vault, Admins use Inbound Email Addresses to create and manage Vault-owned email addresses. */
export interface Inboundemailaddress {
  label?: string;
  active?: boolean;
  authenticate_sender?: boolean;
  email_username: string;
  email_processor_user?: string;
  description?: string;
  email_processor: string;
  allowed_sender_groups?: string;
  allowed_senders: string;
  run_as_processor_user?: boolean;
  send_bounce_notification?: boolean;
}

/** Integration Rules provide platform configuration and SDK services for developers building Vault to Vault integration features. */
export interface Integrationrule {
  label?: string;
  active?: boolean;
  integration_point?: string;
  primary_query_object?: string;
  connection: string;
}

export interface Integrationrule_Queryobjectrule {
  label?: string;
  active?: boolean;
  query_object: string;
  filter_clause?: string;
}

export interface Integrationrule_Fieldrule {
  label?: string;
  active?: boolean;
  target_field_lookup?: Integrationrule_Fieldrule_Target_field_lookup[];
  query_field_select?: Integrationrule_Fieldrule_Query_field_select[];
  query_field_type?: string;
  query_object?: string;
  query_field?: string;
  target_object?: string;
  target_object_field?: Integrationrule_Fieldrule_Target_object_field[];
  target_docfield?: string;
  field_default?: string;
}

/** Vault's job scheduler provides a simpler, streamlined way for Admins to create a batch operation and schedule it to repeatedly execute. */
export interface Job {
  label: string;
  active: boolean;
  schedule?: string;
  type?: string;
  owner?: string;
  priority?: string;
  timezone?: string;
  time?: string;
  repeat_on_days?: string;
  hourly_interval?: number;
  day_of_month?: number;
  week_number?: number;
  week_day?: string;
  month_repeat_type?: string;
  trigger_date?: string;
  usages?: string;
  active_after_clone?: boolean;
}

export interface Job_Sendnotificationjobaction {
  label: string;
  primary_action?: boolean;
  conditions?: string;
  recipients?: string;
  template?: string;
}

export interface Job_Changeobjectstatejobaction {
  label: string;
  primary_action?: boolean;
  conditions: string;
  object: string;
  terminate_existing_workflows: boolean;
  destination_state: string;
}

export interface Job_Sdkbatchjobaction {
  label: string;
  primary_action?: boolean;
  job_metadata: string;
}

export interface Job_Flashreportjobaction {
  label: string;
  primary_action?: boolean;
  report_name?: string;
  send_notifications_on_completion?: boolean;
  custom_flash_email?: boolean;
}

export interface Job_Scheduleddataexportjobaction {
  label: string;
  primary_action?: boolean;
  selected_entities: string;
  data_storage_option: string;
  enable_full_data_export: boolean;
}

export interface Job_Httppostjobaction {
  label: string;
  primary_action?: boolean;
  post_session_creds?: boolean;
  run_as_user?: string;
  external_url?: string;
}

/** The Jobmetadata component is used to define the metadata of an SDK Job. */
export interface Jobmetadata {
  label?: string;
  active?: boolean;
  chunk_size?: number;
  single_instance_states?: string;
  description?: string;
  job_code: string;
  queue?: string;
  timeout_duration?: string;
}

/** Label Sets allow Vault Admins to specify alternative Vault UI labels within a single language. */
export interface Labelset {
  label?: string;
  active?: boolean;
  description?: string;
  language: string;
  language_label?: string;
}

/** Layout profiles are collections of layouts for various objects. */
export interface Layoutprofile {
  label?: string;
  active?: boolean;
  description?: string;
}

export interface Layoutprofile_Objectlayout {
  label?: string;
  active?: boolean;
  record_creation_layout?: string;
  objecttype?: string;
  object: string;
  pagelayouts: string;
}

/** Layout rules streamline the data entry process by dynamically hiding irrelevant fields and sections on an object record detail page. */
export interface Layoutrule {
  label: string;
  active: boolean;
  expression: string;
  description?: string;
  order?: number;
  hide_layout?: boolean;
  focus_on_layout?: boolean;
  hide_fields?: string;
  hide_sections?: string;
  hidden_pages?: string;
  displayed_as_required_fields?: string;
  displayed_as_readonly_fields?: string;
  hide_controls?: string;
}

/** Lifecycle state types are groups of lifecycle states across lifecycles that share common behaviors. */
export interface Lifecyclestatetype {
  label?: string;
  active?: boolean;
  description?: string;
  state_required: boolean;
  lifecycle_scope: string;
  object_lifecycle?: string;
  document_lifecycle?: string;
  lifecycle_type: string;
}

/** Lifecycle state type associations join a state and a lifecycle state type record. */
export interface Lifecyclestatetypeassociation {
  label?: string;
  active?: boolean;
  document_lifecycle?: string;
  state_type: string;
  description?: string;
  object_lifecycle_state?: Lifecyclestatetypeassociation_Object_lifecycle_state[];
  document_lifecycle_state?: string;
  object_lifecycle?: string;
}

/** Defines a Web Action that allows users to execute links to external URLs. */
export interface Link {
  send_session_via_post_message?: boolean;
  url: string;
  component: string;
  label: string;
  target: string;
}

/** Matching Sharing Rules are part of Dynamic Access Control for objects. */
export interface Matchingrule {
  label: string;
  active: boolean;
  description?: string;
  object: string;
  role: string;
  user_role_fields: string;
  data_fields: string;
}

/** Spark message delivery event handlers contain custom logic that executes when a Spark message is delivered to a remote queue. */
export interface Messagedeliveryeventhandler {
  label?: string;
  active?: boolean;
  checksum?: string;
  source_code?: string;
}

/** To provide a localized experience, Vault Java SDK developers can create translations for messages displayed to the end user in Vault. */
export interface Messagegroup {
  label?: string;
  active?: boolean;
}

export interface Messagegroup_Message {
  label?: string;
  active?: boolean;
  default_value: string;
  system_managed?: boolean;
}

/** Spark message processors provide logic to handle Spark messages. */
export interface Messageprocessor {
  label?: string;
  active?: boolean;
  checksum?: string;
  source_code?: string;
}

/** Configuration for Vault Mobile sharing actions that define how files can be shared to Vault from smartphones. */
export interface Mobileshareactionconfig {
  label?: string;
  active?: boolean;
  mobile_share_action: string;
  allowed_for: string;
  description?: string;
  allowed_for_groups?: string;
}

/** A tab within the UI of Vault Mobile. Admins can configure mobile tabs from Mobile Setup in the Vault UI. */
export interface Mobiletab {
  label?: string;
  active?: boolean;
  icon?: string;
  minimum_app_version?: string;
  page?: string;
  mobile_label?: string;
  order?: number;
}

/** Defines a template for email and notification messages sent by Vault. */
export interface Notificationtemplate {
  label: string;
  active: boolean;
  description?: string;
  subject: string;
  notification: string;
  email_body: string;
  entity_type: string;
  email_preferences?: string;
  referenced_component?: string;
  notification_category?: string;
}

/** Objects define the data model for the Vault. */
export interface Object {
  label: string;
  label_plural: string;
  active: boolean;
  description?: string;
  help_content?: string;
  in_menu: boolean;
  enable_esignatures?: boolean;
  allow_attachments?: boolean;
  audit: boolean;
  dynamic_security?: boolean;
  system_managed?: boolean;
  available_lifecycles?: string;
  object_class?: string;
  allow_types?: boolean;
  data_store?: string;
  user_role_setup_object?: string;
  secure_audit_trail?: boolean;
  secure_sharing_settings?: boolean;
  secure_attachments?: boolean;
  prevent_record_overwrite?: boolean;
  multi_select?: boolean;
  enable_merges?: boolean;
  summary_fields?: string;
  security_tree_object?: string;
  secure_copy_record?: boolean;
  relate_records_select_all?: boolean;
}

export interface Object_Field {
  label: string;
  type: string;
  active: boolean;
  required: boolean;
  description?: string;
  unique?: boolean;
  max_length?: number;
  max_value?: number;
  min_value?: number;
  scale?: number;
  format_mask?: string;
  help_content?: string;
  list_column: boolean;
  order?: number;
  multi_value?: boolean;
  object?: string;
  relationship_type?: string;
  relationship_outbound_name?: string;
  relationship_inbound_name?: string;
  relationship_inbound_label?: string;
  controlling_field?: string;
  relationship_deletion?: string;
  relationship_criteria?: string;
  system_managed_name?: boolean;
  start_number?: number;
  value_format?: string;
  component?: string;
  relationship_deep_copy?: boolean;
  picklist?: string;
  lookup_relationship_name?: string;
  lookup_source_field?: string;
  document_version_reference?: string;
  no_copy?: boolean;
  create_object_inline?: boolean;
  sequential_naming?: boolean;
  default_value?: string;
  checkbox?: boolean;
  formula?: string;
  encrypted?: boolean;
  editable?: boolean;
  searchable?: boolean;
  secure_relationship?: boolean;
  subtype?: string;
  rollup?: boolean;
  rollup_relationship_name?: string;
  rollup_function?: string;
  rollup_source_field?: string;
  rollup_filter_criteria?: string;
}

export interface Object_Index {
  label: string;
  fields: string;
  description?: string;
}

/** Object user actions are configurable instances of a Vault Action that you can define for an object and object types. */
export interface Objectaction {
  label: string;
  description?: string;
  parameters?: string;
  action: string;
  active: boolean;
  keyboard_shortcut?: string;
  available_all_states?: boolean;
}

/** Definition of the lifecycle of an Object. */
export interface Objectlifecycle {
  label: string;
  active: boolean;
  starting_state?: string;
  description?: string;
  state_labels?: string;
  labeled_states?: string;
}

export interface Objectlifecycle_Objectlifecyclestate {
  label: string;
  active: boolean;
  record_status?: string;
  description?: string;
  cancel_state?: string;
  skip_cancel_state?: boolean;
  skip_entry_actions_cancel_state?: boolean;
}

export interface Objectlifecycle_Objectlifecyclestateentrycriterion {
  rule: string;
  order?: number;
}

export interface Objectlifecycle_Objectlifecyclestateentryaction {
  rule: string;
  order?: number;
}

export interface Objectlifecycle_Objectlifecyclestateuseraction {
  label: string;
  rule: string;
  order: number;
  keyboard_shortcut?: string;
}

export interface Objectlifecycle_Objectlifecycleevent {
  rule: string;
  event: string;
  order: number;
}

export interface Objectlifecycle_Objectlifecyclerole {
  active: boolean;
  application_role: string;
  permissions?: string;
}

export interface Objectlifecycle_Objectlifecyclepermission {
  states: string;
  role: string;
  permission: string;
}

/** Lifecycle Stages allow organizations to group object lifecycle states into ordered collections of stages. */
export interface Objectlifecyclestagegroup {
  label?: string;
  active?: boolean;
  object_type: string;
  description?: string;
  object_lifecycle: string;
}

export interface Objectlifecyclestagegroup_Objectlifecyclestage {
  label?: string;
  active?: boolean;
  object_lifecycle_states_ref?: string;
  description?: string;
  order: number;
}

/** An object type is a collection of fields that are grouped to capture similar but not identical data within a single object. */
export interface Objecttype {
  label: string;
  label_plural: string;
  active: boolean;
  description?: string;
  additional_type_validations?: string;
  default_type?: boolean;
  summary_fields?: string;
}

export interface Objecttype_Typefield {
  source: string;
  required?: boolean;
  relationship_criteria?: string;
  pickentries?: string;
  default_value?: string;
}

export interface Objecttype_Typeaction {
  action: string;
}

/** Validation rules enhance data quality by comparing user-entered data with Admin-defined rules for an object record. */
export interface Objectvalidation {
  label: string;
  active: boolean;
  expression: string;
  description?: string;
  error_message: string;
  error_location: string;
}

/** Object workflows apply to a single object lifecycle and a single object. */
export interface Objectworkflow {
  label: string;
  active: boolean;
  object_lifecycles: string;
  document_content_lifecycle?: string;
  record_content_lifecycle?: string;
  description?: string;
  auto_start?: boolean;
  envelope_name_format?: string;
  cancellation_comment?: boolean;
  workflow_variables?: string;
  workflow_content_type?: string;
  class?: string;
  cardinality?: string;
  show_fields?: boolean;
  bound_create_draft?: boolean;
  users_can_only_complete_one_task?: boolean;
}

export interface Objectworkflow_Objectworkflowstep {
  label: string;
  type: string;
  next_steps?: string;
  tags?: string;
  step_detail?: string;
  description?: string;
}

export interface Objectworkflow_WorkflowCancelationAction {
  rule: string;
  order: number;
}

/** Defines the configuration of an Outbound Email Domain. Specific to MedInquiry, Clinical, and QMS Vaults with Outbound Email Domains enabled. */
export interface Outboundemaildomain {
  label?: string;
  active?: boolean;
  dns_records?: string;
  description?: string;
  verification_status?: string;
  sender_domain: string;
}

/** Overlays allow you to apply information to a viewable rendition when users download it from Vault. */
export interface Overlaytemplate {
  label?: string;
  active?: boolean;
  allow_download_without_overlay?: boolean;
  footer_01?: string;
  footer_02?: string;
  footer_richtext?: string;
  header_01?: string;
  header_02?: string;
  header_richtext?: string;
  overlay_override_page_match?: boolean;
  schema_version?: number;
  watermark?: string;
  watermark_richtext?: string;
}

export interface Overlaytemplate_Overlayoverride {
  label?: string;
  active?: boolean;
  file?: string;
  footer_richtext?: string;
  header_footer_placement?: string;
  header_richtext?: string;
  orientation?: string;
  overlay_format_type?: string;
  paper_size?: string;
  watermark_richtext?: string;
}

/** This component defines a custom page in Vault. */
export interface Page {
  label?: string;
  active?: boolean;
  client_distribution?: string;
  description?: string;
  disable_chrome?: boolean;
  page_client_code?: string;
  page_controller?: string;
  url_path_name?: string;
}

/** The layout of an object detail page. */
export interface Pagelayout {
  label: string;
  default_layout?: boolean;
  page_markup: string;
  active?: boolean;
  description?: string;
  display_lifecycle_stages?: boolean;
}

/** Represents a page link that navigates users to custom URL or Custom Page destinations. */
export interface Pagelink {
  label: string;
  active?: boolean;
  description?: string;
  url: string;
  mode: string;
  page_type: string;
  object?: string;
  page?: string;
  disable_type_select?: boolean;
}

/** Permission sets are a way to group permissions together. */
export interface Permissionset {
  label: string;
  active?: boolean;
}

/** Picklists allow users to select a value for a field from a range of predefined options. */
export interface Picklist {
  label: string;
  active: boolean;
  can_add_values?: boolean;
  can_reorder_values?: boolean;
  system_managed?: boolean;
  controlling_picklist?: string;
  order_type?: string;
}

export interface Picklist_Picklistentry {
  value: string;
  order: number;
  active: boolean;
  controlling_entries?: string;
}

/** Defines the configuration of a Printable View for Test Scripts, specific to Quality Vaults with Veeva Validation Management enabled. */
export interface Printableviewtestscript {
  label?: string;
  active?: boolean;
  template: string;
  header_logo?: string;
}

/** Specific to Quality family Vaults, supports the Change Control features. Action Paths allow administrators to define Action Paths and Steps for the Change Control Process. */
export interface Qmsactionpathconfiguration {
  label?: string;
  active?: boolean;
  action_object: string;
  allow_in_progress_changes?: boolean;
  description?: string;
  hide_other_actions?: boolean;
  parent_locked_states?: Qmsactionpathconfiguration_Parent_locked_states[];
  parent_object: string;
  parent_object_lifecycle: string;
  parent_object_type: string;
}

export interface Qmsactionpathconfiguration_Qmsactionstepconfiguration {
  label?: string;
  active?: boolean;
  action_finish_states: Qmsactionpathconfiguration_Qmsactionstepconfiguration_Action_finish_states[];
  action_object_lifecycle: string;
  action_start_state: Qmsactionpathconfiguration_Qmsactionstepconfiguration_Action_start_state[];
  included_action_states: Qmsactionpathconfiguration_Qmsactionstepconfiguration_Included_action_states[];
  order: number;
  parent_finish_state: Qmsactionpathconfiguration_Qmsactionstepconfiguration_Parent_finish_state[];
  parent_object_lifecycle: string;
  parent_start_state: Qmsactionpathconfiguration_Qmsactionstepconfiguration_Parent_start_state[];
}

export interface Qmsactionpathconfiguration_Qmsactionpathfieldtodisplay {
  label?: string;
  active?: boolean;
  field_name: string;
}

/** Specific to Quality family Vaults, supports External Audit Collaboration features. Used to automatically create External User accounts for existing Person records. */
export interface Qmsautomationusertemplate {
  label?: string;
  active?: boolean;
  application_role?: string;
  description?: string;
  security_profile?: string;
  username_pattern: string;
}

/** Specific to Quality family Vaults, supports the Batch Release Role Assignment feature. Defines a Batch Disposition Role Assignment. */
export interface Qualitybatchroleassignment {
  label?: string;
  active?: boolean;
  assigned_lifecyclerole?: Qualitybatchroleassignment_Assigned_lifecyclerole[];
  eligible_lifecyclerole?: Qualitybatchroleassignment_Eligible_lifecyclerole[];
  object: string;
  objectlifecycle: string;
  single_user_role?: boolean;
  user_field: Qualitybatchroleassignment_User_field[];
}

/** Specific to Quality family Vaults, supports Curriculum Matching feature. Defines a Curriculum Matching Rule to automatically assign Curricula to Learners. */
export interface Qualitycurriculumsmartmatchrule {
  label?: string;
  active?: boolean;
  description?: string;
  matching_type: string;
}

export interface Qualitycurriculumsmartmatchrule_Qualitycurriculumsmartmatchfield {
  label?: string;
  active?: boolean;
  curriculum_field: Qualitycurriculumsmartmatchrule_Qualitycurriculumsmartmatchfield_Curriculum_field[];
  curriculum_object: string;
  person_field: Qualitycurriculumsmartmatchrule_Qualitycurriculumsmartmatchfield_Person_field[];
  person_object: string;
}

/** Specific to Quality family Vaults, supports External Notifications feature. Defines filters to limit Persons available for inclusion in a Distribution Group. */
export interface Qualitydistributiongroupfilterset {
  label?: string;
  active?: boolean;
  target_object: string;
  target_object_type: string;
}

export interface Qualitydistributiongroupfilterset_Qualitydistributiongroupfilter {
  label?: string;
  active?: boolean;
  person_object: string;
  person_object_match_field: Qualitydistributiongroupfilterset_Qualitydistributiongroupfilter_Person_object_match_field[];
  target_object: string;
  target_object_match_field: Qualitydistributiongroupfilterset_Qualitydistributiongroupfilter_Target_object_match_field[];
  target_object_type: string;
}

/** Specific to Quality family Vaults, supports Dynamic Enrollment feature. Defines a Dynamic Enrollment Rule to automatically assign Learner roles to a Person. */
export interface Qualitydynamicenrollmentrule {
  label?: string;
  active?: boolean;
  description?: string;
  matching_type: string;
}

export interface Qualitydynamicenrollmentrule_Qualitydynamicenrollmentmatchingfield {
  label?: string;
  active?: boolean;
  learner_role_field: Qualitydynamicenrollmentrule_Qualitydynamicenrollmentmatchingfield_Learner_role_field[];
  learner_role_object: string;
  person_field: Qualitydynamicenrollmentrule_Qualitydynamicenrollmentmatchingfield_Person_field[];
  person_object: string;
}

/** Available only in Quality Vaults with QMS enabled. Defines a Quality External Notification for distribution groups, notification templates, emails, and document sharing. */
export interface Qualityexternalnotification {
  label?: string;
  active?: boolean;
  description?: string;
  do_not_include_header_footer: boolean;
  target_object: string;
  target_object_type: string;
}

export interface Qualityexternalnotification_Qualitydistributiongroup {
  label?: string;
  active?: boolean;
  distribution_group_filter_set?: string;
  help_content?: string;
  maximum_recipients: number;
  minimum_recipients: number;
  populated_by_action?: boolean;
  prevent_manual_update?: boolean;
}

/** Available only in Quality Vaults with QMS enabled. Defines a Notification Template for Quality External Notifications. */
export interface Qualityexternalnotificationtemplate {
  label?: string;
  active?: boolean;
  recipient_roles?: string;
  download_option?: string;
  help_content?: string;
  target_object: string;
  notification_template: string;
  document_link_field?: Qualityexternalnotificationtemplate_Document_link_field[];
  object_fields_on_document?: string;
  can_add_or_remove_documents?: boolean;
  expiry_period?: number;
  recipients?: Qualityexternalnotificationtemplate_Recipients[];
  recipient?: Qualityexternalnotificationtemplate_Recipient[];
  document_reference_fields?: Qualityexternalnotificationtemplate_Document_reference_fields[];
  can_add_or_remove_participants: boolean;
  quality_external_notification: string;
}

export interface Qualityexternalnotificationtemplate_Qualityextnotificationtemplatedistgroup {
  label?: string;
  active?: boolean;
  quality_distribution_group: string;
  quality_external_notification: string;
}

/** Specific to Quality family Vaults, supports Supplier Change Notification feature. Defines an Extract Entity Configuration. */
export interface Qualityextractentityconfig {
  label?: string;
  active?: boolean;
  match_field: Qualityextractentityconfig_Match_field[];
  match_object: string;
  match_object_type?: string;
}

/** Specific to Quality family Vaults, supports Complaint Email Ingestion and Supplier Change Notice Email Ingestion features. */
export interface Qualityinboundemailaddressconfiguration {
  label?: string;
  active?: boolean;
  description?: string;
  disable_email_attachments_parsing?: boolean;
  disable_email_body_parsing?: boolean;
  email_thread_detection?: string;
  inbound_email_address: string;
  target_object: string;
  target_object_type: string;
}

/** Available only in Quality Vaults with QMS enabled. Defines a Quality Incident Layout Configuration for intake forms. */
export interface Qualityincidentintakelayout {
  label?: string;
  active?: boolean;
  description?: string;
}

export interface Qualityincidentintakelayout_Qualityincidentintakelayoutfield {
  label?: string;
  active?: boolean;
  target_object: string;
  field_to_display: Qualityincidentintakelayout_Qualityincidentintakelayoutfield_Field_to_display[];
  order: number;
}

/** Defines how records of a configured object type will be duplicated and inverted to support bidirectional relationships. */
export interface Qualityobjectrelationshipconfig {
  label?: string;
  active?: boolean;
  left_field?: Qualityobjectrelationshipconfig_Left_field[];
  left_fields?: Qualityobjectrelationshipconfig_Left_fields[];
  relationship_picklist: Qualityobjectrelationshipconfig_Relationship_picklist[];
  right_field?: Qualityobjectrelationshipconfig_Right_field[];
  right_fields?: Qualityobjectrelationshipconfig_Right_fields[];
  target_object: string;
  target_object_type?: string;
}

export interface Qualityobjectrelationshipconfig_Qualitypicklistvaluemap {
  label?: string;
  active?: boolean;
  inverse_picklist_value: Qualityobjectrelationshipconfig_Qualitypicklistvaluemap_Inverse_picklist_value[];
  picklist_value: Qualityobjectrelationshipconfig_Qualitypicklistvaluemap_Picklist_value[];
  relationship_picklist: string;
}

/** Available only in Quality Vaults with QMS enabled. Defines a Reason for Change configuration to prompt users for a reason upon editing specific fields. */
export interface Qualityreasonforchange {
  label?: string;
  active?: boolean;
  capture_field?: Qualityreasonforchange_Capture_field[];
  capture_type_field?: Qualityreasonforchange_Capture_type_field[];
  description?: string;
  record_trigger?: string;
  target_object: string;
  target_object_type?: string;
}

export interface Qualityreasonforchange_Qualityreasonforchangefieldcriteria {
  label?: string;
  active?: boolean;
}

export interface Qualityreasonforchange_Qualityreasonforchangecommentcriterion {
  label?: string;
  active?: boolean;
  reason_for_change_picklist: string;
  value_requiring_comment: Qualityreasonforchange_Qualityreasonforchangecommentcriterion_Value_requiring_comment[];
}

export interface Qualityreasonforchange_Qualityreasonforchangestate {
  label?: string;
  active?: boolean;
  field_criteria: string;
  lifecycle_state: Qualityreasonforchange_Qualityreasonforchangestate_Lifecycle_state[];
  object_lifecycle: string;
  reason_for_change: string;
}

export interface Qualityreasonforchange_Qualityreasonforchangefield {
  label?: string;
  active?: boolean;
  field: Qualityreasonforchange_Qualityreasonforchangefield_Field[];
  field_criteria: string;
  reason_for_change: string;
  target_object: string;
  target_object_type?: string;
}

/** Defines the configuration for a Quality Record Check for checking recurrences and duplicates. */
export interface Qualityrecordcheck {
  label?: string;
  active?: boolean;
  allow_date_overrides?: boolean;
  allow_match_field_overrides?: boolean;
  check_type: Qualityrecordcheck_Check_type[];
  check_type_picklist: string;
  copy_attachments?: boolean;
  date_field: Qualityrecordcheck_Date_field[];
  date_quantity: number;
  date_unit_entry: Qualityrecordcheck_Date_unit_entry[];
  date_unit_picklist: string;
  description?: string;
  display_fields?: Qualityrecordcheck_Display_fields[];
  editable_fields?: Qualityrecordcheck_Editable_fields[];
  record_check_insight_config?: string;
  relationship_type?: string;
  relationship_type_entries?: Qualityrecordcheck_Relationship_type_entries[];
  relationship_type_picklist?: string;
  search_fields?: Qualityrecordcheck_Search_fields[];
  simple_match_fields?: Qualityrecordcheck_Simple_match_fields[];
  target_object: string;
  target_object_type: string;
}

export interface Qualityrecordcheck_Qualitymatchtier {
  label?: string;
  active?: boolean;
  display: boolean;
  justification_required?: boolean;
  relationship_type_required?: boolean;
  similarity_score: number;
  tier: number;
  tier_color: string;
  tier_label: Qualityrecordcheck_Qualitymatchtier_Tier_label[];
  tier_picklist: string;
}

export interface Qualityrecordcheck_Qualityfieldsuggestion {
  label?: string;
  active?: boolean;
  order: number;
  suggestion_field: Qualityrecordcheck_Qualityfieldsuggestion_Suggestion_field[];
  target_object: string;
}

export interface Qualityrecordcheck_Qualityrelatedrecordmatchfield {
  label?: string;
  active?: boolean;
  filter_object: string;
  quality_record_check?: string;
  record_check_rule?: string;
  related_inbound_field?: Qualityrecordcheck_Qualityrelatedrecordmatchfield_Related_inbound_field[];
  related_match_fields: Qualityrecordcheck_Qualityrelatedrecordmatchfield_Related_match_fields[];
  related_object: string;
  related_outbound_field?: Qualityrecordcheck_Qualityrelatedrecordmatchfield_Related_outbound_field[];
  relationship_type: string;
}

export interface Qualityrecordcheck_Qualityrelationshipstatemap {
  label?: string;
  active?: boolean;
  object_lifecycle: string;
  object_lifecycle_state?: Qualityrecordcheck_Qualityrelationshipstatemap_Object_lifecycle_state[];
  quality_object_relationship_config: string;
  quality_picklist_value_map: string;
}

export interface Qualityrecordcheck_Qualityrelatedrecordsuggestion {
  label?: string;
  active?: boolean;
  join_object: string;
  order: number;
  related_object: string;
  related_object_display_field: Qualityrecordcheck_Qualityrelatedrecordsuggestion_Related_object_display_field[];
  related_object_reference_field: Qualityrecordcheck_Qualityrelatedrecordsuggestion_Related_object_reference_field[];
  target_object_reference_field: Qualityrecordcheck_Qualityrelatedrecordsuggestion_Target_object_reference_field[];
}

export interface Qualityrecordcheck_Qualityrecordcheckrule {
  label?: string;
  active?: boolean;
  order: number;
  search_fields?: Qualityrecordcheck_Qualityrecordcheckrule_Search_fields[];
  simple_match_fields: Qualityrecordcheck_Qualityrecordcheckrule_Simple_match_fields[];
  target_object: string;
}

/** Available only in Quality Vaults with QMS enabled. Defines the configuration for a Quality Record Check Insight. */
export interface Qualityrecordcheckinsight {
  label?: string;
  active?: boolean;
  target_object_type: string;
  probable_similarity_score: number;
  target_object: string;
  description?: string;
  likely_similarity_score: number;
}

export interface Qualityrecordcheckinsight_Qualityinsightrelatedrecordsummaryfield {
  label?: string;
  active?: boolean;
  related_inbound_field: Qualityrecordcheckinsight_Qualityinsightrelatedrecordsummaryfield_Related_inbound_field[];
  summary_field: Qualityrecordcheckinsight_Qualityinsightrelatedrecordsummaryfield_Summary_field[];
  relationship_type: string;
  summary_object: string;
  related_outbound_field?: Qualityrecordcheckinsight_Qualityinsightrelatedrecordsummaryfield_Related_outbound_field[];
  related_object: string;
}

export interface Qualityrecordcheckinsight_Qualityinsightdisplayfield {
  label?: string;
  active?: boolean;
  target_object_type: string;
  target_object: string;
  display_field: Qualityrecordcheckinsight_Qualityinsightdisplayfield_Display_field[];
  order: number;
  display_type_field: Qualityrecordcheckinsight_Qualityinsightdisplayfield_Display_type_field[];
}

export interface Qualityrecordcheckinsight_Qualityinsightsummaryfield {
  label?: string;
  active?: boolean;
  summary_type_field: Qualityrecordcheckinsight_Qualityinsightsummaryfield_Summary_type_field[];
  target_object_type: string;
  summary_field: Qualityrecordcheckinsight_Qualityinsightsummaryfield_Summary_field[];
  target_object: string;
}

/** Available only in Quality Vaults with QMS enabled. Defines the lifecycle state associations for a Quality Record Check. */
export interface Qualityrecordchecklifecycleassociation {
  label?: string;
  active?: boolean;
  lifecycle_states_to_include?: Qualityrecordchecklifecycleassociation_Lifecycle_states_to_include[];
  lifecycle_states_to_not_transition?: Qualityrecordchecklifecycleassociation_Lifecycle_states_to_not_transition[];
  object_lifecycle: string;
  quality_record_check: string;
}

export interface Qualityrecordchecklifecycleassociation_Qualityrelationshipstatemapassociation {
  label?: string;
  active?: boolean;
  object_lifecycle: string;
  object_lifecycle_state?: Qualityrecordchecklifecycleassociation_Qualityrelationshipstatemapassociation_Object_lifecycle_state[];
  quality_object_relationship_config: string;
  quality_picklist_value_map: string;
  quality_record_check: string;
}

/** Available only in Quality Vaults with QMS or VPS enabled. Defines the configuration for a Reportability Decision Tree. */
export interface Qualityreportabilitydecisiontreeconfig {
  label?: string;
  active?: boolean;
  description?: string;
  reportability_assessment_type: string;
}

export interface Qualityreportabilitydecisiontreeconfig_Qualityrdtqualifier {
  label?: string;
  active?: boolean;
  description?: string;
  object: string;
  object_field: Qualityreportabilitydecisiontreeconfig_Qualityrdtqualifier_Object_field[];
  picklist: string;
  picklist_value: Qualityreportabilitydecisiontreeconfig_Qualityrdtqualifier_Picklist_value[];
}

/** Specific to Quality family Vaults, supports Quality Teams feature. Allows users to make individual work assignments to Change Controls, Audits, CAPAs, or other processes. */
export interface Qualityteam {
  label?: string;
  active?: boolean;
  destination_state?: Qualityteam_Destination_state[];
  initial_state?: Qualityteam_Initial_state[];
  locked_in_states?: string;
  object: string;
  objectlifecycle?: string;
  objecttype?: string;
  process_relationships?: string;
  quality_team_member_object_name?: string;
  record_trigger?: string;
  recordrole_trigger?: string;
}

export interface Qualityteam_Qualityteamrolemembershiprestriction {
  label?: string;
  active?: boolean;
  exclusive_role?: string;
  exclusive_team_role?: string;
  quality_team?: string;
  role?: string;
  team_role?: string;
}

export interface Qualityteam_Qualityteamrole {
  label?: string;
  active?: boolean;
  cascade_behavior: string;
  cascade_from_multiple_relationship?: Qualityteam_Qualityteamrole_Cascade_from_multiple_relationship[];
  cascade_from_relationship?: Qualityteam_Qualityteamrole_Cascade_from_relationship[];
  constraining_role?: string;
  display_order: number;
  exclusive_membership?: boolean;
  help_content?: string;
  linked_application_role?: string;
  linked_field?: Qualityteam_Qualityteamrole_Linked_field[];
  locked_in_states?: string;
  maximum_user_in_role: number;
  minimum_user_in_role: number;
  object: string;
  objectlifecycle?: string;
}

export interface Qualityteam_Qualityteamrelatedobjectsecurity {
  label?: string;
  active?: boolean;
  app_security_assignment: string;
  app_security_rule: string;
  application_role: string;
  outbound_reference_field: Qualityteam_Qualityteamrelatedobjectsecurity_Outbound_reference_field[];
  quality_team: string;
  quality_team_role: string;
  related_object: string;
}

/** Available only in Quality Vaults with QMS enabled. Defines the lifecycle associations for a Quality Team. */
export interface Qualityteamlifecycleassociation {
  label?: string;
  active?: boolean;
  destination_state?: Qualityteamlifecycleassociation_Destination_state[];
  initial_state?: Qualityteamlifecycleassociation_Initial_state[];
  locked_in_states?: Qualityteamlifecycleassociation_Locked_in_states[];
  objectlifecycle: string;
  quality_team: string;
}

export interface Qualityteamlifecycleassociation_Qualityteamroleassociation {
  label?: string;
  active?: boolean;
  constraining_role?: string;
  linked_application_role: string;
  locked_in_states?: Qualityteamlifecycleassociation_Qualityteamroleassociation_Locked_in_states[];
  objectlifecycle: string;
  quality_team: string;
  quality_team_role: string;
}

/** Spark messaging is a message notification system that allows for loosely coupled, asynchronous, near real-time integration between Vaults or external applications. */
export interface Queue {
  label?: string;
  active?: boolean;
  sequential?: boolean;
  message_delivery_event_handler?: string;
  description?: string;
  type: string;
  message_processor_user?: string;
  rollback_on_error?: boolean;
  message_processor?: string;
  message_delivery_event_handler_user?: string;
}

export interface Queue_Queueconnection {
  label?: string;
  active?: boolean;
  connection_type?: string;
  connection_name: string;
  deliver_to_queue?: string;
}

/** Custom actions for records, called record actions, are invoked by a user on a specific record from the UI or API. */
export interface Recordaction {
  label?: string;
  active?: boolean;
  user_input_object_type?: string;
  user_input_object?: string;
  icon?: string;
  type?: string;
  checksum?: string;
  usages?: string;
  source_code?: string;
  object?: string;
}

/** Record merge event handlers contain custom logic to execute as a merge starts or after the merge finishes. */
export interface Recordmergeeventhandler {
  label?: string;
  active?: boolean;
  checksum?: string;
  source_code?: string;
  object: string;
}

/** Record role triggers execute custom code whenever roles are directly (manually) added or removed from an object record. */
export interface Recordroletrigger {
  label?: string;
  active?: boolean;
  checksum?: string;
  source_code?: string;
  event_segment?: string;
  events?: string;
  order?: string;
  object: string;
}

/** Record triggers execute custom code when a record Event occurs. */
export interface Recordtrigger {
  label?: string;
  active?: boolean;
  checksum?: string;
  source_code?: string;
  event_segment?: string;
  events?: string;
  order?: string;
  object: string;
}

/** Record workflow actions execute custom code during object workflow events. */
export interface Recordworkflowaction {
  label?: string;
  active?: boolean;
  checksum?: string;
  source_code?: string;
  step_types?: string;
  object?: string;
}

/** Defines a Related Record Configuration for automatically creating related object records when source records enter specific lifecycle states. */
export interface Relatedrecordsetup {
  label?: string;
  active?: boolean;
  copy_attachments?: boolean;
  description?: string;
  outbound_reference_field: Relatedrecordsetup_Outbound_reference_field[];
  related_object: string;
  related_object_type: string;
  source_object: string;
  source_object_type: string;
}

export interface Relatedrecordsetup_Relatedrecordreferencemapping {
  label?: string;
  active?: boolean;
  source_object_field: Relatedrecordsetup_Relatedrecordreferencemapping_Source_object_field[];
  source_reference_field: Relatedrecordsetup_Relatedrecordreferencemapping_Source_reference_field[];
  source_reference_object: string;
  target_object_field: Relatedrecordsetup_Relatedrecordreferencemapping_Target_object_field[];
  target_reference_field: Relatedrecordsetup_Relatedrecordreferencemapping_Target_reference_field[];
  target_reference_object: string;
}

export interface Relatedrecordsetup_Relatedrecordfieldvaluemapping {
  label?: string;
  active?: boolean;
  mapped_source_object_field: Relatedrecordsetup_Relatedrecordfieldvaluemapping_Mapped_source_object_field[];
  related_object_field: Relatedrecordsetup_Relatedrecordfieldvaluemapping_Related_object_field[];
  mapped_source_reference_object_field?: Relatedrecordsetup_Relatedrecordfieldvaluemapping_Mapped_source_reference_object_field[];
  source_object: string;
  mapped_source_reference_object?: string;
  related_object: string;
}

/** Rendition profiles allow you to choose rendition settings that differ from the Vault-wide rendition settings. */
export interface Renditionprofile {
  label?: string;
  active?: boolean;
  bookmark_expansion_level: string;
  embed_fonts: string;
  generate_bookmarks_for_title_style: boolean;
  generate_caption_bookmarks_tree: boolean;
  generate_toc_lofs_bookmarks: boolean;
  include_metadata?: boolean;
  no_mark_up: boolean;
  pdf_format: string;
  protected_pdf?: boolean;
  render_ppt_speaker_notes: boolean;
  set_link_text_to_blue: boolean;
}

/** Renditions are files that users can add to documents. Rendition types describe the kind or purpose of the rendition. */
export interface Renditiontype {
  label?: string;
  active?: boolean;
  density?: string;
  format?: string;
  quality?: string;
  flatten_image?: boolean;
  include_source_doc_name?: boolean;
  automated_image?: boolean;
  colorspace?: string;
  transparency?: boolean;
  resize?: string;
  resample?: boolean;
}

/** Represents a Vault report. */
export interface Report {
  label: string;
  active: boolean;
  description?: string;
  format: string;
  definition: string;
  layout: string;
  content_type: string;
  report_type?: string;
  editors?: string;
  viewers?: string;
  owners?: string;
  class?: string;
  tags?: string;
}

/** Excel Report Templates allow users to create report templates in Microsoft Excel, upload them to Vault, and export report data into those templates. */
export interface Reportexceltemplate {
  label?: string;
  active?: boolean;
  filename: string;
  template_file: string;
}

/** Report Type determines the reporting objects for Vault reports. */
export interface Reporttype {
  label: string;
  active: boolean;
  description?: string;
  primary_object?: string;
  configuration: string;
  class?: string;
}

/** RIM Document Type Configurations define the document types, subtypes, and classifications available for extracting Health Authority Questions and Commitments. */
export interface Rimdoctypeconfig {
  label?: string;
  active?: boolean;
  document_type: string;
  enable_commitment_extraction?: boolean;
  enable_haq_panel?: boolean;
}

/** Event Change Details are the building blocks of Event Change Types in RIM Vaults. */
export interface Rimeventchangedetail {
  label?: string;
  active?: boolean;
  detail_description?: string;
  detail_object: string;
  detail_object_type?: string;
}

export interface Rimeventchangedetail_Rimeventchangedetailfield {
  label?: string;
  active?: boolean;
  detail_field_is_required?: boolean;
  detail_object: string;
  detail_field: Rimeventchangedetail_Rimeventchangedetailfield_Detail_field[];
  detail_field_order: number;
}

/** Event Change Types define the required and optional details for a specific kind of change in the Create Event Details wizard. */
export interface Rimeventchangetype {
  label?: string;
  active?: boolean;
  related_change_type_add: Rimeventchangetype_Related_change_type_add[];
  related_change_type_picklist: string;
  primary_event_change_detail: string;
  related_change_type_replacing: Rimeventchangetype_Related_change_type_replacing[];
  related_change_type_withdraw: Rimeventchangetype_Related_change_type_withdraw[];
  related_change_type_replaced_by: Rimeventchangetype_Related_change_type_replaced_by[];
  type_description?: string;
}

export interface Rimeventchangetype_Rimeventchangetypedetail {
  label?: string;
  active?: boolean;
  type_detail: string;
  detail_order: number;
  detail_is_required?: boolean;
}

/** Rimobjectconfigurations define properties for RIM objects used across Veeva Submissions and Veeva Registrations features. */
export interface Rimobjectconfiguration {
  label?: string;
  active?: boolean;
  object: string;
}

export interface Rimobjectconfiguration_Rimfieldconfig {
  label?: string;
  active?: boolean;
  create_display_order?: number;
  editable_application_detail?: boolean;
  editable_registered_detail?: boolean;
  field: Rimobjectconfiguration_Rimfieldconfig_Field[];
  object: string;
  unique_across_applications?: boolean;
  unique_within_application?: boolean;
  update_display_order?: number;
  use_in_comparison?: boolean;
}

/** RIM Object Mappings define which RIM objects and fields are mapped when data is moved between objects from various RIM processes and wizards. */
export interface Rimobjectmapping {
  label?: string;
  active?: boolean;
  source_object: string;
  target_object: string;
}

export interface Rimobjectmapping_Rimfieldmapping {
  label?: string;
  active?: boolean;
  exclude_from_matching?: boolean;
  source_field: Rimobjectmapping_Rimfieldmapping_Source_field[];
  source_object: string;
  target_field: Rimobjectmapping_Rimfieldmapping_Target_field[];
  target_object: string;
}

/** Role permissions enable Admins to assign permissions to a specific user without affecting security profiles. */
export interface Rolepermissionset {
  label?: string;
  active?: boolean;
  permission_sets: string;
  application_role: string;
}

/** Saved or custom views allow users to save or bookmark current searches, filters, and other view settings. */
export interface Savedview {
  label: string;
  description?: string;
  tab?: string;
  created_by?: string;
  icon?: string;
  search_criteria?: string;
  view_layout?: string;
  mandatory?: boolean;
  view_layout_type?: string;
  managed?: boolean;
  owner?: string;
  status?: string;
  vql_search_criteria?: string;
}

/** Job processors provide logic to process jobs in bulk. */
export interface Sdkjob {
  label?: string;
  active?: boolean;
  admin_cancellable?: boolean;
  admin_configurable?: boolean;
  checksum?: string;
  idempotent?: boolean;
  source_code?: string;
  visible?: boolean;
}

/** Definition of a searchable object Field. */
export interface Searchablefield {
  label?: string;
  active?: boolean;
  field?: Searchablefield_Field[];
  object?: string;
  required?: boolean;
}

/** Search collections allow you to create and edit groups of documents and objects that users can search at the same time as expanded searches. */
export interface Searchcollection {
  label?: string;
  active?: boolean;
  description?: string;
}

export interface Searchcollection_Searchcollectionsection {
  label?: string;
  active?: boolean;
  document_type?: string;
  include_on_tab?: boolean;
  object?: string;
  object_relationships?: string;
  object_type?: string;
  tab_assignments?: string;
  vql_criteria?: string;
}

/** Security profiles are how Vault applies permissions sets to individual users. */
export interface Securityprofile {
  permission_sets?: string;
  active: boolean;
  label: string;
  description?: string;
}

/** Specific to Clinical Operations Vaults, supports Quality Issues feature. Allows QC Specialists to correct specific errors while tracking Quality Issues. */
export interface Selfevidentcorrection {
  label?: string;
  active?: boolean;
  document_fields: string;
  lifecycle_state_type: string;
  qc_issue_type: string;
}

/** Vault uses Sharing Rules to manage users' roles on specific object records by matching rule criteria to specific user assignments. */
export interface Sharingrule {
  criteria?: string;
  description?: string;
  active: boolean;
  label: string;
  object: string;
}

export interface Sharingrule_Sharingrole {
  members?: string;
}

/** Signature pages display the manifested signatures for a Vault document. */
export interface Signaturepage {
  label?: string;
  active?: boolean;
  advanced_mode?: boolean;
  file?: string;
  footer_richtext?: string;
  header_richtext?: string;
  location?: string;
  schema_version?: number;
  signature_block_richtext?: string;
}

/** Specific to Quality family Vaults, supports Configurable Document Metadata feature of Station Manager application. */
export interface Stationconfiguration {
  label?: string;
  active?: boolean;
  description?: string;
}

export interface Stationconfiguration_Stationconfigurationdocfield {
  label?: string;
  active?: boolean;
  display_order: number;
  field: string;
}

/** A Tab within the UI of Vault. */
export interface Tab {
  label: string;
  order: number;
  dashboard?: string;
  object?: string;
  doctype?: string;
  subtype?: string;
  classification?: string;
  url?: string;
  send_session_via_post_message?: boolean;
  parent?: string;
  all_view_label?: string;
  my_view_label?: string;
  recent_view_label?: string;
  object_type?: string;
  active?: boolean;
  filter_criteria?: string;
  prevent_create?: boolean;
  modal_create_record?: boolean;
  page?: string;
  default_view?: string;
}

export interface Tab_Subtab {
  label: string;
  order: number;
  dashboard?: string;
  object?: string;
  doctype?: string;
  subtype?: string;
  classification?: string;
  url?: string;
  send_session_via_post_message?: boolean;
  parent?: string;
  object_type?: string;
  active?: boolean;
  filter_criteria?: string;
  prevent_create?: boolean;
  page?: string;
  default_view?: string;
}

/** Tab collections are groups of related tabs relevant to particular roles or tasks. */
export interface Tabcollection {
  label?: string;
  active?: boolean;
  help_content?: string;
}

export interface Tabcollection_Tabcollectionitem {
  label?: string;
  active?: boolean;
  tab?: string;
  tabs?: string;
  type: string;
  order: number;
}

/** Admins can create tags to label any document with a topic, category, or any other type of identification. */
export interface Tag {
  label: string;
  phrases?: string;
  auto_tag?: boolean;
  include_archive?: boolean;
  active: boolean;
}

/** Internal Vault developers use Tag Security Rules to control access to records. */
export interface Tagsecurityrule {
  label?: string;
  active?: boolean;
  help_content?: string;
  allow_custom_objects: boolean;
  description?: string;
}

export interface Tagsecurityrule_Tagsecurityruleassignment {
  label?: string;
  active?: boolean;
  object: string;
}

/** Allow users to apply custom business logic on objects when a record is created, updated, or deleted. */
export interface Userdefinedclass {
  label?: string;
  active?: boolean;
  checksum?: string;
  source_code?: string;
}

/** Allows users to create reusable data access objects, or models, and annotate their getters and setters as user-defined properties. */
export interface Userdefinedmodel {
  label?: string;
  active?: boolean;
  parent?: string;
  checksum?: string;
  source_code?: string;
  serialization_include?: string;
}

/** Allow users to wrap reusable logic into a service that can be used by other Vault Java SDK code. */
export interface Userdefinedservice {
  label?: string;
  active?: boolean;
  checksum?: string;
  interface?: string;
  source_code?: string;
}

/** This component type stores tokens which can be resolved Vault-wide. */
export interface Vaulttoken {
  label?: string;
  active?: boolean;
  clone_behavior: string;
  type: string;
  system_managed?: boolean;
  value?: string;
}

/** Visual Hierarchies form the backbone of the Process Navigator page in Veeva QualityDocs. Configurations determine which fields are available on Visual Hierarchy detail pages. */
export interface Visualhierarchyconfiguration {
  label?: string;
  active?: boolean;
  visual_hierarchy_type?: string;
}

export interface Visualhierarchyconfiguration_Visualhierarchyconfigurationfield {
  label?: string;
  active?: boolean;
  field: Visualhierarchyconfiguration_Visualhierarchyconfigurationfield_Field[];
  display_order: number;
  object: string;
}

/** Each Customwebapi must be assigned to a Webapigroup. */
export interface Webapigroup {
  label?: string;
  active?: boolean;
  description: string;
}

/** A workflow is a series of steps configured in Vault to align with a specific business purpose. */
export interface Workflow {
  label?: string;
  active?: boolean;
  lifecycle: string;
  description?: string;
  workflow_type?: string;
  bulk_email_message: string;
}

export interface Workflow_Workflowstep {
  label?: string;
  status?: boolean;
  description?: string;
  flow_type: string;
  next_steps?: string;
}
