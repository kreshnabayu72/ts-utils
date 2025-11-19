package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetControlTransaction struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	OriginGroup      model.AttributeEmbedded `json:"originGroup" bson:"originGroup"`
	DestinationGroup model.AttributeEmbedded `json:"destinationGroup" bson:"destinationGroup"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
