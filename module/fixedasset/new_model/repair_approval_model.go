package new_model

import (
	"time"
)

type RepairApproval struct {
	ID              int                   `json:"id"              column:"repair_approvals.id"`
	RepairID        int                   `json:"repairId"        column:"repair_approvals.repair_id"`
	ApproverID      int                   `json:"approverId"      column:"repair_approvals.approver_id"`
	ApprovalRound   int                   `json:"approvalRound"   column:"repair_approvals.approval_round"`
	Level           int                   `json:"level"           column:"repair_approvals.level"`
	Status          ApprovalStatus        `json:"status"          column:"repair_approvals.status"`
	IsApproved      *bool                 `json:"isApproved"      column:"repair_approvals.is_approved"`
	Notes           *string               `json:"notes"           column:"repair_approvals.notes"`
	Type            ApprovalType          `json:"type"            column:"repair_approvals.type"`
	HandledBy       RepairTicketingMethod `json:"handledBy"       column:"repair_approvals.handled_by"`
	ServiceCenterId *int                  `json:"serviceCenterId" column:"repair_approvals.service_center_id"`
	Address         string                `json:"address"         column:"repair_approvals.address"`
	Coordinate      string                `json:"coordinate"      column:"repair_approvals.coordinate"`
	ApprovedAt      *time.Time            `json:"approvedAt"      column:"repair_approvals.approved_at"`
	CreatedAt       time.Time             `json:"createdAt"       column:"repair_approvals.created_at"`
	UpdatedAt       time.Time             `json:"updatedAt"       column:"repair_approvals.updated_at"`
}
