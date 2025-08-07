package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type ImportAssetName struct {
	ID  int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	Key int `json:"key" column:"key" sql:"INTEGER NOT NULL"`

	Name         string `json:"name" column:"name" sql:"VARCHAR NOT NULL"`
	Code         string `json:"code" column:"code" sql:"VARCHAR NOT NULL"`
	Category     string `json:"category" column:"category" sql:"VARCHAR NOT NULL"`
	CategoryCode string `json:"categoryCode" column:"category_code" sql:"VARCHAR NOT NULL"`
	Measurement  string `json:"measurement" column:"measurement" sql:"VARCHAR NOT NULL"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (a *ImportAssetName) TableName() string {
	return "import_asset_names"
}

func (a *ImportAssetName) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(a)

	builder.AddIndex("import_asset_names_name", "LOWER(name)", true, "")
	builder.AddIndex("import_asset_names_code", "LOWER(code)", true, "")
}
