package new_model

import "time"

type ReportAssetNameFilter struct {
	ID          int       `json:"_id"                  column:"reports_filter_asset_name.id"`
	ReportId    int       `json:"reportId"            column:"reports_filter_asset_name.report_id"`
	AssetNameId int       `json:"assetNameId"            column:"reports_filter_asset_name.asset_name_id"`
	CreatedAt   time.Time `json:"createdAt"            column:"reports_filter_asset_name.created_at"`
	UpdatedAt   time.Time `json:"updatedAt"            column:"reports_filter_asset_name.updated_at"`
}
