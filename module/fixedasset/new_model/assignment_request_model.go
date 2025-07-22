package new_model

import "time"

type AssignmentRequest struct {
	ID            int                     `json:"id" column:"id"`
	AssetID       int                     `json:"assetId" column:"asset_id"`
	TransactionID int                     `json:"transactionId" column:"transaction_id"`
	ParentID      *int                    `json:"parentId" column:"parent_id"`
	Status        AssignmentRequestStatus `json:"status" column:"status"`
	IsVerified    bool                    `json:"isVerified" column:"is_verified"`
	CreatedAt     time.Time               `json:"createdAt" column:"created_at"`
	UpdatedAt     time.Time               `json:"updatedAt" column:"updated_at"`
}
