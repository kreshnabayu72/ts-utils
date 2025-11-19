package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DisposalRequest struct {
	ID            primitive.ObjectID      `json:"id,omitempty" bson:"_id,omitempty"`
	Transaction   primitive.ObjectID      `json:"transaction" bson:"transaction"`
	Asset         primitive.ObjectID      `json:"asset" bson:"asset"`
	Manager       model.UserEmbedded      `json:"manager" bson:"manager"`
	OriginGroup   model.AttributeEmbedded `json:"originGroup" bson:"originGroup"`
	DisposalGroup model.AttributeEmbedded `json:"disposalGroup" bson:"disposalGroup"`
	AssetImage    string                  `json:"assetImage" bson:"assetImage"`
	AssetName     AssetNameEmbedAttr      `json:"assetName" bson:"assetName"`
	AssetBrand    model.AttributeEmbedded `json:"assetBrand" bson:"assetBrand"`
	AssetModel    model.AttributeEmbedded `json:"assetModel" bson:"assetModel"`
	Status        DisposalStatus          `json:"status" bson:"status"`
	AssetTagType  string                  `json:"assetTagType" bson:"assetTagType"`
	AssetQr       string                  `json:"assetQr" bson:"assetQr"`
	AssetRfid     string                  `json:"assetRfid" bson:"assetRfid"`
	Method        string                  `json:"method" bson:"method"`

	Reason string `json:"reason" bson:"reason"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
