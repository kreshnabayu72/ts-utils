package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEmbedded struct {
	ID       *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FullName *string             `json:"fullName" bson:"fullName"`
	Key      *int32              `json:"key" bson:"key"`
}
