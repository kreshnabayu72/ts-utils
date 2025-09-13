package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// AssignmentRequestHistory represents the assignment_request_histories table in the database.
type AssignmentRequestHistory struct {
	ID               int                     `json:"id"               column:"assignment_request_histories.id"                 sql:"BIGINT PRIMARY KEY"`
	ParentID         *int                    `json:"parentId"         column:"assignment_request_histories.parent_id"           sql:"BIGINT"`
	AssetID          string                  `json:"assetId"          column:"assignment_request_histories.asset_id"           sql:"VARCHAR"`
	Name             string                  `json:"name"             column:"assignment_request_histories.name"               sql:"VARCHAR"`
	Brand            string                  `json:"brand"            column:"assignment_request_histories.brand"              sql:"VARCHAR"`
	Model            string                  `json:"model"            column:"assignment_request_histories.model"              sql:"VARCHAR"`
	NameWithSequence string                  `json:"nameWithSequence" column:"assignment_request_histories.name_with_sequence" sql:"VARCHAR"`
	SequenceNumber   int                     `json:"sequenceNumber"   column:"assignment_request_histories.sequence_number"    sql:"BIGINT"`
	Image            *string                 `json:"image"            column:"assignment_request_histories.image"              sql:"VARCHAR"`
	Group            string                  `json:"group"            column:"assignment_request_histories.group"              sql:"VARCHAR"`
	GroupID          int                     `json:"groupId"          column:"assignment_request_histories.group_id"           sql:"BIGINT"`
	Transaction      int                     `json:"transaction"      column:"assignment_request_histories.transaction"        sql:"BIGINT"`
	Status           AssignmentRequestStatus `json:"status"           column:"assignment_request_histories.status"             sql:"assignment_request_status_enum"`
	Notes            string                  `json:"notes"            column:"assignment_request_histories.notes"              sql:"VARCHAR"`
	User             string                  `json:"user"             column:"assignment_request_histories.user"               sql:"VARCHAR"`
	Manager          string                  `json:"manager"          column:"assignment_request_histories.manager"            sql:"VARCHAR"`
	CreatedAt        time.Time               `json:"createdAt"        column:"assignment_request_histories.created_at"         sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt        time.Time               `json:"updatedAt"        column:"assignment_request_histories.updated_at"         sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *AssignmentRequestHistory) TableName() string {
	return "assignment_request_histories"
}

// Up defines the migration logic to create the table and its constraints.
func (m *AssignmentRequestHistory) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)
}
