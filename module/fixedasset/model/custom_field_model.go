package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomField struct {
	ID                   primitive.ObjectID        `json:"_id,omitempty"        bson:"_id,omitempty"`
	Name                 string                    `json:"name"                 bson:"name"`
	Type                 string                    `json:"type"                 bson:"type"`
	DataType             string                    `json:"dataType"             bson:"dataType"`
	IsRequired           bool                      `json:"isRequired"           bson:"isRequired"`
	OptionValue          []string                  `json:"optionValue"          bson:"optionValue"`
	Value                any                       `json:"value"                bson:"value"`
	Category             []model.AttributeEmbedded `json:"category"             bson:"category"`
	Status               bool                      `json:"status"               bson:"status"`
	IsDeleted            bool                      `json:"isDeleted"            bson:"isDeleted"`
	DeleteFromDatabaseAt *time.Time                `json:"deleteFromDatabaseAt" bson:"deleteFromDatabaseAt"`
	CreatedAt            time.Time                 `json:"createdAt"            bson:"createdAt"`
	UpdatedAt            time.Time                 `json:"updatedAt"            bson:"updatedAt"`
}
