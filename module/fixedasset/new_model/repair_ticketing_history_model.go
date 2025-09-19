package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RepairTicketingHistory struct {
	ID                 int    `json:"_id"                        column:"repair_ticketing_histories.id"                            sql:"BIGINT PRIMARY KEY"`
	Asset              int    `json:"asset"                    column:"repair_ticketing_histories.asset"                      sql:"BIGINT NOT NULL"`
	ReportedByFullName string `json:"reportedByFullName"               column:"repair_ticketing_histories.reported_by_full_name"                sql:"VARCHAR"`

	AssetNameWithSequence string `json:"assetNameWithSequence"                column:"repair_ticketing_histories.asset_name_with_sequence"                  sql:"VARCHAR"`
	AssetAliasCode        string `json:"assetAliasCode"                column:"repair_ticketing_histories.asset_alias_code"                  sql:"VARCHAR"`
	AssetBrand            string `json:"assetBrand"                column:"repair_ticketing_histories.asset_brand"                  sql:"VARCHAR"`
	AssetModel            string `json:"assetModel"                column:"repair_ticketing_histories.asset_model"                  sql:"VARCHAR"`
	AssetGroupName        string `json:"assetGroup"                column:"repair_ticketing_histories.asset_group_name"                  sql:"VARCHAR"`
	AssetGroupId          int    `json:"assetGroupId"                column:"repair_ticketing_histories.asset_group_id"                  sql:"BIGINT NOT NULL"`

	DamageImage                string                  `json:"damageImage"                column:"repair_ticketing_histories.damage_image"                  sql:"VARCHAR"`
	PreviousStatus             AssetStatus             `json:"previousStatus"             column:"repair_ticketing_histories.previous_status"               sql:"asset_status_enum NOT NULL"`
	TransactionID              string                  `json:"transactionId"              column:"repair_ticketing_histories.transaction_id"                sql:"VARCHAR NOT NULL"`
	Status                     RepairTicketingStatus   `json:"status"                     column:"repair_ticketing_histories.status"                        sql:"repair_ticketing_status_enum NOT NULL"`
	PreviousMethod             []RepairTicketingMethod `json:"previousMethod"             column:"repair_ticketing_histories.previous_method"               sql:"repair_ticketing_method_enum[] NOT NULL"`
	ReportNotes                string                  `json:"reportNotes"                column:"repair_ticketing_histories.report_notes"                  sql:"VARCHAR"`
	RepairNotes                string                  `json:"repairNotes"                column:"repair_ticketing_histories.repair_notes"                  sql:"VARCHAR"`
	IsRepairSuccess            *bool                   `json:"isRepairSuccess"            column:"repair_ticketing_histories.is_repair_success"             sql:"BOOLEAN"`
	DownTimeSeconds            float64                 `json:"downTimeSeconds"            column:"repair_ticketing_histories.down_time_seconds"             sql:"FLOAT"`
	ReportSource               AssetReportSource       `json:"reportSource"               column:"repair_ticketing_histories.report_source"                 sql:"asset_report_source_enum NOT NULL"`
	Details                    string                  `json:"details"                    column:"repair_ticketing_histories.details"                       sql:"VARCHAR"`
	HandledBy                  RepairTicketingMethod   `json:"handledBy"                  column:"repair_ticketing_histories.handled_by"                    sql:"repair_ticketing_method_enum"`
	ApprovalRound              int                     `json:"approvalRound"              column:"repair_ticketings.approval_round"                sql:"INTEGER NOT NULL DEFAULT 0"`
	AdditionalCostTotalPrice   float32                 `json:"additionalCostTotalPrice"   column:"repair_ticketing_histories.additional_cost_total_price"   sql:"FLOAT DEFAULT 0"`
	AdditionalCostCurrency     string                  `json:"additionalCostCurrency"     column:"repair_ticketing_histories.additional_cost_currency"      sql:"VARCHAR"`
	AdditionalCostPaymentProof string                  `json:"additionalCostPaymentProof" column:"repair_ticketing_histories.additional_cost_payment_proof" sql:"VARCHAR"`
	Address                    string                  `json:"address"         column:"repair_ticketing_histories.address"           sql:"VARCHAR NULL DEFAULT NULL"`
	Coordinate                 string                  `json:"coordinate"      column:"repair_ticketing_histories.coordinate"        sql:"VARCHAR NULL DEFAULT NULL"`
	StaffFullName              string                  `json:"staffFullName"                    column:"repair_ticketing_histories.staff_full_name"                      sql:"VARCHAR"`
	ServiceCenterName          *string                 `json:"serviceCenterName"            column:"repair_ticketing_histories.service_center_name"             sql:"VARCHAR NULL"`
	RepairStartAt              *time.Time              `json:"repairStartAt"              column:"repair_ticketing_histories.repair_start_at"               sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt                  time.Time               `json:"createdAt"                  column:"repair_ticketing_histories.created_at"                    sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt                  time.Time               `json:"updatedAt"                  column:"repair_ticketing_histories.updated_at"                    sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *RepairTicketingHistory) TableName() string {
	return "repair_ticketing_histories"
}

func (m *RepairTicketingHistory) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)
}
