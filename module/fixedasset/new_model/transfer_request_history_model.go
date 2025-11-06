package new_model

import (
	"time"
)

// TransferRequestHistory represents the transfer_request_histories table in the database.
type TransferRequestHistory struct {
	ID                 int                            `json:"id"                 column:"transfer_request_histories.id"`
	ParentID           *int                           `json:"parentId"           column:"transfer_request_histories.parent_id"`
	TagType            TagType                        `json:"tagType"            column:"transfer_request_histories.tag_type"`
	Asset              int                            `json:"asset"              column:"transfer_request_histories.asset"`
	AssetID            string                         `json:"assetId"            column:"transfer_request_histories.asset_id"`
	AliasCode          *string                        `json:"aliasCode"                column:"transfer_request_histories.alias_code"`
	AssetStatus        *TransferRequestAssetStatus    `json:"assetStatus"        column:"transfer_request_histories.asset_status"`
	Name               string                         `json:"name"               column:"transfer_request_histories.name"`
	Brand              string                         `json:"brand"              column:"transfer_request_histories.brand"`
	Model              string                         `json:"model"              column:"transfer_request_histories.model"`
	NameWithSequence   string                         `json:"nameWithSequence"   column:"transfer_request_histories.name_with_sequence"`
	SequenceNumber     int                            `json:"sequenceNumber"     column:"transfer_request_histories.sequence_number"`
	Image              *string                        `json:"image"              column:"transfer_request_histories.image"`
	Transaction        *int                           `json:"transaction"        column:"transfer_request_histories.transaction"`
	TransactionID      string                         `json:"transactionId"             column:"transfer_request_histories.transaction_id"`
	Status             TransferRequestStatus          `json:"status"             column:"transfer_request_histories.status"`
	TagStatus          *TransferRequestAssetTagStatus `json:"tagStatus"          column:"transfer_request_histories.tag_status"`
	OriginGroup        string                         `json:"originGroup"        column:"transfer_request_histories.origin_group"`
	OriginGroupID      int                            `json:"originGroupId"      column:"transfer_request_histories.origin_group_id"`
	DestinationGroup   *string                        `json:"destinationGroup"   column:"transfer_request_histories.destination_group"`
	DestinationGroupID *int                           `json:"destinationGroupId" column:"transfer_request_histories.destination_group_id"`
	Manager            *string                        `json:"manager"            column:"transfer_request_histories.manager"`
	Notes              string                         `json:"notes"              column:"transfer_request_histories.notes"`
	CreatedAt          time.Time                      `json:"createdAt"          column:"transfer_request_histories.created_at"`
	UpdatedAt          time.Time                      `json:"updatedAt"          column:"transfer_request_histories.updated_at"`
}
