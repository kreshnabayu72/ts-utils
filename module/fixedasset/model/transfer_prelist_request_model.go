package model

import (
	"time"
)

// TransferPrelistRequest represents the transfer_prelist_requests table in the database.
type TransferPrelistRequest struct {
	ID        int       `json:"id"        column:"transfer_prelist_requests.id"`
	AssetID   int       `json:"assetId"   column:"transfer_prelist_requests.asset_id"`
	PrelistID int       `json:"prelistId" column:"transfer_prelist_requests.prelist_id"`
	CreatedAt time.Time `json:"createdAt" column:"transfer_prelist_requests.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"transfer_prelist_requests.updated_at"`
}
