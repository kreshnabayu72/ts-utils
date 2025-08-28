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

type BorrowingApproval struct {
	ID            int            `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	RequestID     int            `json:"requestId" column:"request_id" sql:"BIGINT"`
	ApproverID    int            `json:"approverId" column:"approver_id" sql:"BIGINT"`
	Level         int            `json:"level" column:"level" sql:"INTEGER"`
	ApprovalRound int            `json:"approval_round" column:"approval_round" sql:"INTEGER DEFAULT 1"`
	Status        ApprovalStatus `json:"status" column:"status" sql:"approval_status_enum"`
	IsApproved    *bool          `json:"isApproved" column:"is_approved" sql:"BOOLEAN NULL DEFAULT NULL"`
	Notes         *string        `json:"notes" column:"notes" sql:"VARCHAR"`
	Type          ApprovalType   `json:"type" column:"type" sql:"approval_type_enum NOT NULL"`
	ApprovedAt    *time.Time     `json:"approvedAt" column:"approved_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt     time.Time      `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time      `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *BorrowingApproval) TableName() string {
	return "borrowing_approvals"
}

func (m *BorrowingApproval) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_borrowing_approvals_request",
		"request_id",
		"borrowing_requests",
		"id",
		"",
	)

	builder.AddForeignKey(
		"fk_borrowing_approvals_approver",
		"approver_id",
		"users",
		"id",
		"",
	)

	// Insert default data
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)
	path := filepath.Join(baseDir, "migration/borrowing_approval_post_query.sql")

	postQuery, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error path: " + path)
		log.Fatal(err)
	}
	builder.SetPostQuery(string(postQuery))
}
