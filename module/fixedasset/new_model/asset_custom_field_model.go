package new_model

import "time"

type AssetCustomField struct {
	ID int `json:"_id" column:"id"`

	AssetId       int `json:"assetId"    column:"asset_id"`
	CustomFieldId int `json:"customFieldId" column:"custom_field_id"`
	Value         any `json:"value" column:"value"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}
