package new_model

import (
	"time"
)

// DisposalRequestHistory represents the disposal_request_histories table in the database.
type DisposalRequestHistory struct {
	ID               int                   `json:"id"                 column:"disposal_request_histories.id"`
	Asset            int                   `json:"asset"              column:"disposal_request_histories.asset"`
	AssetID          string                `json:"assetId"            column:"disposal_request_histories.asset_id"`
	TagType          *TagType              `json:"tagType"            column:"disposal_request_histories.tag_type"`
	Name             string                `json:"name"               column:"disposal_request_histories.name"`
	Brand            string                `json:"brand"              column:"disposal_request_histories.brand"`
	Model            string                `json:"model"              column:"disposal_request_histories.model"`
	NameWithSequence string                `json:"nameWithSequence"   column:"disposal_request_histories.name_with_sequence"`
	SequenceNumber   int                   `json:"sequenceNumber"     column:"disposal_request_histories.sequence_number"`
	Image            *string               `json:"image"              column:"disposal_request_histories.image"`
	OriginGroup      string                `json:"originGroup"        column:"disposal_request_histories.origin_group"`
	OriginGroupID    int                   `json:"originGroupId"      column:"disposal_request_histories.origin_group_id"`
	DisposalGroup    string                `json:"disposalGroup"      column:"disposal_request_histories.disposal_group"`
	DisposalGroupID  int                   `json:"disposalGroupId"    column:"disposal_request_histories.disposal_group_id"`
	Transaction      int                   `json:"transaction"        column:"disposal_request_histories.transaction"`
	Status           DisposalRequestStatus `json:"status"             column:"disposal_request_histories.status"`
	SalvageableParts []string              `json:"salvageableParts"   column:"disposal_request_histories.salvageable_parts"`
	SalePrice        string                `json:"salePrice"          column:"disposal_request_histories.sale_price"`
	LastCondition    DisposalLastCondition `json:"lastCondition"      column:"disposal_request_histories.last_condition"`
	Method           DisposalMethod        `json:"method"             column:"disposal_request_histories.method"`
	Notes            string                `json:"notes"              column:"disposal_request_histories.notes"`
	Manager          string                `json:"manager"            column:"disposal_request_histories.manager"`
	CreatedAt        time.Time             `json:"createdAt"          column:"disposal_request_histories.created_at"`
	UpdatedAt        time.Time             `json:"updatedAt"          column:"disposal_request_histories.updated_at"`
}
