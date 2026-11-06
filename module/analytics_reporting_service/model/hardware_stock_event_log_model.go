// Package model defines the shared analytics-reporting-service wire models for the
// admin-console Hardware Stock module.
package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HardwareStockEventLog records one append-only Add Stock or Delete Stock session
// of one SKU; units are snapshots taken when the session committed.
type HardwareStockEventLog struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`

	SkuID      string                      `json:"skuId" bson:"skuId"`
	EntityType EntityType                  `json:"entityType" bson:"entityType"`
	EntityID   string                      `json:"entityId" bson:"entityId"`
	Action     HardwareStockEventLogAction `json:"action" bson:"action"`

	SuccessCount int `json:"successCount" bson:"successCount"`
	// WaitingApprovalCount is nil for Delete Stock sessions.
	WaitingApprovalCount *int                        `json:"waitingApprovalCount" bson:"waitingApprovalCount"`
	Units                []HardwareStockEventLogUnit `json:"units" bson:"units"`

	ModifiedByID        *string `json:"modifiedById" bson:"modifiedById"`
	ModifiedByFirstName *string `json:"modifiedByFirstName" bson:"modifiedByFirstName"`
	ModifiedByLastName  *string `json:"modifiedByLastName" bson:"modifiedByLastName"`
	ModifiedBySnapshot  string  `json:"modifiedBySnapshot" bson:"modifiedBySnapshot"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

// HardwareStockEventLogUnit is one unit of a session, in stored order.
type HardwareStockEventLogUnit struct {
	Section      HardwareStockEventLogSection `json:"section" bson:"section"`
	DeviceName   string                       `json:"deviceName" bson:"deviceName"`
	SerialNumber string                       `json:"serialNumber" bson:"serialNumber"`
}

type HardwareStockEventLogAction string

const (
	HardwareStockEventLogActionAddStock    HardwareStockEventLogAction = "Add Stock"
	HardwareStockEventLogActionDeleteStock HardwareStockEventLogAction = "Delete Stock"
)

type HardwareStockEventLogSection string

const (
	HardwareStockEventLogSectionAdded           HardwareStockEventLogSection = "added"
	HardwareStockEventLogSectionWaitingApproval HardwareStockEventLogSection = "waiting_approval"
	HardwareStockEventLogSectionDeleted         HardwareStockEventLogSection = "deleted"
)
