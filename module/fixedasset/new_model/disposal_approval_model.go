package new_model

import "time"

type DisposalApproval struct {
	ID            int               `json:"id" column:"id"`
	TransactionID *int              `json:"transactionId" column:"transaction_id"`
	RequestID     *int              `json:"requestId" column:"request_id"`
	ApproverID    int               `json:"approverId" column:"approver_id"`
	ApprovalRound int               `json:"approvalRound" column:"approval_round"`
	Level         int               `json:"level" column:"level"`
	Status        ApprovalStatus    `json:"status" column:"status"`
	IsApproved    *bool             `json:"isApproved" column:"is_approved"`
	Notes         *string           `json:"notes" column:"notes"`
	Type          ApprovalType      `json:"type" column:"type"`
	GroupType     DisposalGroupType `json:"groupType" column:"group_type"`
	ApprovedAt    *time.Time        `json:"approvedAt" column:"approved_at"`
	CreatedAt     time.Time         `json:"createdAt" column:"created_at"`
	UpdatedAt     time.Time         `json:"updatedAt" column:"updated_at"`
}
