package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type ReplaceTagPrelist struct {
	ID        int                   `json:"_id"       column:"replace_tag_prelists.id"         sql:"BIGINT PRIMARY KEY"`
	AssetId   int                   `json:"assetId"   column:"replace_tag_prelists.asset_id"   sql:"BIGINT"`
	Type      ReplaceTagPrelistType `json:"type"   column:"replace_tag_prelists.type"   sql:"VARCHAR"`
	CreatedAt time.Time             `json:"createdAt" column:"replace_tag_prelists.created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time             `json:"updatedAt" column:"replace_tag_prelists.updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (r *ReplaceTagPrelist) TableName() string {
	return "replace_tag_prelists"
}

func (r *ReplaceTagPrelist) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(r)

	builder.AddForeignKey(
		"fk_replace_tag_prelist_asset_id",
		"asset_id",
		"assets",
		"id",
		"CASCADE",
	)
}
