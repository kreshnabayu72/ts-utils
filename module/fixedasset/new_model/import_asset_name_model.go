package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type ImportAssetName struct {
	ID           int       `json:"_id"          column:"import_asset_names.id"            sql:"BIGINT PRIMARY KEY"`
	Key          int       `json:"key"          column:"import_asset_names.key"           sql:"INTEGER NOT NULL"`
	Name         string    `json:"name"         column:"import_asset_names.name"          sql:"VARCHAR NOT NULL"`
	Code         string    `json:"code"         column:"import_asset_names.code"          sql:"VARCHAR NOT NULL"`
	Category     string    `json:"category"     column:"import_asset_names.category"      sql:"VARCHAR NOT NULL"`
	CategoryCode string    `json:"categoryCode" column:"import_asset_names.category_code" sql:"VARCHAR NOT NULL"`
	Measurement  string    `json:"measurement"  column:"import_asset_names.measurement"   sql:"VARCHAR NOT NULL"`
	CreatedAt    time.Time `json:"createdAt"    column:"import_asset_names.created_at"    sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt    time.Time `json:"updatedAt"    column:"import_asset_names.updated_at"    sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (a *ImportAssetName) TableName() string {
	return "import_asset_names"
}

func (a *ImportAssetName) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(a)

	builder.AddIndex("import_asset_names_name", "LOWER(name)", true, "")
	builder.AddIndex("import_asset_names_code", "LOWER(code)", true, "")
}
