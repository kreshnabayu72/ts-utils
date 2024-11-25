package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DepreciationGroup struct {
	ID                   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name                 string             `json:"name" bson:"name"`
	UsefulLife           uint               `json:"usefulLife" bson:"usefulLife"`
	StraightLineRate     float32            `json:"straightLineRate" bson:"straightLineRate"`
	DoubleDecreaseRate   float32            `json:"doubleDecreaseRate" bson:"doubleDecreaseRate"`
	IsDeleted            bool               `json:"isDeleted" bson:"isDeleted"`
	Key                  int                `json:"key" bson:"key"`
	DeleteFromDatabaseAt *time.Time         `json:"deleteFromDatabaseAt" bson:"deleteFromDatabaseAt"`
	CreatedAt            time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt            time.Time          `json:"updatedAt" bson:"updatedAt"`
}
