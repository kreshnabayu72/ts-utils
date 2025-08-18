package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type BorrowingRequest struct {
	ID int `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`

	// Foreign key to assets
	Asset int `json:"asset" column:"asset" sql:"BIGINT"`
	// Asset's unique id from alias code
	AssetID string `json:"assetId" column:"asset_id" sql:"VARCHAR"`
	// Foreign key to related transaction
	Transaction int `json:"transaction" column:"transaction" sql:"BIGINT"`
	// Transaction's unique code based on creation time
	TransactionID string `json:"transactionId" column:"transaction_id" sql:"VARCHAR"`

	ParentID         *int                   `json:"parentId" column:"parent_id" sql:"BIGINT"`
	Status           BorrowingRequestStatus `json:"status" column:"status" sql:"borrowing_request_status_enum"`
	IsVerified       bool                   `json:"isVerified" column:"is_verified" sql:"BOOLEAN"`
	Notes            string                 `json:"notes" column:"notes" sql:"VARCHAR"`
	ActualReturnDate *time.Time             `json:"actualReturnDate" column:"actual_return_date" sql:"TIMESTAMPTZ DEFAULT NULL"`
	DueReturnDate    time.Time              `json:"dueReturnDate" column:"due_return_date" sql:"TIMESTAMPTZ NOT NULL"`
	CreatedAt        time.Time              `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt        time.Time              `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *BorrowingRequest) TableName() string {
	return "borrowing_requests"
}

func (m *BorrowingRequest) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_borrowing_request_asset",
		"asset_id",
		"assets",
		"id",
		"",
	)

	builder.AddForeignKey(
		"fk_borrowing_request_transaction",
		"transaction_id",
		"borrowing_transactions",
		"id",
		"",
	)

	builder.AddForeignKey(
		"fk_borrowing_request_parent",
		"parent_id",
		"borrowing_requests",
		"id",
		"",
	)
}
