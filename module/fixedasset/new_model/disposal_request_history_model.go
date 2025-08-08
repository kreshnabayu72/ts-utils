package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// DisposalRequDisposalRequestHistoryest represents the disposal_request_histories table in the database.
type DisposalRequestHistory struct {
	ID int `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`

	Asset            string  `json:"asset"            column:"asset"           sql:"VARCHAR"`
	AssetID          string  `json:"assetId"            column:"asset_id"           sql:"VARCHAR"`
	TagType          TagType `json:"tagType" column:"tag_type" sql:"tag_type_enum"`
	Name             string  `json:"name"             column:"name"            sql:"VARCHAR"`
	Brand            string  `json:"brand"             column:"brand"            sql:"VARCHAR"`
	Model            string  `json:"model"             column:"model"            sql:"VARCHAR"`
	NameWithSequence string  `json:"nameWithSequence"   column:"name_with_sequence" sql:"VARCHAR"`
	SequenceNumber   int     `json:"sequenceNumber"     column:"sequence_number"    sql:"BIGINT"`
	Image            *string `json:"image"         column:"image"        sql:"VARCHAR"`

	OriginGroup     string `json:"originGroup"   column:"origin_group" sql:"VARCHAR"`
	OriginGroupID   int    `json:"originGroupId"   column:"origin_group_id" sql:"BIGINT"`
	DisposalGroup   string `json:"disposalGroup"   column:"disposal_group" sql:"VARCHAR"`
	DisposalGroupID int    `json:"disposalGroupId"   column:"disposal_group_id" sql:"BIGINT"`

	Transaction      int                   `json:"transaction" column:"transaction" sql:"BIGINT"`
	Status           DisposalRequestStatus `json:"status" column:"status" sql:"disposal_request_status_enum"`
	SalvageableParts []string              `json:"salvageableParts" column:"salvageable_parts" sql:"TEXT[]"`
	SalePrice        string                `json:"salePrice" column:"sale_price" sql:"VARCHAR"`
	LastCondition    DisposalLastCondition `json:"lastCondition" column:"last_condition" sql:"disposal_last_condition_enum"`
	Method           DisposalMethod        `json:"method" column:"method" sql:"disposal_method_enum"`
	Notes            string                `json:"notes" column:"notes" sql:"VARCHAR"`
	Manager          string                `json:"manager"   column:"manager" sql:"VARCHAR"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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
