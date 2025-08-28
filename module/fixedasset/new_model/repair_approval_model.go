package new_model

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RepairApproval struct {
	ID            int            `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	RepairID      int            `json:"repairId" column:"repair_id" sql:"BIGINT NOT NULL"`
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
func (m *RepairApproval) TableName() string {
	return "routine_approvals"
}

// Up defines the migration logic to create the table and its constraints.
func (m *RepairApproval) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_repair_approval_repair_id",
		"repair_id",
		"repair_ticketings",
		"id",
		"CASCADE",
	)
	builder.AddForeignKey(
		"fk_repair_approval_approver_id",
		"approver_id",
		"users",
		"id",
		"CASCADE",
	)

	// Insert default data
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)
	path := filepath.Join(baseDir, "migration/repair_approval_post_query.sql")

	postQuery, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error path: " + path)
		log.Fatal(err)
	}
	builder.SetPostQuery(string(postQuery))
}
