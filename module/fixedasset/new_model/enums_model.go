package new_model

type AssetDefaultImage string

const (
	AssetDefaultImageFirstImage  = AssetDefaultImage("first_image")
	AssetDefaultImageSecondImage = AssetDefaultImage("second_image")
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
	ModuleAccessSettingsPortal = ModuleAccess("Settings Portal")
	ModuleAccessAdminConsole   = ModuleAccess("Admin Console")
	ModuleAccessSupplyAsset    = ModuleAccess("Supply Asset")
)

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
	AssignmentTransactionStatusWaitingForApproval AssignmentTransactionStatus = "Waiting for Approval"
	AssignmentTransactionStatusWaitingForHandover AssignmentTransactionStatus = "Waiting for Handover"
	AssignmentTransactionStatusAssigned           AssignmentTransactionStatus = "Assigned"
	AssignmentTransactionStatusUnassigned         AssignmentTransactionStatus = "Unassigned"
	AssignmentTransactionStatusRejected           AssignmentTransactionStatus = "Rejected"
)

type AssignmentRequestStatus string

const (
	AssignmentRequestStatusWaitingForApproval    AssignmentRequestStatus = "Waiting for Approval"
	AssignmentRequestStatusWaitingForHandover    AssignmentRequestStatus = "Waiting for Handover"
	AssignmentRequestStatusAssigned              AssignmentRequestStatus = "Assigned"
	AssignmentRequestStatusUnassigned            AssignmentRequestStatus = "Unassigned"
	AssignmentRequestStatusApproved              AssignmentRequestStatus = "Approved"
	AssignmentRequestStatusRejected              AssignmentRequestStatus = "Rejected"
	AssignmentRequestStatusDamaged               AssignmentRequestStatus = "Damaged"
	AssignmentRequestStatusMissing               AssignmentRequestStatus = "Missing"
	AssignmentRequestStatusReportedDamagedByUser AssignmentRequestStatus = "Reported Damaged By User"
	AssignmentRequestStatusReportedMissingByUser AssignmentRequestStatus = "Reported Missing By User"
	AssignmentRequestStatusReportedDisposal      AssignmentRequestStatus = "Reported Disposal"
	AssignmentRequestStatusOnDisposalProcess     AssignmentRequestStatus = "On Disposal Process"
	AssignmentRequestStatusDisposed              AssignmentRequestStatus = "Disposed"
)

type BorrowingTransactionStatus string

const (
	BorrowingTransactionStatusWaitingForApproval BorrowingTransactionStatus = "Waiting for Approval"
	BorrowingTransactionStatusWaitingForHandover BorrowingTransactionStatus = "Waiting for Handover"
	BorrowingTransactionStatusBorrowed           BorrowingTransactionStatus = "Borrowed"
	BorrowingTransactionStatusReturned           BorrowingTransactionStatus = "Returned"
	BorrowingTransactionStatusExtendApproval     BorrowingTransactionStatus = "Extend Approval"
	BorrowingTransactionStatusRequestExtension   BorrowingTransactionStatus = "Request Extension"
	BorrowingTransactionStatusRejected           BorrowingTransactionStatus = "Rejected"
)

type BorrowingRequestStatus string

const (
	BorrowingRequestStatusWaitingForApproval        BorrowingRequestStatus = "Waiting for Approval"
	BorrowingRequestStatusWaitingForHandover        BorrowingRequestStatus = "Waiting for Handover"
	BorrowingRequestStatusBorrowed                  BorrowingRequestStatus = "Borrowed"
	BorrowingRequestStatusReturned                  BorrowingRequestStatus = "Returned"
	BorrowingRequestStatusApproved                  BorrowingRequestStatus = "Approved"
	BorrowingRequestStatusRejected                  BorrowingRequestStatus = "Rejected"
	BorrowingRequestStatusDamaged                   BorrowingRequestStatus = "Damaged"
	BorrowingRequestStatusMissing                   BorrowingRequestStatus = "Missing"
	BorrowingRequestStatusReportedDamagedByBorrower BorrowingRequestStatus = "Reported Damaged By Borrower"
	BorrowingRequestStatusReportedMissingByBorrower BorrowingRequestStatus = "Reported Missing By Borrower"
	BorrowingRequestStatusReportedDisposal          BorrowingRequestStatus = "Reported Disposal"
	BorrowingRequestStatusOnDisposalProcess         BorrowingRequestStatus = "On Disposal Process"
	BorrowingRequestStatusDisposed                  BorrowingRequestStatus = "Disposed"
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
	AssignmentPrelistTypeOr  AssignmentPrelistType = "Assign"
	AssignmentPrelistTypeAnd AssignmentPrelistType = "Unassign"
)

type TransferTransactionStatus string

const (
	TransferTransactionStatusWaitingForApproval TransferTransactionStatus = "Waiting for Approval"
	TransferTransactionStatusWaitingForHandover TransferTransactionStatus = "Waiting for Handover"
	TransferTransactionStatusCompleted          TransferTransactionStatus = "Completed"
	TransferTransactionStatusRejected           TransferTransactionStatus = "Rejected"
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

type DisposalMethod string

const (
	DisposalMethodDonate    DisposalMethod = "Donate"
	DisposalMethodSell      DisposalMethod = "Sell"
	DisposalMethodDestroy   DisposalMethod = "Destroy"
	DisposalMethodClearance DisposalMethod = "Clearance"
)

type DisposalTransactionStatus string

const (
	DisposalTransactionStatusWaitingForApproval     DisposalTransactionStatus = "Waiting for Approval"
	DisposalTransactionStatusWaitingForVerification DisposalTransactionStatus = "Waiting for Verification"
	DisposalTransactionStatusOnProgress             DisposalTransactionStatus = "On Progress"
	DisposalTransactionStatusCompletionApproval     DisposalTransactionStatus = "Completion Approval"
	DisposalTransactionStatusCompleted              DisposalTransactionStatus = "Completed"
	DisposalTransactionStatusRejected               DisposalTransactionStatus = "Rejected"
)

type DisposalRequestStatus string

const (
	DisposalRequestStatusWaitingForApproval     DisposalRequestStatus = "Waiting for Approval"
	DisposalRequestStatusWaitingForVerification DisposalRequestStatus = "Waiting for Verification"
	DisposalRequestStatusOnProgress             DisposalRequestStatus = "On Progress"
	DisposalRequestStatusCompleted              DisposalRequestStatus = "Completed"
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
