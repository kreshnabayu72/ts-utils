package new_model

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
