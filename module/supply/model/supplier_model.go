package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SuppliedItem struct {
	ItemName primitive.ObjectID   `json:"itemName" bson:"itemName"`
	ItemSKU  []primitive.ObjectID `json:"itemSKU" bson:"itemSKU"`
}

type Supplier struct {
	ID                 primitive.ObjectID  `json:"_id" bson:"_id,omitempty"`
	IsActive           bool                `json:"isActive" bson:"isActive"`
	Name               string              `json:"name" bson:"name"`
	Items              []SuppliedItem      `json:"items" bson:"items"`
	Country            *string             `json:"country" bson:"country"`
	State              *string             `json:"state" bson:"state"`
	City               *string             `json:"city" bson:"city"`
	ZipCode            *string             `json:"zipCode" bson:"zipCode"`
	CompanyAddress     *string             `json:"companyAddress" bson:"companyAddress"`
	CompanyPhoneNumber *string             `json:"companyPhoneNumber" bson:"companyPhoneNumber"`
	PICName            *string             `json:"picName"  bson:"picName"`
	PICEmail           *string             `json:"picEmail" bson:"picEmail"`
	PICContactNumber   *string             `json:"picContactNumber" bson:"picContactNumber"`
	ModifiedBy         *primitive.ObjectID `json:"modifiedBy" bson:"modifiedBy"`
	UpdatedAt          time.Time           `json:"updatedAt" bson:"updatedAt"`
	CreatedAt          time.Time           `json:"createdAt" bson:"createdAt"`
}
