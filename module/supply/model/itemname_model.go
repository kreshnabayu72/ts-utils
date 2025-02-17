package model

import (
	"time"
	
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemName struct {
	ID          primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string               `json:"name"          bson:"name"`
	Key         int                  `json:"key"           bson:"key"`
	Measurement primitive.ObjectID   `json:"measurement"   bson:"measurement"`
	Category    primitive.ObjectID   `json:"category"      bson:"category"`
	Brand       []primitive.ObjectID `json:"brand"         bson:"brand"`
	Model       []primitive.ObjectID `json:"model"         bson:"model"`
	ModifiedBy  primitive.ObjectID   `json:"modifiedBy"    bson:"modifiedBy"`
	CreatedAt   time.Time            `json:"createdAt"     bson:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"     bson:"updatedAt"`
}
