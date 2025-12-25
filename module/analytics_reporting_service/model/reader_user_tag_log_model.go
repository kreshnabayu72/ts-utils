package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReaderUserTagLog struct {
	ID           primitive.ObjectID `json:"_id"                    bson:"_id"`
	UserID       string             `json:"userId"                 bson:"userId"`
	User         string             `json:"user"                   bson:"user"`
	ModifiedByID string             `json:"modifiedById"           bson:"modifiedById"`
	ModifiedBy   string             `json:"modifiedBy"             bson:"modifiedBy"`
	CreatedAt    time.Time          `json:"createdAt"              bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"              bson:"updatedAt"`
}
