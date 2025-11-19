package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuditAssetCondition = string

const (
	AuditAssetNormal               AuditAssetCondition = "Normal"
	AuditAssetReportedDamaged      AuditAssetCondition = "Reported Damaged"
	AuditAssetReportedMissing      AuditAssetCondition = "Reported Missing"
	AuditAssetTagFoundAssetMissing AuditAssetCondition = "TAG Found Asset Missing"
)

type AuditReportedImageAttr struct {
	Path         string `json:"path" bson:"path"`
	OriginalName string `json:"originalName" bson:"originalName"`
	Mimetype     string `json:"mimetype" bson:"mimetype"`
}

type AuditAssetFoundByAttr struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	FullName string             `json:"fullName" bson:"fullName"`
	Key      int                `json:"key" bson:"key"`
}

type AuditScheduledAsset struct {
	ID            primitive.ObjectID      `json:"_id,omitempty" bson:"_id,omitempty"`
	TransactionID string                  `json:"transactionId" bson:"transactionId"`
	Note          string                  `json:"note" bson:"note"`
	TagType       TagType                 `json:"tagType" bson:"tagType"`
	Task          primitive.ObjectID      `json:"task" bson:"task"`
	Asset         primitive.ObjectID      `json:"asset" bson:"asset"`
	AssetGroup    model.AttributeEmbedded `json:"assetGroup" bson:"assetGroup"`
	AssetImage    *string                 `json:"assetImage" bson:"assetImage"`
	Address       *string                 `json:"address" bson:"address"`
	Coordinate    *string                 `json:"coordinate" bson:"coordinate"`
	ReportedImage *string                 `json:"reportedImage" bson:"reportedImage"`
	// ReportedImage          AuditReportedImageAttr  `json:"reportedImage" bson:"reportedImage"`
	AssetName              AssetNameEmbedAttr      `json:"assetName" bson:"assetName"`
	AssetModel             model.AttributeEmbedded `json:"assetModel" bson:"assetModel"`
	AssetBrand             model.AttributeEmbedded `json:"assetBrand" bson:"assetBrand"`
	AssetStatus            string                  `json:"assetStatus" bson:"assetStatus"`
	Condition              *AuditAssetCondition    `json:"condition" bson:"condition"`
	AssetOriginalCondition string                  `json:"assetOriginalCondition" bson:"assetOriginalCondition"`
	LocationFoundGroup     model.AttributeEmbedded `json:"locationFoundGroup" bson:"locationFoundGroup"`
	FoundBy                AuditAssetFoundByAttr   `json:"foundBy" bson:"foundBy"`
	AuditedOn              *time.Time              `json:"auditedOn" bson:"auditedOn"`
	FoundDate              *time.Time              `json:"foundDate" bson:"foundDate"`
	SerialNumber           *string                 `json:"serialNumber" bson:"serialNumber"`
	CreatedAt              time.Time               `json:"createdAt" bson:"createdAt"`
	UpdatedAt              time.Time               `json:"updatedAt" bson:"updatedAt"`
}
