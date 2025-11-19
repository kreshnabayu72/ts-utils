package model

import (
	"time"
)

type ServiceCenterBrand struct {
	ID              int       `json:"_id"             column:"service_center_brands.id"`
	ServiceCenterId int       `json:"serviceCenterId" column:"service_center_brands.service_center_id"`
	BrandCategoryId int       `json:"brandCategoryId" column:"service_center_brands.brand_category_id"`
	CreatedAt       time.Time `json:"createdAt"       column:"service_center_brands.created_at"`
	UpdatedAt       time.Time `json:"updatedAt"       column:"service_center_brands.updated_at"`
}
