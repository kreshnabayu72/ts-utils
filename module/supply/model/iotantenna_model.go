package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
)

type IotAntenna struct {
	ID               primitive.ObjectID      `bson:"_id"                        json:"_id"`
	ImageSmall       *string                 `bson:"imageSmall,omitempty"       json:"imageSmall,omitempty"`
	ImageMedium      *string                 `bson:"imageMedium,omitempty"      json:"imageMedium,omitempty"`
	ImageBig         *string                 `bson:"imageBig,omitempty"         json:"imageBig,omitempty"`
	Name             model.NameEmbedded      `bson:"name"                       json:"name"`
	SerialNumber     string                  `bson:"serialNumber"               json:"serialNumber"               validate:"required"`
	UniqueIdentifier *string                 `bson:"uniqueIdentifier,omitempty" json:"uniqueIdentifier,omitempty"`
	Brand            model.AttributeEmbedded `bson:"brand"                      json:"brand"`
	Model            model.AttributeEmbedded `bson:"model"                      json:"model"`
	Manager          model.UserEmbedded      `bson:"manager"                    json:"manager"`
	ReportedDate     *time.Time              `bson:"reportedDate,omitempty"     json:"reportedDate,omitempty"`
	Status           string                  `bson:"status"                     json:"status"                     validate:"oneof=Available Damaged Missing"`
	CreatedAt        time.Time               `bson:"createdAt"                  json:"createdAt"`
	UpdatedAt        time.Time               `bson:"updatedAt"                  json:"updatedAt"`
}
