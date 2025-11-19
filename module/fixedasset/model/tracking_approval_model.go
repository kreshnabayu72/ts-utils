package model

import (
	"time"
)

type TrackingApproval struct {
	ID            int            `json:"id"            column:"tracking_approvals.id"`
	TrackingID    int            `json:"trackingId"    column:"tracking_approvals.tracking_id"`
	ApproverID    int            `json:"approverId"    column:"tracking_approvals.approver_id"`
	ApprovalRound int            `json:"approvalRound" column:"tracking_approvals.approval_round"`
	Level         int            `json:"level"         column:"tracking_approvals.level"`
	Status        ApprovalStatus `json:"status"        column:"tracking_approvals.status"`
	IsApproved    *bool          `json:"isApproved"    column:"tracking_approvals.is_approved"`
	Notes         *string        `json:"notes"         column:"tracking_approvals.notes"`
	Type          ApprovalType   `json:"type"          column:"tracking_approvals.type"`
	ApprovedAt    *time.Time     `json:"approvedAt"    column:"tracking_approvals.approved_at"`
	CreatedAt     time.Time      `json:"createdAt"     column:"tracking_approvals.created_at"`
	UpdatedAt     time.Time      `json:"updatedAt"     column:"tracking_approvals.updated_at"`
}
