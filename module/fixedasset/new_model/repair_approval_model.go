package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	fixed_asset_migration "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model/migration"
)

type RepairApproval struct {
	ID            int            `json:"id"            column:"repair_approvals.id"             sql:"BIGINT PRIMARY KEY"`
	RepairID      int            `json:"repairId"      column:"repair_approvals.repair_id"      sql:"BIGINT NOT NULL"`
	ApproverID    int            `json:"approverId"    column:"repair_approvals.approver_id"    sql:"BIGINT"`
	ApprovalRound int            `json:"approvalRound" column:"repair_approvals.approval_round" sql:"INTEGER NOT NULL"`
	Level         int            `json:"level"         column:"repair_approvals.level"          sql:"INTEGER"`
	Status        ApprovalStatus `json:"status"        column:"repair_approvals.status"         sql:"approval_status_enum"`
	IsApproved    *bool          `json:"isApproved"    column:"repair_approvals.is_approved"    sql:"BOOLEAN NULL DEFAULT NULL"`
	Notes         *string        `json:"notes"         column:"repair_approvals.notes"          sql:"VARCHAR"`
	Type          ApprovalType   `json:"type"          column:"repair_approvals.type"           sql:"approval_type_enum NOT NULL"`
	ApprovedAt    *time.Time     `json:"approvedAt"    column:"repair_approvals.approved_at"    sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt     time.Time      `json:"createdAt"     column:"repair_approvals.created_at"     sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time      `json:"updatedAt"     column:"repair_approvals.updated_at"     sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *RepairApproval) TableName() string {
	return "repair_approvals"
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

	// ─────── Insert default data ───────
	// NOTE: reading from sql file has been tried and doesnt work in server, so use string instead
	builder.SetPostQuery(fixed_asset_migration.RepairApprovalPostQuery)
}
