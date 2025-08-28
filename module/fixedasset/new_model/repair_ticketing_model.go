package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RepairTicketing struct {
	ID                         int                     `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	AssetID                    int                     `json:"assetId" column:"asset_id" sql:"BIGINT NOT NULL"`
	ReportedByID               int                     `json:"reportedById" column:"reported_by_id" sql:"BIGINT NOT NULL"`
	DamageImage                string                  `json:"damageImage" column:"damage_image" sql:"VARCHAR"`
	PreviousStatus             AssetStatus             `json:"previousStatus" column:"previous_status" sql:"asset_status_enum NOT NULL"`
	TransactionID              string                  `json:"transactionId" column:"transaction_id" sql:"VARCHAR NOT NULL"`
	Status                     RepairTicketingStatus   `json:"status" column:"status" sql:"repair_ticketing_status_enum NOT NULL"`
	PreviousMethod             []RepairTicketingMethod `json:"previousMethod" column:"previous_method" sql:"repair_ticketing_method_enum[] NOT NULL"`
	ReportNotes                string                  `json:"reportNotes" column:"report_notes" sql:"VARCHAR"`
	RepairNotes                string                  `json:"repairNotes" column:"repair_notes" sql:"VARCHAR"`
	IsRepairSuccess            bool                    `json:"isRepairSuccess" column:"is_repair_success" sql:"BOOLEAN"`
	DownTimeSeconds            float64                 `json:"downTimeSeconds" column:"down_time_seconds" sql:"FLOAT"`
	ReportSource               AssetReportSource       `json:"reportSource" column:"report_source" sql:"asset_report_source_enum NOT NULL"`
	ApprovalRound              int                     `json:"approvalRound" column:"approval_round" sql:"INTEGER NOT NULL DEFAULT 0"`
	Details                    string                  `json:"details" column:"details" sql:"VARCHAR"`
	HandledBy                  RepairTicketingMethod   `json:"handledBy" column:"handled_by" sql:"repair_ticketing_method_enum"`
	AdditionalCostTotalPrice   float32                 `json:"additionalCostTotalPrice" column:"additional_cost_total_price" sql:"FLOAT DEFAULT 0"`
	AdditionalCostCurrency     string                  `json:"additionalCostCurrency" column:"additional_cost_currency" sql:"VARCHAR"`
	AdditionalCostPaymentProof string                  `json:"additionalCostPaymentProof" column:"additional_cost_payment_proof" sql:"VARCHAR"`
	StaffID                    *int                    `json:"staffId" column:"staff_id" sql:"BIGINT NULL DEFAULT NULL"`
	IsDeleted                  bool                    `json:"isDeleted" column:"is_deleted" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	RepairStartAt              *time.Time              `json:"repairStartAt" column:"repair_start_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt                  time.Time               `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt                  time.Time               `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *RepairTicketing) TableName() string {
	return "repair_ticketings"
}

func (m *RepairTicketing) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_repair_ticketing_asset_id",
		"asset_id",
		"assets",
		"id",
		"CASCADE",
	)
	builder.AddForeignKey(
		"fk_repair_ticketing_reported_by_id",
		"reported_by_id",
		"users",
		"id",
		"SET NULL",
	)
	builder.AddForeignKey(
		"fk_repair_ticketing_staff_id",
		"staff_id",
		"users",
		"id",
		"SET NULL",
	)
}
