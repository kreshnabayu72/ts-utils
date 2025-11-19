package model

import (
	"encoding/json"
	"time"
)

type UserEmbedded struct {
	ID        *string `json:"_id,omitempty"       bson:"_id,omitempty"`
	FullName  *string `json:"fullName"            bson:"fullName"`
	Key       *int    `json:"key"                 bson:"key"`
	Position  *string `json:"position,omitempty"  bson:"position,omitempty"`
	Division  *string `json:"division,omitempty"  bson:"division,omitempty"`
	IsDefault *bool   `json:"isDefault,omitempty" bson:"isDefault,omitempty"`
}
type AttributeEmbedded struct {
	ID   *string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name *string `json:"name"          bson:"name"`
	Key  *int    `json:"key"           bson:"key"`
}

type NameEmbedded struct {
	ID               string `bson:"_id"              json:"_id"`
	NameWithSequence string `bson:"nameWithSequence" json:"nameWithSequence"`
	Name             string `bson:"name"             json:"name"`
	Key              int    `bson:"key"              json:"key"`
	NameNumber       *int   `bson:"nameNumber"       json:"nameNumber"`
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

// JSONBArrayString is a custom type for handling a Go []string that is stored as
// a JSONB array of strings in a PostgreSQL database.
// It ensures correct marshaling and unmarshaling, particularly for handling nil/null states
// when interacting with the database or API payloads.
type JSONBArrayString []string

// MarshalJSON implements the json.Marshaler interface.
// It converts the Go slice to a JSON array string (e.g., '["a","b"]').
// It explicitly returns 'null' for a nil Go slice, rather than '[]'.
func (s JSONBArrayString) MarshalJSON() ([]byte, error) {
	if s == nil {
		return json.Marshal(nil)
	}
	return json.Marshal([]string(s))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It converts incoming JSON data (from an API or database query result) back into a Go slice.
// It handles the JSON 'null' value gracefully by setting the slice to nil.
func (s *JSONBArrayString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*s = nil
		return nil
	}

	var temp []string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*s = temp
	return nil
}

type TimeRange struct {
	Start time.Time
	End   time.Time
}
