package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// TransferPrelist represents the transfer_prelists table in the database.
type TransferPrelist struct {
	ID                 int       `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	OriginGroupID      int       `json:"originGroupId" column:"origin_group_id" sql:"BIGINT"`
	DestinationGroupID int       `json:"destinationGroupId" column:"destination_group_id" sql:"BIGINT"`
	CreatedAt          time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt          time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *TransferPrelist) TableName() string {
	return "transfer_prelists"
}

// Up defines the migration logic to create the table and its constraints.
func (m *TransferPrelist) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_transfer_prelists_origin_group_id",
		"origin_group_id",
		"groups",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_transfer_prelists_destination_group_id",
		"destination_group_id",
		"groups",
		"id",
		"",
	)

	// Add the unique index for origin group id and destination group id
	builder.AddIndex(
		"transfer_prelists_origin_group_id_destination_group_id",
		"origin_group_id,destination_group_id",
		true,
		"",
	)
}
