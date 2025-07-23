package new_model

import "time"

type AssignmentApproval struct {
	ID         int            `json:"id" column:"id"`
	RequestID  int            `json:"requestId" column:"request_id"`
	ApproverID int            `json:"approverId" column:"approver_id"`
	Level      int            `json:"level" column:"level"`
	Status     ApprovalStatus `json:"status" column:"status"`
	IsApproved *bool          `json:"isApproved" column:"is_approved"`
	Notes      *string        `json:"notes" column:"notes"`
	Type       ApprovalType   `json:"type" column:"type"`
	ApprovedAt *time.Time     `json:"approvedAt" column:"approved_at"`
	CreatedAt  time.Time      `json:"createdAt" column:"created_at"`
	UpdatedAt  time.Time      `json:"updatedAt" column:"updated_at"`
}
