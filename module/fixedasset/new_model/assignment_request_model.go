package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// AssignmentRequest represents the assignment_requests table in the database.
type AssignmentRequest struct {
	ID             int                     `json:"id"             column:"assignment_requests.id"              sql:"BIGINT PRIMARY KEY"`
	Asset          int                     `json:"asset"          column:"assignment_requests.asset"           sql:"BIGINT"`
	Transaction    int                     `json:"transaction"    column:"assignment_requests.transaction"     sql:"BIGINT"`
	ParentID       *int                    `json:"parentId"       column:"assignment_requests.parent_id"       sql:"BIGINT"`
	RepairID       *int                    `json:"repairId"       column:"assignment_requests.repair_id"       sql:"BIGINT"`
	Status         AssignmentRequestStatus `json:"status"         column:"assignment_requests.status"          sql:"assignment_request_status_enum"`
	PreviousStatus AssignmentRequestStatus `json:"previousStatus" column:"assignment_requests.previous_status" sql:"assignment_request_status_enum NULL DEFAULT NULL"`
	IsVerified     bool                    `json:"isVerified"     column:"assignment_requests.is_verified"     sql:"BOOLEAN"`
	Notes          string                  `json:"notes"          column:"assignment_requests.notes"           sql:"VARCHAR"`
	SerialNumber   *string                 `json:"serialNumber"   column:"assignment_requests.serial_number"   sql:"VARCHAR"`
	CreatedAt      time.Time               `json:"createdAt"      column:"assignment_requests.created_at"      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt      time.Time               `json:"updatedAt"      column:"assignment_requests.updated_at"      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *AssignmentRequest) TableName() string {
	return "assignment_requests"
}

// Up defines the migration logic to create the table and its constraints.
func (m *AssignmentRequest) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_assignment_request_asset",
		"asset",
		"assets",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_assignment_request_transaction",
		"transaction",
		"assignment_transactions",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_assignment_request_parent",
		"parent_id",
		"assignment_requests",
		"id",
		"",
	)
}
