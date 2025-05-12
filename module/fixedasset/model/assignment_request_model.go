package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssignmentStatus string

const (
	AssignStatusAssigned = AssignmentStatus("Assigned")
	AssignStatusDamaged  = AssignmentStatus("Reported Damaged")
	AssignStatusMissing  = AssignmentStatus("Reported Missing")
	AssignStatusReturned = AssignmentStatus("Returned")
)

type AssignmentRequest struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	Transaction      primitive.ObjectID      `json:"transaction" bson:"transaction"`
	TransactionId    string                  `json:"transactionId" bson:"transactionId"`
	Asset            primitive.ObjectID      `json:"asset" bson:"asset"`
	AssetId          string                  `json:"assetId" bson:"assetId"`
	AssetParent      *primitive.ObjectID     `json:"assetParent" bson:"assetParent"`
	AssetImageSmall  *string                 `json:"assetImageSmall" bson:"assetImageSmall"`
	AssetImageMedium *string                 `json:"assetImageMedium" bson:"assetImageMedium"`
	AssetImageBig    *string                 `json:"assetImageBig" bson:"assetImageBig"`
	Status           AssignmentStatus        `json:"status" bson:"status"`
	AssetName        AssetNameEmbedAttr      `json:"assetName" bson:"assetName"`
	AssetCategory    model.AttributeEmbedded `json:"assetCategory" bson:"assetCategory"`
	AssetGroup       model.AttributeEmbedded `json:"assetGroup" bson:"assetGroup"`
	AssetBrand       model.AttributeEmbedded `json:"assetBrand" bson:"assetBrand"`
	AssetModel       model.AttributeEmbedded `json:"assetModel" bson:"assetModel"`

	Manager    model.UserEmbedded `json:"manager" bson:"manager"`
	AssignedTo model.UserEmbedded `json:"assignedTo" bson:"assignedTo"`

	Notes string `json:"notes" bson:"notes"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
