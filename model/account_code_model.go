package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountCode struct {
	ID                   primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`
	Code                 string             `json:"code"                 bson:"code"`
	Name                 string             `json:"name"                 bson:"name"`
	IsDeleted            bool               `json:"isDeleted"            bson:"isDeleted"`
	Key                  int                `json:"key"                  bson:"key"`
	DeleteFromDatabaseAt *time.Time         `json:"deleteFromDatabaseAt" bson:"deleteFromDatabaseAt"`
	CreatedAt            time.Time          `json:"createdAt"            bson:"createdAt"`
	UpdatedAt            time.Time          `json:"updatedAt"            bson:"updatedAt"`
}
