package new_model

import (
	"time"
)

type AssetControlRequest struct {
	ID int `json:"id"           column:"asset_control_requests.id"`

	TransferRequestId     int  `json:"transferRequestId"      column:"asset_control_requests.transfer_request_id"`
	TransferTransactionId int  `json:"transferTransactionId"      column:"asset_control_requests.transfer_transaction_id"`
	AssetId               int  `json:"assetId"      column:"asset_control_requests.asset_id"`
	OriginGroupId         int  `json:"originGroupId"      column:"asset_control_requests.origin_group_id"`
	DestinationGroupId    *int `json:"destinationGroupId"      column:"asset_control_requests.destination_group_id"`

	CreatedAt time.Time `json:"createdAt"    column:"asset_control_requests.created_at"`
	UpdatedAt time.Time `json:"updatedAt"    column:"asset_control_requests.updated_at"`
}
