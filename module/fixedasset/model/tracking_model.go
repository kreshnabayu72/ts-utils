package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TrackingStatus = string

const (
	TrackingOnTracking               TrackingStatus = "On Tracking"
	TrackingTrackedByUser            TrackingStatus = "Tracked by User"
	TrackingTrackedByMaintenance     TrackingStatus = "Tracked by Maintenance"
	TrackingTrackedByAudit           TrackingStatus = "Tracked by Audit"
	TrackingTrackedByIOT             TrackingStatus = "Tracked by IOT"
	TrackingTrackedByTracking        TrackingStatus = "Tracked by Tracking"
	TrackingReportedPermanentMissing TrackingStatus = "Reported as Permanently Missing"
	TrackingPermanentlyMissing       TrackingStatus = "Permanently Missing"
	TrackingReportedDisposal         TrackingStatus = "Reported Disposal"
	TrackingOnDisposalProcess        TrackingStatus = "On Disposal Process"
	TrackingDisposed                 TrackingStatus = "Disposed"
	TrackingTAGReported              TrackingStatus = "TAG Reported"
)

type TrackingCondition = string

const (
	TrackingNormal  TrackingCondition = "Asset found, normal condition"
	TrackingDamaged TrackingCondition = "Asset found, damaged condition"
	TrackingMissing TrackingCondition = "TAG found, asset missing"
)

type ReportSource = string

const (
	ReportSourceAudit       ReportSource = "Audit Asset"
	ReportSourceAvailable   ReportSource = "Available Asset"
	ReportSourceBorrowed    ReportSource = "Borrowed Asset"
	ReportSourceAssigned    ReportSource = "Assigned Asset"
	ReportSourceTransfer    ReportSource = "Transfer Asset"
	ReportSourceMaintenance ReportSource = "Maintenance Asset"
	ReportSourceTracking    ReportSource = "Tracking Asset"
)

type Tracking struct {
	ID              primitive.ObjectID      `json:"_id,omitempty" bson:"_id,omitempty"`
	ReportedBy      model.UserEmbedded      `json:"reportedBy" bson:"reportedBy"`
	Asset           primitive.ObjectID      `json:"asset" bson:"asset"`
	AssetImage      string                  `json:"assetImage" bson:"assetImage"`
	Address         string                  `json:"address" bson:"address"`
	Coordinate      *string                 `json:"coordinate" bson:"coordinate"`
	AssetName       AssetNameEmbedAttr      `json:"assetName" bson:"assetName"`
	AssetAliasName  string                  `json:"assetAliasName" bson:"assetAliasName"`
	AssetGroup      model.AttributeEmbedded `json:"assetGroup" bson:"assetGroup"`
	AssetBrand      model.AttributeEmbedded `json:"assetBrand" bson:"assetBrand"`
	AssetModel      model.AttributeEmbedded `json:"assetModel" bson:"assetModel"`
	FoundGroup      model.AttributeEmbedded `json:"foundGroup" bson:"foundGroup"`
	Status          TrackingStatus          `json:"status" bson:"status"`
	AssetTagType    TagType                 `json:"assetTagType" bson:"assetTagType"`
	AssetQR         string                  `json:"assetQr" bson:"assetQr"`
	AssetRFID       string                  `json:"assetRfid" bson:"assetRfid"`
	PreviousStatus  string                  `json:"previousStatus" bson:"previousStatus"`
	FoundBy         model.UserEmbedded      `json:"foundBy" bson:"foundBy"`
	Condition       *TrackingCondition      `json:"condition" bson:"condition"`
	Notes           string                  `json:"notes" bson:"notes"`
	DamagedImage    string                  `json:"damagedImage" bson:"damagedImage"`
	FoundDate       time.Time               `json:"foundDate" bson:"foundDate"`
	IsMisplaced     bool                    `json:"isMisplaced" bson:"isMisplaced"`
	ReportSource    ReportSource            `json:"reportSource" bson:"reportSource"`
	DownTimeSeconds *int                    `json:"downTimeSeconds" bson:"downTimeSeconds"`
	TagReported     TagReported             `json:"tagReported" bson:"tagReported"`
	ApprovalRound   int                     `json:"approvalRound" bson:"approvalRound"`
	CreatedAt       time.Time               `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time               `json:"updatedAt" bson:"updatedAt"`
}
