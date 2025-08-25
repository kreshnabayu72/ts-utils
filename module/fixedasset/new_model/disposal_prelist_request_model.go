package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// DisposalPrelistRequest represents the disposal_prelist_requests table in the database.
type DisposalPrelistRequest struct {
	ID        int       `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	AssetID   int       `json:"assetId" column:"asset_id" sql:"BIGINT"`
	PrelistID int       `json:"prelistId" column:"prelist_id" sql:"BIGINT"`
	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *DisposalPrelistRequest) TableName() string {
	return "disposal_prelist_requests"
}

// Up defines the migration logic to create the table and its constraints.
func (m *DisposalPrelistRequest) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_disposal_prelist_req_asset",
		"asset_id",
		"assets",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_disposal_prelist_req_prelist",
		"prelist_id",
		"disposal_prelists",
		"id",
		"CASCADE",
	)
}
