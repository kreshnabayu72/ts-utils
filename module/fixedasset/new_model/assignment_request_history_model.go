package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// AssignmentRequestHistory represents the assignment_request_histories table in the database.
type AssignmentRequestHistory struct {
	ID int `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`

	// Foreign key to assets
	Asset int `json:"asset"            column:"asset"           sql:"BIGINT"`
	// Asset ID based on alias code
	AssetID          string  `json:"assetId"            column:"asset_id"           sql:"VARCHAR"`
	Name             string  `json:"name"             column:"name"            sql:"VARCHAR"`
	Brand            string  `json:"brand"             column:"brand"            sql:"VARCHAR"`
	Model            string  `json:"model"             column:"model"            sql:"VARCHAR"`
	NameWithSequence string  `json:"nameWithSequence"   column:"name_with_sequence" sql:"VARCHAR"`
	SequenceNumber   int     `json:"sequenceNumber"     column:"sequence_number"    sql:"BIGINT"`
	Image            *string `json:"image"         column:"image"        sql:"VARCHAR"`
	Group            string  `json:"group"   column:"group" sql:"VARCHAR"`
	GroupID          int     `json:"groupId"   column:"group_id" sql:"BIGINT"`

	Transaction int                     `json:"transaction" column:"transaction" sql:"BIGINT"`
	Status      AssignmentRequestStatus `json:"status" column:"status" sql:"assignment_request_status_enum"`
	Notes       string                  `json:"notes" column:"notes" sql:"VARCHAR"`
	User        string                  `json:"user"   column:"user" sql:"VARCHAR"`
	Manager     string                  `json:"manager"   column:"manager" sql:"VARCHAR"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *AssignmentRequestHistory) TableName() string {
	return "assignment_request_histories"
}

// Up defines the migration logic to create the table and its constraints.
func (m *AssignmentRequestHistory) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_assignment_history_asset",
		"asset",
		"assets",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_assignment_request_history_transaction",
		"transaction_id",
		"assignment_transaction_histories",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_assignment_request_history_parent",
		"parent_id",
		"assignment_request_histories",
		"id",
		"",
	)
}
