package new_model

import "time"

type ReplaceTagPrelist struct {
	ID int `json:"_id" column:"id"`

	AssetId int `json:"assetId"    column:"asset_id"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}
