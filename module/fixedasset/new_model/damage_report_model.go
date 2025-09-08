package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type DamageReport struct {
	ID             int                `json:"_id"            column:"damage_reports.id"              sql:"BIGINT PRIMARY KEY"`
	Asset          int                `json:"asset"          column:"damage_reports.asset"           sql:"BIGINT"`
	ReportedById   int                `json:"reportedById"   column:"damage_reports.reported_by_id"  sql:"BIGINT"`
	DamageImage    string             `json:"damageImage"    column:"damage_reports.damage_image"    sql:"VARCHAR NOT NULL CHECK (char_length(damage_image) > 0)"`
	Status         DamageReportStatus `json:"status"         column:"damage_reports.status"          sql:"damage_report_status_enum NOT NULL"`
	PreviousStatus AssetStatus        `json:"previousStatus" column:"damage_reports.previous_status" sql:"asset_status_enum NOT NULL"`
	ReportSource   AssetReportSource  `json:"reportSource"   column:"damage_reports.report_source"   sql:"asset_report_source_enum NOT NULL"`
	Notes          string             `json:"notes"          column:"damage_reports.notes"           sql:"VARCHAR"`
	RepairedAt     *time.Time         `json:"repairedAt"     column:"damage_reports.repaired_at"     sql:"TIMESTAMPTZ DEFAULT NULL"`
	CreatedAt      time.Time          `json:"createdAt"      column:"damage_reports.created_at"      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt      time.Time          `json:"updatedAt"      column:"damage_reports.updated_at"      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (d *DamageReport) TableName() string {
	return "damage_reports"
}

func (d *DamageReport) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(d)

	builder.AddForeignKey(
		"fk_damage_reports_asset_id",
		"asset_id",
		"assets",
		"id",
		"CASCADE",
	)

	builder.AddForeignKey(
		"fk_damage_reports_reported_by_id",
		"reported_by_id",
		"users",
		"id",
		"CASCADE",
	)
}
