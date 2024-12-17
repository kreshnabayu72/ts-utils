package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ImportBrand struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // MongoDB ObjectID
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Category  string             `bson:"category,omitempty" json:"category,omitempty"`
	Models    string             `bson:"models,omitempty" json:"models,omitempty"`
	Key       int                `bson:"key,omitempty" json:"key,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
