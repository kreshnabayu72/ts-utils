package replica

import "time"

type AssetCustomField struct {
	ID            int       `json:"_id"           column:"asset_custom_fields.id"`
	AssetID       int       `json:"assetId"       column:"asset_custom_fields.asset_id"`
	CustomFieldID int       `json:"customFieldId" column:"asset_custom_fields.custom_field_id"`
	Value         any       `json:"value"         column:"asset_custom_fields.value"`
	CreatedAt     time.Time `json:"createdAt"     column:"asset_custom_fields.created_at"`
	UpdatedAt     time.Time `json:"updatedAt"     column:"asset_custom_fields.updated_at"`
}

func (a *AssetCustomField) GetTableName() string {
	return "asset_custom_fields"
}

func (a *AssetCustomField) GetTableNameWithSchema() string {
	return "public.asset_custom_fields"
}

func (a *AssetCustomField) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":              "BIGINT",
		"asset_id":        "BIGINT",
		"custom_field_id": "BIGINT",
		"value":           "JSONB",
		"created_at":      "TIMESTAMPTZ",
		"updated_at":      "TIMESTAMPTZ",
	}
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
