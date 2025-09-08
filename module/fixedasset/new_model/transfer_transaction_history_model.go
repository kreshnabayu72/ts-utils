package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// TransferTransactionHistory represents the transfer_transactions table in the database.
type TransferTransactionHistory struct {
	ID                 int                       `json:"_id"                column:"transfer_transaction_histories.id"                   sql:"BIGINT PRIMARY KEY"`
	TransactionID      string                    `json:"transactionId"      column:"transfer_transaction_histories.transaction_id"       sql:"VARCHAR UNIQUE"`
	Manager            string                    `json:"manager"            column:"transfer_transaction_histories.manager"              sql:"VARCHAR"`
	OriginGroupID      int                       `json:"originGroupId"      column:"transfer_transaction_histories.origin_group_id"      sql:"BIGINT"`
	DestinationGroupID int                       `json:"destinationGroupId" column:"transfer_transaction_histories.destination_group_id" sql:"BIGINT"`
	OriginGroup        string                    `json:"originGroup"        column:"transfer_transaction_histories.origin_group"         sql:"VARCHAR"`
	DestinationGroup   string                    `json:"destinationGroup"   column:"transfer_transaction_histories.destination_group"    sql:"VARCHAR"`
	Status             TransferTransactionStatus `json:"status"             column:"transfer_transaction_histories.status"               sql:"transfer_transaction_status_enum"`
	TotalAssets        int                       `json:"totalAssets"        column:"transfer_transaction_histories.total_assets"         sql:"BIGINT"`
	CreatedAt          time.Time                 `json:"createdAt"          column:"transfer_transaction_histories.created_at"           sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt          time.Time                 `json:"updatedAt"          column:"transfer_transaction_histories.updated_at"           sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *TransferTransactionHistory) TableName() string {
	return "transfer_transaction_histories"
}

// Up defines the migration logic to create the table and its constraints.
func (m *TransferTransactionHistory) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add the unique index for transaction_id
	builder.AddIndex("transfer_transaction_histories_transaction_id", "transaction_id", true, "")
}
