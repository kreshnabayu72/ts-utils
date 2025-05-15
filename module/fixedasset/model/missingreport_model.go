package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MissingReport struct {
	ID               primitive.ObjectID      `bson:"_id" json:"_id"`
	Asset            primitive.ObjectID      `bson:"asset" json:"asset"`
	AssetName        AssetNameAttribute      `bson:"assetName" json:"assetName"`
	AssetAliasName   *string                 `bson:"assetAliasName" json:"assetAliasName"`
	AssetBrand       model.AttributeEmbedded `bson:"assetBrand" json:"assetBrand"`
	AssetModel       model.AttributeEmbedded `bson:"assetModel" json:"assetModel"`
	AssetGroup       AssetTreeAttribute      `bson:"assetGroup" json:"assetGroup"`
	AssetQR          *string                 `bson:"assetQr" json:"assetQr"`
	AssetRFID        *string                 `bson:"assetRfid" json:"assetRfid"`
	AssetImageBig    *string                 `bson:"assetImageBig" json:"assetImageBig"`
	AssetImageMedium *string                 `bson:"assetImageMedium" json:"assetImageMedium"`
	AssetImageSmall  *string                 `bson:"assetImageSmall" json:"assetImageSmall"`
	PreviousStatus   *string                 `bson:"previousStatus" json:"previousStatus"`
	Status           string                  `bson:"status" json:"status"`
	Notes            string                  `bson:"notes" json:"notes"`
	ReportedBy       model.UserEmbedded      `bson:"reportedBy" json:"reportedBy"`
	TagType          string                  `bson:"tagType" json:"tagType"`
	ReportSource     string                  `bson:"reportSource" json:"reportSource"`
	ReportedAt       time.Time               `bson:"reportedAt" json:"reportedAt"`
	FoundAt          *time.Time              `bson:"foundAt" json:"foundAt"`
	CreatedAt        time.Time               `bson:"createdAt" json:"createdAt"`
	UpdatedAt        time.Time               `bson:"updatedAt" json:"updatedAt"`
}
