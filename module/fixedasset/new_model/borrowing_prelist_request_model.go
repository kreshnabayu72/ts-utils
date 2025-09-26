package new_model

import (
	"time"
)

type BorrowingPrelistRequest struct {
	ID        int       `json:"id"        column:"borrowing_prelist_requests.id"`
	AssetId   int       `json:"assetId"   column:"borrowing_prelist_requests.asset_id"`
	PrelistId int       `json:"prelistId" column:"borrowing_prelist_requests.prelist_id"`
	CreatedAt time.Time `json:"createdAt" column:"borrowing_prelist_requests.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"borrowing_prelist_requests.updated_at"`
}
