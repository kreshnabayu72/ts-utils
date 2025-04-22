package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationSeverity string

const (
	NotificationSeverityInfo    NotificationSeverity = "info"
	NotificationSeverityDanger  NotificationSeverity = "danger"
	NotificationSeverityWarning NotificationSeverity = "warning"
	NotificationSeveritySuccess NotificationSeverity = "success"
)

type NotificationCategory string

const (
	NotificationCategoryCompleted  NotificationCategory = "completed"
	NotificationCategoryRejected   NotificationCategory = "rejected"
	NotificationCategoryOnProgress NotificationCategory = "onProgress"
	NotificationCategoryWarning    NotificationCategory = "warning"
)

type NotificationModule string

const (
	NotificationModuleReception      NotificationModule = "Reception"
	NotificationModuleSupplierReturn NotificationModule = "Supplier Return"
	NotificationModuleTransfer       NotificationModule = "Transfer"
	NotificationModuleCheckout       NotificationModule = "Checkout"
	NotificationModuleInternalReturn NotificationModule = "Internal Return"
	NotificationModuleExpired        NotificationModule = "Expired"
	NotificationModuleDefect         NotificationModule = "Defect"
	NotificationModuleDisposal       NotificationModule = "Disposal"
	NotificationModuleMissing        NotificationModule = "Missing"
	NotificationModuleStockOpname    NotificationModule = "Stock Opname"
	NotificationModuleTracking       NotificationModule = "Tracking"
	NotificationModuleTAG            NotificationModule = "TAG"
	NotificationModuleGeneralItem    NotificationModule = "General Item"
)

type Notification struct {
	ID        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	User      primitive.ObjectID   `json:"user"          bson:"user"`
	Title     string               `json:"title"         bson:"title"`
	Detail    string               `json:"detail"        bson:"detail"`
	IsRead    bool                 `json:"isRead"        bson:"isRead"`
	Module    NotificationModule   `json:"module"        bson:"module"`
	Severity  NotificationSeverity `json:"severity"      bson:"severity"`
	Category  NotificationCategory `json:"category"      bson:"category"`
	Data      map[string]any       `json:"data"          bson:"data"`
	CreatedAt time.Time            `json:"createdAt"     bson:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt"     bson:"updatedAt"`
}
