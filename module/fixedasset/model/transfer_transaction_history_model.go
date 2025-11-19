package model

import (
	"time"
)

// TransferTransactionHistory represents the transfer_transactions table in the database.
type TransferTransactionHistory struct {
	ID                 int                       `json:"_id"                column:"transfer_transaction_histories.id"`
	TransactionID      string                    `json:"transactionId"      column:"transfer_transaction_histories.transaction_id"`
	Manager            *string                   `json:"manager"            column:"transfer_transaction_histories.manager"`
	OriginGroupID      int                       `json:"originGroupId"      column:"transfer_transaction_histories.origin_group_id"`
	DestinationGroupID *int                      `json:"destinationGroupId" column:"transfer_transaction_histories.destination_group_id"`
	OriginGroup        string                    `json:"originGroup"        column:"transfer_transaction_histories.origin_group"`
	DestinationGroup   *string                   `json:"destinationGroup"   column:"transfer_transaction_histories.destination_group"`
	Status             TransferTransactionStatus `json:"status"             column:"transfer_transaction_histories.status"`
	TotalAssets        int                       `json:"totalAssets"        column:"transfer_transaction_histories.total_assets"`
	CreatedAt          time.Time                 `json:"createdAt"          column:"transfer_transaction_histories.created_at"`
	UpdatedAt          time.Time                 `json:"updatedAt"          column:"transfer_transaction_histories.updated_at"`
}
