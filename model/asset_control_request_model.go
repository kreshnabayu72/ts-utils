package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetControlTransaction struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Transaction primitive.ObjectID `json:"transaction" bson:"transaction"`

	OriginGroup      AttributeEmbedded `json:"originGroup" bson:"originGroup"`
	DestinationGroup AttributeEmbedded `json:"destinationGroup" bson:"destinationGroup"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
