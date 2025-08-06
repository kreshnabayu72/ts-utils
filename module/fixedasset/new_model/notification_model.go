package new_model

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
	NotificationModuleDamage             = NotificationModule("Damage")
	NotificationModuleMissing            = NotificationModule("Missing")
	NotificationModuleTracking           = NotificationModule("Tracking")
	NotificationModuleTAG                = NotificationModule("TAG")
	NotificationModuleGeneralAsset       = NotificationModule("General Asset")
)

type NotificationApp string

const (
	NotificationAppFixedAsset NotificationModule = "Fixed Asset"
)

type Notification struct {
	ID         primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	UserKey    int                    `json:"user" bson:"user"`
	Title      string                 `json:"title"         bson:"title"`
	Detail     string                 `json:"detail"        bson:"detail"`
	Plan       string                 `json:"plan"        bson:"plan"`
	IsRead     bool                   `json:"isRead"        bson:"isRead"`
	IsReadOnly bool                   `json:"isReadOnly"    bson:"isReadOnly"`
	IsManager  bool                   `json:"isManager"     bson:"isManager"`
	App        NotificationApp        `json:"app"        bson:"app"`
	Module     NotificationModule     `json:"module"        bson:"module"`
	Status     string                 `json:"status"        bson:"status"`
	Severity   NotificationSeverity   `json:"severity"      bson:"severity"`
	Category   NotificationCategory   `json:"category"      bson:"category"`
	Type       string                 `json:"type"          bson:"type"`
	FCMToken   string                 `json:"fcmToken"    bson:"fcmToken"`
	Data       map[string]interface{} `json:"data"          bson:"data"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
