package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Asset struct {
	ID                   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name                 string             `json:"name" bson:"name"`
	Parent               primitive.ObjectID `json:"parent" bson:"parent"`
	FullPath             string             `json:"fullPath" bson:"fullPath"`
	IsDefault            bool               `json:"isDefault" bson:"isDefault"`
	IsDeleted            bool               `json:"isDeleted" bson:"isDeleted"`
	Key                  int32              `json:"key" bson:"key"`
	DeleteFromDatabaseAt *time.Time         `json:"deleteFromDatabaseAt" bson:"deleteFromDatabaseAt"`
	CreatedAt            time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt            time.Time          `json:"updatedAt" bson:"updatedAt"`
}
