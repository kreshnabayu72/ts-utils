package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// AssignmentTransactionHistory represents the assignment_transaction_histories table in the database.
type AssignmentTransactionHistory struct {
	ID            int                         `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	TransactionID string                      `json:"transactionId" column:"transaction_id" sql:"VARCHAR UNIQUE"`
	GroupID       int                         `json:"groupId" column:"group_id" sql:"BIGINT"`
	AssignedTo    string                      `json:"assignedTo" column:"assigned_to" sql:"VARCHAR"`
	Manager       string                      `json:"manager" column:"manager" sql:"VARCHAR"`
	TotalAssets   int                         `json:"totalAssets" column:"total_assets" sql:"BIGINT"`
	Status        AssignmentTransactionStatus `json:"status" column:"status" sql:"assignment_transaction_status_enum"`
	CreatedAt     time.Time                   `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time                   `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *AssignmentTransactionHistory) TableName() string {
	return "assignment_transaction_histories"
}

// Up defines the migration logic to create the table and its constraints.
func (m *AssignmentTransactionHistory) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_assignment_transaction_assigned_to",
		"assigned_to_id",
		"users",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_assignment_transaction_manager",
		"manager_id",
		"users",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_assignment_transaction_group",
		"group_id",
		"groups",
		"id",
		"",
	)

	// Add the unique index for transaction_id
	builder.AddIndex("assignment_transactions_transaction_id", "transaction_id", true, "")
}
