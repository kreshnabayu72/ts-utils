package new_model

type AssetStatus string

const (
	AssetStatusNonTransactional   AssetStatus = "Non-Transactional"
	AssetStatusAvailable          AssetStatus = "Available"
	AssetStatusAssigned           AssetStatus = "Assigned"
	AssetStatusBorrowed           AssetStatus = "Borrowed"
	AssetStatusDisposed           AssetStatus = "Disposed"
	AssetStatusDamaged            AssetStatus = "Damaged"
	AssetStatusMissing            AssetStatus = "Missing"
	AssetStatusOnAssignProcess    AssetStatus = "On Assign Process"
	AssetStatusOnBorrowingProcess AssetStatus = "On Borrowing Process"
	AssetStatusOnDisposalProcess  AssetStatus = "On Disposal Process"
	AssetStatusReportedDisposal   AssetStatus = "Reported Disposal"
	AssetStatusOnTransfer         AssetStatus = "On Transfer"
	AssetStatusExpiredLicense     AssetStatus = "Expired License"
	AssetStatusTAGReported        AssetStatus = "TAG Reported"
	AssetStatusTAGNotSet          AssetStatus = "TAG Not Set"
)

type CustomFieldType string

const (
	CustomFieldTypeGeneralInformation = CustomFieldType("General Information")
	CustomFieldTypePurchase           = CustomFieldType("Purchase")
	CustomFieldTypeAccounting         = CustomFieldType("Accounting")
	CustomFieldTypeMaintenanceRoutine = CustomFieldType("Maintenance Routine")
	CustomFieldTypeAudit              = CustomFieldType("Audit")
)

type CustomFieldDataType string

const (
	CustomFieldDataTypeText          = CustomFieldDataType("Text")
	CustomFieldDataTypeTextarea      = CustomFieldDataType("Textarea")
	CustomFieldDataTypeDropdown      = CustomFieldDataType("Dropdown")
	CustomFieldDataTypeMultidropdown = CustomFieldDataType("Multidropdown")
	CustomFieldDataTypeRadioButton   = CustomFieldDataType("Radio button")
	CustomFieldDataTypeCheckbox      = CustomFieldDataType("Checkbox")
	CustomFieldDataTypeDate          = CustomFieldDataType("Date")
	CustomFieldDataTypeDatetime      = CustomFieldDataType("Datetime")
	CustomFieldDataTypeCurrency      = CustomFieldDataType("Currency")
	CustomFieldDataTypeNumeric       = CustomFieldDataType("Numeric")
	CustomFieldDataTypePhone         = CustomFieldDataType("Phone")
	CustomFieldDataTypePercentage    = CustomFieldDataType("Percentage")
	CustomFieldDataTypeEmail         = CustomFieldDataType("Email")
	CustomFieldDataTypeURL           = CustomFieldDataType("URL")
	CustomFieldDataTypeDocument      = CustomFieldDataType("Document")
)

type TagType string

const (
	TagTypeQR     = TagType("QR")
	TagTypeRFID   = TagType("RFID")
	TagTypeRFIDQR = TagType("RFID & QR")
	TagTypeNonTAG = TagType("Non TAG")
)

type SystemRoleName string

const (
	SystemRoleNameImportData                SystemRoleName = "Import data"
	SystemRoleNameManageAssetAttribute      SystemRoleName = "Manage asset names, brands, categories, and custom fields"
	SystemRoleNameManageAssetPolicy         SystemRoleName = "Manage asset policy"
	SystemRoleNameManageDepreciationGroup   SystemRoleName = "Manage depreciation group"
	SystemRoleNameManageDepreciationMethod  SystemRoleName = "Manage depreciation method"
	SystemRoleNameManageDocumentDeletion    SystemRoleName = "Manage document deletion"
	SystemRoleNameManageGroup               SystemRoleName = "Manage group"
	SystemRoleNameManageUserAndRole         SystemRoleName = "Manage user and role"
	SystemRoleNameManageIoTAndReader        SystemRoleName = "Manage IOT and reader"
	SystemRoleNameManageLicense             SystemRoleName = "Manage license"
	SystemRoleNameManagePurchaseDocument    SystemRoleName = "Manage purchase document"
	SystemRoleNameManagePurchaseInformation SystemRoleName = "Manage purchase information"
	SystemRoleNameManageTag                 SystemRoleName = "Manage TAG"
	SystemRoleNameRegisterAssets            SystemRoleName = "Register assets"
	SystemRoleNameManageTagType             SystemRoleName = "Manage TAG type"
	SystemRoleNameOpenAPI                   SystemRoleName = "Open API"
)

type SystemRoleAttribute string

const (
	SystemRoleAttributeImportData         SystemRoleAttribute = "importData"
	SystemRoleAttributeAssetAttribute     SystemRoleAttribute = "assetAttribute"
	SystemRoleAttributeAssetPolicies      SystemRoleAttribute = "assetPolicies"
	SystemRoleAttributeDepreciationGroup  SystemRoleAttribute = "depreciationGroup"
	SystemRoleAttributeDepreciationMethod SystemRoleAttribute = "depreciationMethod"
	SystemRoleAttributeDocumentDeletion   SystemRoleAttribute = "documentDeletion"
	SystemRoleAttributeGroup              SystemRoleAttribute = "group"
	SystemRoleAttributeUserAndRole        SystemRoleAttribute = "userAndRole"
	SystemRoleAttributeIoTReader          SystemRoleAttribute = "iotReader"
	SystemRoleAttributeLicense            SystemRoleAttribute = "license"
	SystemRoleAttributePurchaseDocument   SystemRoleAttribute = "purchaseDocument"
	SystemRoleAttributePurchaseInfo       SystemRoleAttribute = "purchaseInformation"
	SystemRoleAttributeTag                SystemRoleAttribute = "tag"
	SystemRoleAttributeRegisterAsset      SystemRoleAttribute = "registerAsset"
	SystemRoleAttributeTagType            SystemRoleAttribute = "tagType"
	SystemRoleAttributeOpenAPI            SystemRoleAttribute = "openAPI"
)

type TransactionRoleAttribute string

const (
	TransactionRoleAttributeBorrowing    TransactionRoleAttribute = "borrowingRole"
	TransactionRoleAttributeDisposal     TransactionRoleAttribute = "disposalRole"
	TransactionRoleAttributeTransfer     TransactionRoleAttribute = "transferRole"
	TransactionRoleAttributeAudit        TransactionRoleAttribute = "auditRole"
	TransactionRoleAttributeMaintenance  TransactionRoleAttribute = "maintenanceRole"
	TransactionRoleAttributeRepair       TransactionRoleAttribute = "repairRole"
	TransactionRoleAttributeAssetControl TransactionRoleAttribute = "assetControlRole"
	TransactionRoleAttributeTracking     TransactionRoleAttribute = "trackingRole"
	TransactionRoleAttributeDamaged      TransactionRoleAttribute = "damagedRole"
	TransactionRoleAttributeMissing      TransactionRoleAttribute = "missingRole"
)

type TransactionRoleName string

const (
	TransactionRoleNameBorrowingAssignment TransactionRoleName = "Borrowing/Assignment"
	TransactionRoleNameDisposal            TransactionRoleName = "Disposal"
	TransactionRoleNameTransfer            TransactionRoleName = "Transfer"
	TransactionRoleNameAudit               TransactionRoleName = "Audit"
	TransactionRoleNameMaintenanceRoutine  TransactionRoleName = "Maintenance Routine"
	TransactionRoleNameRepairTicketing     TransactionRoleName = "Repair Ticketing"
	TransactionRoleNameAssetControl        TransactionRoleName = "Asset Control"
	TransactionRoleNameTracking            TransactionRoleName = "Tracking"
	TransactionRoleNameDamaged             TransactionRoleName = "Damaged"
	TransactionRoleNameMissing             TransactionRoleName = "Missing"
)

type TransactionRoleType string

const (
	TransactionRoleTypeManager    TransactionRoleType = "Manager"
	TransactionRoleTypeMonitoring TransactionRoleType = "Monitoring"
	TransactionRoleTypeApproval   TransactionRoleType = "Approval"
	TransactionRoleTypeStaff      TransactionRoleType = "Staff"
)

type PlanTypes string

const (
	PlanBasic      = PlanTypes("Basic")
	PlanStandard   = PlanTypes("Standard")
	PlanEnterprise = PlanTypes("Enterprise")
)

var PlanEnum = map[string]PlanTypes{
	"Basic":      PlanBasic,
	"Standard":   PlanStandard,
	"Enterprise": PlanEnterprise,
}

type CompanyStatus string

const (
	CompanyStatusActive   CompanyStatus = "Active"
	CompanyStatusExpired  CompanyStatus = "Expired"
	CompanyStatusArchived CompanyStatus = "Archived"
)

type CompanyOpenAPIStatus string

const (
	CompanyOpenAPIStatusApproved CompanyOpenAPIStatus = "Approved"
	CompanyOpenAPIStatusWaiting  CompanyOpenAPIStatus = "Waiting for Approval"
)

type CompanyDepreciationMethod string

const (
	CompanyDepreciationMethodStraightLine           CompanyDepreciationMethod = "Straight Line"
	CompanyDepreciationMethodDoubleDecliningBalance CompanyDepreciationMethod = "Double-Declining Balance"
)

type AttachmentType string

const (
	AttachmentTypePurchase AttachmentType = "Purchase"
	AttachmentTypeProduct  AttachmentType = "Product"
)

type FileType string

const (
	FileTypeAttachment         FileType = "Attachment"
	FileTypeDamage             FileType = "Damage"
	FileTypeDisposal           FileType = "Disposal"
	FileTypeRepairTicketing    FileType = "Repair Ticketing"
	FileTypeMaintenanceRoutine FileType = "Maintenance Routine"
)
