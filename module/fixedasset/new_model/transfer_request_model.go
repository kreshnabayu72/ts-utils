package new_model

import (
	"time"
)

// TransferRequest represents the transfer_requests table in the database.
type TransferRequest struct {
	ID             int                           `json:"id"             column:"transfer_requests.id"`
	Asset          int                           `json:"asset"          column:"transfer_requests.asset"`
	Transaction    int                           `json:"transaction"    column:"transfer_requests.transaction"`
	ParentID       *int                          `json:"parentId"       column:"transfer_requests.parent_id"`
	Status         TransferRequestStatus         `json:"status"         column:"transfer_requests.status"`
	AssetTagStatus TransferRequestAssetTagStatus `json:"assetTagStatus" column:"transfer_requests.asset_tag_status"`
	AssetStatus    *TransferRequestAssetStatus   `json:"assetStatus"    column:"transfer_requests.asset_status"`
	IsVerified     bool                          `json:"isVerified"     column:"transfer_requests.is_verified"`
	Notes          string                        `json:"notes"          column:"transfer_requests.notes"`
	CreatedAt      time.Time                     `json:"createdAt"      column:"transfer_requests.created_at"`
	UpdatedAt      time.Time                     `json:"updatedAt"      column:"transfer_requests.updated_at"`
}
