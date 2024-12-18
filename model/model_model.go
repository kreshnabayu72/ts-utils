package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`           // MongoDB ObjectID
	Name      string              `bson:"name" json:"name"`                             // Field `name` wajib
	Key       int                 `bson:"key" json:"key"`                               // Field `key` wajib
	Brand     *primitive.ObjectID `bson:"brand,omitempty" json:"brand,omitempty"`       // Default null, indexed
	Category  *primitive.ObjectID `bson:"category,omitempty" json:"category,omitempty"` // Default null
	CreatedAt time.Time           `bson:"createdAt" json:"createdAt"`                   // Otomatis disimpan MongoDB
	UpdatedAt time.Time           `bson:"updatedAt" json:"updatedAt"`                   // Otomatis disimpan MongoDB
}
