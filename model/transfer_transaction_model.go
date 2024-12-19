package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransferTransaction struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	TransactionId    string             `json:"transactionId" bson:"transactionId"`
	Manager          *UserEmbedded      `json:"manager" bson:"manager"`
	OriginGroup      AttributeEmbedded  `json:"originGroup" bson:"originGroup"`
	DestinationGroup AttributeEmbedded  `json:"destinationGroup" bson:"destinationGroup"`
	TotalAssets      int32              `json:"totalAssets" bson:"totalAssets"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt" bson:"updatedAt"`
}
