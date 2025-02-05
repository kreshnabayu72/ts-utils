package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserToken struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Token     string             `json:"token" bson:"token"`
	Email     string             `json:"email" bson:"email"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
