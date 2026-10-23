package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TagActivateLog is one Event Log row of the Activate TAG tab: an Activate or a Delete session.
type TagActivateLog struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`

	// SessionID is the hardware-tagging session id and the upsert key.
	SessionID string                 `json:"sessionId" bson:"sessionId"`
	Activity  TagActivateLogActivity `json:"activity" bson:"activity"`
	TagType   string                 `json:"tagType" bson:"tagType"`
	Success   int                    `json:"success" bson:"success"`

	// WaitingForApproval is empty for Delete sessions.
	WaitingForApproval *int `json:"waitingForApproval" bson:"waitingForApproval"`
	Failed             int  `json:"failed" bson:"failed"`

	// Hardware* are empty when no handheld reader was used.
	HardwareID       *string `json:"hardwareId" bson:"hardwareId"`
	HardwareSnapshot *string `json:"hardwareSnapshot" bson:"hardwareSnapshot"`

	Rows []TagActivateLogRow `json:"rows" bson:"rows"`

	CreatedByID       string `json:"createdById" bson:"createdById"`
	CreatedBySnapshot string `json:"createdBySnapshot" bson:"createdBySnapshot"`

	// CreatedAt is the session start.
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type TagActivateLogActivity string

const (
	TagActivateLogActivityActivate TagActivateLogActivity = "Activate"
	TagActivateLogActivityDelete   TagActivateLogActivity = "Delete"
)

// TagActivateLogRow is one detail row; catalog snapshots for Activate, StatusBefore for Delete.
type TagActivateLogRow struct {
	TagID        string                   `json:"tagId" bson:"tagId"`
	Section      TagActivateLogRowSection `json:"section" bson:"section"`
	TagType      string                   `json:"tagType" bson:"tagType"`
	Code         string                   `json:"code" bson:"code"`
	DeviceName   *string                  `json:"deviceName" bson:"deviceName"`
	SKU          *string                  `json:"sku" bson:"sku"`
	StatusBefore *string                  `json:"statusBefore" bson:"statusBefore"`
}

type TagActivateLogRowSection string

const (
	TagActivateLogRowSectionActivated       TagActivateLogRowSection = "activated"
	TagActivateLogRowSectionWaitingApproval TagActivateLogRowSection = "waiting_approval"
	TagActivateLogRowSectionDeleted         TagActivateLogRowSection = "deleted"
	TagActivateLogRowSectionFailed          TagActivateLogRowSection = "failed"
)
