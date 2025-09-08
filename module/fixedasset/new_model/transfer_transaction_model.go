package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// TransferTransaction represents the transfer_transactions table in the database.
type TransferTransaction struct {
	ID                   int                       `json:"_id"                  column:"transfer_transactions.id"                     sql:"BIGINT PRIMARY KEY"`
	TransactionID        string                    `json:"transactionId"        column:"transfer_transactions.transaction_id"         sql:"VARCHAR UNIQUE"`
	ManagerID            int                       `json:"managerId"            column:"transfer_transactions.manager_id"             sql:"BIGINT"`
	OriginGroupID        int                       `json:"originGroupId"        column:"transfer_transactions.origin_group_id"        sql:"BIGINT"`
	DestinationGroupID   int                       `json:"destinationGroupId"   column:"transfer_transactions.destination_group_id"   sql:"BIGINT"`
	Status               TransferTransactionStatus `json:"status"               column:"transfer_transactions.status"                 sql:"transfer_transaction_status_enum"`
	IsVerified           bool                      `json:"isVerified"           column:"transfer_transactions.is_verified"            sql:"BOOLEAN"`
	CurrentApprovalType  TransferGroupType         `json:"currentApprovalType"  column:"transfer_transactions.current_approval_type"  sql:"transfer_group_type_enum NOT NULL"`
	CurrentApprovalLevel int                       `json:"currentApprovalLevel" column:"transfer_transactions.current_approval_level" sql:"INTEGER DEFAULT 1"`
	CurrentApprovalRound int                       `json:"currentApprovalRound" column:"transfer_transactions.current_approval_round" sql:"INTEGER DEFAULT 1"`
	CreatedAt            time.Time                 `json:"createdAt"            column:"transfer_transactions.created_at"             sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt            time.Time                 `json:"updatedAt"            column:"transfer_transactions.updated_at"             sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *TransferTransaction) TableName() string {
	return "transfer_transactions"
}

// Up defines the migration logic to create the table and its constraints.
func (m *TransferTransaction) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_transfer_transaction_manager",
		"manager_id",
		"users",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_transfer_transaction_origin_group",
		"origin_group_id",
		"groups",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_transfer_transaction_destination_group",
		"destination_group_id",
		"groups",
		"id",
		"",
	)

	// Add the unique index for transaction_id
	builder.AddIndex("transfer_transactions_transaction_id", "transaction_id", true, "")
}
