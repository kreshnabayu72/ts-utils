package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionLogType string

const (
	TransactionLogTypeReception TransactionLogType = "Reception"
	TransactionLogTypeExpired   TransactionLogType = "Expired"
)

type TransactionLogDetail struct {
	PIC       UserAttributeEmbedded `json:"pic"       bson:"pic"`
	Detail    map[string]any        `json:"detail"    bson:"detail"`
	Timestamp time.Time             `json:"timestamp" bson:"timestamp"`
}

type TransactionLog[O any] struct {
	ID            primitive.ObjectID      `json:"_id,omitempty" bson:"_id,omitempty"`
	Type          TransactionLogType      `json:"type"          bson:"type"`
	Transaction   *primitive.ObjectID     `json:"transaction"   bson:"transaction"`
	TransactionId *string                 `json:"transactionId" bson:"transactionId"`
	RequestId     *primitive.ObjectID     `json:"requestId"     bson:"requestId"`
	ItemName      *AttributeEmbedded      `json:"itemName"      bson:"itemName"`
	ItemSKU       *SKUAttributeEmbedded   `json:"itemSKU"       bson:"itemSKU"`
	Stock         *StockAttributeEmbedded `json:"stock"         bson:"stock"`
	Data          *O                      `json:"data"          bson:"data"`
	Details       []TransactionLogDetail  `json:"details"       bson:"details"`
	PIC           UserAttributeEmbedded   `json:"pic"           bson:"pic"`
	UpdatedAt     time.Time               `json:"updatedAt"     bson:"updatedAt"`
	CreatedAt     time.Time               `json:"createdAt"     bson:"createdAt"`
}
