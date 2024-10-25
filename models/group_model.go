package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Key  int32              `json:"key" bson:"key"`
}
