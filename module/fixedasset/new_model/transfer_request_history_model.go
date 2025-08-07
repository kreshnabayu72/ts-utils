package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// TransferRequestHistory represents the transfer_request_histories table in the database.
type TransferRequestHistory struct {
	ID int `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`

	Asset            string                     `json:"asset"            column:"asset"           sql:"VARCHAR"`
	AssetID          string                     `json:"assetId"            column:"asset_id"           sql:"VARCHAR"`
	AssetStatus      TransferRequestAssetStatus `json:"assetStatus" column:"asset_status" sql:"transfer_request_asset_status_enum"`
	Name             string                     `json:"name"             column:"name"            sql:"VARCHAR"`
	NameWithSequence string                     `json:"nameWithSequence"   column:"name_with_sequence" sql:"VARCHAR"`
	SequenceNumber   int                        `json:"sequenceNumber"     column:"sequence_number"    sql:"BIGINT"`
	Image            *string                    `json:"image"         column:"image"        sql:"VARCHAR"`

	Transaction      int                           `json:"transaction" column:"transaction" sql:"BIGINT"`
	Status           TransferRequestStatus         `json:"status" column:"status" sql:"transfer_request_status_enum"`
	TagStatus        TransferRequestAssetTagStatus `json:"tagStatus" column:"tag_status" sql:"transfer_request_asset_tag_status_enum"`
	OriginGroup      string                        `json:"originGroup"   column:"origin_group" sql:"VARCHAR"`
	DestinationGroup string                        `json:"destinationGroup"   column:"destination_group" sql:"VARCHAR"`
	Manager          string                        `json:"manager"   column:"manager" sql:"VARCHAR"`

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
}
