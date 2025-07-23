package new_model

import "time"

type DisposalTransaction struct {
	ID              int                       `json:"_id" column:"id"`
	TransactionID   string                    `json:"transactionId" column:"transaction_id"`
	ManagerID       int                       `json:"managerId" column:"manager_id"`
	OriginGroupID   int                       `json:"originGroupId" column:"origin_group_id"`
	DisposalGroupID int                       `json:"disposalGroupId" column:"disposal_group_id"`
	Status          DisposalTransactionStatus `json:"status" column:"status"`
	IsVerified      bool                      `json:"isVerified" column:"is_verified"`

	Notes              string                `json:"notes" column:"notes"`
	CompletionNotes    string                `json:"completionNotes" column:"completion_notes"`
	LastCondition      DisposalLastCondition `json:"lastCondition" column:"last_condition"`
	Method             DisposalMethod        `json:"method" column:"method"`
	Organization       string                `json:"organization" column:"organization"`
	DocumentationPhoto string                `json:"documentationPhoto" column:"documentation_photo"`

	CurrentApprovalType  DisposalGroupType `json:"currentApprovalType" column:"current_approval_type"`
	CurrentApprovalLevel int               `json:"currentApprovalLevel" column:"current_approval_level"`
	CurrentApprovalRound int               `json:"currentApprovalRound" column:"current_approval_round"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}
