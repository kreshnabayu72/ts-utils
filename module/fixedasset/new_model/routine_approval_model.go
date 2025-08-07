package new_model

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RoutineApproval struct {
	ID            int            `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	TaskID        int            `json:"taskId" column:"task_id" sql:"BIGINT NOT NULL"`
	ApproverID    int            `json:"approverId" column:"approver_id" sql:"BIGINT"`
	ApprovalRound int            `json:"approvalRound" column:"approval_round" sql:"INTEGER NOT NULL"`
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
func (m *RoutineApproval) TableName() string {
	return "routine_approvals"
}

// Up defines the migration logic to create the table and its constraints.
func (m *RoutineApproval) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_routine_approval_task_id",
		"task_id",
		"routine_tasks",
		"id",
		"CASCADE",
	)
	builder.AddForeignKey(
		"fk_routine_approval_approver_id",
		"approver_id",
		"users",
		"id",
		"CASCADE",
	)

	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)
	postQuery, err := os.ReadFile(
		filepath.Join(baseDir, "..", "migration/routine_approval_post_query.sql"),
	)
	if err != nil {
		log.Fatal(err)
	}
	builder.SetPostQuery(string(postQuery))
}
