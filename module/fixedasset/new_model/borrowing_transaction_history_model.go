package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type BorrowingTransactionHistory struct {
	ID            int                        `json:"_id"           column:"borrowing_transaction_histories.id"             sql:"BIGINT PRIMARY KEY"`
	TransactionID string                     `json:"transactionId" column:"borrowing_transaction_histories.transaction_id" sql:"VARCHAR"`
	Borrower      string                     `json:"borrower"      column:"borrowing_transaction_histories.borrower"       sql:"VARCHAR"`
	Manager       string                     `json:"manager"       column:"borrowing_transaction_histories.manager"        sql:"VARCHAR"`
	GroupID       int                        `json:"groupId"       column:"borrowing_transaction_histories.group_id"       sql:"BIGINT"`
	Group         string                     `json:"group"         column:"borrowing_transaction_histories.group"          sql:"VARCHAR"`
	Status        BorrowingTransactionStatus `json:"status"        column:"borrowing_transaction_histories.status"         sql:"borrowing_transaction_status_enum"`
	TotalAssets   int                        `json:"totalAssets"   column:"borrowing_transaction_histories.total_assets"   sql:"BIGINT"`
	CreatedAt     time.Time                  `json:"createdAt"     column:"borrowing_transaction_histories.created_at"     sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time                  `json:"updatedAt"     column:"borrowing_transaction_histories.updated_at"     sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (b *BorrowingTransactionHistory) TableName() string {
	return "borrowing_transaction_histories"
}

func (b *BorrowingTransactionHistory) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(b)

	// Add the unique index for transaction_id
	builder.AddIndex("borrowing_transactions_transaction_id", "transaction_id", true, "")
}
