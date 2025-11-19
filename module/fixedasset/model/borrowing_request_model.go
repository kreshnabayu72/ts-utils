package model

import (
	"time"
)

type BorrowingRequest struct {
	ID               int                    `json:"id"               column:"borrowing_requests.id"`
	Asset            int                    `json:"asset"            column:"borrowing_requests.asset"`
	Transaction      int                    `json:"transaction"      column:"borrowing_requests.transaction"`
	ApprovalRound    int                    `json:"approval_round"   column:"borrowing_requests.approval_round"`
	ParentID         *int                   `json:"parentId"         column:"borrowing_requests.parent_id"`
	RepairID         *int                   `json:"repairId"         column:"borrowing_requests.repair_id"`
	TrackingID       *int                   `json:"trackingId"       column:"borrowing_requests.tracking_id"`
	Status           BorrowingRequestStatus `json:"status"           column:"borrowing_requests.status"`
	PreviousStatus   BorrowingRequestStatus `json:"previousStatus"   column:"borrowing_requests.previous_status"`
	IsVerified       bool                   `json:"isVerified"       column:"borrowing_requests.is_verified"`
	Notes            string                 `json:"notes"            column:"borrowing_requests.notes"`
	Duration         int                    `json:"duration"         column:"borrowing_requests.duration"`
	PreviousDuration int                    `json:"previousDuration" column:"borrowing_requests.previous_duration"`
	ExtendDuration   *int                   `json:"extendDuration"   column:"borrowing_requests.extend_duration"`
	ActualReturnDate *time.Time             `json:"actualReturnDate" column:"borrowing_requests.actual_return_date"`
	DueReturnDate    time.Time              `json:"dueReturnDate"    column:"borrowing_requests.due_return_date"`
	IsOverdue        bool                   `json:"isOverdue"        column:"borrowing_requests.is_overdue"`
	SerialNumber     *string                `json:"serialNumber"     column:"borrowing_requests.serial_number"`
	CreatedAt        time.Time              `json:"createdAt"        column:"borrowing_requests.created_at"`
	UpdatedAt        time.Time              `json:"updatedAt"        column:"borrowing_requests.updated_at"`
}
