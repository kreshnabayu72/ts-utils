package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DamageReport struct {
	ID             primitive.ObjectID      `bson:"_id" json:"_id"`
	Asset          primitive.ObjectID      `bson:"asset" json:"asset"`
	AssetName      AssetNameAttribute      `bson:"assetName" json:"assetName"`
	AssetBrand     model.AttributeEmbedded `bson:"assetBrand" json:"assetBrand"`
	AssetModel     model.AttributeEmbedded `bson:"assetModel" json:"assetModel"`
	AssetGroup     AssetTreeAttribute      `bson:"assetGroup" json:"assetGroup"`
	AssetQR        *string                 `bson:"assetQr" json:"assetQr"`
	AssetRFID      *string                 `bson:"assetRfid" json:"assetRfid"`
	AssetImage     *string                 `bson:"assetImage" json:"assetImage"`
	DamageImage    *string                 `bson:"damageImage" json:"damageImage"`
	PreviousStatus *string                 `bson:"previousStatus" json:"previousStatus"`
	Status         string                  `bson:"status" json:"status"`
	Notes          string                  `bson:"notes" json:"notes"`
	ReportedBy     model.UserEmbedded      `bson:"reportedBy" json:"reportedBy"`
	TagType        string                  `bson:"tagType" json:"tagType"`
	ReportSource   string                  `bson:"reportSource" json:"reportSource"`
	ReportedAt     time.Time               `bson:"reportedAt" json:"reportedAt"`
	RepairedAt     *time.Time              `bson:"repairedAt" json:"repairedAt"`
	CreatedAt      time.Time               `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time               `bson:"updatedAt" json:"updatedAt"`
}
