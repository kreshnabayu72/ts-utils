package new_model

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// TransferApproval represents the transfer_approvals table in the database.
type TransferApproval struct {
	ID          int               `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	Transaction int               `json:"transaction" column:"transaction" sql:"BIGINT"`
	RequestID   int               `json:"requestId" column:"request_id" sql:"BIGINT"`
	ApproverID  int               `json:"approverId" column:"approver_id" sql:"BIGINT"`
	Level       int               `json:"level" column:"level" sql:"INTEGER"`
	Status      ApprovalStatus    `json:"status" column:"status" sql:"approval_status_enum"`
	IsApproved  *bool             `json:"isApproved" column:"is_approved" sql:"BOOLEAN NULL DEFAULT NULL"`
	Notes       *string           `json:"notes" column:"notes" sql:"VARCHAR"`
	Type        ApprovalType      `json:"type" column:"type" sql:"approval_type_enum NOT NULL"`
	GroupType   TransferGroupType `json:"groupType" column:"group_type" sql:"transfer_group_type_enum NOT NULL"`
	ApprovedAt  *time.Time        `json:"approvedAt" column:"approved_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt   time.Time         `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt   time.Time         `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *TransferApproval) TableName() string {
	return "transfer_approvals"
}

// Up defines the migration logic to create the table and its constraints.
func (m *TransferApproval) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_transfer_approvals_request",
		"request_id",
		"transfer_requests",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_transfer_approvals_approver",
		"approver_id",
		"users",
		"id",
		"",
	)

	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)
	postQuery, err := os.ReadFile(
		filepath.Join(baseDir, "..", "migration/transfer_approval_post_query.sql"),
	)
	if err != nil {
		log.Fatal(err)
	}
	builder.SetPostQuery(string(postQuery))
}
