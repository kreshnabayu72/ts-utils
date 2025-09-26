package new_model

import (
	"time"
)

// AssignmentApproval represents the assignment_approvals table in the database.
type AssignmentApproval struct {
	ID         int            `json:"id"         column:"assignment_approvals.id"`
	RequestID  int            `json:"requestId"  column:"assignment_approvals.request_id"`
	ApproverID int            `json:"approverId" column:"assignment_approvals.approver_id"`
	Level      int            `json:"level"      column:"assignment_approvals.level"`
	Status     ApprovalStatus `json:"status"     column:"assignment_approvals.status"`
	IsApproved *bool          `json:"isApproved" column:"assignment_approvals.is_approved"`
	Notes      *string        `json:"notes"      column:"assignment_approvals.notes"`
	Type       ApprovalType   `json:"type"       column:"assignment_approvals.type"`
	ApprovedAt *time.Time     `json:"approvedAt" column:"assignment_approvals.approved_at"`
	CreatedAt  time.Time      `json:"createdAt"  column:"assignment_approvals.created_at"`
	UpdatedAt  time.Time      `json:"updatedAt"  column:"assignment_approvals.updated_at"`
}
