package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	fixed_asset_migration "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model/migration"
)

type BorrowingApproval struct {
	ID            int            `json:"id"             column:"borrowing_approvals.id"             sql:"BIGINT PRIMARY KEY"`
	RequestID     int            `json:"requestId"      column:"borrowing_approvals.request_id"     sql:"BIGINT"`
	ApproverID    int            `json:"approverId"     column:"borrowing_approvals.approver_id"    sql:"BIGINT"`
	Level         int            `json:"level"          column:"borrowing_approvals.level"          sql:"INTEGER"`
	ApprovalRound int            `json:"approval_round" column:"borrowing_approvals.approval_round" sql:"INTEGER DEFAULT 1"`
	Status        ApprovalStatus `json:"status"         column:"borrowing_approvals.status"         sql:"approval_status_enum"`
	IsApproved    *bool          `json:"isApproved"     column:"borrowing_approvals.is_approved"    sql:"BOOLEAN NULL DEFAULT NULL"`
	Notes         *string        `json:"notes"          column:"borrowing_approvals.notes"          sql:"VARCHAR"`
	Type          ApprovalType   `json:"type"           column:"borrowing_approvals.type"           sql:"approval_type_enum NOT NULL"`
	ApprovedAt    *time.Time     `json:"approvedAt"     column:"borrowing_approvals.approved_at"    sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt     time.Time      `json:"createdAt"      column:"borrowing_approvals.created_at"     sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time      `json:"updatedAt"      column:"borrowing_approvals.updated_at"     sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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

	// ─────── Insert default data ───────
	// NOTE: reading from sql file has been tried and doesnt work in server, so use string instead
	builder.SetPostQuery(fixed_asset_migration.BorrowingApprovalPostQuery)
}
