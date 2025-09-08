package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type MissingReport struct {
	ID             int                 `json:"_id"            column:"missing_reports.id"              sql:"BIGINT PRIMARY KEY"`
	Asset          int                 `json:"asset"          column:"missing_reports.asset"           sql:"BIGINT"`
	ReportedById   int                 `json:"reportedById"   column:"missing_reports.reported_by_id"  sql:"BIGINT"`
	Status         MissingReportStatus `json:"status"         column:"missing_reports.status"          sql:"missing_report_status_enum NOT NULL"`
	PreviousStatus AssetStatus         `json:"previousStatus" column:"missing_reports.previous_status" sql:"asset_status_enum NOT NULL"`
	ReportSource   AssetReportSource   `json:"reportSource"   column:"missing_reports.report_source"   sql:"asset_report_source_enum NOT NULL"`
	Notes          string              `json:"notes"          column:"missing_reports.notes"           sql:"VARCHAR"`
	FoundAt        *time.Time          `json:"foundAt"        column:"missing_reports.found_at"        sql:"TIMESTAMPTZ DEFAULT NULL"`
	CreatedAt      time.Time           `json:"createdAt"      column:"missing_reports.created_at"      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt      time.Time           `json:"updatedAt"      column:"missing_reports.updated_at"      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *MissingReport) TableName() string {
	return "missing_reports"
}

func (m *MissingReport) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_missing_reports_asset_id",
		"asset_id",
		"assets",
		"id",
		"CASCADE",
	)

	builder.AddForeignKey(
		"fk_missing_reports_reported_by_id",
		"reported_by_id",
		"users",
		"id",
		"CASCADE",
	)
}
