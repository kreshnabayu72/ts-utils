package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TagAuditLog is one Event Log row of the Audit TAG tab.
type TagAuditLog struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`

	// SessionID is the hardware-tagging session id and the upsert key.
	SessionID string `json:"sessionId" bson:"sessionId"`
	TagType   string `json:"tagType" bson:"tagType"`
	Found     int    `json:"found" bson:"found"`
	NotFound  int    `json:"notFound" bson:"notFound"`

	// NotFoundAction is empty when NotFound is zero.
	NotFoundAction *TagAuditLogNotFoundAction `json:"notFoundAction" bson:"notFoundAction"`

	HardwareID       *string `json:"hardwareId" bson:"hardwareId"`
	HardwareSnapshot *string `json:"hardwareSnapshot" bson:"hardwareSnapshot"`

	Rows []TagAuditLogRow `json:"rows" bson:"rows"`

	CreatedByID       string    `json:"createdById" bson:"createdById"`
	CreatedBySnapshot string    `json:"createdBySnapshot" bson:"createdBySnapshot"`
	CreatedAt         time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt" bson:"updatedAt"`
}

type TagAuditLogNotFoundAction string

const (
	TagAuditLogNotFoundActionMarkedDamagedMissing TagAuditLogNotFoundAction = "marked_damaged_missing"
	TagAuditLogNotFoundActionKeptPreviousStatus   TagAuditLogNotFoundAction = "kept_previous_status"
)

type TagAuditLogRow struct {
	TagID       string                `json:"tagId" bson:"tagId"`
	Section     TagAuditLogRowSection `json:"section" bson:"section"`
	TagType     string                `json:"tagType" bson:"tagType"`
	Code        string                `json:"code" bson:"code"`
	DeviceName  *string               `json:"deviceName" bson:"deviceName"`
	SKU         *string               `json:"sku" bson:"sku"`
	StatusAfter string                `json:"statusAfter" bson:"statusAfter"`
}

type TagAuditLogRowSection string

const (
	TagAuditLogRowSectionFound    TagAuditLogRowSection = "found"
	TagAuditLogRowSectionNotFound TagAuditLogRowSection = "not_found"
)
