package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// AssignmentPrelist represents the assignment_prelists table in the database.
type AssignmentPrelist struct {
	ID           int                   `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	GroupID      int                   `json:"groupId" column:"group_id" sql:"BIGINT"`
	AssignedToID int                   `json:"assignedToId" column:"assigned_to_id" sql:"BIGINT"`
	Type         AssignmentPrelistType `json:"type" column:"type" sql:"assignment_prelist_type_enum"`
	CreatedAt    time.Time             `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt    time.Time             `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *AssignmentPrelist) TableName() string {
	return "assignment_prelists"
}

// Up defines the migration logic to create the table and its constraints.
func (m *AssignmentPrelist) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_assignment_prelists_group",
		"group_id",
		"groups",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_assignment_prelists_assigned_to",
		"assigned_to_id",
		"users",
		"id",
		"",
	)
}
