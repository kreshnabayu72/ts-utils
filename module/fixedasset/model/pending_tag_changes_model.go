package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PendingTagTypeGroup struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Key  int                `json:"key" bson:"key"`
}

type PendingTagTypeAssetName struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Name             string             `json:"name" bson:"name"`
	Key              int                `json:"key" bson:"key"`
	NameWithSequence string             `json:"nameWithSequence" bson:"nameWithSequence"`
}

type PendingTagChange struct {
	ID                     primitive.ObjectID      `json:"_id,omitempty" bson:"_id,omitempty"`
	Type                   string                  `json:"type" bson:"type"`
	TagType                string                  `json:"tagType" bson:"tagType"`
	RFID                   string                  `json:"rfid" bson:"rfid"`
	QR                     string                  `json:"qr" bson:"qr"`
	AssetStatus            string                  `json:"assetStatus" bson:"assetStatus"`
	AssetMaintenanceStatus string                  `json:"assetMaintenanceStatus" bson:"assetMaintenanceStatus"`
	AssetAuditStatus       string                  `json:"assetAuditStatus" bson:"assetAuditStatus"`
	Group                  PendingTagTypeGroup     `json:"group" bson:"group"`
	AssetName              PendingTagTypeAssetName `json:"assetName" bson:"assetName"`
	Asset                  primitive.ObjectID      `json:"asset" bson:"asset"`
	CreatedAt              time.Time               `json:"createdAt" bson:"createdAt"`
	UpdatedAt              time.Time               `json:"updatedAt" bson:"updatedAt"`
}
