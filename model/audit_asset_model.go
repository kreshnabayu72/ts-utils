package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuditAsset struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Schedule         primitive.ObjectID `json:"schedule" bson:"schedule"`
	Asset            primitive.ObjectID `json:"asset" bson:"asset"`
	TagType          TagType            `json:"tagType" bson:"tagType"`
	AssetImageSmall  *string            `json:"assetImageSmall" bson:"assetImageSmall"`
	AssetImageMedium *string            `json:"assetImageMedium" bson:"assetImageMedium"`
	AssetImageBig    *string            `json:"assetImageBig" bson:"assetImageBig"`
	AssetName        AssetNameEmbedAttr `json:"assetName" bson:"assetName"`
	AssetGroup       AttributeEmbedded  `json:"assetGroup" bson:"assetGroup"`
	AssetBrand       AttributeEmbedded  `json:"assetBrand" bson:"assetBrand"`
	AssetModel       AttributeEmbedded  `json:"assetModel" bson:"assetModel"`
	AssetCategory    AttributeEmbedded  `json:"assetCategory" bson:"assetCategory"`
	IsActive         bool               `json:"isActive" bson:"isActive"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt" bson:"updatedAt"`
}
