package new_model

import "time"

type TransferTransaction struct {
	ID                 int                       `json:"_id" column:"id"`
	TransactionID      string                    `json:"transactionId" column:"transaction_id"`
	ManagerID          int                       `json:"managerId" column:"manager_id"`
	OriginGroupID      int                       `json:"originGroupId" column:"origin_group_id"`
	DestinationGroupID int                       `json:"destinationGroupId" column:"destination_group_id"`
	Status             TransferTransactionStatus `json:"status" column:"status"`
	IsVerified         bool                      `json:"isVerified" column:"is_verified"`

	CurrentApprovalType  TransferGroupType `json:"currentApprovalType" column:"current_approval_type"`
	CurrentApprovalLevel int               `json:"currentApprovalLevel" column:"current_approval_level"`
	CurrentApprovalRound int               `json:"currentApprovalRound" column:"current_approval_round"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}
