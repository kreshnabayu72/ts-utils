package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEmbedded struct {
	ID        *primitive.ObjectID `json:"_id,omitempty"       bson:"_id,omitempty"`
	FullName  *string             `json:"fullName"            bson:"fullName"`
	Key       *int                `json:"key"                 bson:"key"`
	Position  *string             `json:"position,omitempty"  bson:"position,omitempty"`
	Division  *string             `json:"division,omitempty"  bson:"division,omitempty"`
	IsDefault *bool               `json:"isDefault,omitempty" bson:"isDefault,omitempty"`
}
type AttributeEmbedded struct {
	ID   *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name *string             `json:"name"          bson:"name"`
	Key  *int                `json:"key"           bson:"key"`
}

type NameEmbedded struct {
	ID               primitive.ObjectID `bson:"_id"              json:"_id"`
	NameWithSequence string             `bson:"nameWithSequence" json:"nameWithSequence"`
	Name             string             `bson:"name"             json:"name"`
	Key              int                `bson:"key"              json:"key"`
	NameNumber       *int               `bson:"nameNumber"       json:"nameNumber"`
}

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

type AppModule string

const (
	FixedAssetModule  AppModule = "Fixed Asset"
	SupplyAssetModule AppModule = "Supply Asset"
)
