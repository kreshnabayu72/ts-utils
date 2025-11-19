package model

import "time"

type ReportBrandFilter struct {
	ID        int       `json:"_id"                  column:"reports_filter_brand.id"`
	ReportId  int       `json:"reportId"            column:"reports_filter_brand.report_id"`
	BrandId   int       `json:"brandId"            column:"reports_filter_brand.brand_id"`
	CreatedAt time.Time `json:"createdAt"            column:"reports_filter_brand.created_at"`
	UpdatedAt time.Time `json:"updatedAt"            column:"reports_filter_brand.updated_at"`
}
