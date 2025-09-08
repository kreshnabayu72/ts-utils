package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// TransferRequest represents the transfer_requests table in the database.
type TransferRequest struct {
	ID             int                           `json:"id"             column:"transfer_requests.id"               sql:"BIGINT PRIMARY KEY"`
	Asset          int                           `json:"asset"          column:"transfer_requests.asset"            sql:"BIGINT"`
	Transaction    int                           `json:"transaction"    column:"transfer_requests.transaction"      sql:"BIGINT"`
	ParentID       *int                          `json:"parentId"       column:"transfer_requests.parent_id"        sql:"BIGINT"`
	Status         TransferRequestStatus         `json:"status"         column:"transfer_requests.status"           sql:"transfer_request_status_enum"`
	AssetTagStatus TransferRequestAssetTagStatus `json:"assetTagStatus" column:"transfer_requests.asset_tag_status" sql:"transfer_request_asset_tag_status_enum"`
	AssetStatus    *TransferRequestAssetStatus   `json:"assetStatus"    column:"transfer_requests.asset_status"     sql:"transfer_request_asset_status_enum"`
	IsVerified     bool                          `json:"isVerified"     column:"transfer_requests.is_verified"      sql:"BOOLEAN"`
	Notes          string                        `json:"notes"          column:"transfer_requests.notes"            sql:"VARCHAR"`
	CreatedAt      time.Time                     `json:"createdAt"      column:"transfer_requests.created_at"       sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt      time.Time                     `json:"updatedAt"      column:"transfer_requests.updated_at"       sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *TransferRequest) TableName() string {
	return "transfer_requests"
}

// Up defines the migration logic to create the table and its constraints.
func (m *TransferRequest) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_transfer_request_asset",
		"asset",
		"assets",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_transfer_request_transaction",
		"transaction",
		"transfer_transactions",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_transfer_request_parent",
		"parent_id",
		"transfer_requests",
		"id",
		"",
	)
}
