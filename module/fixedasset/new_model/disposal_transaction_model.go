package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// DisposalTransaction represents the disposal_transactions table in the database.
type DisposalTransaction struct {
	ID                   int                       `json:"_id"                  column:"disposal_transactions.id"                     sql:"BIGINT PRIMARY KEY"`
	TransactionID        string                    `json:"transactionId"        column:"disposal_transactions.transaction_id"         sql:"VARCHAR UNIQUE"`
	ManagerID            int                       `json:"managerId"            column:"disposal_transactions.manager_id"             sql:"BIGINT"`
	OriginGroupID        int                       `json:"originGroupId"        column:"disposal_transactions.origin_group_id"        sql:"BIGINT"`
	DisposalGroupID      int                       `json:"disposalGroupId"      column:"disposal_transactions.disposal_group_id"      sql:"BIGINT"`
	Status               DisposalTransactionStatus `json:"status"               column:"disposal_transactions.status"                 sql:"disposal_transaction_status_enum"`
	IsVerified           bool                      `json:"isVerified"           column:"disposal_transactions.is_verified"            sql:"BOOLEAN"`
	Notes                string                    `json:"notes"                column:"disposal_transactions.notes"                  sql:"VARCHAR"`
	CompletionNotes      string                    `json:"completionNotes"      column:"disposal_transactions.completion_notes"       sql:"VARCHAR"`
	LastCondition        DisposalLastCondition     `json:"lastCondition"        column:"disposal_transactions.last_condition"         sql:"disposal_last_condition_enum"`
	Method               DisposalMethod            `json:"method"               column:"disposal_transactions.method"                 sql:"disposal_method_enum"`
	Organization         string                    `json:"organization"         column:"disposal_transactions.organization"           sql:"VARCHAR"`
	DocumentationPhoto   string                    `json:"documentationPhoto"   column:"disposal_transactions.documentation_photo"    sql:"VARCHAR"`
	CurrentApprovalType  DisposalGroupType         `json:"currentApprovalType"  column:"disposal_transactions.current_approval_type"  sql:"disposal_group_type_enum NOT NULL"`
	CurrentApprovalLevel int                       `json:"currentApprovalLevel" column:"disposal_transactions.current_approval_level" sql:"INTEGER DEFAULT 1"`
	CurrentApprovalRound int                       `json:"currentApprovalRound" column:"disposal_transactions.current_approval_round" sql:"INTEGER DEFAULT 1"`
	CreatedAt            time.Time                 `json:"createdAt"            column:"disposal_transactions.created_at"             sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt            time.Time                 `json:"updatedAt"            column:"disposal_transactions.updated_at"             sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *DisposalTransaction) TableName() string {
	return "disposal_transactions"
}

// Up defines the migration logic to create the table and its constraints.
func (m *DisposalTransaction) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_disposal_transaction_manager",
		"manager_id",
		"users",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_disposal_transaction_origin_group",
		"origin_group_id",
		"groups",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_disposal_transaction_disposal_group",
		"disposal_group_id",
		"groups",
		"id",
		"",
	)
}
