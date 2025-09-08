package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	fixed_asset_migration "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model/migration"
)

// AssignmentApproval represents the assignment_approvals table in the database.
type AssignmentApproval struct {
	ID         int            `json:"id"         column:"assignment_approvals.id"          sql:"BIGINT PRIMARY KEY"`
	RequestID  int            `json:"requestId"  column:"assignment_approvals.request_id"  sql:"BIGINT"`
	ApproverID int            `json:"approverId" column:"assignment_approvals.approver_id" sql:"BIGINT"`
	Level      int            `json:"level"      column:"assignment_approvals.level"       sql:"INTEGER"`
	Status     ApprovalStatus `json:"status"     column:"assignment_approvals.status"      sql:"approval_status_enum"`
	IsApproved *bool          `json:"isApproved" column:"assignment_approvals.is_approved" sql:"BOOLEAN NULL DEFAULT NULL"`
	Notes      *string        `json:"notes"      column:"assignment_approvals.notes"       sql:"VARCHAR"`
	Type       ApprovalType   `json:"type"       column:"assignment_approvals.type"        sql:"approval_type_enum NOT NULL"`
	ApprovedAt *time.Time     `json:"approvedAt" column:"assignment_approvals.approved_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt  time.Time      `json:"createdAt"  column:"assignment_approvals.created_at"  sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt  time.Time      `json:"updatedAt"  column:"assignment_approvals.updated_at"  sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *AssignmentApproval) TableName() string {
	return "assignment_approvals"
}

// Up defines the migration logic to create the table and its constraints.
func (m *AssignmentApproval) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_assignment_approvals_request",
		"request_id",
		"assignment_requests",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_assignment_approvals_approver",
		"approver_id",
		"users",
		"id",
		"",
	)

	// ─────── Insert default data ───────
	// NOTE: reading from sql file has been tried and doesnt work in server, so use string instead

	builder.SetPostQuery(fixed_asset_migration.AssignmentApprovalPostQuery)
}
