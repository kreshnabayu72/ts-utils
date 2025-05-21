package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BatchNumberStatus string

const (
	BatchNumberStatusActive   BatchNumberStatus = "Active"
	BatchNumberStatusOnHold   BatchNumberStatus = "On Hold"
	BatchNumberStatusExpired  BatchNumberStatus = "Expired"
	BatchNumberStatusDefect   BatchNumberStatus = "Defect"
	BatchNumberStatusDisposed BatchNumberStatus = "Disposed"
)

type BatchNumber struct {
	ID             primitive.ObjectID `json:"_id,omitempty"     bson:"_id,omitempty"`
	BatchNumber    string             `json:"batchNumber"       bson:"batchNumber"`
	ItemSKU        primitive.ObjectID `json:"itemSKU"           bson:"itemSKU"`
	Supplier       primitive.ObjectID `json:"supplier"          bson:"supplier"`
	Status         BatchNumberStatus  `json:"status"            bson:"status"`
	ProductionDate time.Time          `json:"productionDate"    bson:"productionDate"`
	ExpiredDate    time.Time          `json:"expiredDate"       bson:"expiredDate"`
	CreatedAt      time.Time          `json:"createdAt"         bson:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"         bson:"updatedAt"`
}
