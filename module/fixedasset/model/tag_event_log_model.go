package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TagEventLogType string

const (
	TELPairTag    TagEventLogType = "Pair TAG"
	TELUnpairTag  TagEventLogType = "Unpair TAG"
	TELReplaceTag TagEventLogType = "Replace TAG"
	TELChangeTag  TagEventLogType = "Change TAG"
	TELReportTag  TagEventLogType = "Report TAG"
)

type TagEventLog struct {
	ID              primitive.ObjectID `json:"_id,omitempty"   bson:"_id,omitempty"`
	Group           GroupAttribute     `json:"group"          bson:"group"`
	Type            string             `json:"type"            bson:"type"`
	TotalAssets     int                `json:"totalAssets"          bson:"totalAssets"`
	PreviousTagType TagType            `json:"previousTagType"          bson:"previousTagType"`
	TagType         TagType            `json:"tagType"          bson:"tagType"`
	Rfid            string             `json:"rfid"          bson:"rfid"`
	Qr              string             `json:"qr"          bson:"qr"`
	AssetName       AssetNameAttribute `json:"assetName" bson:"assetName"`
	UserFullName    string             `json:"userFullName"          bson:"userFullName"`
	ReportedBy      string             `json:"reportedBy"          bson:"reportedBy"`
	ReportedAt      *time.Time         `json:"reportedAt"       bson:"reportedAt"`
	FinishedAt      *time.Time         `json:"finishedAt"       bson:"finishedAt"`
	CreatedAt       time.Time          `json:"createdAt"       bson:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt"       bson:"updatedAt"`
}
