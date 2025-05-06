package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ImportCustomField struct {
	ID primitive.ObjectID `bson:"_id,omitempty"       json:"_id,omitempty"` // MongoDB ObjectID

	Name         string `bson:"name"                json:"name"`
	Type         string `bson:"type"        json:"type"`
	DataType     string `bson:"dataType"        json:"dataType"`
	OptionValue  string `bson:"optionValue"        json:"optionValue"`
	IsRequired   string `bson:"isRequired"        json:"isRequired"`
	Category     string `bson:"category"        json:"category"`
	CategoryCode string `bson:"categoryCode"        json:"categoryCode"` //not used in basic

	Key       int       `bson:"key,omitempty"       json:"key,omitempty"`
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
