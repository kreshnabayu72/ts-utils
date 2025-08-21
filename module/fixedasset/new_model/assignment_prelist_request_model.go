package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// AssignmentPrelistRequest represents the assignment_prelist_requests table in the database.
type AssignmentPrelistRequest struct {
	ID        int       `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	AssetID   int       `json:"assetId" column:"asset_id" sql:"BIGINT"`
	PrelistID int       `json:"prelistId" column:"prelist_id" sql:"BIGINT"`
	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *AssignmentPrelistRequest) TableName() string {
	return "assignment_prelist_requests"
}

// Up defines the migration logic to create the table and its constraints.
func (m *AssignmentPrelistRequest) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_assignment_prelist_req_asset",
		"asset_id",
		"assets",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_assignment_prelist_req_prelist",
		"prelist_id",
		"assignment_prelists",
		"id",
		"",
	)

	// Add the unique index for prelist id and asset id
	builder.AddIndex(
		"assignment_prelist_requests_prelist_id_asset_id",
		"prelist_id,asset_id",
		true,
		"",
	)
}
