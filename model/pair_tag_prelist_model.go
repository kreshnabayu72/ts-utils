package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PairTagPrelistGroup struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Key  int                `json:"key" bson:"key"`
}

type PairTagPrelistAssetName = PairTagPrelistGroup

type PairTagPrelist struct {
	ID        primitive.ObjectID      `json:"_id,omitempty" bson:"_id,omitempty"`
	TagType   string                  `json:"tagType" bson:"tagType"`
	Group     PairTagPrelistGroup     `json:"group" bson:"group"`
	AssetName PairTagPrelistAssetName `json:"assetName" bson:"assetName"`
	CreatedAt time.Time               `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time               `json:"updatedAt" bson:"updatedAt"`
}
