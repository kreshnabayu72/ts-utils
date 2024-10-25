package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransferRequest struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Transaction      primitive.ObjectID `json:"transaction" bson:"transaction"`
	AssetParent      primitive.ObjectID `json:"assetParent" bson:"assetParent"`
	Asset            primitive.ObjectID `json:"asset" bson:"asset"`
	Manager          UserEmbedded       `json:"manager" bson:"manager"`
	OriginGroup      AssetAttribute     `json:"originGroup" bson:"originGroup"`
	DestinationGroup AssetAttribute     `json:"destinationGroup" bson:"destinationGroup"`
	AssetImageBig    string             `json:"assetImageBig" bson:"assetImageBig"`
	AssetImageMedium string             `json:"assetImageMedium" bson:"assetImageMedium"`
	AssetImageSmall  string             `json:"assetImageSmall" bson:"assetImageSmall"`
	AssetName        AssetName          `json:"assetName" bson:"assetName"`
	AssetBrand       AssetAttribute     `json:"assetBrand" bson:"assetBrand"`
	AssetModel       AssetAttribute     `json:"assetModel" bson:"assetModel"`
	Status           string             `json:"status" bson:"status"`
	AssetTagType     string             `json:"assetTagType" bson:"assetTagType"`
	AssetQr          string             `json:"assetQr" bson:"assetQr"`
	AssetRfid        string             `json:"assetRfid" bson:"assetRfid"`
	HasChildren      bool               `json:"hasChildren" bson:"hasChildren"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt" bson:"updatedAt"`
}
