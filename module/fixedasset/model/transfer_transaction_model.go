package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransferStatus string

const (
	TransferStatusCompleted = TransferStatus("Completed")
)

type TransferTransaction struct {
	ID               primitive.ObjectID      `json:"id,omitempty" bson:"_id,omitempty"`
	TransactionId    string                  `json:"transactionId" bson:"transactionId"`
	Manager          *model.UserEmbedded     `json:"manager" bson:"manager"`
	OriginGroup      model.AttributeEmbedded `json:"originGroup" bson:"originGroup"`
	DestinationGroup model.AttributeEmbedded `json:"destinationGroup" bson:"destinationGroup"`
	TotalAssets      int                     `json:"totalAssets" bson:"totalAssets"`
	Status           TransferStatus          `json:"status" bson:"status"`
	CreatedAt        time.Time               `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time               `json:"updatedAt" bson:"updatedAt"`
}
