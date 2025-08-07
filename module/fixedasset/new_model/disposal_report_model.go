package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// DisposalReport represents the disposal_reports table in the database.
type DisposalReport struct {
	ID            int                   `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	AssetID       int                   `json:"assetId" column:"asset_id" sql:"BIGINT"`
	ReportedByID  int                   `json:"reportedById" column:"reported_by_id" sql:"BIGINT"`
	LastCondition DisposalLastCondition `json:"lastCondition" column:"last_condition" sql:"disposal_last_condition_enum"`
	IsDeleted     bool                  `json:"isDeleted" column:"is_deleted" sql:"BOOLEAN"`
	CreatedAt     time.Time             `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time             `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *DisposalReport) TableName() string {
	return "disposal_reports"
}

// Up defines the migration logic to create the table and its constraints.
func (m *DisposalReport) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_disposal_reports_asset",
		"asset_id",
		"assets",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_disposal_reports_prelist",
		"reported_by_id",
		"users",
		"id",
		"",
	)
}
