package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// TransferRequestHistory represents the transfer_request_histories table in the database.
type TransferRequestHistory struct {
	ID                 int                            `json:"id"                 column:"transfer_request_histories.id"                   sql:"BIGINT PRIMARY KEY"`
	ParentID           *int                           `json:"parentId"           column:"transfer_request_histories.parent_id"            sql:"BIGINT"`
	Asset              int                            `json:"asset"              column:"transfer_request_histories.asset"              sql:"BIGINT"`
	AssetID            string                         `json:"assetId"            column:"transfer_request_histories.asset_id"             sql:"VARCHAR"`
	AssetStatus        TransferRequestAssetStatus     `json:"assetStatus"        column:"transfer_request_histories.asset_status"         sql:"transfer_request_asset_status_enum"`
	Name               string                         `json:"name"               column:"transfer_request_histories.name"                 sql:"VARCHAR"`
	Brand              string                         `json:"brand"              column:"transfer_request_histories.brand"                sql:"VARCHAR"`
	Model              string                         `json:"model"              column:"transfer_request_histories.model"                sql:"VARCHAR"`
	NameWithSequence   string                         `json:"nameWithSequence"   column:"transfer_request_histories.name_with_sequence"   sql:"VARCHAR"`
	SequenceNumber     int                            `json:"sequenceNumber"     column:"transfer_request_histories.sequence_number"      sql:"BIGINT"`
	Image              *string                        `json:"image"              column:"transfer_request_histories.image"                sql:"VARCHAR"`
	Transaction        int                            `json:"transaction"        column:"transfer_request_histories.transaction"          sql:"BIGINT"`
	Status             TransferRequestStatus          `json:"status"             column:"transfer_request_histories.status"               sql:"transfer_request_status_enum"`
	TagStatus          *TransferRequestAssetTagStatus `json:"tagStatus"          column:"transfer_request_histories.tag_status"           sql:"transfer_request_asset_tag_status_enum"`
	OriginGroup        string                         `json:"originGroup"        column:"transfer_request_histories.origin_group"         sql:"VARCHAR"`
	OriginGroupID      int                            `json:"originGroupId"      column:"transfer_request_histories.origin_group_id"      sql:"BIGINT"`
	DestinationGroup   string                         `json:"destinationGroup"   column:"transfer_request_histories.destination_group"    sql:"VARCHAR"`
	DestinationGroupID int                            `json:"destinationGroupId" column:"transfer_request_histories.destination_group_id" sql:"BIGINT"`
	Manager            string                         `json:"manager"            column:"transfer_request_histories.manager"              sql:"VARCHAR"`
	Notes              string                         `json:"notes"              column:"transfer_request_histories.notes"                sql:"VARCHAR"`
	CreatedAt          time.Time                      `json:"createdAt"          column:"transfer_request_histories.created_at"           sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt          time.Time                      `json:"updatedAt"          column:"transfer_request_histories.updated_at"           sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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
