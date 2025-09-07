package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// TransferRequestHistory represents the transfer_request_histories table in the database.
type TransferRequestHistory struct {
	ID int `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	// Foreign key to parent request history
	ParentID *int `json:"parentId"            column:"parent_id"           sql:"BIGINT"`

	// Foreign key to assets
	Asset int `json:"asset"            column:"asset"           sql:"BIGINT"`
	// Asset ID based on alias code
	AssetID          string                     `json:"assetId"            column:"asset_id"           sql:"VARCHAR"`
	AssetStatus      TransferRequestAssetStatus `json:"assetStatus" column:"asset_status" sql:"transfer_request_asset_status_enum"`
	Name             string                     `json:"name"             column:"name"            sql:"VARCHAR"`
	Brand            string                     `json:"brand"             column:"brand"            sql:"VARCHAR"`
	Model            string                     `json:"model"             column:"model"            sql:"VARCHAR"`
	NameWithSequence string                     `json:"nameWithSequence"   column:"name_with_sequence" sql:"VARCHAR"`
	SequenceNumber   int                        `json:"sequenceNumber"     column:"sequence_number"    sql:"BIGINT"`
	Image            *string                    `json:"image"         column:"image"        sql:"VARCHAR"`

	// Foreign key to transfer_transaction_histories
	Transaction        int                            `json:"transaction" column:"transaction" sql:"BIGINT"`
	Status             TransferRequestStatus          `json:"status" column:"status" sql:"transfer_request_status_enum"`
	TagStatus          *TransferRequestAssetTagStatus `json:"tagStatus" column:"tag_status" sql:"transfer_request_asset_tag_status_enum"`
	OriginGroup        string                         `json:"originGroup"   column:"origin_group" sql:"VARCHAR"`
	OriginGroupID      int                            `json:"originGroupId"   column:"origin_group_id" sql:"BIGINT"`
	DestinationGroup   string                         `json:"destinationGroup"   column:"destination_group" sql:"VARCHAR"`
	DestinationGroupID int                            `json:"destinationGroupId"   column:"destination_group_id" sql:"BIGINT"`
	Manager            string                         `json:"manager"   column:"manager" sql:"VARCHAR"`
	Notes              string                         `json:"notes" column:"notes" sql:"VARCHAR"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *TransferRequestHistory) TableName() string {
	return "transfer_request_histories"
}

// Up defines the migration logic to create the table and its constraints.
func (m *TransferRequestHistory) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add the index for asset name with sequence
	builder.AddIndex("transfer_transaction_histories_name_with_sequence", "name_with_sequence", false, "")

	builder.AddForeignKey(
		"fk_transfer_request_history_parent",
		"parent_id",
		"transfer_request_histories",
		"id",
		"",
	)
}
