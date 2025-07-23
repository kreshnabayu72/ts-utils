package new_model

import "time"

type TransferRequest struct {
	ID             int                           `json:"id" column:"id"`
	AssetID        int                           `json:"assetId" column:"asset_id"`
	TransactionID  int                           `json:"transactionId" column:"transaction_id"`
	ParentID       *int                          `json:"parentId" column:"parent_id"`
	Status         TransferRequestStatus         `json:"status" column:"status"`
	AssetTagStatus TransferRequestAssetTagStatus `json:"assetTagStatus" column:"asset_tag_status"`
	AssetStatus    TransferRequestAssetStatus    `json:"assetStatus" column:"asset_status"`
	IsVerified     bool                          `json:"isVerified" column:"is_verified"`
	CreatedAt      time.Time                     `json:"createdAt" column:"created_at"`
	UpdatedAt      time.Time                     `json:"updatedAt" column:"updated_at"`
}
