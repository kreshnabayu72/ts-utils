package new_model

import (
	"time"
)

// DisposalReport represents the disposal_reports table in the database.
type DisposalReport struct {
	ID            int                   `json:"id"            column:"disposal_reports.id"`
	AssetID       int                   `json:"assetId"       column:"disposal_reports.asset_id"`
	ReportedByID  int                   `json:"reportedById"  column:"disposal_reports.reported_by_id"`
	LastCondition DisposalLastCondition `json:"lastCondition" column:"disposal_reports.last_condition"`
	Note          string                `json:"note"          column:"disposal_reports.note"`
	IsDeleted     bool                  `json:"isDeleted"     column:"disposal_reports.is_deleted"`
	CreatedAt     time.Time             `json:"createdAt"     column:"disposal_reports.created_at"`
	UpdatedAt     time.Time             `json:"updatedAt"     column:"disposal_reports.updated_at"`
}
