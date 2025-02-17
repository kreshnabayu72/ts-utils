package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomField struct {
	ID                   primitive.ObjectID   `json:"_id,omitempty"        bson:"_id,omitempty"`
	Name                 string               `json:"name"                 bson:"name"`
	Type                 string               `json:"type"                 bson:"type"` // Global, Specific
	DataType             string               `json:"dataType"             bson:"dataType"`
	IsRequired           bool                 `json:"isRequired"           bson:"isRequired"`
	OptionValue          []string             `json:"optionValue"          bson:"optionValue"`
	Value                any                  `json:"value"                bson:"value"`
	ItemName             []primitive.ObjectID `json:"itemName"             bson:"itemName"`
	IsActive             bool                 `json:"isActive"             bson:"isActive"`
	IsDeleted            bool                 `json:"isDeleted"            bson:"isDeleted"`
	DeleteFromDatabaseAt *time.Time           `json:"deleteFromDatabaseAt" bson:"deleteFromDatabaseAt"`
	CreatedAt            time.Time            `json:"createdAt"            bson:"createdAt"`
	UpdatedAt            time.Time            `json:"updatedAt"            bson:"updatedAt"`
}
