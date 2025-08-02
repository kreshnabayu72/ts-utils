package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type ReplaceTagPrelist struct {
	ID      int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	AssetId int `json:"assetId" column:"asset_id" sql:"BIGINT"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
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
