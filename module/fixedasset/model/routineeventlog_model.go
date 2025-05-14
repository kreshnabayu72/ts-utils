package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoutineEventLog struct {
	ID                    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	TaskName              string             `json:"taskName" bson:"taskName"`
	AssetName             string             `json:"assetName" bson:"assetName"`
	AssetNameWithSequence string             `json:"assetNameWithSequence" bson:"assetNameWithSequence"`
	Action                string             `json:"action" bson:"action"`
	Group                 string             `json:"group" bson:"group"`
	ModifiedBy            string             `json:"modifiedBy" bson:"modifiedBy"`
	CreatedAt             time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt             time.Time          `json:"updatedAt" bson:"updatedAt"`
}
