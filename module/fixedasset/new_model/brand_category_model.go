package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type BrandCategory struct {
	ID         int        `json:"_id"        column:"brand_categories.id"          sql:"BIGINT PRIMARY KEY"`
	BrandId    int        `json:"brandId"    column:"brand_categories.brand_id"    sql:"BIGINT"`
	CategoryId int        `json:"categoryId" column:"brand_categories.category_id" sql:"BIGINT"`
	CreatedAt  time.Time  `json:"createdAt"  column:"brand_categories.created_at"  sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt  time.Time  `json:"updatedAt"  column:"brand_categories.updated_at"  sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	IsDeleted  bool       `json:"isDeleted"  column:"brand_categories.is_deleted"  sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeletedAt  *time.Time `json:"deletedAt"  column:"brand_categories.deleted_at"  sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
}

func (m *BrandCategory) TableName() string {
	return "brand_categories"
}

func (m *BrandCategory) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_brand_categories_category",
		"category_id",
		"categories",
		"id",
		"CASCADE",
	)
	builder.AddForeignKey(
		"fk_brand_categories_brand",
		"brand_id",
		"brands",
		"id",
		"CASCADE",
	)

}
