package new_model

import "time"

type AssignmentPrelistRequest struct {
	ID        int       `json:"id" column:"id"`
	AssetId   int       `json:"assetId" column:"asset_id"`
	PrelistId int       `json:"prelistId" column:"prelist_id"`
	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}
