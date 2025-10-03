package new_model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventLogType string

const (
	EvengLogTypeRoutine EventLogType = "Routine"
	EvengLogTypeAudit   EventLogType = "Audit"
)

type EventLog struct {
	ID                    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	TaskID                string             `json:"taskId" bson:"taskId"`
	TaskName              string             `json:"taskName" bson:"taskName"`
	AssetName             *string            `json:"assetName" bson:"assetName"`
	AssetNameWithSequence *string            `json:"assetNameWithSequence" bson:"assetNameWithSequence"`
	Action                string             `json:"action" bson:"action"`
	Group                 string             `json:"group" bson:"group"`
	Type                  EventLogType       `json:"type" bson:"type"`
	AuditType             *AuditType         `json:"auditType" bson:"auditType"`
	ModifiedBy            string             `json:"modifiedBy" bson:"modifiedBy"`
	CreatedAt             time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt             time.Time          `json:"updatedAt" bson:"updatedAt"`
}
