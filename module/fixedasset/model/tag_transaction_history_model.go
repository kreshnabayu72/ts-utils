package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TagTransactionHistory struct {
	ID           primitive.ObjectID        `json:"id,omitempty" bson:"_id,omitempty"`
	Group        GroupAttribute            `json:"group" bson:"group"`
	Type         TagTransactionHistoryType `json:"type" bson:"type"`
	TagType      TagType                   `json:"tagType" bson:"tagType"`
	Rfid         string                    `json:"rfid" bson:"rfid"`
	Qr           string                    `json:"qr" bson:"qr"`
	AssetName    AssetNameAttribute        `json:"assetName" bson:"assetName"`
	UserFullName string                    `json:"userFullName" bson:"userFullName"`
	CreatedAt    time.Time                 `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time                 `json:"updatedAt" bson:"updatedAt"`
}
