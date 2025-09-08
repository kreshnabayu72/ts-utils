package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// DisposalRequest represents the disposal_requests table in the database.
type DisposalRequest struct {
	ID                      int                   `json:"id"                      column:"disposal_requests.id"                        sql:"BIGINT PRIMARY KEY"`
	Asset                   int                   `json:"asset"                   column:"disposal_requests.asset"                     sql:"BIGINT"`
	Transaction             int                   `json:"transaction"             column:"disposal_requests.transaction"               sql:"BIGINT"`
	ReportID                *int                  `json:"reportId"                column:"disposal_requests.report_id"                 sql:"BIGINT"`
	PreviousStatus          AssetStatus           `json:"previousStatus"          column:"disposal_requests.previous_status"           sql:"asset_status_enum"`
	Status                  DisposalRequestStatus `json:"status"                  column:"disposal_requests.status"                    sql:"disposal_request_status_enum"`
	SalvageableParts        []string              `json:"salvageableParts"        column:"disposal_requests.salvageable_parts"         sql:"TEXT[]"`
	Notes                   string                `json:"notes"                   column:"disposal_requests.notes"                     sql:"VARCHAR"`
	CompletionApprovalRound *int                  `json:"completionApprovalRound" column:"disposal_requests.completion_approval_round" sql:"INTEGER NULL"`
	IsVerified              bool                  `json:"isVerified"              column:"disposal_requests.is_verified"               sql:"BOOLEAN"`
	CreatedAt               time.Time             `json:"createdAt"               column:"disposal_requests.created_at"                sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt               time.Time             `json:"updatedAt"               column:"disposal_requests.updated_at"                sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *DisposalRequest) TableName() string {
	return "disposal_requests"
}

// Up defines the migration logic to create the table and its constraints.
func (m *DisposalRequest) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_disposal_request_asset",
		"asset",
		"assets",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_disposal_request_transaction",
		"transaction",
		"disposal_transactions",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_disposal_request_report",
		"report_id",
		"disposal_reports",
		"id",
		"",
	)
}
