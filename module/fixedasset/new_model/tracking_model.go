package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Tracking struct {
	ID              int               `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	AssetID         int               `json:"assetId" column:"asset_id" sql:"BIGINT NOT NULL"`
	ReportedByID    int               `json:"reportedById" column:"reported_by_id" sql:"BIGINT NOT NULL"`
	DamageImage     string            `json:"damageImage" column:"damage_image" sql:"VARCHAR"`
	PreviousStatus  AssetStatus       `json:"previousStatus" column:"previous_status" sql:"asset_status_enum NOT NULL"`
	TransactionID   string            `json:"transactionId" column:"transaction_id" sql:"VARCHAR NOT NULL"`
	Status          TrackingStatus    `json:"status" column:"status" sql:"tracking_status_enum NOT NULL"`
	Condition       TrackingCondition `json:"condition" column:"condition" sql:"tracking_condition_enum NOT NULL"`
	Notes           string            `json:"notes" column:"notes" sql:"VARCHAR"`
	Address         string            `json:"address" column:"address" sql:"VARCHAR"`
	Coordinate      string            `json:"coordinate" column:"coordinate" sql:"VARCHAR"`
	DownTimeSeconds float64           `json:"downTimeSeconds" column:"down_time_seconds" sql:"FLOAT"`
	ReportSource    AssetReportSource `json:"reportSource" column:"report_source" sql:"asset_report_source_enum NOT NULL"`
	ApprovalRound   int               `json:"approvalRound" column:"approval_round" sql:"INTEGER NOT NULL DEFAULT 0"`
	CreatedAt       time.Time         `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt       time.Time         `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *Tracking) TableName() string {
	return "trackings"
}

func (m *Tracking) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_tracking_asset_id",
		"asset_id",
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
