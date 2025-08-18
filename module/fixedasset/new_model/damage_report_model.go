package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type DamageReport struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	// Foreign key to assets
	Asset int `json:"asset" column:"asset" sql:"BIGINT"`
	// Asset's unique id from alias code
	AssetID string `json:"assetId" column:"asset_id" sql:"VARCHAR"`

	ReportedById int `json:"reportedById" column:"reported_by_id" sql:"BIGINT"`

	DamageImage    string             `json:"damageImage" column:"damage_image" sql:"VARCHAR NOT NULL CHECK (char_length(damage_image) > 0)"`
	Status         DamageReportStatus `json:"status" column:"status" sql:"damage_report_status_enum NOT NULL"`
	PreviousStatus AssetStatus        `json:"previousStatus" column:"previous_status" sql:"asset_status_enum NOT NULL"`
	ReportSource   AssetReportSource  `json:"reportSource" column:"report_source" sql:"asset_report_source_enum NOT NULL"`
	Notes          string             `json:"notes" column:"notes" sql:"VARCHAR"`

	RepairedAt *time.Time `json:"repairedAt" column:"repaired_at" sql:"TIMESTAMPTZ DEFAULT NULL"`
	CreatedAt  time.Time  `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt  time.Time  `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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
