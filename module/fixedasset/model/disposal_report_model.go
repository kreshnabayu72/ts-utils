package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DisposalLastCondition string

const (
	DisposalLastConditionNoSalvageValue = DisposalLastCondition("No Salvage Value")
	DisposalLastConditionDamaged        = DisposalLastCondition("Damaged")
	DisposalLastConditionMissing        = DisposalLastCondition("Missing")
)

type DisposalReport struct {
	ID             primitive.ObjectID      `json:"id,omitempty" bson:"_id,omitempty"`
	Asset          primitive.ObjectID      `json:"asset" bson:"asset"`
	AssetImage     string                  `json:"assetImage" bson:"assetImage"`
	PreviousStatus string                  `json:"previousStatus" bson:"previousStatus"`
	IsDeleted      bool                    `json:"isDeleted" bson:"isDeleted"`
	AssetName      AssetNameEmbedAttr      `json:"assetName" bson:"assetName"`
	AssetBrand     model.AttributeEmbedded `json:"assetBrand" bson:"assetBrand"`
	AssetModel     model.AttributeEmbedded `json:"assetModel" bson:"assetModel"`
	OriginGroup    model.AttributeEmbedded `json:"originGroup" bson:"originGroup"`
	LastCondition  DisposalLastCondition   `json:"lastCondition" bson:"lastCondition"`
	ReportedBy     model.UserEmbedded      `bson:"reportedBy" json:"reportedBy"`

	Notes  string `json:"notes" bson:"notes"`
	Reason string `json:"reason" bson:"reason"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
