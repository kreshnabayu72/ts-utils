package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetName struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name             string             `json:"name" bson:"name"`
	NameWithSequence string             `json:"nameWithSequence" bson:"nameWithSequence"`
	Key              int32              `json:"key" bson:"key"`
}

type AssetAttribute struct {
	ID   *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name *string             `json:"name" bson:"name"`
	Key  *int32              `json:"key" bson:"key"`
}

type AssetTag struct {
	Code      *string `json:"code" bson:"code"`
	ShortCode *string `json:"shortCode" bson:"shortCode"`
}

type Asset struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      AssetName          `json:"name" bson:"name"`
	AssetId   string             `json:"assetId" bson:"assetId"`
	Model     AssetAttribute     `json:"model" bson:"model"`
	Brand     AssetAttribute     `json:"brand" bson:"brand"`
	Group     AssetAttribute     `json:"group" bson:"group"`
	Rfid      AssetTag           `json:"rfid" bson:"rfid"`
	Qr        AssetTag           `json:"qr" bson:"qr"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
