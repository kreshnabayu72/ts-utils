package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
)

// ReceptionRequest represents the individual request of an item in the ReceptionTransaction
type ReceptionRequest struct {
	ID             primitive.ObjectID         `json:"_id,omitempty"  bson:"_id,omitempty"`
	TransactionId  primitive.ObjectID         `json:"transactionId"  bson:"transactionId"`
	Status         ReceptionTransactionStatus `json:"status"         bson:"status"`
	ItemName       primitive.ObjectID         `json:"itemName"       bson:"itemName"`
	ItemSKU        primitive.ObjectID         `json:"itemSKU"        bson:"itemSKU"`
	Brand          primitive.ObjectID         `json:"brand"          bson:"brand"`
	Model          primitive.ObjectID         `json:"model"          bson:"model"`
	ReceivingGroup primitive.ObjectID         `json:"receivingGroup" bson:"receivingGroup"`
	PIC            primitive.ObjectID         `json:"pic"            bson:"pic"`
	Stocks         []primitive.ObjectID       `json:"stocks"         bson:"stocks"`
	TagType        model.TagType              `json:"tagType"        bson:"tagType"`
	Quantity       int                        `json:"quantity"       bson:"quantity"`
	UpdatedAt      time.Time                  `json:"updatedAt"      bson:"updatedAt"`
	CreatedAt      time.Time                  `json:"createdAt"      bson:"createdAt"`
}

// ReceptionRequestHistory represents the canceled/completed request of ReceptionRequest
type ReceptionRequestHistory struct {
	ID             primitive.ObjectID         `json:"_id,omitempty"  bson:"_id,omitempty"`
	TransactionId  primitive.ObjectID         `json:"transactionId"  bson:"transactionId"`
	Status         ReceptionTransactionStatus `json:"status"         bson:"status"`
	ItemName       AttributeEmbedded          `json:"itemName"       bson:"itemName"`
	ItemSKU        SKUAttributeEmbedded       `json:"itemSKU"        bson:"itemSKU"`
	Brand          AttributeEmbedded          `json:"brand"          bson:"brand"`
	Model          AttributeEmbedded          `json:"model"          bson:"model"`
	ReceivingGroup TreeAttributeEmbedded      `json:"receivingGroup" bson:"receivingGroup"`
	PIC            UserAttributeEmbedded      `json:"pic"            bson:"pic"`
	Stocks         []StockAttributeEmbedded   `json:"stocks"         bson:"stocks"`
	TagType        model.TagType              `json:"tagType"        bson:"tagType"`
	Quantity       int                        `json:"quantity"       bson:"quantity"`
	UpdatedAt      time.Time                  `json:"updatedAt"      bson:"updatedAt"`
	CreatedAt      time.Time                  `json:"createdAt"      bson:"createdAt"`
}
