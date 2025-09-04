package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// AssignmentRequest represents the assignment_requests table in the database.
type AssignmentRequest struct {
	ID int `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`

	// Foreign key to assets
	Asset int `json:"asset" column:"asset" sql:"BIGINT"`
	// Foreign key to related transaction
	Transaction int `json:"transaction" column:"transaction" sql:"BIGINT"`

	ParentID       *int                    `json:"parentId" column:"parent_id" sql:"BIGINT"`
	Status         AssignmentRequestStatus `json:"status" column:"status" sql:"assignment_request_status_enum"`
	PreviousStatus AssignmentRequestStatus `json:"previousStatus" column:"previous_status" sql:"assignment_request_status_enum NULL DEFAULT NULL"`

	IsVerified bool      `json:"isVerified" column:"is_verified" sql:"BOOLEAN"`
	Notes      string    `json:"notes" column:"notes" sql:"VARCHAR"`
	CreatedAt  time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt  time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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
