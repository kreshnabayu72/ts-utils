package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DisposalStatus string

const (
	DisposalStatusCompleted = DisposalStatus("Completed")
	DisposalStatusDisposed  = DisposalStatus("Disposed")
)

type DisposalTransaction struct {
	ID            primitive.ObjectID      `json:"id,omitempty" bson:"_id,omitempty"`
	TransactionId string                  `json:"transactionId" bson:"transactionId"`
	Method        string                  `json:"method" bson:"method"`
	Organization  string                  `json:"organization" bson:"organization"`
	Manager       *model.UserEmbedded     `json:"manager" bson:"manager"`
	OriginGroup   model.AttributeEmbedded `json:"originGroup" bson:"originGroup"`
	DisposalGroup model.AttributeEmbedded `json:"disposalGroup" bson:"disposalGroup"`
	Status        DisposalStatus          `json:"status" bson:"status"`
	CreatedAt     time.Time               `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time               `json:"updatedAt" bson:"updatedAt"`
}
