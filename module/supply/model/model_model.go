package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string              `json:"name"          bson:"name"`
	Key       int                 `json:"key"           bson:"key"`
	IsDeleted bool                `bson:"isDeleted"     json:"isDeleted"`
	Brand     primitive.ObjectID  `json:"brand"         bson:"brand"`
	Category  *primitive.ObjectID `json:"category"      bson:"category"`
	CreatedAt time.Time           `json:"createdAt"     bson:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt"     bson:"updatedAt"`
}
