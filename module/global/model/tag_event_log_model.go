package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TagEventLog struct {
	ID primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`

	Status          string `json:"status"        bson:"status"`
	TotalTagScanned int    `json:"totalTagScanned"        bson:"totalTagScanned"`
	ModifiedBy      string `json:"modifiedBy"        bson:"modifiedBy"`
	Detail          any    `json:"detail"        bson:"detail"` //TODO: Change type

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}
