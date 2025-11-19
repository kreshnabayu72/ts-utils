package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reader struct {
	ID               primitive.ObjectID       `bson:"_id"                        json:"_id"`
	Image            *string                  `bson:"image,omitempty"            json:"image,omitempty"`
	AliasName        *string                  `bson:"aliasName,omitempty"        json:"aliasName,omitempty"`
	Name             model.NameEmbedded       `bson:"name"                       json:"name"`
	SerialNumber     string                   `bson:"serialNumber"               json:"serialNumber"`
	FixedAssetGroup  *model.AttributeEmbedded `bson:"fixedAssetGroup,omitempty"  json:"fixedAssetGroup,omitempty"`
	SupplyAssetGroup *model.AttributeEmbedded `bson:"supplyAssetGroup,omitempty" json:"supplyAssetGroup,omitempty"`
	Brand            model.AttributeEmbedded  `bson:"brand"                      json:"brand"`
	Model            model.AttributeEmbedded  `bson:"model"                      json:"model"`
	Manager          model.AttributeEmbedded  `bson:"manager"                    json:"manager"`
	Type             string                   `bson:"type"                       json:"type"`
	SKU              string                   `bson:"sku"                        json:"sku"`
	Key              int                      `bson:"key"                        json:"key"`
	Status           string                   `bson:"status"                     json:"status"`
	LastReportReason *string                  `bson:"lastReportReason"           json:"lastReportReason"`
	QRCode           string                   `bson:"qrCode"                     json:"qrCode"`
	LastReportDate   *time.Time               `bson:"lastReportDate,omitempty"   json:"lastReportDate,omitempty"`
	ReportedBy       *model.UserEmbedded      `bson:"reportedBy"                 json:"reportedBy"`
	CreatedAt        time.Time                `bson:"createdAt"                  json:"createdAt"`
	UpdatedAt        time.Time                `bson:"updatedAt"                  json:"updatedAt"`
}
