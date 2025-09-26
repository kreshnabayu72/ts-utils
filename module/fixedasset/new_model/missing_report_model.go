package new_model

import (
	"time"
)

type MissingReport struct {
	ID             int                 `json:"_id"            column:"missing_reports.id"`
	Asset          int                 `json:"asset"          column:"missing_reports.asset"`
	ReportedById   int                 `json:"reportedById"   column:"missing_reports.reported_by_id"`
	Status         MissingReportStatus `json:"status"         column:"missing_reports.status"`
	PreviousStatus AssetStatus         `json:"previousStatus" column:"missing_reports.previous_status"`
	ReportSource   AssetReportSource   `json:"reportSource"   column:"missing_reports.report_source"`
	Notes          string              `json:"notes"          column:"missing_reports.notes"`
	FoundAt        *time.Time          `json:"foundAt"        column:"missing_reports.found_at"`
	CreatedAt      time.Time           `json:"createdAt"      column:"missing_reports.created_at"`
	UpdatedAt      time.Time           `json:"updatedAt"      column:"missing_reports.updated_at"`
}
