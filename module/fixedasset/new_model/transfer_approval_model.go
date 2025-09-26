package new_model

import (
	"time"
)

// TransferApproval represents the transfer_approvals table in the database.
type TransferApproval struct {
	ID         int               `json:"id"         column:"transfer_approvals.id"`
	RequestID  int               `json:"requestId"  column:"transfer_approvals.request_id"`
	ApproverID int               `json:"approverId" column:"transfer_approvals.approver_id"`
	Level      int               `json:"level"      column:"transfer_approvals.level"`
	Status     ApprovalStatus    `json:"status"     column:"transfer_approvals.status"`
	IsApproved *bool             `json:"isApproved" column:"transfer_approvals.is_approved"`
	Notes      *string           `json:"notes"      column:"transfer_approvals.notes"`
	Type       ApprovalType      `json:"type"       column:"transfer_approvals.type"`
	GroupType  TransferGroupType `json:"groupType"  column:"transfer_approvals.group_type"`
	ApprovedAt *time.Time        `json:"approvedAt" column:"transfer_approvals.approved_at"`
	CreatedAt  time.Time         `json:"createdAt"  column:"transfer_approvals.created_at"`
	UpdatedAt  time.Time         `json:"updatedAt"  column:"transfer_approvals.updated_at"`
}
