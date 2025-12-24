package model

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HardwareChangeLog struct {
	ID                     primitive.ObjectID `json:"_id"                    bson:"_id"`
	Type                   enum.HardwareType  `json:"type"                   bson:"type"`
	Action                 string             `json:"action"                 bson:"action"`
	DeviceID               string             `json:"deviceId"               bson:"deviceId"`
	DeviceName             string             `json:"deviceName"             bson:"deviceName"`
	DeviceNameWithSequence *string            `json:"deviceNameWithSequence" bson:"deviceNameWithSequence"`
	OldValue               *string            `json:"oldValue"               bson:"oldValue"`
	NewValue               *string            `json:"newValue"               bson:"newValue"`
	SKU                    string             `json:"sku"                    bson:"sku"`
	SerialNumber           string             `json:"serialNumber"           bson:"serialNumber"`
	UserID                 string             `json:"userId"                 bson:"userId"`
	UserFullName           string             `json:"userFullName"           bson:"userFullName"`
	CreatedAt              time.Time          `json:"createdAt"              bson:"createdAt"`
	UpdatedAt              time.Time          `json:"updatedAt"              bson:"updatedAt"`
}
