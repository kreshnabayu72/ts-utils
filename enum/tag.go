package enum

type TagStatus string

const (
	TagStatusAvailable     TagStatus = "Available"
	TagStatusPaired        TagStatus = "Paired"
	TagStatusReported      TagStatus = "Reported"
	TagStatusDamageMissing TagStatus = "Damage / Missing"
	TagStatusToBeReturned  TagStatus = "To Be Returned"
)

type TagUsageType string

const (
	TagUsageTypeUser  TagUsageType = "User TAG"
	TagUsageTypeAsset TagUsageType = "Asset TAG"
)

type TagType string

const (
	TagTypeNonTag          TagType = "Non TAG"
	TagTypeRFID            TagType = "RFID"
	TagTypeQR              TagType = "QR"
	TagTypeNFC             TagType = "NFC"
	TagTypeRFIDAndQR       TagType = "RFID & QR"
	TagTypeRFIDAndNFC      TagType = "RFID & NFC"
	TagTypeNFCAndQR        TagType = "NFC & QR"
	TagTypeRFIDAndNFCAndQR TagType = "RFID & NFC & QR"
)
