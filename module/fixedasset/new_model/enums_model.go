package new_model

type AssetDefaultImage string

const (
	AssetDefaultImageFirstImage  = AssetDefaultImage("firstImage")
	AssetDefaultImageSecondImage = AssetDefaultImage("secondImage")
)

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
	AssetStatusNotPairedYet       AssetStatus = "Not Paired Yet"
)

type AssetLicenseStatus string

const (
	AssetLicenseStatusManaged           AssetLicenseStatus = "Managed"
	AssetLicenseStatusUnmanaged         AssetLicenseStatus = "Unmanaged"
	AssetLicenseStatusPendingActivation AssetLicenseStatus = "Pending Activation"
)

type AssetLicenseType string

const (
	AssetLicenseTypeFixedAsset      AssetLicenseType = "Fixed Asset"
	AssetLicenseTypeMaintenance     AssetLicenseType = "Maintenance"
	AssetLicenseTypeTracking        AssetLicenseType = "Tracking"
	AssetLicenseTypeAssetControl    AssetLicenseType = "Asset Control"
	AssetLicenseTypeAudit           AssetLicenseType = "Audit"
	AssetLicenseTypeRepairTicketing AssetLicenseType = "Repair Ticketing"
)

type AuditType string

const (
	AuditTypeExistence AuditType = "Existence"
	AuditTypeCondition AuditType = "Condition"
)

type AuditTaskAssetStatus string

const (
	AuditTaskAssetStatusInAudit     AuditTaskAssetStatus = "In Audit"
	AuditTaskAssetStatusAudited     AuditTaskAssetStatus = "Audited"
	AuditTaskAssetStatusTAGReported AuditTaskAssetStatus = "TAG Reported"
	AuditTaskAssetStatusOverdue     AuditTaskAssetStatus = "Overdue"
	AuditTaskAssetStatusMisplaced   AuditTaskAssetStatus = "Misplaced"
	AuditTaskAssetStatusFound       AuditTaskAssetStatus = "Found"
)

type AuditTaskAssetCondition string

const (
	AuditTaskAssetConditionNormal               AuditTaskAssetCondition = "Normal"
	AuditTaskAssetConditionReportedDamaged      AuditTaskAssetCondition = "Reported Damaged"
	AuditTaskAssetConditionReportedMissing      AuditTaskAssetCondition = "Reported Missing"
	AuditTaskAssetConditionTAGFoundAssetMissing AuditTaskAssetCondition = "TAG Found Asset Missing"
)

type AuditTaskStatus string

const (
	AuditTaskStatusOnSchedule         AuditTaskStatus = "On Schedule"
	AuditTaskStatusInAudit            AuditTaskStatus = "In Audit"
	AuditTaskStatusWaitingForApproval AuditTaskStatus = "Waiting for Approval"
	AuditTaskStatusGracePeriod        AuditTaskStatus = "Grace Period"
	AuditTaskStatusOverdue            AuditTaskStatus = "Overdue"
	AuditTaskStatusCompleted          AuditTaskStatus = "Completed"
	AuditTaskStatusCanceled           AuditTaskStatus = "Canceled"
	AuditTaskStatusStopped            AuditTaskStatus = "Stopped"
)

type AuditOtherTransaction string

const (
	AuditOtherTransactionTransfer  AuditOtherTransaction = "Transfer"
	AuditOtherTransactionDisposal  AuditOtherTransaction = "Disposal"
	AuditOtherTransactionTAGChange AuditOtherTransaction = "TAG Change"
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
	CustomFieldDataTypeMultiDropdown = CustomFieldDataType("MultiDropdown")
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

type GroupQuotaType string

const (
	GroupQuotaTypeFixedAsset      = GroupQuotaType("Fixed Asset")
	GroupQuotaTypeRepairTicketing = GroupQuotaType("Repair Ticketing")
	GroupQuotaTypeTracking        = GroupQuotaType("Tracking")
	GroupQuotaTypeAudit           = GroupQuotaType("Audit")
	GroupQuotaTypeMaintenance     = GroupQuotaType("Maintenance")
	GroupQuotaTypeAssetControl    = GroupQuotaType("Asset Control")
)

var GroupQuotaTypesArray = []string{
	string(GroupQuotaTypeFixedAsset),
	string(GroupQuotaTypeRepairTicketing),
	string(GroupQuotaTypeTracking),
	string(GroupQuotaTypeAudit),
	string(GroupQuotaTypeMaintenance),
	string(GroupQuotaTypeAssetControl),
}

type GroupTransactionType string

const (
	GroupTransactionTypeBorrowingAssignment = GroupTransactionType("Borrowing/Assignment")
	GroupTransactionTypeTransfer            = GroupTransactionType("Transfer")
	GroupTransactionTypeDisposal            = GroupTransactionType("Disposal")
	GroupTransactionTypeDamage              = GroupTransactionType("Damage")
	GroupTransactionTypeMissing             = GroupTransactionType("Missing")
	GroupTransactionTypeRepairTicketing     = GroupTransactionType("Repair Ticketing")
	GroupTransactionTypeTracking            = GroupTransactionType("Tracking")
	GroupTransactionTypeAudit               = GroupTransactionType("Audit")
	GroupTransactionTypeMaintenance         = GroupTransactionType("Maintenance")
	GroupTransactionTypeAssetControl        = GroupTransactionType("Asset Control")
)

var GroupTransactionTypesArray = []string{
	string(GroupTransactionTypeBorrowingAssignment),
	string(GroupTransactionTypeTransfer),
	string(GroupTransactionTypeDisposal),
	string(GroupTransactionTypeDamage),
	string(GroupTransactionTypeMissing),
	string(GroupTransactionTypeRepairTicketing),
	string(GroupTransactionTypeTracking),
	string(GroupTransactionTypeAudit),
	string(GroupTransactionTypeMaintenance),
	string(GroupTransactionTypeAssetControl),
}

type TagType string

const (
	TagTypeQR     = TagType("QR")
	TagTypeRFID   = TagType("RFID")
	TagTypeRFIDQR = TagType("RFID & QR")
	TagTypeNonTAG = TagType("Non TAG")
)

type ReplaceTagPrelistType string

const (
	ReplaceTagPrelistTypeRFID   = TagType("Replace RFID TAG")
	ReplaceTagPrelistTypeQR     = TagType("Replace QR TAG")
	ReplaceTagPrelistTypeRFIDQR = TagType("Replace RFID & QR TAG")
)

type Permission struct {
	Create bool
	View   bool
	Update bool
	Delete bool
}

type PermissionType string

const (
	PermissionCreate PermissionType = "create"
	PermissionUpdate PermissionType = "update"
	PermissionDelete PermissionType = "delete"
	PermissionView   PermissionType = "view"
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

func GetSystemRoleNameFromAttribute(attr SystemRoleAttribute) SystemRoleName {
	switch attr {
	case SystemRoleAttributeImportData:
		return SystemRoleNameImportData
	case SystemRoleAttributeAssetAttribute:
		return SystemRoleNameManageAssetAttribute
	case SystemRoleAttributeAssetPolicies:
		return SystemRoleNameManageAssetPolicy
	case SystemRoleAttributeDepreciationGroup:
		return SystemRoleNameManageDepreciationGroup
	case SystemRoleAttributeDepreciationMethod:
		return SystemRoleNameManageDepreciationMethod
	case SystemRoleAttributeDocumentDeletion:
		return SystemRoleNameManageDocumentDeletion
	case SystemRoleAttributeGroup:
		return SystemRoleNameManageGroup
	case SystemRoleAttributeUserAndRole:
		return SystemRoleNameManageUserAndRole
	case SystemRoleAttributeIoTReader:
		return SystemRoleNameManageIoTAndReader
	case SystemRoleAttributeLicense:
		return SystemRoleNameManageLicense
	case SystemRoleAttributePurchaseDocument:
		return SystemRoleNameManagePurchaseDocument
	case SystemRoleAttributePurchaseInfo:
		return SystemRoleNameManagePurchaseInformation
	case SystemRoleAttributeTag:
		return SystemRoleNameManageTag
	case SystemRoleAttributeRegisterAsset:
		return SystemRoleNameRegisterAssets
	case SystemRoleAttributeTagType:
		return SystemRoleNameManageTagType
	case SystemRoleAttributeOpenAPI:
		return SystemRoleNameOpenAPI
	default:
		return SystemRoleName(attr) // fallback: cast raw string to SystemRoleName
	}
}

type SystemRoleType string

const (
	SystemRoleTypeGroup  SystemRoleType = "Group"
	SystemRoleTypeGlobal SystemRoleType = "Global"
)

var SystemRoleTypeEnum = map[string]SystemRoleType{
	"Group":  SystemRoleTypeGroup,
	"Global": SystemRoleTypeGlobal,
}

func GetSystemRoleTypeFromAttribute(attr SystemRoleAttribute) SystemRoleType {
	switch attr {
	case
		SystemRoleAttributeImportData,
		SystemRoleAttributeAssetAttribute,
		SystemRoleAttributeAssetPolicies,
		SystemRoleAttributeDepreciationGroup,
		SystemRoleAttributeDepreciationMethod,
		SystemRoleAttributeTagType,
		SystemRoleAttributeOpenAPI:
		return SystemRoleTypeGlobal
	default:
		return SystemRoleTypeGroup
	}
}

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

var SystemRoleAttributeEnum = map[string]SystemRoleAttribute{
	"importData":          SystemRoleAttributeImportData,
	"assetAttribute":      SystemRoleAttributeAssetAttribute,
	"assetPolicies":       SystemRoleAttributeAssetPolicies,
	"depreciationGroup":   SystemRoleAttributeDepreciationGroup,
	"depreciationMethod":  SystemRoleAttributeDepreciationMethod,
	"documentDeletion":    SystemRoleAttributeDocumentDeletion,
	"group":               SystemRoleAttributeGroup,
	"userAndRole":         SystemRoleAttributeUserAndRole,
	"iotReader":           SystemRoleAttributeIoTReader,
	"license":             SystemRoleAttributeLicense,
	"purchaseDocument":    SystemRoleAttributePurchaseDocument,
	"purchaseInformation": SystemRoleAttributePurchaseInfo,
	"tag":                 SystemRoleAttributeTag,
	"registerAsset":       SystemRoleAttributeRegisterAsset,
	"tagType":             SystemRoleAttributeTagType,
	"openAPI":             SystemRoleAttributeOpenAPI,
}

var SystemRoleAttributeArray = []string{
	string(SystemRoleAttributeImportData),
	string(SystemRoleAttributeAssetAttribute),
	string(SystemRoleAttributeAssetPolicies),
	string(SystemRoleAttributeDepreciationGroup),
	string(SystemRoleAttributeDepreciationMethod),
	string(SystemRoleAttributeDocumentDeletion),
	string(SystemRoleAttributeGroup),
	string(SystemRoleAttributeUserAndRole),
	string(SystemRoleAttributeIoTReader),
	string(SystemRoleAttributeLicense),
	string(SystemRoleAttributePurchaseDocument),
	string(SystemRoleAttributePurchaseInfo),
	string(SystemRoleAttributeTag),
	string(SystemRoleAttributeRegisterAsset),
	string(SystemRoleAttributeTagType),
	string(SystemRoleAttributeOpenAPI),
}

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

var TransactionRoleAttributeEnum = map[string]TransactionRoleAttribute{
	"borrowingRole":    TransactionRoleAttributeBorrowing,
	"disposalRole":     TransactionRoleAttributeDisposal,
	"transferRole":     TransactionRoleAttributeTransfer,
	"auditRole":        TransactionRoleAttributeAudit,
	"maintenanceRole":  TransactionRoleAttributeMaintenance,
	"repairRole":       TransactionRoleAttributeRepair,
	"assetControlRole": TransactionRoleAttributeAssetControl,
	"trackingRole":     TransactionRoleAttributeTracking,
	"damagedRole":      TransactionRoleAttributeDamaged,
	"missingRole":      TransactionRoleAttributeMissing,
}

var TransactionRoleAttributeArray = []string{
	string(TransactionRoleAttributeBorrowing),
	string(TransactionRoleAttributeDisposal),
	string(TransactionRoleAttributeTransfer),
	string(TransactionRoleAttributeAudit),
	string(TransactionRoleAttributeMaintenance),
	string(TransactionRoleAttributeRepair),
	string(TransactionRoleAttributeAssetControl),
	string(TransactionRoleAttributeTracking),
	string(TransactionRoleAttributeDamaged),
	string(TransactionRoleAttributeMissing),
}

var TransactionRoleNameArray = []string{
	string(TransactionRoleNameBorrowingAssignment),
	string(TransactionRoleNameDisposal),
	string(TransactionRoleNameTransfer),
	string(TransactionRoleNameAudit),
	string(TransactionRoleNameMaintenanceRoutine),
	string(TransactionRoleNameRepairTicketing),
	string(TransactionRoleNameAssetControl),
	string(TransactionRoleNameTracking),
	string(TransactionRoleNameDamaged),
	string(TransactionRoleNameMissing),
}

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

func GetTransactionRoleNameFromAttribute(attr TransactionRoleAttribute) TransactionRoleName {
	switch attr {
	case TransactionRoleAttributeBorrowing:
		return TransactionRoleNameBorrowingAssignment
	case TransactionRoleAttributeDisposal:
		return TransactionRoleNameDisposal
	case TransactionRoleAttributeTransfer:
		return TransactionRoleNameTransfer
	case TransactionRoleAttributeAudit:
		return TransactionRoleNameAudit
	case TransactionRoleAttributeMaintenance:
		return TransactionRoleNameMaintenanceRoutine
	case TransactionRoleAttributeRepair:
		return TransactionRoleNameRepairTicketing
	case TransactionRoleAttributeAssetControl:
		return TransactionRoleNameAssetControl
	case TransactionRoleAttributeTracking:
		return TransactionRoleNameTracking
	case TransactionRoleAttributeDamaged:
		return TransactionRoleNameDamaged
	case TransactionRoleAttributeMissing:
		return TransactionRoleNameMissing
	default:
		return TransactionRoleName(attr) // fallback: cast attribute as-is
	}
}

func GetTransactionRoleAttributeFromName(name TransactionRoleName) TransactionRoleAttribute {
	switch name {
	case TransactionRoleNameBorrowingAssignment:
		return TransactionRoleAttributeBorrowing
	case TransactionRoleNameDisposal:
		return TransactionRoleAttributeDisposal
	case TransactionRoleNameTransfer:
		return TransactionRoleAttributeTransfer
	case TransactionRoleNameAudit:
		return TransactionRoleAttributeAudit
	case TransactionRoleNameMaintenanceRoutine:
		return TransactionRoleAttributeMaintenance
	case TransactionRoleNameRepairTicketing:
		return TransactionRoleAttributeRepair
	case TransactionRoleNameAssetControl:
		return TransactionRoleAttributeAssetControl
	case TransactionRoleNameTracking:
		return TransactionRoleAttributeTracking
	case TransactionRoleNameDamaged:
		return TransactionRoleAttributeDamaged
	case TransactionRoleNameMissing:
		return TransactionRoleAttributeMissing
	default:
		return TransactionRoleAttribute(name) // fallback: cast name as-is
	}
}

type TransactionRoleType string

const (
	TransactionRoleTypeManager    TransactionRoleType = "Manager"
	TransactionRoleTypeMonitoring TransactionRoleType = "Monitoring"
	TransactionRoleTypeApproval   TransactionRoleType = "Approval"
	TransactionRoleTypeStaff      TransactionRoleType = "Staff"
)

var TransactionRoleTypeEnum = map[string]TransactionRoleType{
	"Manager":    TransactionRoleTypeManager,
	"Monitoring": TransactionRoleTypeMonitoring,
	"Approval":   TransactionRoleTypeApproval,
	"Staff":      TransactionRoleTypeStaff,
}

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

type TagStatus string

const (
	TagStatusAvailable      = TagStatus("Available")
	TagStatusPaired         = TagStatus("Paired")
	TagStatusReported       = TagStatus("Reported")
	TagStatusDamagedMissing = TagStatus("Damaged / Missing")
	TagStatusToBeReturned   = TagStatus("To be Returned")
)

type ModuleAccess string

const (
	ModuleAccessFixedAsset     = ModuleAccess("Fixed Asset")
	ModuleAccessGlobalSettings = ModuleAccess("Global Settings")
	ModuleAccessAdminConsole   = ModuleAccess("Admin Console")
	ModuleAccessSupplyAsset    = ModuleAccess("Supply Asset")
)

var ModuleAccessEnum = map[string]ModuleAccess{
	"Fixed Asset":     ModuleAccessFixedAsset,
	"Global Settings": ModuleAccessGlobalSettings,
	"Admin Console":   ModuleAccessAdminConsole,
	"Supply Asset":    ModuleAccessSupplyAsset,
}

type LicenseModuleType string

const (
	LicenseModuleTypeFixedAsset         = LicenseModuleType("Fixed Asset")
	LicenseModuleTypeAssetControl       = LicenseModuleType("+Asset Control")
	LicenseModuleTypeMaintenanceRoutine = LicenseModuleType("+Maintenance Routine")
	LicenseModuleTypeRepairTicketing    = LicenseModuleType("+Repair Ticketing")
	LicenseModuleTypeTracking           = LicenseModuleType("+Tracking")
)

type LicensePurchaseType string

const (
	LicensePurchaseTypePurchase = LicensePurchaseType("Purchase")
	LicensePurchaseTypeTrial    = LicensePurchaseType("Trial")
)

type LicenseUseType string

const (
	LicenseUseTypeExtend     = LicenseUseType("Extend")
	LicenseUseTypeAdditional = LicenseUseType("Additional")
	LicenseUseTypeNew        = LicenseUseType("New")
)

type ReaderStatus string

const (
	ReaderStatusAvailable = ReaderStatus("Available")
	ReaderStatusDamaged   = ReaderStatus("Damaged")
	ReaderStatusMissing   = ReaderStatus("Missing")
)

type DamageReportStatus string

const (
	DamageReportStatusDamaged           = DamageReportStatus("Damaged")
	DamageReportStatusRepaired          = DamageReportStatus("Repaired")
	DamageReportStatusReportedDisposal  = DamageReportStatus("Reported Disposal")
	DamageReportStatusOnDisposalProcess = DamageReportStatus("On Disposal Process")
	DamageReportStatusDisposed          = DamageReportStatus("Disposed")
)

type MissingReportStatus string

const (
	MissingReportStatusMissing           MissingReportStatus = "Missing"
	MissingReportStatusFound             MissingReportStatus = "Found"
	MissingReportStatusReportedDisposal  MissingReportStatus = "Reported Disposal"
	MissingReportStatusOnDisposalProcess MissingReportStatus = "On Disposal Process"
	MissingReportStatusDisposed          MissingReportStatus = "Disposed"
)

type AssetReportSource string

const (
	AssetReportSourceAvailable   AssetReportSource = "Available Asset"
	AssetReportSourceAudit       AssetReportSource = "Audit Asset"
	AssetReportSourceBorrowed    AssetReportSource = "Borrowed Asset"
	AssetReportSourceAssigned    AssetReportSource = "Assigned Asset"
	AssetReportSourceTransfer    AssetReportSource = "Transfer Asset"
	AssetReportSourceTracking    AssetReportSource = "Tracking Asset"
	AssetReportSourceMaintenance AssetReportSource = "Maintenance Asset"
)

type AssignmentTransactionStatus string

const (
	AssignmentTransactionStatusWaitingForApproval     AssignmentTransactionStatus = "Waiting for Approval"
	AssignmentTransactionStatusWaitingForHandover     AssignmentTransactionStatus = "Waiting for Handover"
	AssignmentTransactionStatusWaitingForConfirmation AssignmentTransactionStatus = "Waiting for Confirmation"
	AssignmentTransactionStatusWaitingForVerification AssignmentTransactionStatus = "Waiting for Verification"
	AssignmentTransactionStatusAssigned               AssignmentTransactionStatus = "Assigned"
	AssignmentTransactionStatusCanceled               AssignmentTransactionStatus = "Canceled"
	AssignmentTransactionStatusUnassigned             AssignmentTransactionStatus = "Unassigned"
	AssignmentTransactionStatusRejected               AssignmentTransactionStatus = "Rejected"
)

type AssignmentRequestStatus string

const (
	AssignmentRequestStatusWaitingForApproval     AssignmentRequestStatus = "Waiting for Approval"
	AssignmentRequestStatusWaitingForHandover     AssignmentRequestStatus = "Waiting for Handover"
	AssignmentRequestStatusWaitingForConfirmation AssignmentRequestStatus = "Waiting for Confirmation"
	AssignmentRequestStatusWaitingForVerification AssignmentRequestStatus = "Waiting for Verification"
	AssignmentRequestStatusAssigned               AssignmentRequestStatus = "Assigned"
	AssignmentRequestStatusCanceled               AssignmentRequestStatus = "Canceled"
	AssignmentRequestStatusUnassigned             AssignmentRequestStatus = "Unassigned"
	AssignmentRequestStatusApproved               AssignmentRequestStatus = "Approved"
	AssignmentRequestStatusRejected               AssignmentRequestStatus = "Rejected"
	AssignmentRequestStatusDamaged                AssignmentRequestStatus = "Reported Damaged"
	AssignmentRequestStatusMissing                AssignmentRequestStatus = "Reported Missing"
	AssignmentRequestStatusReportedDisposal       AssignmentRequestStatus = "Reported Disposal"
	AssignmentRequestStatusOnDisposalProcess      AssignmentRequestStatus = "On Disposal Process"
	AssignmentRequestStatusDisposed               AssignmentRequestStatus = "Disposed"
	AssignmentRequestStatusTAGReported            AssignmentRequestStatus = "TAG Reported"
)

type BorrowingTransactionStatus string

const (
	BorrowingTransactionStatusWaitingForApproval     BorrowingTransactionStatus = "Waiting for Approval"
	BorrowingTransactionStatusWaitingForVerification BorrowingTransactionStatus = "Waiting for Verification"
	BorrowingTransactionStatusWaitingForConfirmation BorrowingTransactionStatus = "Waiting for Confirmation"
	BorrowingTransactionStatusWaitingForHandover     BorrowingTransactionStatus = "Waiting for Handover"
	BorrowingTransactionStatusBorrowed               BorrowingTransactionStatus = "Borrowed"
	BorrowingTransactionStatusCanceled               BorrowingTransactionStatus = "Canceled"
	BorrowingTransactionStatusReturned               BorrowingTransactionStatus = "Returned"
	BorrowingTransactionStatusExtendApproval         BorrowingTransactionStatus = "Extend Approval"
	BorrowingTransactionStatusRequestExtension       BorrowingTransactionStatus = "Request Extension"
	BorrowingTransactionStatusRejected               BorrowingTransactionStatus = "Rejected"
)

type BorrowingRequestStatus string

const (
	BorrowingRequestStatusWaitingForApproval     BorrowingRequestStatus = "Waiting for Approval"
	BorrowingRequestStatusWaitingForVerification BorrowingRequestStatus = "Waiting for Verification"
	BorrowingRequestStatusWaitingForConfirmation BorrowingRequestStatus = "Waiting for Confirmation"
	BorrowingRequestStatusWaitingForHandover     BorrowingRequestStatus = "Waiting for Handover"
	BorrowingRequestStatusBorrowed               BorrowingRequestStatus = "Borrowed"
	BorrowingRequestStatusCanceled               BorrowingRequestStatus = "Canceled"
	BorrowingRequestStatusReturned               BorrowingRequestStatus = "Returned"
	BorrowingRequestStatusApproved               BorrowingRequestStatus = "Approved"
	BorrowingRequestStatusRejected               BorrowingRequestStatus = "Rejected"
	BorrowingRequestStatusDamaged                BorrowingRequestStatus = "Reported Damaged"
	BorrowingRequestStatusMissing                BorrowingRequestStatus = "Reported Missing"
	BorrowingRequestStatusReportedDisposal       BorrowingRequestStatus = "Reported Disposal"
	BorrowingRequestStatusOnDisposalProcess      BorrowingRequestStatus = "On Disposal Process"
	BorrowingRequestStatusDisposed               BorrowingRequestStatus = "Disposed"
	BorrowingRequestStatusTAGReported            BorrowingRequestStatus = "TAG Reported"
	BorrowingRequestStatusRequestExtension       BorrowingRequestStatus = "Request Extension"
	BorrowingRequestStatusExtendApproval         BorrowingRequestStatus = "Extend Approval"
)

type ApprovalStatus string

const (
	ApprovalStatusNeedApproval     ApprovalStatus = "Need Approval"
	ApprovalStatusFinishedApproval ApprovalStatus = "Finished Approval"
	ApprovalStatusPending          ApprovalStatus = "Pending"
)

type ApprovalType string

const (
	ApprovalTypeOr  ApprovalType = "Or"
	ApprovalTypeAnd ApprovalType = "And"
)

type AssignmentPrelistType string

const (
	AssignmentPrelistTypeAssign   AssignmentPrelistType = "Assign"
	AssignmentPrelistTypeUnassign AssignmentPrelistType = "Unassign"
)

type TransferTransactionStatus string

const (
	TransferTransactionStatusWaitingForApproval TransferTransactionStatus = "Waiting for Approval"
	TransferTransactionStatusWaitingForHandover TransferTransactionStatus = "Waiting for Handover"
	TransferTransactionStatusCompleted          TransferTransactionStatus = "Completed"
	TransferTransactionStatusRejected           TransferTransactionStatus = "Rejected"
	TransferTransactionStatusCanceled           TransferTransactionStatus = "Canceled"
)

type TransferRequestStatus string

const (
	TransferRequestStatusWaitingForApproval TransferRequestStatus = "Waiting for Approval"
	TransferRequestStatusWaitingForHandover TransferRequestStatus = "Waiting for Handover"
	TransferRequestStatusCompleted          TransferRequestStatus = "Completed"
	TransferRequestStatusApproved           TransferRequestStatus = "Approved"
	TransferRequestStatusRejected           TransferRequestStatus = "Rejected"
	TransferRequestStatusCanceled           TransferRequestStatus = "Canceled"
	TransferRequestStatusReportedMissing    TransferRequestStatus = "Reported Missing"
)

type TransferRequestAssetTagStatus string

const (
	TransferRequestAssetTagStatusVerified    TransferRequestAssetTagStatus = "Verified"
	TransferRequestAssetTagStatusUnverified  TransferRequestAssetTagStatus = "Unverified"
	TransferRequestAssetTagStatusTAGReported TransferRequestAssetTagStatus = "TAG Reported"
)

type TransferRequestAssetStatus string

const (
	TransferRequestAssetStatusNormal  TransferRequestAssetStatus = "Normal"
	TransferRequestAssetStatusDamaged TransferRequestAssetStatus = "Damaged"
	TransferRequestAssetStatusMissing TransferRequestAssetStatus = "Missing"
)

type TransferGroupType string

const (
	TransferGroupTypeOrigin      TransferGroupType = "Origin"
	TransferGroupTypeDestination TransferGroupType = "Destination"
)

type DisposalLastCondition string

const (
	DisposalLastConditionNoSalvageValue DisposalLastCondition = "No Salvage Value"
	DisposalLastConditionMissing        DisposalLastCondition = "Missing"
	DisposalLastConditionDamaged        DisposalLastCondition = "Damaged"
)

var DisposalLastCoditionEnum = map[string]DisposalLastCondition{
	"No Salvage Value": DisposalLastConditionNoSalvageValue,
	"Missing":          DisposalLastConditionMissing,
	"Damaged":          DisposalLastConditionDamaged,
}

type DisposalMethod string

const (
	DisposalMethodDonate    DisposalMethod = "Donate"
	DisposalMethodSell      DisposalMethod = "Sell"
	DisposalMethodDestroy   DisposalMethod = "Destroy"
	DisposalMethodClearance DisposalMethod = "Clearance"
	DisposalMethodOther     DisposalMethod = "Other"
)

type DisposalTransactionStatus string

const (
	DisposalTransactionStatusWaitingForApproval     DisposalTransactionStatus = "Waiting for Approval"
	DisposalTransactionStatusWaitingForVerification DisposalTransactionStatus = "Waiting for Verification"
	DisposalTransactionStatusOnProgress             DisposalTransactionStatus = "On Progress"
	DisposalTransactionStatusCompletionApproval     DisposalTransactionStatus = "Completion Approval"
	DisposalTransactionStatusCompleted              DisposalTransactionStatus = "Completed"
	DisposalTransactionStatusCanceled               DisposalTransactionStatus = "Canceled"
	DisposalTransactionStatusRejected               DisposalTransactionStatus = "Rejected"
)

type DisposalRequestStatus string

const (
	DisposalRequestStatusWaitingForApproval     DisposalRequestStatus = "Waiting for Approval"
	DisposalRequestStatusWaitingForVerification DisposalRequestStatus = "Waiting for Verification"
	DisposalRequestStatusOnProgress             DisposalRequestStatus = "On Progress"
	DisposalRequestStatusCommpletionApproval    DisposalRequestStatus = "Completion Approval"
	DisposalRequestStatusCompleted              DisposalRequestStatus = "Completed"
	DisposalRequestStatusCanceled               DisposalRequestStatus = "Canceled"
	DisposalRequestStatusApproved               DisposalRequestStatus = "Approved"
	DisposalRequestStatusRejected               DisposalRequestStatus = "Rejected"
)

type DisposalGroupType string

const (
	DisposalGroupTypeOrigin   DisposalGroupType = "Origin"
	DisposalGroupTypeDisposal DisposalGroupType = "Disposal"
)

type RoutineScheduleFrequency string

const (
	RoutineScheduleFrequencyOnce     RoutineScheduleFrequency = "Once"
	RoutineScheduleFrequencyDaily    RoutineScheduleFrequency = "Daily"
	RoutineScheduleFrequencyWeekly   RoutineScheduleFrequency = "Weekly"
	RoutineScheduleFrequencyMonthly  RoutineScheduleFrequency = "Monthly"
	RoutineScheduleFrequencyAnnually RoutineScheduleFrequency = "Annually"
	RoutineScheduleFrequencyCustom   RoutineScheduleFrequency = "Custom"
)

type Frequency string

const (
	FrequencyOnce     Frequency = "Once"
	FrequencyDaily    Frequency = "Daily"
	FrequencyWeekly   Frequency = "Weekly"
	FrequencyMonthly  Frequency = "Monthly"
	FrequencyAnnually Frequency = "Annually"
	FrequencyCustom   Frequency = "Custom"
)

type IndicatorType string

const (
	IndicatorTypeSelection   IndicatorType = "Selection"
	IndicatorTypeSingleValue IndicatorType = "Single Value"
	IndicatorTypeRangeValue  IndicatorType = "Range Value"
)

type IndicatorDataType string

const (
	IndicatorDataTypeText   IndicatorDataType = "Text"
	IndicatorDataTypeNumber IndicatorDataType = "Number"
)

type RoutineTaskStatus string

const (
	RoutineTaskStatusOnSchedule           RoutineTaskStatus = "On Schedule"
	RoutineTaskStatusWaitingForAssignment RoutineTaskStatus = "Waiting for Assignment"
	RoutineTaskStatusUnderMaintenance     RoutineTaskStatus = "Under Maintenance"
	RoutineTaskStatusWaitingForReview     RoutineTaskStatus = "Waiting for Review"
	RoutineTaskStatusWaitingForApproval   RoutineTaskStatus = "Waiting for Approval"
	RoutineTaskStatusOverdue              RoutineTaskStatus = "Overdue"
	RoutineTaskStatusCompleted            RoutineTaskStatus = "Completed"
	RoutineTaskStatusCanceled             RoutineTaskStatus = "Canceled"
	RoutineTaskStatusReportedDamaged      RoutineTaskStatus = "Reported Damaged"
	RoutineTaskStatusStopped              RoutineTaskStatus = "Stopped"
)

type RepairTicketingStatus string

const (
	RepairTicketingStatusReportedDamaged        RepairTicketingStatus = "Reported Damaged"
	RepairTicketingStatusUnderRepair            RepairTicketingStatus = "Under Repair"
	RepairTicketingStatusWaitingForConfirmation RepairTicketingStatus = "Waiting for Confirmation"
	RepairTicketingStatusWaitingForApproval     RepairTicketingStatus = "Waiting for Approval"
	RepairTicketingStatusRepaired               RepairTicketingStatus = "Repaired"
	RepairTicketingStatusPermanentlyDamaged     RepairTicketingStatus = "Permanently Damaged"
	RepairTicketingStatusReportedDisposal       RepairTicketingStatus = "Reported Disposal"
	RepairTicketingStatusOnDisposalProcess      RepairTicketingStatus = "On Disposal Process"
	RepairTicketingStatusDisposed               RepairTicketingStatus = "Disposed"
)

var RepairTicketingStatusEnum = map[string]RepairTicketingStatus{
	"Reported Damaged":         RepairTicketingStatusReportedDamaged,
	"Under Repair":             RepairTicketingStatusUnderRepair,
	"Waiting for Confirmation": RepairTicketingStatusWaitingForConfirmation,
	"Waiting for Approval":     RepairTicketingStatusWaitingForApproval,
	"Repaired":                 RepairTicketingStatusRepaired,
	"Permanently Damaged":      RepairTicketingStatusPermanentlyDamaged,
	"Reported Disposal":        RepairTicketingStatusReportedDisposal,
	"On Disposal Process":      RepairTicketingStatusOnDisposalProcess,
	"Disposed":                 RepairTicketingStatusDisposed,
}

type RepairTicketingMethod string

const (
	RepairTicketingMethodInternalRepair RepairTicketingMethod = "Internal Repair"
	RepairTicketingMethodThirdParty     RepairTicketingMethod = "Third Party"
	RepairTicketingMethodWarrantyClaim  RepairTicketingMethod = "Warranty Claim"
)

var RepairTicketingMethodEnum = map[string]RepairTicketingMethod{
	"Internal Repair": RepairTicketingMethodInternalRepair,
	"Third Party":     RepairTicketingMethodThirdParty,
	"Warranty Claim":  RepairTicketingMethodWarrantyClaim,
}

type TrackingStatus string

const (
	TrackingStatusOnTracking                   TrackingStatus = "On Tracking"
	TrackingStatusTrackedByUser                TrackingStatus = "Tracked by User"
	TrackingStatusTrackedByMaintenance         TrackingStatus = "Tracked by Maintenance"
	TrackingStatusTrackedByAudit               TrackingStatus = "Tracked by Audit"
	TrackingStatusTrackedByIOT                 TrackingStatus = "Tracked by IOT"
	TrackingStatusTrackedByTracking            TrackingStatus = "Tracked by Tracking"
	TrackingStatusReportedAsPermanentlyMissing TrackingStatus = "Reported as Permanently Missing"
	TrackingStatusPermanentlyMissing           TrackingStatus = "Permanently Missing"
	TrackingStatusReportedDisposal             TrackingStatus = "Reported Disposal"
	TrackingStatusOnDisposalProcess            TrackingStatus = "On Disposal Process"
	TrackingStatusDisposed                     TrackingStatus = "Disposed"
	TrackingStatusTAGReported                  TrackingStatus = "TAG Reported"
)

var TrackingStatusEnum = map[string]TrackingStatus{
	"On Tracking":                     TrackingStatusOnTracking,
	"Tracked by User":                 TrackingStatusTrackedByUser,
	"Tracked by Maintenance":          TrackingStatusTrackedByMaintenance,
	"Tracked by Audit":                TrackingStatusTrackedByAudit,
	"Tracked by IOT":                  TrackingStatusTrackedByIOT,
	"Tracked by Tracking":             TrackingStatusTrackedByTracking,
	"Reported as Permanently Missing": TrackingStatusReportedAsPermanentlyMissing,
	"Permanently Missing":             TrackingStatusPermanentlyMissing,
	"Reported Disposal":               TrackingStatusReportedDisposal,
	"On Disposal Process":             TrackingStatusOnDisposalProcess,
	"Disposed":                        TrackingStatusDisposed,
}

type TrackingCondition string

const (
	TrackingConditionNormal   TrackingCondition = "Asset found, normal condition"
	TrackingConditionDamaged  TrackingCondition = "Asset found, damaged condition"
	TrackingConditionTAGFound TrackingCondition = "TAG found, asset missing"
)

var TrackingConditionEnum = map[string]TrackingCondition{
	"Asset found, normal condition":  TrackingConditionNormal,
	"Asset found, damaged condition": TrackingConditionDamaged,
	"TAG found, asset missing":       TrackingConditionTAGFound,
}

type ReportType string

const (
	ReportTypeCustomReport             ReportType = "Custom Report"
	ReportTypeSummaryOverview          ReportType = "Summary Overview"
	ReportTypeDepreciationReport       ReportType = "Depreciation Report"
	ReportTypeMaintenanceRoutineReport ReportType = "Maintenance Routine Report"
	ReportTypeRepairmentReport         ReportType = "Repairment Report"
)

type ReportListedBy string

const (
	ListedByGroup    ReportListedBy = "Group"
	ListedByCategory ReportListedBy = "Category"
)

type ReportAssetInformationType string

const (
	ReportAssetInfoBasicInformation   ReportAssetInformationType = "Basic Information"
	ReportAssetInfoBorrow             ReportAssetInformationType = "Borrow"
	ReportAssetInfoTransfer           ReportAssetInformationType = "Transfer"
	ReportAssetInfoAssignment         ReportAssetInformationType = "Assignment"
	ReportAssetInfoTracking           ReportAssetInformationType = "Tracking"
	ReportAssetInfoAudit              ReportAssetInformationType = "Audit"
	ReportAssetInfoDisposal           ReportAssetInformationType = "Disposal"
	ReportAssetInfoMaintenanceRoutine ReportAssetInformationType = "Maintenance Routine"
	ReportAssetInfoRepairTicketing    ReportAssetInformationType = "Repair Ticketing"
	ReportAssetInfoTagging            ReportAssetInformationType = "Tagging"
	ReportAssetInfoLicense            ReportAssetInformationType = "License"
	ReportAssetInfoAssetPolicy        ReportAssetInformationType = "Asset Policy"
	ReportAssetInfoDamage             ReportAssetInformationType = "Damage"
	ReportAssetInfoMissing            ReportAssetInformationType = "Missing"
)
