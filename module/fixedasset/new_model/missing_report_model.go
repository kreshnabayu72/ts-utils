package new_model

import "time"

type MissingReport struct {
	ID int `json:"_id"         column:"id"`

	AssetId      int `json:"assetId"               column:"asset_id"`
	ReportedById int `json:"reportedById"               column:"reported_by_id"`

	Status         MissingReportStatus `json:"status"               column:"status"`
	PreviousStatus AssetStatus         `json:"previousStatus"               column:"previous_status"`
	ReportSource   AssetReportSource   `json:"reportSource"               column:"report_source"`
	Notes          string              `json:"notes"               column:"notes"`

	CreatedAt time.Time `json:"createdAt"            column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"            column:"updated_at"`
}
