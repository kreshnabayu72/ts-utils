package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reader struct {
	ID             primitive.ObjectID       `bson:"_id"                        json:"_id"`
	Image          *string                  `bson:"image,omitempty"       json:"image,omitempty"`
	AliasName      *string                  `bson:"aliasName,omitempty"        json:"aliasName,omitempty"`
	Name           model.NameEmbedded       `bson:"name"                       json:"name"`
	SerialNumber   string                   `bson:"serialNumber"               json:"serialNumber"`
	Group          *model.AttributeEmbedded `bson:"group,omitempty"            json:"group,omitempty"`
	Brand          model.AttributeEmbedded  `bson:"brand"                      json:"brand"`
	Model          model.AttributeEmbedded  `bson:"model"                      json:"model"`
	Manager        model.AttributeEmbedded  `bson:"manager"                    json:"manager"`
	LastReportDate *time.Time               `bson:"lastReportDate,omitempty"     json:"lastReportDate,omitempty"`
	Status         string                   `bson:"status"                     json:"status"`
	Key            int                      `bson:"key"                        json:"key"`
	CreatedAt      time.Time                `bson:"createdAt"                  json:"createdAt"`
	UpdatedAt      time.Time                `bson:"updatedAt"                  json:"updatedAt"`
}
