package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
)

type Reader struct {
	ID               primitive.ObjectID      `json:"_id,omitempty"    bson:"_id,omitempty"`
	ImageSmall       *string                 `json:"imageSmall"       bson:"imageSmall"`
	ImageMedium      *string                 `json:"imageMedium"      bson:"imageMedium"`
	ImageBig         *string                 `json:"imageBig"         bson:"imageBig"`
	AliasName        *string                 `json:"aliasName"        bson:"aliasName"`
	SerialNumber     string                  `json:"serialNumber"     bson:"serialNumber"`
	Name             model.NameEmbedded      `json:"name"             bson:"name"`
	Group            model.AttributeEmbedded `json:"group"            bson:"group"`
	Manager          model.UserEmbedded      `json:"manager"          bson:"manager"`
	Type             string                  `json:"type"             bson:"type"`
	SKU              string                  `json:"sku"              bson:"sku"`
	Key              int                     `json:"key"              bson:"key"`
	Status           string                  `json:"status"           bson:"status"`           // Available, Damaged, Missing
	LastReportReason string                  `json:"lastReportReason" bson:"lastReportReason"` // Damaged, Missing
	NameNumber       int                     `json:"nameNumber"       bson:"nameNumber"`
	QRCode           string                  `json:"qrCode"           bson:"qrCode"`
	ReportedBy       *model.UserEmbedded     `json:"reportedBy"       bson:"reportedBy"`
	ReportedAt       *time.Time              `json:"reportedAt"       bson:"reportedAt"`
	CreatedAt        time.Time               `json:"createdAt"        bson:"createdAt"`
	UpdatedAt        time.Time               `json:"updatedAt"        bson:"updatedAt"`
}
