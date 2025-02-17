package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
)

type Port struct {
	Antenna  *primitive.ObjectID      `bson:"antenna,omitempty" json:"antenna,omitempty"`
	Group    *model.AttributeEmbedded `bson:"group,omitempty"   json:"group,omitempty"`
	Power    int                      `bson:"power"             json:"power"`
	IsActive bool                     `bson:"isActive"          json:"isActive"`
}

type AssetControl struct {
	Active         bool                     `bson:"active"           json:"active"`
	RoleRestricted bool                     `bson:"roleRestricted"   json:"roleRestricted"`
	InFrom         *model.AttributeEmbedded `bson:"inFrom,omitempty" json:"inFrom,omitempty"`
	OutTo          *model.AttributeEmbedded `bson:"outTo,omitempty"  json:"outTo,omitempty"`
}

type IotReader struct {
	ID               primitive.ObjectID         `bson:"_id"                        json:"_id"`
	ImageSmall       *string                    `bson:"imageSmall,omitempty"       json:"imageSmall,omitempty"`
	ImageMedium      *string                    `bson:"imageMedium,omitempty"      json:"imageMedium,omitempty"`
	ImageBig         *string                    `bson:"imageBig,omitempty"         json:"imageBig,omitempty"`
	AliasName        *string                    `bson:"aliasName,omitempty"        json:"aliasName,omitempty"`
	Name             model.NameEmbedded         `bson:"name"                       json:"name"`
	Type             string                     `json:"type"                       bson:"type"`
	SKU              string                     `json:"sku"                        bson:"sku"`
	SerialNumber     string                     `bson:"serialNumber"               json:"serialNumber"`
	TotalPort        int                        `bson:"totalPort"                  json:"totalPort"`
	InputTopic       *string                    `bson:"inputTopic,omitempty"       json:"inputTopic,omitempty"`
	OutputTopic      *string                    `bson:"outputTopic,omitempty"      json:"outputTopic,omitempty"`
	UniqueIdentifier *string                    `bson:"uniqueIdentifier,omitempty" json:"uniqueIdentifier,omitempty"`
	Group            *model.AttributeEmbedded   `bson:"group"                      json:"group"`
	AntennaGroup     *[]model.AttributeEmbedded `bson:"antennaGroup"               json:"antennaGroup"`
	Brand            model.AttributeEmbedded    `bson:"brand"                      json:"brand"`
	Model            model.AttributeEmbedded    `bson:"model"                      json:"model"`
	Manager          model.UserEmbedded         `bson:"manager"                    json:"manager"`
	Status           string                     `bson:"status"                     json:"status"` // Available, Damaged, Missing
	Key              int                        `bson:"key"                        json:"key"`
	Ports            []Port                     `bson:"ports"                      json:"ports"`
	NetworkStatus    string                     `bson:"networkStatus"              json:"networkStatus"` // Online, Offline
	ResponseID       *string                    `bson:"responseId,omitempty"       json:"responseId,omitempty"`
	AssetControl     *AssetControl              `bson:"assetControl,omitempty"     json:"assetControl,omitempty"`
	ReportedBy       *model.UserEmbedded        `bson:"reportedBy"                 json:"reportedBy"`
	ReportedAt       *time.Time                 `bson:"reportedAt"                 json:"reportedAt"`
	CreatedAt        time.Time                  `bson:"createdAt"                  json:"createdAt"`
	UpdatedAt        time.Time                  `bson:"updatedAt"                  json:"updatedAt"`
}
