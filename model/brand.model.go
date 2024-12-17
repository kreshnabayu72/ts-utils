package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BrandAttributeCategory struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string             `bson:"name" json:"name"`
	Key  int                `bson:"key" json:"key"`
}

type Brand struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Key                int                `bson:"key" json:"key"`
	Name               string             `bson:"name" json:"name"`
	Category           []BrandAttributeCategory         `bson:"category" json:"category"`
	TotalAssets        int                `bson:"totalAssets,omitempty" json:"totalAssets,omitempty"`
	IsDeleted          bool               `bson:"isDeleted,omitempty" json:"isDeleted,omitempty"`
	DeleteFromDatabase *time.Time         `bson:"deleteFromDatabaseAt,omitempty" json:"deleteFromDatabaseAt,omitempty"`
	CreatedAt          time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt          time.Time          `bson:"updatedAt" json:"updatedAt"`
}