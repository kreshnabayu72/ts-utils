package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type BorrowingRequest struct {
	ID            int                    `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	AssetID       int                    `json:"assetId" column:"asset_id" sql:"BIGINT"`
	TransactionID int                    `json:"transactionId" column:"transaction_id" sql:"BIGINT"`
	ParentID      *int                   `json:"parentId" column:"parent_id" sql:"BIGINT"`
	Status        BorrowingRequestStatus `json:"status" column:"status" sql:"borrowing_request_status_enum"`
	IsVerified    bool                   `json:"isVerified" column:"is_verified" sql:"BOOLEAN"`
	Notes         string                 `json:"notes" column:"notes" sql:"VARCHAR"`
	CreatedAt     time.Time              `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time              `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
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
