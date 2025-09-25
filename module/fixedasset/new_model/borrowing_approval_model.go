package new_model

import (
	"time"
)

type BorrowingApproval struct {
	ID               int            `json:"id"               column:"borrowing_approvals.id"`
	RequestID        int            `json:"requestId"        column:"borrowing_approvals.request_id"`
	ApproverID       int            `json:"approverId"       column:"borrowing_approvals.approver_id"`
	Level            int            `json:"level"            column:"borrowing_approvals.level"`
	ApprovalRound    int            `json:"approval_round"   column:"borrowing_approvals.approval_round"`
	Status           ApprovalStatus `json:"status"           column:"borrowing_approvals.status"`
	IsApproved       *bool          `json:"isApproved"       column:"borrowing_approvals.is_approved"`
	Notes            *string        `json:"notes"            column:"borrowing_approvals.notes"`
	Type             ApprovalType   `json:"type"             column:"borrowing_approvals.type"`
	ApprovedDuration *int           `json:"approvedDuration" column:"borrowing_approvals.approved_duration"`
	ApprovedAt       *time.Time     `json:"approvedAt"       column:"borrowing_approvals.approved_at"`
	CreatedAt        time.Time      `json:"createdAt"        column:"borrowing_approvals.created_at"`
	UpdatedAt        time.Time      `json:"updatedAt"        column:"borrowing_approvals.updated_at"`
}
