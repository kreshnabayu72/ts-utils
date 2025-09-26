package new_model

import (
	"time"
)

type AssetCustomField struct {
	ID            int       `json:"_id"           column:"asset_custom_fields.id"`
	AssetId       int       `json:"assetId"       column:"asset_custom_fields.asset_id"`
	CustomFieldId int       `json:"customFieldId" column:"asset_custom_fields.custom_field_id"`
	Value         *string   `json:"value"         column:"asset_custom_fields.value"`
	CreatedAt     time.Time `json:"createdAt"     column:"asset_custom_fields.created_at"`
	UpdatedAt     time.Time `json:"updatedAt"     column:"asset_custom_fields.updated_at"`
}
