package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type BorrowingTransactionHistory struct {
	ID            int    `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	TransactionID string `json:"transactionId" column:"transaction_id" sql:"VARCHAR"`

	Borrower string `json:"borrower" column:"borrower" sql:"VARCHAR"`
	Manager  string `json:"manager" column:"manager" sql:"VARCHAR"`
	GroupID  int    `json:"groupId" column:"group_id" sql:"BIGINT"`
	Group    int    `json:"group" column:"group" sql:"VARCHAR"`

	Status      BorrowingTransactionStatus `json:"status" column:"status" sql:"borrowing_transaction_status_enum"`
	TotalAssets int                        `json:"totalAssets" column:"total_assets" sql:"BIGINT"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (b *BorrowingTransactionHistory) TableName() string {
	return "borrowing_transaction_histories"
}

func (b *BorrowingTransactionHistory) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(b)
}
