package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Model struct {
	ID              int        `json:"_id"             column:"models.id"                sql:"BIGINT PRIMARY KEY"`
	Name            string     `json:"name"            column:"models.name"              sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(name) > 0)"`
	BrandCategoryId int        `json:"brandCategoryId" column:"models.brand_category_id" sql:"BIGINT"`
	IsDeleted       bool       `json:"isDeleted"       column:"models.is_deleted"        sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeletedAt       *time.Time `json:"deletedAt"       column:"models.deleted_at"        sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt       time.Time  `json:"createdAt"       column:"models.created_at"        sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt       time.Time  `json:"updatedAt"       column:"models.updated_at"        sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *Model) TableName() string {
	return "models"
}

func (m *Model) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("idx_models_name", "name", false, "")

	builder.AddForeignKey(
		"fk_models_brand_category",
		"brand_category_id",
		"brand_categories",
		"id",
		"CASCADE",
	)
}
