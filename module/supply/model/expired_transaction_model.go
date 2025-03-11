package model

import (
	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpiredTransactionStatus string

const (
	ExpiredTransactionStatusReported           ExpiredTransactionStatus = "Reported"
	ExpiredTransactionStatusWaitingForApproval ExpiredTransactionStatus = "Waiting for Approval"
	ExpiredTransactionStatusCompleted          ExpiredTransactionStatus = "Completed"
)

// ExpiredTransaction represents the active expired transaction that groups ExpiredRequest
//
// Expired transaction is a compiled request of ExpiredRequest
type ExpiredTransaction struct {
	ID                  primitive.ObjectID       `json:"_id,omitempty"       bson:"_id,omitempty"`
	Status              ExpiredTransactionStatus `json:"status"              bson:"status"`
	ItemName            primitive.ObjectID       `json:"itemName"            bson:"itemName"`
	ItemSKU             primitive.ObjectID       `json:"itemSKU"             bson:"itemSKU"`
	Brand               primitive.ObjectID       `json:"brand"               bson:"brand"`
	Model               primitive.ObjectID       `json:"model"               bson:"model"`
	TagType             model.TagType            `json:"tagType"             bson:"tagType"`
	AdministrativeGroup primitive.ObjectID       `json:"administrativeGroup" bson:"administrativeGroup"`
	TotalStock          int                      `json:"totalStock"          bson:"totalStock"`
	UpdatedAt           time.Time                `json:"updatedAt"           bson:"updatedAt"`
	CreatedAt           time.Time                `json:"createdAt"           bson:"createdAt"`
}

// ExpiredTransactionHistory represents the completed transaction of ExpiredTransaction
type ExpiredTransactionHistory struct {
	ID                  primitive.ObjectID       `json:"_id,omitempty"       bson:"_id,omitempty"`
	Status              ExpiredTransactionStatus `json:"status"              bson:"status"`
	ItemName            AttributeEmbedded        `json:"itemName"            bson:"itemName"`
	ItemSKU             SKUAttributeEmbedded     `json:"itemSKU"             bson:"itemSKU"`
	Brand               AttributeEmbedded        `json:"brand"               bson:"brand"`
	Model               AttributeEmbedded        `json:"model"               bson:"model"`
	TagType             model.TagType            `json:"tagType"             bson:"tagType"`
	AdministrativeGroup TreeAttributeEmbedded    `json:"administrativeGroup" bson:"administrativeGroup"`
	TotalStock          int                      `json:"totalStock"          bson:"totalStock"`
	UpdatedAt           time.Time                `json:"updatedAt"           bson:"updatedAt"`
	CreatedAt           time.Time                `json:"createdAt"           bson:"createdAt"`
}
