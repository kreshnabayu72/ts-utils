package new_model

import "time"

type ReportCategoryFilter struct {
	ID         int       `json:"_id"                  column:"reports_filter_category.id"`
	ReportId   int       `json:"reportId"            column:"reports_filter_category.report_id"`
	CategoryId int       `json:"categoryId"            column:"reports_filter_category.category_id"`
	CreatedAt  time.Time `json:"createdAt"            column:"reports_filter_category.created_at"`
	UpdatedAt  time.Time `json:"updatedAt"            column:"reports_filter_category.updated_at"`
}
