package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogDetail struct {
	Type string `json:"type"        bson:"type"`

	Rfid      int `json:"rfid,omitempty"        bson:"rfid,omitempty"`
	RfidNfc   int `json:"rfidNfc,omitempty"        bson:"rfidNfc,omitempty"`
	RfidQr    int `json:"rfidQr,omitempty"        bson:"rfidQr,omitempty"`
	RfidNfcQr int `json:"rfidNfcQr,omitempty"        bson:"rfidNfcQr,omitempty"`

	IsChecked bool   `json:"isChecked,omitempty"        bson:"isChecked,omitempty"`
	GroupName string `json:"groupName,omitempty"        bson:"groupName,omitempty"`
}
type TagEventLog struct {
	ID   primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`
	Type string             `json:"type"        bson:"type"`

	Status          string      `json:"status"        bson:"status"`
	TotalTagScanned int         `json:"totalTagScanned"        bson:"totalTagScanned"`
	ModifiedBy      string      `json:"modifiedBy"        bson:"modifiedBy"`
	Detail          []LogDetail `json:"detail"        bson:"detail"`

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}
