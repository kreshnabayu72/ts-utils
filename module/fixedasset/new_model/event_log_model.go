package new_model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoutineTaskEventLog struct {
	ID                    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	TaskID                *string            `json:"taskId" bson:"taskId"`
	TaskName              string             `json:"taskName" bson:"taskName"`
	AssetName             *string            `json:"assetName" bson:"assetName"`
	AssetNameWithSequence *string            `json:"assetNameWithSequence" bson:"assetNameWithSequence"`
	Action                string             `json:"action" bson:"action"`
	GroupID               string             `json:"groupId" bson:"groupId"`
	GroupName             string             `json:"groupName" bson:"groupName"`
	ModifiedByID          string             `json:"modifiedById" bson:"modifiedById"`
	ModifiedByFullName    string             `json:"modifiedByFullName" bson:"modifiedByFullName"`
	CreatedAt             time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt             time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type AuditTaskEventLog struct {
	ID                 primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	TaskID             string             `json:"taskId" bson:"taskId"`
	TaskName           string             `json:"taskName" bson:"taskName"`
	Action             string             `json:"action" bson:"action"`
	GroupID            string             `json:"groupId" bson:"groupId"`
	GroupName          string             `json:"groupName" bson:"groupName"`
	AuditType          AuditType          `json:"auditType" bson:"auditType"`
	ModifiedByID       string             `json:"modifiedById" bson:"modifiedById"`
	ModifiedByFullName string             `json:"modifiedByFullName" bson:"modifiedByFullName"`
	CreatedAt          time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt          time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type AuditEventLog struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Note         *string            `json:"note" bson:"note"`
	Duration     *int               `json:"duration" bson:"duration"`
	TaskID       string             `json:"taskId" bson:"taskId"`
	Action       string             `json:"action" bson:"action"`
	UserID       string             `json:"userId" bson:"userId"`
	UserFullName string             `json:"userFullName" bson:"userFullName"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}
