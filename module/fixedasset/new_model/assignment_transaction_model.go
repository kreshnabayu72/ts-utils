package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// AssignmentTransaction represents the assignment_transactions table in the database.
type AssignmentTransaction struct {
	ID                           int                         `json:"_id"                          column:"assignment_transactions.id"                              sql:"BIGINT PRIMARY KEY"`
	TransactionID                string                      `json:"transactionId"                column:"assignment_transactions.transaction_id"                  sql:"VARCHAR UNIQUE"`
	AssignedToID                 int                         `json:"assignedToId"                 column:"assignment_transactions.assigned_to_id"                  sql:"BIGINT"`
	ManagerID                    int                         `json:"managerId"                    column:"assignment_transactions.manager_id"                      sql:"BIGINT"`
	GroupID                      int                         `json:"groupId"                      column:"assignment_transactions.group_id"                        sql:"BIGINT"`
	Status                       AssignmentTransactionStatus `json:"status"                       column:"assignment_transactions.status"                          sql:"assignment_transaction_status_enum"`
	IsVerified                   bool                        `json:"isVerified"                   column:"assignment_transactions.is_verified"                     sql:"BOOLEAN"`
	ConfirmationEmail            string                      `json:"confirmationEmail"            column:"assignment_transactions.confirmation_email"              sql:"VARCHAR"`
	ConfirmationEmailConfirmed   bool                        `json:"confirmationEmailConfirmed"   column:"assignment_transactions.confirmation_email_confirmed"    sql:"BOOLEAN"`
	ConfirmationEmailSentAt      *time.Time                  `json:"confirmationEmailSentAt"      column:"assignment_transactions.confirmation_email_sent_at"      sql:"TIMESTAMPTZ"`
	ConfirmationEmailExpiredAt   *time.Time                  `json:"confirmationEmailExpiredAt"   column:"assignment_transactions.confirmation_email_expired_at"   sql:"TIMESTAMPTZ"`
	ConfirmationEmailToken       string                      `json:"confirmationEmailToken"       column:"assignment_transactions.confirmation_email_token"        sql:"VARCHAR"`
	ConfirmationEmailConfirmedAt *time.Time                  `json:"confirmationEmailConfirmedAt" column:"assignment_transactions.confirmation_email_confirmed_at" sql:"TIMESTAMPTZ"`
	CurrentApprovalLevel         int                         `json:"currentApprovalLevel"         column:"assignment_transactions.current_approval_level"          sql:"INTEGER DEFAULT 1"`
	AssignDate                   time.Time                   `json:"assignDate" column:"assignment_transactions.assign_date" sql:"TIMESTAMPTZ"`
	CreatedAt                    time.Time                   `json:"createdAt"                    column:"assignment_transactions.created_at"                      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt                    time.Time                   `json:"updatedAt"                    column:"assignment_transactions.updated_at"                      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *AssignmentTransaction) TableName() string {
	return "assignment_transactions"
}

// Up defines the migration logic to create the table and its constraints.
func (m *AssignmentTransaction) Up(builder *migration.MigrationBuilder) {
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
