package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReaderTaggingLog struct {
	ID           primitive.ObjectID `json:"_id"                    bson:"_id"`
	Quantity     int                `json:"quantity"                   bson:"quantity"`
	Action       string             `json:"action"                 bson:"action"`
	HardwareID   string             `json:"hardwareId" bson:"hardwareId"`
	ModifiedByID string             `json:"modifiedById"                 bson:"modifiedById"`
	ModifiedBy   string             `json:"modifiedBy"           bson:"modifiedBy"`
	CreatedAt    time.Time          `json:"createdAt"              bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"              bson:"updatedAt"`
}
