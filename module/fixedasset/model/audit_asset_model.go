package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuditAsset struct {
	ID            primitive.ObjectID      `json:"_id,omitempty" bson:"_id,omitempty"`
	Schedule      primitive.ObjectID      `json:"schedule" bson:"schedule"`
	Asset         primitive.ObjectID      `json:"asset" bson:"asset"`
	TagType       TagType                 `json:"tagType" bson:"tagType"`
	AssetImage    *string                 `json:"assetImage" bson:"assetImage"`
	AssetName     AssetNameEmbedAttr      `json:"assetName" bson:"assetName"`
	AssetGroup    model.AttributeEmbedded `json:"assetGroup" bson:"assetGroup"`
	AssetBrand    model.AttributeEmbedded `json:"assetBrand" bson:"assetBrand"`
	AssetModel    model.AttributeEmbedded `json:"assetModel" bson:"assetModel"`
	AssetCategory model.AttributeEmbedded `json:"assetCategory" bson:"assetCategory"`
	IsActive      bool                    `json:"isActive" bson:"isActive"`
	CreatedAt     time.Time               `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time               `json:"updatedAt" bson:"updatedAt"`
}
