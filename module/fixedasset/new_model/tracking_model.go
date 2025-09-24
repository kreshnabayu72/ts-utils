package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Tracking struct {
	ID              int                `json:"_id"             column:"trackings.id"                sql:"BIGINT PRIMARY KEY"`
	Asset           int                `json:"asset"         column:"trackings.asset"          sql:"BIGINT NOT NULL"`
	ReportedByID    int                `json:"reportedById"    column:"trackings.reported_by_id"    sql:"BIGINT NOT NULL"`
	DamageImage     string             `json:"damageImage"     column:"trackings.damage_image"      sql:"VARCHAR"`
	PreviousStatus  AssetStatus        `json:"previousStatus"  column:"trackings.previous_status"   sql:"asset_status_enum NOT NULL"`
	TransactionID   string             `json:"transactionId"   column:"trackings.transaction_id"    sql:"VARCHAR NOT NULL"`
	Status          TrackingStatus     `json:"status"          column:"trackings.status"            sql:"tracking_status_enum NOT NULL"`
	Condition       *TrackingCondition `json:"condition"       column:"trackings.condition"         sql:"VARCHAR NULL"`
	Notes           string             `json:"notes"           column:"trackings.notes"             sql:"VARCHAR"`
	Address         string             `json:"address"         column:"trackings.address"           sql:"VARCHAR NULL DEFAULT NULL"`
	Coordinate      string             `json:"coordinate"      column:"trackings.coordinate"        sql:"VARCHAR NULL DEFAULT NULL"`
	DownTimeSeconds float64            `json:"downTimeSeconds" column:"trackings.down_time_seconds" sql:"FLOAT"`
	ReportSource    AssetReportSource  `json:"reportSource"    column:"trackings.report_source"     sql:"asset_report_source_enum NOT NULL"`
	ApprovalRound   int                `json:"approvalRound"   column:"trackings.approval_round"    sql:"INTEGER NOT NULL DEFAULT 0"`

	FoundAt    time.Time `json:"foundAt" column:"trackings.found_at"`
	FoundBy    int       `json:"foundBy" column:"trackings.found_by_id"`
	FoundGroup int       `json:"found_group" column:"trackings.found_group_id"`

	CreatedAt time.Time `json:"createdAt"       column:"trackings.created_at"        sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt"       column:"trackings.updated_at"        sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *Tracking) TableName() string {
	return "trackings"
}

func (m *Tracking) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_tracking_found_group_id",
		"found_group_id",
		"groups",
		"id",
		"CASCADE",
	)
	builder.AddForeignKey(
		"fk_tracking_found_by_id",
		"found_by_id",
		"users",
		"id",
		"CASCADE",
	)
	builder.AddForeignKey(
		"fk_tracking_asset",
		"asset",
		"assets",
		"id",
		"CASCADE",
	)
	builder.AddForeignKey(
		"fk_tracking_repored_by_id",
		"reported_by_id",
		"users",
		"id",
		"SET NULL",
	)
}
