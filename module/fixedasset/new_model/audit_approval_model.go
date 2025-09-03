package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// AuditApproval represents the audit_approvals table in the database.
type AuditApproval struct {
	ID            int            `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	TaskId        int            `json:"taskId" column:"task_id" sql:"BIGINT"`
	ApproverID    int            `json:"approverId" column:"approver_id" sql:"BIGINT"`
	ApprovalRound int            `json:"approvalRound" column:"approval_round" sql:"INTEGER"`
	Level         int            `json:"level" column:"level" sql:"INTEGER"`
	Status        ApprovalStatus `json:"status" column:"status" sql:"approval_status_enum"`
	IsApproved    *bool          `json:"isApproved" column:"is_approved" sql:"BOOLEAN NULL DEFAULT NULL"`
	Notes         *string        `json:"notes" column:"notes" sql:"VARCHAR"`
	Type          ApprovalType   `json:"type" column:"type" sql:"approval_type_enum NOT NULL"`
	ApprovedAt    *time.Time     `json:"approvedAt" column:"approved_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt     time.Time      `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time      `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *AuditApproval) TableName() string {
	return "audit_approvals"
}

// Up defines the migration logic to create the table and its constraints.
func (m *AuditApproval) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_audit_approvals_task",
		"task_id",
		"audit_tasks",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_audit_approvals_approver",
		"approver_id",
		"users",
		"id",
		"",
	)
}
