package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brand struct {
	ID                 primitive.ObjectID   `bson:"_id,omitempty"        json:"_id,omitempty"`
	Key                int                  `bson:"key"                  json:"key"`
	Name               string               `bson:"name"                 json:"name"`
	Category           []primitive.ObjectID `bson:"category"             json:"category"`
	TotalAssets        int                  `bson:"totalAssets"          json:"totalAssets"`
	IsDeleted          bool                 `bson:"isDeleted"            json:"isDeleted"`
	DeleteFromDatabase *time.Time           `bson:"deleteFromDatabaseAt" json:"deleteFromDatabaseAt"`
	CreatedAt          time.Time            `bson:"createdAt"            json:"createdAt"`
	UpdatedAt          time.Time            `bson:"updatedAt"            json:"updatedAt"`
}
