package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type ServiceCenterBrand struct {
	ID              int       `json:"_id"             column:"service_center_brands.id"                sql:"BIGINT PRIMARY KEY"`
	ServiceCenterId int       `json:"serviceCenterId" column:"service_center_brands.service_center_id" sql:"BIGINT NOT NULL"`
	BrandCategoryId int       `json:"brandCategoryId" column:"service_center_brands.brand_category_id" sql:"BIGINT NOT NULL"`
	CreatedAt       time.Time `json:"createdAt"       column:"service_center_brands.created_at"        sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt       time.Time `json:"updatedAt"       column:"service_center_brands.updated_at"        sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *ServiceCenterBrand) TableName() string {
	return "service_center_brands"
}

func (m *ServiceCenterBrand) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_service_center_brands_service_center",
		"service_center_id",
		"service_centers",
		"id",
		"CASCADE",
	)
	builder.AddForeignKey(
		"fk_service_center_brands_brand_category",
		"brand_category_id",
		"brand_categories",
		"id",
		"CASCADE",
	)
}
