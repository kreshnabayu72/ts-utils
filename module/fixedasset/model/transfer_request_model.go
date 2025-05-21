package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransferRequest struct {
	ID               primitive.ObjectID      `json:"id,omitempty" bson:"_id,omitempty"`
	Transaction      primitive.ObjectID      `json:"transaction" bson:"transaction"`
	AssetParent      *primitive.ObjectID     `json:"assetParent" bson:"assetParent"`
	Asset            primitive.ObjectID      `json:"asset" bson:"asset"`
	Manager          model.UserEmbedded      `json:"manager" bson:"manager"`
	OriginGroup      model.AttributeEmbedded `json:"originGroup" bson:"originGroup"`
	DestinationGroup model.AttributeEmbedded `json:"destinationGroup" bson:"destinationGroup"`
	AssetImageBig    string                  `json:"assetImageBig" bson:"assetImageBig"`
	AssetImageMedium string                  `json:"assetImageMedium" bson:"assetImageMedium"`
	AssetImageSmall  string                  `json:"assetImageSmall" bson:"assetImageSmall"`
	AssetName        AssetNameEmbedAttr      `json:"assetName" bson:"assetName"`
	AssetBrand       model.AttributeEmbedded `json:"assetBrand" bson:"assetBrand"`
	AssetModel       model.AttributeEmbedded `json:"assetModel" bson:"assetModel"`
	Status           TransferStatus          `json:"status" bson:"status"`
	AssetTagType     string                  `json:"assetTagType" bson:"assetTagType"`
	AssetQr          string                  `json:"assetQr" bson:"assetQr"`
	AssetRfid        string                  `json:"assetRfid" bson:"assetRfid"`
	Notes            string                  `json:"notes" bson:"notes"`
	HasChildren      bool                    `json:"hasChildren" bson:"hasChildren"`
	CreatedAt        time.Time               `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time               `json:"updatedAt" bson:"updatedAt"`
}
