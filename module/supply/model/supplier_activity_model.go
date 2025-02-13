package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SupplierActivity struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	TransactionID string             `json:"transactionId" bson:"transactionId"`
	Supplier      primitive.ObjectID `json:"supplier" bson:"supplier"`
	PIC           primitive.ObjectID `json:"pic" bson:"pic"`
	Status        string             `json:"status" bson:"status"`
	Transaction   string             `json:"transaction" bson:"transaction"`
	Group         primitive.ObjectID `json:"group" bson:"group"`
	TotalItem     int                `json:"totalItem" bson:"totalItem"`
	TotalStock    int                `json:"totalStock" bson:"totalStock"`
	FinishedAt    *time.Time         `json:"finishedAt" bson:"finishedAt"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
}
