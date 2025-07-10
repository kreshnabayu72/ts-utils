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
