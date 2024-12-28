package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoutineStatus = string

const (
	RoutineOnSchedule        RoutineStatus = "On Schedule"
	RoutineWaitingAssignment RoutineStatus = "Waiting for Assignment"
	RoutineMaintenance       RoutineStatus = "Under Maintenance"
	RoutineWaitingReview     RoutineStatus = "Waiting for Review"
	RoutineWaitingApproval   RoutineStatus = "Waiting for Approval"
	RoutineOverdue           RoutineStatus = "Overdue"
	RoutineCompleted         RoutineStatus = "Completed"
	RoutineCanceled          RoutineStatus = "Canceled"
	RoutineReportedDamaged   RoutineStatus = "Reported Damaged"
	RoutineStopped           RoutineStatus = "Stopped"
)

type IndicatorType = string

const (
	IndicatorSelection IndicatorType = "Selection"
	IndicatorSingle    IndicatorType = "Single Value"
	IndicatorRange     IndicatorType = "Range Value"
)

type RoutineIndicatorStandardValue struct {
	Type  IndicatorType `json:"type" bson:"type"`
	Value string        `json:"value" bson:"value"`
}

type RoutineIndicatorGroup = string

const (
	IndicatorGroupNormal        RoutineIndicatorGroup = "Normal"
	IndicatorGroupAboveStandard RoutineIndicatorGroup = "Above The Standard"
	IndicatorGroupBelowStandard RoutineIndicatorGroup = "Below The Standard"
)

type RoutineIndicator struct {
	ID                  primitive.ObjectID            `json:"_id,omitempty" bson:"_id,omitempty"`
	Name                string                        `json:"name" bson:"name"`
	Result              string                        `json:"result" bson:"result"`
	DataType            string                        `json:"dataType" bson:"dataType"`
	Standard            string                        `json:"standard" bson:"standard"`
	Value               string                        `json:"value" bson:"value"`
	IsMandatory         bool                          `json:"isMandatory" bson:"isMandatory"`
	IsHaveStandardValue bool                          `json:"isHaveStandardValue" bson:"isHaveStandardValue"`
	Options             []string                      `json:"options" bson:"options"`
	StandardValue       RoutineIndicatorStandardValue `json:"standardValue" bson:"standardValue"`
	Measurement         string                        `json:"measurement" bson:"measurement"`
	Group               *RoutineIndicatorGroup        `json:"group" bson:"group"`
}

type RoutineTaskAdditionalCost struct {
	TotalPrice   int    `json:"totalPrice" bson:"totalPrice"`
	PaymentProof string `json:"paymentProof"`
	Currency     string `json:"currency" bson:"currency"`
}

type RoutineTask struct {
	ID                     primitive.ObjectID        `json:"_id,omitempty" bson:"_id,omitempty"`
	AssetImageSmall        string                    `json:"assetImageSmall" bson:"assetImageSmall"`
	AssetImageMedium       string                    `json:"assetImageMedium" bson:"assetImageMedium"`
	AssetImageBig          string                    `json:"assetImageBig" bson:"assetImageBig"`
	Address                string                    `json:"address" bson:"address"`
	Coordinate             string                    `json:"coordinate" bson:"coordinate"`
	AssetName              AssetNameEmbedAttr        `json:"assetName" bson:"assetName"`
	AssetGroup             AttributeEmbedded         `json:"assetGroup" bson:"assetGroup"`
	AssetBrand             AttributeEmbedded         `json:"assetBrand" bson:"assetBrand"`
	AssetModel             AttributeEmbedded         `json:"assetModel" bson:"assetModel"`
	Notes                  string                    `json:"notes" bson:"notes"`
	Status                 RoutineStatus             `json:"status" bson:"status"`
	Manager                UserEmbedded              `json:"manager" bson:"manager"`
	Staff                  []UserEmbedded            `json:"staff" bson:"staff"`
	Schedule               primitive.ObjectID        `json:"schedule" bson:"schedule"`
	TaskName               string                    `json:"taskName" bson:"taskName"`
	MaintenanceIndicator   []RoutineIndicator        `json:"maintenanceIndicator" bson:"maintenanceIndicator"`
	Asset                  primitive.ObjectID        `json:"asset" bson:"asset"`
	Reviewer               primitive.ObjectID        `json:"reviewer" bson:"reviewer"`
	ReviewedAt             time.Time                 `json:"reviewedAt" bson:"reviewedAt"`
	AssetQR                bool                      `json:"assetQr,omitempty" bson:"assetQr,omitempty"`
	AssetRFID              bool                      `json:"assetRfid,omitempty" bson:"assetRfid,omitempty"`
	ServiceCenterID        primitive.ObjectID        `json:"serviceCenterId" bson:"serviceCenterId"`
	ServiceCenterName      string                    `json:"serviceCenterName" bson:"serviceCenterName"`
	MaintenanceImageSmall  string                    `json:"maintenanceImageSmall" bson:"maintenanceImageSmall"`
	MaintenanceImageMedium string                    `json:"maintenanceImageMedium" bson:"maintenanceImageMedium"`
	MaintenanceImageBig    string                    `json:"maintenanceImageBig" bson:"maintenanceImageBig"`
	WarrantyEndDate        time.Time                 `json:"warrantyEndDate" bson:"warrantyEndDate"`
	StartAt                time.Time                 `json:"startAt" bson:"startAt"`
	TagReported            TagReported               `json:"tagReported" bson:"tagReported"`
	Duration               int                       `json:"duration" bson:"duration"`
	AdditionalCost         RoutineTaskAdditionalCost `json:"additionalCost" bson:"additionalCost"`
	IsOverdue              bool                      `json:"isOverdue" bson:"isOverdue"`
	ApprovalRound          int                       `json:"approvalRound" bson:"approvalRound"`
	SerialNumber           *string                   `json:"serialNumber" bson:"serialNumber"`
	CreatedAt              time.Time                 `json:"createdAt" bson:"createdAt"`
	UpdatedAt              time.Time                 `json:"updatedAt" bson:"updatedAt"`
}
