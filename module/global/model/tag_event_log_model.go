package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogDetail struct {
	Type string `json:"type"        bson:"type"`

	Rfid      int `json:"rfid"        bson:"rfid"`
	RfidNfc   int `json:"rfidNfc"        bson:"rfidNfc"`
	RfidQr    int `json:"rfidQr"        bson:"rfidQr"`
	RfidNfcQr int `json:"rfidNfcQr"        bson:"rfidNfcQr"`
	Total     int `json:"total"        bson:"total"`

	IsChecked bool   `json:"isChecked,omitempty"        bson:"isChecked,omitempty"`
	GroupName string `json:"groupName,omitempty"        bson:"groupName,omitempty"`
}
type TagEventLog struct {
	ID     primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`
	Type   string             `json:"type"        bson:"type"`
	Module string             `json:"module"        bson:"module"`

	Status          string      `json:"status"        bson:"status"`
	TotalTagScanned int         `json:"totalTagScanned"        bson:"totalTagScanned"`
	ModifiedBy      string      `json:"modifiedBy"        bson:"modifiedBy"`
	Detail          []LogDetail `json:"detail"        bson:"detail"`

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}
