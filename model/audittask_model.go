package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuditTask struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	TransactionID   string              `bson:"transactionId,omitempty"   json:"transactionId,omitempty"`
	Schedule        primitive.ObjectID  `bson:"schedule"                  json:"schedule"                  validate:"required"`
	TaskName        string              `bson:"taskName,omitempty"        json:"taskName,omitempty"`
	Group           AttributeEmbedded   `bson:"group"                     json:"group"`
	AuditType       string              `bson:"auditType,omitempty"       json:"auditType,omitempty"       validate:"omitempty,oneof=Existence Condition"`
	TagType         string              `bson:"tagType,omitempty"         json:"tagType,omitempty"         validate:"omitempty,oneof='Non TAG' QR RFID 'RFID & QR'"`
	Category        []AttributeEmbedded `bson:"category,omitempty"        json:"category,omitempty"`
	Status          string              `bson:"status,omitempty"          json:"status,omitempty"          validate:"omitempty,oneof='On Schedule' 'In Audit' 'Waiting for Approval' 'Grace Period' Overdue Completed Canceled Stopped"`
	CreatedBy       UserEmbedded        `bson:"createdBy"                 json:"createdBy"`
	Duration        int                 `bson:"duration,omitempty"        json:"duration,omitempty"`
	AuditedAsset    int                 `bson:"auditedAsset,omitempty"    json:"auditedAsset,omitempty"`
	TotalAsset      int                 `bson:"totalAsset,omitempty"      json:"totalAsset,omitempty"`
	StartDate       *time.Time          `bson:"startDate,omitempty"       json:"startDate,omitempty"`
	FinishedDate    *time.Time          `bson:"finishedDate,omitempty"    json:"finishedDate,omitempty"`
	GracePeriodDate *time.Time          `bson:"gracePeriodDate,omitempty" json:"gracePeriodDate,omitempty"`
	ApprovalRound   int                 `bson:"approvalRound,omitempty"   json:"approvalRound,omitempty"                                                                                                                                default:"1"`
}
