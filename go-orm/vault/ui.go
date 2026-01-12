package vault

// =============================================================================
// TAB
// =============================================================================

// Tab represents a navigation tab (MDL: Tab)
type Tab struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`
	Order       int     `gorm:"default:0" json:"order" example:"1"`
	URL         *string `gorm:"size:1500" json:"url,omitempty"`

	// Target (object or document type)
	TargetObject  *string `gorm:"size:100" json:"object,omitempty"`
	TargetDoctype *string `gorm:"size:100" json:"doctype,omitempty"`

	// Child relationships
	Subtabs []Subtab `gorm:"foreignKey:TabID" json:"subtabs,omitempty"`
}

// TableName returns the table name for Tab
func (Tab) TableName() string {
	return "vault_tabs"
}

// Subtab represents a subtab within a tab (MDL: Subtab)
type Subtab struct {
	BaseModel
	ComponentBase

	// Parent tab
	TabID uint `gorm:"not null;index" json:"tab_id"`
	Tab   *Tab `gorm:"foreignKey:TabID" json:"-"`

	// Properties
	Order int     `gorm:"default:0" json:"order" example:"1"`
	URL   *string `gorm:"size:1500" json:"url,omitempty"`
}

// TableName returns the table name for Subtab
func (Subtab) TableName() string {
	return "vault_subtabs"
}

// =============================================================================
// TAB COLLECTION
// =============================================================================

// Tabcollection represents a tab collection (MDL: Tabcollection)
type Tabcollection struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Child relationships
	Items []TabcollectionItem `gorm:"foreignKey:CollectionID" json:"items,omitempty"`
}

// TableName returns the table name for Tabcollection
func (Tabcollection) TableName() string {
	return "vault_tabcollections"
}

// TabcollectionItem represents an item in a tab collection (MDL: Tabcollectionitem)
type TabcollectionItem struct {
	BaseModel

	// Parent collection
	CollectionID uint           `gorm:"not null;index" json:"collection_id"`
	Collection   *Tabcollection `gorm:"foreignKey:CollectionID" json:"-"`

	// Tab reference
	TabName string `gorm:"size:100;not null" json:"tab" example:"products__c"`
	Order   int    `gorm:"default:0" json:"order" example:"1"`
}

// TableName returns the table name for TabcollectionItem
func (TabcollectionItem) TableName() string {
	return "vault_tabcollection_items"
}

// =============================================================================
// DASHBOARD
// =============================================================================

// Dashboard represents a dashboard (MDL: Dashboard)
type Dashboard struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:255" json:"description,omitempty"`

	// Permissions
	Editors StringSlice `gorm:"type:json" json:"editors,omitempty"`
	Viewers StringSlice `gorm:"type:json" json:"viewers,omitempty"`
	Owners  StringSlice `gorm:"type:json" json:"owners,omitempty"`

	// Dashboard definition
	DashboardMarkup XMLString `gorm:"type:text" json:"dashboard_markup"`

	// Tags
	Tags StringSlice `gorm:"type:json" json:"tags,omitempty"`
}

// TableName returns the table name for Dashboard
func (Dashboard) TableName() string {
	return "vault_dashboards"
}

// =============================================================================
// REPORT
// =============================================================================

// Report represents a report (MDL: Report)
type Report struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Report type reference
	ReportTypeName *string `gorm:"size:100" json:"report_type,omitempty"`
}

// TableName returns the table name for Report
func (Report) TableName() string {
	return "vault_reports"
}

// Reporttype represents a report type (MDL: Reporttype)
type Reporttype struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`
}

// TableName returns the table name for Reporttype
func (Reporttype) TableName() string {
	return "vault_reporttypes"
}

// =============================================================================
// PAGE - Custom Pages
// =============================================================================

// Page represents a custom page (MDL: Page)
type Page struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`
	URL         *string `gorm:"size:1500" json:"url,omitempty"`
}

// TableName returns the table name for Page
func (Page) TableName() string {
	return "vault_pages"
}

// Pagelayout represents a page layout (MDL: Pagelayout)
type Pagelayout struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Target object
	ObjectName *string `gorm:"size:100" json:"object,omitempty"`
}

// TableName returns the table name for Pagelayout
func (Pagelayout) TableName() string {
	return "vault_pagelayouts"
}

// Pagelink represents a page link (MDL: Pagelink)
type Pagelink struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`
	URL         string  `gorm:"size:1500;not null" json:"url"`
}

// TableName returns the table name for Pagelink
func (Pagelink) TableName() string {
	return "vault_pagelinks"
}

// =============================================================================
// LINK - Web Actions
// =============================================================================

// Link represents a web action link (MDL: Link)
type Link struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`
	URL         string  `gorm:"size:1500;not null" json:"url"`
}

// TableName returns the table name for Link
func (Link) TableName() string {
	return "vault_links"
}

// =============================================================================
// SAVED VIEW
// =============================================================================

// Savedview represents a saved view (MDL: Savedview)
type Savedview struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Target object
	ObjectName *string `gorm:"size:100" json:"object,omitempty"`
}

// TableName returns the table name for Savedview
func (Savedview) TableName() string {
	return "vault_savedviews"
}

// =============================================================================
// SEARCH COLLECTION
// =============================================================================

// Searchcollection represents a search collection (MDL: Searchcollection)
type Searchcollection struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Child relationships
	Sections []SearchcollectionSection `gorm:"foreignKey:CollectionID" json:"sections,omitempty"`
}

// TableName returns the table name for Searchcollection
func (Searchcollection) TableName() string {
	return "vault_searchcollections"
}

// SearchcollectionSection represents a section in a search collection (MDL: Searchcollectionsection)
type SearchcollectionSection struct {
	BaseModel
	ComponentBase

	// Parent collection
	CollectionID uint              `gorm:"not null;index" json:"collection_id"`
	Collection   *Searchcollection `gorm:"foreignKey:CollectionID" json:"-"`

	// Properties
	Order int `gorm:"default:0" json:"order" example:"1"`
}

// TableName returns the table name for SearchcollectionSection
func (SearchcollectionSection) TableName() string {
	return "vault_searchcollection_sections"
}

// =============================================================================
// TAG
// =============================================================================

// Tag represents a tag (MDL: Tag)
type Tag struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:255" json:"description,omitempty"`
}

// TableName returns the table name for Tag
func (Tag) TableName() string {
	return "vault_tags"
}

// =============================================================================
// LAYOUT PROFILE
// =============================================================================

// Layoutprofile represents a layout profile (MDL: Layoutprofile)
type Layoutprofile struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`

	// Child relationships
	ObjectLayouts []Objectlayout `gorm:"foreignKey:ProfileID" json:"object_layouts,omitempty"`
}

// TableName returns the table name for Layoutprofile
func (Layoutprofile) TableName() string {
	return "vault_layoutprofiles"
}

// Objectlayout represents an object layout in a profile (MDL: Objectlayout)
type Objectlayout struct {
	BaseModel
	ComponentBase

	// Parent profile
	ProfileID uint           `gorm:"not null;index" json:"profile_id"`
	Profile   *Layoutprofile `gorm:"foreignKey:ProfileID" json:"-"`

	// Properties
	ObjectName string `gorm:"size:100;not null" json:"object" example:"product__c"`
}

// TableName returns the table name for Objectlayout
func (Objectlayout) TableName() string {
	return "vault_objectlayouts"
}

// =============================================================================
// LABEL SET
// =============================================================================

// Labelset represents a label set for translations (MDL: Labelset)
type Labelset struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`
}

// TableName returns the table name for Labelset
func (Labelset) TableName() string {
	return "vault_labelsets"
}

// =============================================================================
// RENDITION PROFILE
// =============================================================================

// Renditionprofile represents a rendition profile (MDL: Renditionprofile)
type Renditionprofile struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`
}

// TableName returns the table name for Renditionprofile
func (Renditionprofile) TableName() string {
	return "vault_renditionprofiles"
}

// Renditiontype represents a rendition type (MDL: Renditiontype)
type Renditiontype struct {
	BaseModel
	ComponentBase

	// Properties
	Description *string `gorm:"size:500" json:"description,omitempty"`
}

// TableName returns the table name for Renditiontype
func (Renditiontype) TableName() string {
	return "vault_renditiontypes"
}
