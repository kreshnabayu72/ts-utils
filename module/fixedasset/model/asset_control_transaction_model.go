package model

import (
	"time"
)

type AssetControlTransaction struct {
	ID int `json:"id"           column:"asset_control_transactions.id"`

	TransferTransactionId int  `json:"transferTransactionId"      column:"asset_control_transactions.transfer_transaction_id"`
	OriginGroupID         int  `json:"originGroupId"      column:"asset_control_transactions.origin_group_id"`
	DestinationGroupID    *int `json:"destinationGroupId"      column:"asset_control_transactions.destination_group_id"`

	CreatedAt time.Time `json:"createdAt"    column:"asset_control_transactions.created_at"`
	UpdatedAt time.Time `json:"updatedAt"    column:"asset_control_transactions.updated_at"`
}
