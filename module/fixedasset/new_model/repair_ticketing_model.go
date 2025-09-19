package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RepairTicketing struct {
	ID                         int                     `json:"_id"                        column:"repair_ticketings.id"                            sql:"BIGINT PRIMARY KEY"`
	Asset                      int                     `json:"asset"                    column:"repair_ticketings.asset"                      sql:"BIGINT NOT NULL"`
	ReportedByID               int                     `json:"reportedById"               column:"repair_ticketings.reported_by_id"                sql:"BIGINT NOT NULL"`
	DamageImage                string                  `json:"damageImage"                column:"repair_ticketings.damage_image"                  sql:"VARCHAR"`
	PreviousStatus             AssetStatus             `json:"previousStatus"             column:"repair_ticketings.previous_status"               sql:"asset_status_enum NOT NULL"`
	TransactionID              string                  `json:"transactionId"              column:"repair_ticketings.transaction_id"                sql:"VARCHAR NOT NULL"`
	Status                     RepairTicketingStatus   `json:"status"                     column:"repair_ticketings.status"                        sql:"repair_ticketing_status_enum NOT NULL"`
	PreviousMethod             []RepairTicketingMethod `json:"previousMethod"             column:"repair_ticketings.previous_method"               sql:"repair_ticketing_method_enum[] NOT NULL"`
	ReportNotes                string                  `json:"reportNotes"                column:"repair_ticketings.report_notes"                  sql:"VARCHAR"`
	RepairNotes                string                  `json:"repairNotes"                column:"repair_ticketings.repair_notes"                  sql:"VARCHAR"`
	IsRepairSuccess            *bool                   `json:"isRepairSuccess"            column:"repair_ticketings.is_repair_success"             sql:"BOOLEAN NULL DEFAULT NULL"`
	DownTimeSeconds            float64                 `json:"downTimeSeconds"            column:"repair_ticketings.down_time_seconds"             sql:"FLOAT"`
	ReportSource               AssetReportSource       `json:"reportSource"               column:"repair_ticketings.report_source"                 sql:"asset_report_source_enum NOT NULL"`
	ApprovalRound              int                     `json:"approvalRound"              column:"repair_ticketings.approval_round"                sql:"INTEGER NOT NULL DEFAULT 0"`
	Details                    string                  `json:"details"                    column:"repair_ticketings.details"                       sql:"VARCHAR"`
	HandledBy                  RepairTicketingMethod   `json:"handledBy"                  column:"repair_ticketings.handled_by"                    sql:"repair_ticketing_method_enum"`
	AdditionalCostTotalPrice   float32                 `json:"additionalCostTotalPrice"   column:"repair_ticketings.additional_cost_total_price"   sql:"FLOAT DEFAULT 0"`
	AdditionalCostCurrency     string                  `json:"additionalCostCurrency"     column:"repair_ticketings.additional_cost_currency"      sql:"VARCHAR"`
	AdditionalCostPaymentProof string                  `json:"additionalCostPaymentProof" column:"repair_ticketings.additional_cost_payment_proof" sql:"VARCHAR"`
	StaffID                    *int                    `json:"staffId"                    column:"repair_ticketings.staff_id"                      sql:"BIGINT NULL DEFAULT NULL"`
	Address                    string                  `json:"address"         column:"repair_ticketings.address"           sql:"VARCHAR"`
	Coordinate                 string                  `json:"coordinate"      column:"repair_ticketings.coordinate"        sql:"VARCHAR"`
	ServiceCenterId            *int                    `json:"serviceCenterId"            column:"repair_ticketings.service_center_id"             sql:"BIGINT NULL DEFAULT NULL"`
	RepairStartAt              *time.Time              `json:"repairStartAt"              column:"repair_ticketings.repair_start_at"               sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt                  time.Time               `json:"createdAt"                  column:"repair_ticketings.created_at"                    sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt                  time.Time               `json:"updatedAt"                  column:"repair_ticketings.updated_at"                    sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *RepairTicketing) TableName() string {
	return "repair_ticketings"
}

func (m *RepairTicketing) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_repair_ticketing_asset",
		"asset",
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
