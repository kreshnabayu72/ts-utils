package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ImportAsset struct {
	ID primitive.ObjectID `bson:"_id,omitempty"       json:"_id,omitempty"` // MongoDB ObjectID

	Group       string `bson:"group"                json:"group"`
	Category    string `bson:"category"        json:"category"`
	Name        string `bson:"name"        json:"name"`
	Qty         string `bson:"qty"        json:"qty"`
	Brand       string `bson:"brand"        json:"brand"`
	Model       string `bson:"model"        json:"model"`
	Measurement string `bson:"measurement"        json:"measurement"`

	//not used in basic
	GroupCode          string `bson:"groupCode"        json:"groupCode"`
	CategoryCode       string `bson:"categoryCode"        json:"categoryCode"`
	NameCode           string `bson:"nameCode"        json:"nameCode"`
	Currency           string `bson:"currency"        json:"currency"`
	PurchaseCost       string `bson:"purchaseCost"        json:"purchaseCost"`
	TangibleAssetGroup string `bson:"tangibleAssetGroup"        json:"tangibleAssetGroup"`
	UsefulLife         string `bson:"usefulLife"        json:"usefulLife"`
	PurchaseDate       string `bson:"purchaseDate"        json:"purchaseDate"`
	AliasName          string `bson:"aliasName"        json:"aliasName"`

	Key       int       `bson:"key,omitempty"       json:"key,omitempty"`
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
