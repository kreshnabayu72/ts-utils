package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// AssignmentTransaction represents the assignment_transactions table in the database.
type AssignmentTransaction struct {
	ID                           int                         `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	TransactionID                string                      `json:"transactionId" column:"transaction_id" sql:"VARCHAR UNIQUE"`
	AssignedToID                 int                         `json:"assignedToId" column:"assigned_to_id" sql:"BIGINT"`
	ManagerID                    int                         `json:"managerId" column:"manager_id" sql:"BIGINT"`
	GroupID                      int                         `json:"groupId" column:"group_id" sql:"BIGINT"`
	Status                       AssignmentTransactionStatus `json:"status" column:"status" sql:"assignment_transaction_status_enum"`
	IsVerified                   bool                        `json:"isVerified" column:"is_verified" sql:"BOOLEAN"`
	ConfirmationEmail            string                      `json:"confirmationEmail" column:"confirmation_email" sql:"VARCHAR"`
	ConfirmationEmailConfirmed   bool                        `json:"confirmationEmailConfirmed" column:"confirmation_email_confirmed" sql:"BOOLEAN"`
	ConfirmationEmailSentAt      time.Time                   `json:"confirmationEmailSentAt" column:"confirmation_email_sent_at" sql:"TIMESTAMP"`
	ConfirmationEmailExpiredAt   time.Time                   `json:"confirmationEmailExpiredAt" column:"confirmation_email_expired_at" sql:"TIMESTAMP"`
	ConfirmationEmailToken       string                      `json:"confirmationEmailToken" column:"confirmation_email_token" sql:"VARCHAR"`
	ConfirmationEmailConfirmedAt time.Time                   `json:"confirmationEmailConfirmedAt" column:"confirmation_email_confirmed_at" sql:"TIMESTAMP"`
	CurrentApprovalLevel         int                         `json:"currentApprovalLevel" column:"current_approval_level" sql:"INTEGER DEFAULT 1"`
	CreatedAt                    time.Time                   `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
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
