package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type MissingReport struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	// Foreign key to assets
	Asset int `json:"asset" column:"asset" sql:"BIGINT"`

	ReportedById int `json:"reportedById" column:"reported_by_id" sql:"BIGINT"`

	Status         MissingReportStatus `json:"status" column:"status" sql:"missing_report_status_enum NOT NULL"`
	PreviousStatus AssetStatus         `json:"previousStatus" column:"previous_status" sql:"asset_status_enum NOT NULL"`
	ReportSource   AssetReportSource   `json:"reportSource" column:"report_source" sql:"asset_report_source_enum NOT NULL"`
	Notes          string              `json:"notes" column:"notes" sql:"VARCHAR"`

	FoundAt   *time.Time `json:"foundAt" column:"found_at" sql:"TIMESTAMPTZ DEFAULT NULL"`
	CreatedAt time.Time  `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time  `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *MissingReport) TableName() string {
	return "missing_reports"
}

func (m *MissingReport) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_missing_reports_asset_id",
		"asset",
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
