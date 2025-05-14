package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationSeverity string

const (
	NotificationSeverityInfo    = NotificationSeverity("info")
	NotificationSeverityDanger  = NotificationSeverity("danger")
	NotificationSeverityWarning = NotificationSeverity("warning")
	NotificationSeveritySuccess = NotificationSeverity("success")
)

type NotificationCategory string

const (
	NotificationCategoryCompleted  = NotificationCategory("completed")
	NotificationCategoryRejected   = NotificationCategory("rejected")
	NotificationCategoryOnProgress = NotificationCategory("onProgress")
	NotificationCategoryWarning    = NotificationCategory("warning")
)

type NotificationModule string

const (
	NotificationModuleBorrow             = NotificationModule("Borrow")
	NotificationModuleAssignment         = NotificationModule("Assignment")
	NotificationModuleTransfer           = NotificationModule("Transfer")
	NotificationModuleDisposal           = NotificationModule("Disposal")
	NotificationModuleAudit              = NotificationModule("Audit")
	NotificationModuleMaintenanceRoutine = NotificationModule("Maintenance Routine")
	NotificationModuleRepairTicketing    = NotificationModule("Repair Ticketing")
	NotificationModuleTracking           = NotificationModule("Tracking")
	NotificationModuleTAG                = NotificationModule("TAG")
	NotificationModuleGeneralAsset       = NotificationModule("General Asset")
)

type Notification struct {
	ID         primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	User       primitive.ObjectID     `json:"user" bson:"user"`
	Title      string                 `json:"title"         bson:"title"`
	Detail     string                 `json:"detail"        bson:"detail"`
	IsRead     bool                   `json:"isRead"        bson:"isRead"`
	IsReadOnly bool                   `json:"isReadOnly"    bson:"isReadOnly"`
	IsManager  bool                   `json:"isManager"     bson:"isManager"`
	Module     NotificationModule     `json:"module"        bson:"module"`
	Status     string                 `json:"status"        bson:"status"`
	Severity   NotificationSeverity   `json:"severity"      bson:"severity"`
	Category   NotificationCategory   `json:"category"      bson:"category"`
	Type       string                 `json:"type"          bson:"type"`
	Data       map[string]interface{} `json:"data"          bson:"data"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
