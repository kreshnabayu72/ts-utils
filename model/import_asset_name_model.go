package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ImportAssetName struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`      // MongoDB ObjectID
	Name        string             `bson:"name" json:"name" validate:"required"`   // Name, required and unique
	AliasCode   *string            `bson:"aliasCode,omitempty" json:"aliasCode,omitempty"` // Alias code, optional
	Category    string             `bson:"category" json:"category" validate:"required"`   // Category, required
	Measurement string             `bson:"measurement" json:"measurement" validate:"required"` // Measurement, required
	Key         int                `bson:"key" json:"key" validate:"required"`     // Key, required
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"` // Timestamp for creation
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"` // Timestamp for last update
}
