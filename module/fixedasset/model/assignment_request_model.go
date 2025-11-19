package model

import (
	"time"
)

// AssignmentRequest represents the assignment_requests table in the database.
type AssignmentRequest struct {
	ID             int                     `json:"id"             column:"assignment_requests.id"`
	Asset          int                     `json:"asset"          column:"assignment_requests.asset"`
	Transaction    int                     `json:"transaction"    column:"assignment_requests.transaction"`
	ParentID       *int                    `json:"parentId"       column:"assignment_requests.parent_id"`
	RepairID       *int                    `json:"repairId"       column:"assignment_requests.repair_id"`
	TrackingID     *int                    `json:"trackingId"     column:"assignment_requests.tracking_id"`
	Status         AssignmentRequestStatus `json:"status"         column:"assignment_requests.status"`
	PreviousStatus AssignmentRequestStatus `json:"previousStatus" column:"assignment_requests.previous_status"`
	IsVerified     bool                    `json:"isVerified"     column:"assignment_requests.is_verified"`
	Notes          string                  `json:"notes"          column:"assignment_requests.notes"`
	SerialNumber   *string                 `json:"serialNumber"   column:"assignment_requests.serial_number"`
	CreatedAt      time.Time               `json:"createdAt"      column:"assignment_requests.created_at"`
	UpdatedAt      time.Time               `json:"updatedAt"      column:"assignment_requests.updated_at"`
}
