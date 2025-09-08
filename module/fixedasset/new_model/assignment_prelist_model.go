package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// AssignmentPrelist represents the assignment_prelists table in the database.
type AssignmentPrelist struct {
	ID           int                   `json:"id"           column:"assignment_prelists.id"             sql:"BIGINT PRIMARY KEY"`
	GroupID      int                   `json:"groupId"      column:"assignment_prelists.group_id"       sql:"BIGINT"`
	AssignedToID int                   `json:"assignedToId" column:"assignment_prelists.assigned_to_id" sql:"BIGINT"`
	Type         AssignmentPrelistType `json:"type"         column:"assignment_prelists.type"           sql:"assignment_prelist_type_enum"`
	CreatedAt    time.Time             `json:"createdAt"    column:"assignment_prelists.created_at"     sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt    time.Time             `json:"updatedAt"    column:"assignment_prelists.updated_at"     sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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

	// Add the unique index for user id, group id, and type
	builder.AddIndex(
		"assignment_prelists_group_id_assigned_to_id_type",
		"group_id,assigned_to_id,type",
		true,
		"",
	)
}
