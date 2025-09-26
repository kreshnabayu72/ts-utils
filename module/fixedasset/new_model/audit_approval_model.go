package new_model

import (
	"time"
)

// AuditApproval represents the audit_approvals table in the database.
type AuditApproval struct {
	ID            int            `json:"id"            column:"audit_approvals.id"`
	TaskId        int            `json:"taskId"        column:"audit_approvals.task_id"`
	ApproverID    int            `json:"approverId"    column:"audit_approvals.approver_id"`
	ApprovalRound int            `json:"approvalRound" column:"audit_approvals.approval_round"`
	Level         int            `json:"level"         column:"audit_approvals.level"`
	Status        ApprovalStatus `json:"status"        column:"audit_approvals.status"`
	IsApproved    *bool          `json:"isApproved"    column:"audit_approvals.is_approved"`
	Notes         *string        `json:"notes"         column:"audit_approvals.notes"`
	Type          ApprovalType   `json:"type"          column:"audit_approvals.type"`
	Address       string         `json:"address"       column:"audit_approvals.address"`
	Coordinate    string         `json:"coordinate"    column:"audit_approvals.coordinate"`
	ApprovedAt    *time.Time     `json:"approvedAt"    column:"audit_approvals.approved_at"`
	CreatedAt     time.Time      `json:"createdAt"     column:"audit_approvals.created_at"`
	UpdatedAt     time.Time      `json:"updatedAt"     column:"audit_approvals.updated_at"`
}
