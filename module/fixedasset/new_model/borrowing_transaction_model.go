package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type BorrowingTransaction struct {
	ID                           int                        `json:"_id"                          column:"borrowing_transactions.id"                              sql:"BIGINT PRIMARY KEY"`
	TransactionID                string                     `json:"transactionId"                column:"borrowing_transactions.transaction_id"                  sql:"VARCHAR"`
	BorrowerID                   int                        `json:"borrowerId"                   column:"borrowing_transactions.borrower_id"                     sql:"BIGINT"`
	ManagerID                    int                        `json:"managerId"                    column:"borrowing_transactions.manager_id"                      sql:"BIGINT"`
	GroupID                      int                        `json:"groupId"                      column:"borrowing_transactions.group_id"                        sql:"BIGINT"`
	Status                       BorrowingTransactionStatus `json:"status"                       column:"borrowing_transactions.status"                          sql:"borrowing_transaction_status_enum"`
	IsSelfService                bool                       `json:"isSelfService"                column:"borrowing_transactions.is_self_service"                 sql:"BOOLEAN"`
	IsVerified                   bool                       `json:"isVerified"                   column:"borrowing_transactions.is_verified"                     sql:"BOOLEAN"`
	ConfirmationEmail            string                     `json:"confirmationEmail"            column:"borrowing_transactions.confirmation_email"              sql:"VARCHAR"`
	ConfirmationEmailConfirmed   bool                       `json:"confirmationEmailConfirmed"   column:"borrowing_transactions.confirmation_email_confirmed"    sql:"BOOLEAN"`
	ConfirmationEmailSentAt      *time.Time                 `json:"confirmationEmailSentAt"      column:"borrowing_transactions.confirmation_email_sent_at"      sql:"TIMESTAMPTZ"`
	ConfirmationEmailExpiredAt   *time.Time                 `json:"confirmationEmailExpiredAt"   column:"borrowing_transactions.confirmation_email_expired_at"   sql:"TIMESTAMPTZ"`
	ConfirmationEmailToken       string                     `json:"confirmationEmailToken"       column:"borrowing_transactions.confirmation_email_token"        sql:"VARCHAR"`
	ConfirmationEmailConfirmedAt *time.Time                 `json:"confirmationEmailConfirmedAt" column:"borrowing_transactions.confirmation_email_confirmed_at" sql:"TIMESTAMPTZ"`
	CurrentApprovalLevel         int                        `json:"currentApprovalLevel"         column:"borrowing_transactions.current_approval_level"          sql:"INTEGER DEFAULT 1"`
	CreatedAt                    time.Time                  `json:"createdAt"                    column:"borrowing_transactions.created_at"                      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt                    time.Time                  `json:"updatedAt"                    column:"borrowing_transactions.updated_at"                      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (b *BorrowingTransaction) TableName() string {
	return "borrowing_transactions"
}

func (b *BorrowingTransaction) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(b)

	builder.AddForeignKey(
		"fk_borrowing_transaction_borrower",
		"borrower_id",
		"users",
		"id",
		"",
	)

	builder.AddForeignKey(
		"fk_borrowing_transaction_manager",
		"manager_id",
		"users",
		"id",
		"",
	)

	builder.AddForeignKey(
		"fk_borrowing_transaction_group",
		"group_id",
		"groups",
		"id",
		"",
	)

	builder.AddIndex(
		"borrowing_transactions_transaction_id",
		"transaction_id",
		true,
		"",
	)
}
