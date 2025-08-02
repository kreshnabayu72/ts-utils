package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type ImportBrand struct {
	ID  int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	Key int `json:"key" column:"key" sql:"INTEGER NOT NULL"`

	Name         string `json:"name" column:"name" sql:"VARCHAR NOT NULL"`
	Category     string `json:"category" column:"category" sql:"VARCHAR NOT NULL"`
	CategoryCode string `json:"categoryCode" column:"category_code" sql:"VARCHAR NOT NULL"`
	Models       string `json:"models" column:"models" sql:"VARCHAR NOT NULL"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
}

func (b *ImportBrand) TableName() string {
	return "import_brands"
}

func (b *ImportBrand) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(b)

	builder.AddIndex("import_brands_name", "LOWER(name)", true, "")
}
