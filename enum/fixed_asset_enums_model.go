package enum

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

type FixedAssetSystemRoleKey string

const (
	FixedAssetSystemRoleKeyImportData          FixedAssetSystemRoleKey = "importData"
	FixedAssetSystemRoleKeyAssetMasterData     FixedAssetSystemRoleKey = "assetMasterData"
	FixedAssetSystemRoleKeyDepreciationGroup   FixedAssetSystemRoleKey = "depreciationGroup"
	FixedAssetSystemRoleKeyDepreciationMethod  FixedAssetSystemRoleKey = "depreciationMethod"
	FixedAssetSystemRoleKeyAssetPolicies       FixedAssetSystemRoleKey = "assetPolicies"
	FixedAssetSystemRoleKeyAssets              FixedAssetSystemRoleKey = "assets"
	FixedAssetSystemRoleKeyFileManager         FixedAssetSystemRoleKey = "fileManager"
	FixedAssetSystemRoleKeyGroup               FixedAssetSystemRoleKey = "group"
	FixedAssetSystemRoleKeyUserAndRole         FixedAssetSystemRoleKey = "userAndRole"
	FixedAssetSystemRoleKeyIotReader           FixedAssetSystemRoleKey = "iotReader"
	FixedAssetSystemRoleKeyLicense             FixedAssetSystemRoleKey = "license"
	FixedAssetSystemRoleKeyPlanningStrategy    FixedAssetSystemRoleKey = "planningStrategy"
	FixedAssetSystemRoleKeyPurchaseDocument    FixedAssetSystemRoleKey = "purchaseDocument"
	FixedAssetSystemRoleKeyPurchaseInformation FixedAssetSystemRoleKey = "purchaseInformation"
	FixedAssetSystemRoleKeyTag                 FixedAssetSystemRoleKey = "tag"
)

type FixedAssetSystemRoleName string

const (
	FixedAssetSystemRoleNameImportData                FixedAssetSystemRoleName = "Import data"
	FixedAssetSystemRoleNameManageAssetMasterData     FixedAssetSystemRoleName = "Manage asset master data"
	FixedAssetSystemRoleNameManageDepreciationGroup   FixedAssetSystemRoleName = "Manage depreciation group"
	FixedAssetSystemRoleNameManageDepreciationMethod  FixedAssetSystemRoleName = "Manage depreciation method"
	FixedAssetSystemRoleNameManageAssetPolicy         FixedAssetSystemRoleName = "Manage asset policy"
	FixedAssetSystemRoleNameManageAssets              FixedAssetSystemRoleName = "Manage assets"
	FixedAssetSystemRoleNameFileManager               FixedAssetSystemRoleName = "File Manager"
	FixedAssetSystemRoleNameManageGroup               FixedAssetSystemRoleName = "Manage group"
	FixedAssetSystemRoleNameManageUserAndRole         FixedAssetSystemRoleName = "Manage user and role"
	FixedAssetSystemRoleNameManageIOTAndReader        FixedAssetSystemRoleName = "Manage IOT and reader"
	FixedAssetSystemRoleNameManageLicense             FixedAssetSystemRoleName = "Manage license"
	FixedAssetSystemRoleNameManagePlanningAndStrategy FixedAssetSystemRoleName = "Manage planning & strategy"
	FixedAssetSystemRoleNameManagePurchaseDocument    FixedAssetSystemRoleName = "Manage purchase document"
	FixedAssetSystemRoleNameManagePurchaseInformation FixedAssetSystemRoleName = "Manage purchase information"
	FixedAssetSystemRoleNameManageTAG                 FixedAssetSystemRoleName = "Manage TAG"
)

type FixedAssetTransactionRoleTransactionKey string

const (
	FixedAssetTransactionRoleTransactionKeyBorrowAssignRole       FixedAssetTransactionRoleTransactionKey = "borrowAssignRole"
	FixedAssetTransactionRoleTransactionKeyTransferRole           FixedAssetTransactionRoleTransactionKey = "transferRole"
	FixedAssetTransactionRoleTransactionKeyDisposalRole           FixedAssetTransactionRoleTransactionKey = "disposalRole"
	FixedAssetTransactionRoleTransactionKeyAuditRole              FixedAssetTransactionRoleTransactionKey = "auditRole"
	FixedAssetTransactionRoleTransactionKeyMaintenanceRoutineRole FixedAssetTransactionRoleTransactionKey = "maintenanceRoutineRole"
	FixedAssetTransactionRoleTransactionKeyRepairTicketingRole    FixedAssetTransactionRoleTransactionKey = "repairTicketingRole"
	FixedAssetTransactionRoleTransactionKeyTrackingRole           FixedAssetTransactionRoleTransactionKey = "trackingRole"
)

type FixedAssetTransactionRoleTransaction string

const (
	FixedAssetTransactionRoleTransactionBorrowingAssignment FixedAssetTransactionRoleTransaction = "Borrowing/Assignment"
	FixedAssetTransactionRoleTransactionTransfer            FixedAssetTransactionRoleTransaction = "Transfer"
	FixedAssetTransactionRoleTransactionDisposal            FixedAssetTransactionRoleTransaction = "Disposal"
	FixedAssetTransactionRoleTransactionAudit               FixedAssetTransactionRoleTransaction = "Audit"
	FixedAssetTransactionRoleTransactionMaintenanceRoutine  FixedAssetTransactionRoleTransaction = "Maintenance Routine"
	FixedAssetTransactionRoleTransactionRepairTicketing     FixedAssetTransactionRoleTransaction = "Repair Ticketing"
	FixedAssetTransactionRoleTransactionTracking            FixedAssetTransactionRoleTransaction = "Tracking"
)
