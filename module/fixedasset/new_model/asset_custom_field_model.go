package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AssetCustomField struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	AssetId       int `json:"assetId" column:"asset_id" sql:"BIGINT"`
	CustomFieldId int `json:"customFieldId" column:"custom_field_id" sql:"BIGINT"`
	Value         any `json:"value" column:"value" sql:"VARCHAR DEFAULT NULL"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
}

func (a *AssetCustomField) TableName() string {
	return "asset_custom_fields"
}

func (a *AssetCustomField) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(a)

	builder.AddForeignKey(
		"fk_asset_custom_fields_custom_field_id",
		"custom_field_id",
		"custom_fields",
		"id",
		"CASCADE",
	)

	builder.AddForeignKey(
		"fk_asset_custom_fields_asset_id",
		"asset_id",
		"assets",
		"id",
		"SET NULL",
	)

	builder.AddIndex("idx_acf_custom_field_active", "custom_field_id", false, "")
	builder.AddIndex("idx_acf_asset_active", "asset_id", false, "")
	builder.AddIndex("idx_acf_custom_field_asset_active", "custom_field_id, asset_id", false, "")
}
