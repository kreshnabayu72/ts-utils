package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BatchNumberStatus string

const (
	BatchNumberStatusActive      BatchNumberStatus = "Active"
	BatchNumberStatusOnHold      BatchNumberStatus = "On Hold"
	BatchNumberStatusExpired     BatchNumberStatus = "Expired"
	BatchNumberStatusDefect      BatchNumberStatus = "Defect"
	BatchNumberStatusDisposed    BatchNumberStatus = "Disposed"
	BatchNumberStatusUnderReview BatchNumberStatus = "Under Review"
)

type BatchNumber struct {
	ID             primitive.ObjectID  `json:"_id,omitempty"     bson:"_id,omitempty"`
	BatchNumber    string              `json:"batchNumber"       bson:"batchNumber"`
	ItemSKU        primitive.ObjectID  `json:"itemSKU"           bson:"itemSKU"`
	Supplier       *primitive.ObjectID `json:"supplier"          bson:"supplier"`
	Status         BatchNumberStatus   `json:"status"            bson:"status"`
	ProductionDate time.Time           `json:"productionDate"    bson:"productionDate"`
	ExpiredDate    time.Time           `json:"expiredDate"       bson:"expiredDate"`
	CreatedBy      primitive.ObjectID  `json:"createdBy"         bson:"createdBy"`
	ModifiedBy     primitive.ObjectID  `json:"modifiedBy"        bson:"modifiedBy"`
	Source         string              `json:"source"            bson:"source"`
	CreatedAt      time.Time           `json:"createdAt"         bson:"createdAt"`
	UpdatedAt      time.Time           `json:"updatedAt"         bson:"updatedAt"`
	IsCompleted    bool                `json:"isCompleted"       bson:"isCompleted"`
	IsDeleted      bool                `json:"isDeleted"         bson:"isDeleted"`
}
