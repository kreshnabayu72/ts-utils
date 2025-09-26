package new_model

import (
	"time"
)

// TransferTransaction represents the transfer_transactions table in the database.
type TransferTransaction struct {
	ID                   int                       `json:"_id"                  column:"transfer_transactions.id"`
	TransactionID        string                    `json:"transactionId"        column:"transfer_transactions.transaction_id"`
	ManagerID            int                       `json:"managerId"            column:"transfer_transactions.manager_id"`
	OriginGroupID        int                       `json:"originGroupId"        column:"transfer_transactions.origin_group_id"`
	DestinationGroupID   int                       `json:"destinationGroupId"   column:"transfer_transactions.destination_group_id"`
	Status               TransferTransactionStatus `json:"status"               column:"transfer_transactions.status"`
	IsVerified           bool                      `json:"isVerified"           column:"transfer_transactions.is_verified"`
	CurrentApprovalType  TransferGroupType         `json:"currentApprovalType"  column:"transfer_transactions.current_approval_type"`
	CurrentApprovalLevel int                       `json:"currentApprovalLevel" column:"transfer_transactions.current_approval_level"`
	CurrentApprovalRound int                       `json:"currentApprovalRound" column:"transfer_transactions.current_approval_round"`
	CreatedAt            time.Time                 `json:"createdAt"            column:"transfer_transactions.created_at"`
	UpdatedAt            time.Time                 `json:"updatedAt"            column:"transfer_transactions.updated_at"`
}
