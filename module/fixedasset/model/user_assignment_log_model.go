package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAssignmentLog struct {
	ID                    primitive.ObjectID `json:"_id,omitempty"         bson:"_id,omitempty"`
	TransactionID         string             `json:"transactionId"         bson:"transactionId"`
	AssetName             string             `json:"assetName"             bson:"assetName"`
	AssetNameWithSequence string             `json:"assetNameWithSequence" bson:"assetNameWithSequence"`
	Action                string             `json:"action"                bson:"action"`
	UserID                string             `json:"userId"                bson:"userId"`
	UserFullName          string             `json:"userFullName"          bson:"userFullName"`
	CreatedAt             time.Time          `json:"createdAt"             bson:"createdAt"`
	UpdatedAt             time.Time          `json:"updatedAt"             bson:"updatedAt"`
}
