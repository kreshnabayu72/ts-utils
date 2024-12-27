package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type auditScheduleStatus = string

const (
	AuditScheduleStatus_Active   auditScheduleStatus = "Active"
	AuditScheduleStatus_Inactive auditScheduleStatus = "Inactive"
)

type AuditScheduleCustomValue struct {
	Value     int    `json:"value" bson:"value"`
	Frequency string `json:"frequency" bson:"frequency"`
}

type AuditSchedule struct {
	ID             primitive.ObjectID       `json:"_id,omitempty" bson:"_id,omitempty"`
	Status         auditScheduleStatus      `json:"status" bson:"status"`
	TaskName       string                   `json:"taskName" bson:"taskName"`
	Group          AttributeEmbedded        `json:"group" bson:"group"`
	Category       []AttributeEmbedded      `json:"category" bson:"category"`
	TagType        TagType                  `json:"tagType" bson:"tagType"`
	AuditType      string                   `json:"auditType" bson:"auditType"`
	Frequency      string                   `json:"frequency" bson:"frequency"`
	CustomValue    AuditScheduleCustomValue `json:"customValue" bson:"customValue"`
	Duration       int                      `json:"duration" bson:"duration"`
	NextAudit      time.Time                `json:"nextAudit" bson:"nextAudit"`
	AuditableAsset int                      `json:"auditableAsset" bson:"auditableAsset"`
	ActiveAsset    int                      `json:"activeAsset" bson:"activeAsset"`
	ModifiedBy     UserEmbedded             `json:"modifiedBy" bson:"modifiedBy"`
	CreatedBy      UserEmbedded             `json:"createdBy" bson:"createdBy"`
	CreatedAt      time.Time                `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time                `json:"updatedAt" bson:"updatedAt"`
}
