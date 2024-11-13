package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Measurement struct {
	ID                   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name                 string             `json:"name" bson:"name"`
	IsDeleted            bool               `json:"isDeleted" bson:"isDeleted"`
	DeleteFromDatabaseAt time.Time          `json:"deleteFromDatabaseAt" bson:"deleteFromDatabaseAt"`
	CreatedAt            time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt            time.Time          `json:"updatedAt" bson:"updatedAt"`
}
