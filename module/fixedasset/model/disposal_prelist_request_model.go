package model

import (
	"time"
)

// DisposalPrelistRequest represents the disposal_prelist_requests table in the database.
type DisposalPrelistRequest struct {
	ID        int       `json:"id"        column:"disposal_prelist_requests.id"`
	AssetID   int       `json:"assetId"   column:"disposal_prelist_requests.asset_id"`
	PrelistID int       `json:"prelistId" column:"disposal_prelist_requests.prelist_id"`
	Note      string    `json:"note"      column:"disposal_prelist_requests.note"`
	CreatedAt time.Time `json:"createdAt" column:"disposal_prelist_requests.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"disposal_prelist_requests.updated_at"`
}
