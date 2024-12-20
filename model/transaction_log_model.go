package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionLog struct {
	ID                    primitive.ObjectID     `json:"id,omitempty" bson:"_id,omitempty"`
	Type                  string                 `json:"type" bson:"type"`
	Transaction           primitive.ObjectID     `json:"transaction" bson:"transaction"`
	AssetName             string                 `json:"assetName" bson:"assetName"`
	AssetId               primitive.ObjectID     `json:"assetId" bson:"assetId"`
	AssetNameWithSequence string                 `json:"assetNameWithSequence" bson:"assetNameWithSequence"`
	Action                string                 `json:"action" bson:"action"`
	UserId                *primitive.ObjectID    `json:"userId" bson:"userId"`
	UserFullName          *string                `json:"userFullName" bson:"userFullName"`
	Detail                map[string]interface{} `json:"detail" bson:"detail"`
	OriginGroup           string                 `json:"originGroup" bson:"originGroup"`
	DestinationGroup      string                 `json:"destinationGroup" bson:"destinationGroup"`
	ReaderId              primitive.ObjectID     `json:"readerId" bson:"readerId"`
	ReaderName            string                 `json:"readerName" bson:"readerName"`
	ReaderType            string                 `json:"readerType" bson:"readerType"`
	CreatedAt             time.Time              `json:"createdAt" bson:"createdAt"`
	UpdatedAt             time.Time              `json:"updatedAt" bson:"updatedAt"`
}
