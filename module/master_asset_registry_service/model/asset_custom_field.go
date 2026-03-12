package model

import "time"

type AssetCustomField struct {
	ID            int       `json:"_id"           column:"asset_custom_fields.id"`
	AssetID       int       `json:"assetId"       column:"asset_custom_fields.asset_id"`
	CustomFieldID int       `json:"customFieldId" column:"asset_custom_fields.custom_field_id"`
	Value         any       `json:"value"         column:"asset_custom_fields.value"`
	CreatedAt     time.Time `json:"createdAt"     column:"asset_custom_fields.created_at"`
	UpdatedAt     time.Time `json:"updatedAt"     column:"asset_custom_fields.updated_at"`
}
func (a *AssetCustomField) GetColumns() []string {
	return []string{
		"id",
		"asset_id",
		"custom_field_id",
		"value",
		"created_at",
		"updated_at",
	}
}
