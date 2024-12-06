package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ImportCategory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // MongoDB ObjectID
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Key       int                `bson:"key,omitempty" json:"key,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}
