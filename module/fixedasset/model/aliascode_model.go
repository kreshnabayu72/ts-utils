package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AliasCodeData struct {
	Name      string `json:"name" bson:"name"`
	IsChecked bool   `json:"isChecked" bson:"isChecked"`
	Value     string `json:"value" bson:"value"`
}

type AliasCode struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	IsActive       bool               `bson:"isActive" json:"isActive"`
	CompanyInitial string             `bson:"companyInitial" json:"companyInitial"`
	Separator      string             `bson:"separator" json:"separator"`
	Data           []AliasCodeData    `bson:"data" json:"data"`
	CreatedAt      time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt" json:"updatedAt"`
}
