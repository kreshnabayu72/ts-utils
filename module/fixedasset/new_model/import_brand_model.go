package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type ImportBrand struct {
	ID           int       `json:"_id"          column:"import_brands.id"            sql:"BIGINT PRIMARY KEY"`
	Key          int       `json:"key"          column:"import_brands.key"           sql:"INTEGER NOT NULL"`
	Name         string    `json:"name"         column:"import_brands.name"          sql:"VARCHAR NOT NULL"`
	Category     string    `json:"category"     column:"import_brands.category"      sql:"VARCHAR NOT NULL"`
	CategoryCode string    `json:"categoryCode" column:"import_brands.category_code" sql:"VARCHAR NOT NULL"`
	Models       string    `json:"models"       column:"import_brands.models"        sql:"VARCHAR NOT NULL"`
	CreatedAt    time.Time `json:"createdAt"    column:"import_brands.created_at"    sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt    time.Time `json:"updatedAt"    column:"import_brands.updated_at"    sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (b *ImportBrand) TableName() string {
	return "import_brands"
}

func (b *ImportBrand) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(b)

	builder.AddIndex("import_brands_name", "LOWER(name)", true, "")
}
