package model

type TagType string

const (
	TagTypeNonTag    TagType = "Non TAG"
	TagTypeRFID      TagType = "RFID"
	TagTypeQR        TagType = "QR"
	TagTypeNFC       TagType = "NFC"
	TagTypeRFIDQR    TagType = "RFID & QR"
	TagTypeNFCQR     TagType = "NFC & QR"
	TagTypeRFIDNFC   TagType = "RFID & NFC"
	TagTypeRFIDNFCQR TagType = "RFID & NFC & QR"
)

type Tag struct {
	QRCode        *string `json:"qrCode"        bson:"qrCode"`
	ShortQRCode   *string `json:"shortQRCode"   bson:"shortQRCode"`
	RFIDCode      *string `json:"rfidCode"      bson:"rfidCode"`
	ShortRFIDCode *string `json:"shortRFIDCode" bson:"shortRFIDCode"`
	NFCCode       *string `json:"nfcCode"       bson:"nfcCode"`
	ShortNFCCode  *string `json:"shortNFCCode"  bson:"shortNFCCode"`
}
