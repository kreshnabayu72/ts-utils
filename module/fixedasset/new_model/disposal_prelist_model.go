package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// DisposalPrelist represents the disposal_prelists table in the database.
type DisposalPrelist struct {
	ID              int            `json:"id"              column:"disposal_prelists.id"                sql:"BIGINT PRIMARY KEY"`
	DisposalGroupID int            `json:"disposalGroupId" column:"disposal_prelists.disposal_group_id" sql:"BIGINT"`
	Method          DisposalMethod `json:"method"          column:"disposal_prelists.method"            sql:"disposal_method_enum"`
	Organization    string         `json:"organization"    column:"disposal_prelists.organization"      sql:"VARCHAR"`
	CreatedAt       time.Time      `json:"createdAt"       column:"disposal_prelists.created_at"        sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt       time.Time      `json:"updatedAt"       column:"disposal_prelists.updated_at"        sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *DisposalPrelist) TableName() string {
	return "disposal_prelists"
}

// Up defines the migration logic to create the table and its constraints.
func (m *DisposalPrelist) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_disposal_prelists_disposal_group_id",
		"disposal_group_id",
		"groups",
		"id",
		"",
	)
}
