package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// DisposalRequDisposalRequestHistoryest represents the disposal_request_histories table in the database.
type DisposalRequestHistory struct {
	ID               int                   `json:"id"                 column:"disposal_request_histories.id"                 sql:"BIGINT PRIMARY KEY"`
	AssetID          string                `json:"assetId"            column:"disposal_request_histories.asset_id"           sql:"VARCHAR"`
	TagType          *TagType              `json:"tagType"            column:"disposal_request_histories.tag_type"           sql:"tag_type_enum"`
	Name             string                `json:"name"               column:"disposal_request_histories.name"               sql:"VARCHAR"`
	Brand            string                `json:"brand"              column:"disposal_request_histories.brand"              sql:"VARCHAR"`
	Model            string                `json:"model"              column:"disposal_request_histories.model"              sql:"VARCHAR"`
	NameWithSequence string                `json:"nameWithSequence"   column:"disposal_request_histories.name_with_sequence" sql:"VARCHAR"`
	SequenceNumber   int                   `json:"sequenceNumber"     column:"disposal_request_histories.sequence_number"    sql:"BIGINT"`
	Image            *string               `json:"image"              column:"disposal_request_histories.image"              sql:"VARCHAR"`
	OriginGroup      string                `json:"originGroup"        column:"disposal_request_histories.origin_group"       sql:"VARCHAR"`
	OriginGroupID    int                   `json:"originGroupId"      column:"disposal_request_histories.origin_group_id"    sql:"BIGINT"`
	DisposalGroup    string                `json:"disposalGroup"      column:"disposal_request_histories.disposal_group"     sql:"VARCHAR"`
	DisposalGroupID  int                   `json:"disposalGroupId"    column:"disposal_request_histories.disposal_group_id"  sql:"BIGINT"`
	Transaction      int                   `json:"transaction"        column:"disposal_request_histories.transaction"        sql:"BIGINT"`
	Status           DisposalRequestStatus `json:"status"             column:"disposal_request_histories.status"             sql:"disposal_request_status_enum"`
	SalvageableParts []string              `json:"salvageableParts"   column:"disposal_request_histories.salvageable_parts"  sql:"TEXT[]"`
	SalePrice        string                `json:"salePrice"          column:"disposal_request_histories.sale_price"         sql:"VARCHAR"`
	LastCondition    DisposalLastCondition `json:"lastCondition"      column:"disposal_request_histories.last_condition"     sql:"disposal_last_condition_enum"`
	Method           DisposalMethod        `json:"method"             column:"disposal_request_histories.method"             sql:"disposal_method_enum"`
	Notes            string                `json:"notes"              column:"disposal_request_histories.notes"              sql:"VARCHAR"`
	Manager          string                `json:"manager"            column:"disposal_request_histories.manager"            sql:"VARCHAR"`
	CreatedAt        time.Time             `json:"createdAt"          column:"disposal_request_histories.created_at"         sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt        time.Time             `json:"updatedAt"          column:"disposal_request_histories.updated_at"         sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *DisposalRequestHistory) TableName() string {
	return "disposal_request_histories"
}

// Up defines the migration logic to create the table and its constraints.
func (m *DisposalRequestHistory) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)
}
