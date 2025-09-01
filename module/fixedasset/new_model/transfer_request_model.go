package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// TransferRequest represents the transfer_requests table in the database.
type TransferRequest struct {
	ID             int                           `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	Asset          int                           `json:"asset" column:"asset" sql:"BIGINT"`
	Transaction    int                           `json:"transaction" column:"transaction" sql:"BIGINT"`
	ParentID       *int                          `json:"parentId" column:"parent_id" sql:"BIGINT"`
	Status         TransferRequestStatus         `json:"status" column:"status" sql:"transfer_request_status_enum"`
	AssetTagStatus TransferRequestAssetTagStatus `json:"assetTagStatus" column:"asset_tag_status" sql:"transfer_request_asset_tag_status_enum"`
	AssetStatus    *TransferRequestAssetStatus   `json:"assetStatus" column:"asset_status" sql:"transfer_request_asset_status_enum"`
	IsVerified     bool                          `json:"isVerified" column:"is_verified" sql:"BOOLEAN"`
	Notes          string                        `json:"notes" column:"notes" sql:"VARCHAR"`
	CreatedAt      time.Time                     `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt      time.Time                     `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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
