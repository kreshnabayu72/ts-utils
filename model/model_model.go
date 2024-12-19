package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID        primitive.ObjectID  `json:"_id,omitempty"      bson:"_id,omitempty"`      // MongoDB ObjectID
	Name      string              `json:"name"               bson:"name"`               // Field `name` wajib
	Key       int                 `json:"key"                bson:"key"`                // Field `key` wajib
	Brand     *primitive.ObjectID `json:"brand,omitempty"    bson:"brand,omitempty"`    // Default null, indexed
	Category  *primitive.ObjectID `json:"category,omitempty" bson:"category,omitempty"` // Default null
	CreatedAt time.Time           `json:"createdAt"          bson:"createdAt"`          // Otomatis disimpan MongoDB
	UpdatedAt time.Time           `json:"updatedAt"          bson:"updatedAt"`          // Otomatis disimpan MongoDB
}
