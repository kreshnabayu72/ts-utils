package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID                   primitive.ObjectID  `json:"_id,omitempty"        bson:"_id,omitempty"`
	Name                 string              `json:"name"                 bson:"name"`
	FullPath             string              `json:"fullPath"             bson:"fullPath"`
	Parent               *primitive.ObjectID `json:"parent"               bson:"parent"`
	IsDefault            bool                `json:"isDefault"            bson:"isDefault"`
	IsDeleted            bool                `json:"isDeleted"            bson:"isDeleted"`
	IsDeletable			 bool				 `json:"isDeletable"          bson:"isDeletable"`
	IsLowest             bool				 `json:"isLowest"             bson:"isLowest"`
	Key                  int                 `json:"key"                  bson:"key"`
	DeleteFromDatabaseAt *time.Time          `json:"deleteFromDatabaseAt" bson:"deleteFromDatabaseAt"`
	CreatedAt            time.Time           `json:"createdAt"            bson:"createdAt"`
	UpdatedAt            time.Time           `json:"updatedAt"            bson:"updatedAt"`
}
