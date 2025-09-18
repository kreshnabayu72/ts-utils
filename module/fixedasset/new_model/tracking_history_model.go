package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type TrackingHistory struct {
	ID                 int    `json:"_id"             column:"tracking_histories.id"                sql:"BIGINT PRIMARY KEY"`
	Asset              int    `json:"asset"         column:"tracking_histories.asset"          sql:"BIGINT NOT NULL"`
	ReportedByFullName string `json:"reportedByFullName"               column:"tracking_histories.reported_by_full_name"                sql:"VARCHAR"`

	AssetNameWithSequence string `json:"assetNameWithSequence"                column:"tracking_histories.asset_name_with_sequence"                  sql:"VARCHAR"`
	AssetAliasCode        string `json:"assetAliasCode"                column:"tracking_histories.asset_alias_code"                  sql:"VARCHAR"`
	AssetBrand            string `json:"assetBrand"                column:"tracking_histories.asset_brand"                  sql:"VARCHAR"`
	AssetModel            string `json:"assetModel"                column:"tracking_histories.asset_model"                  sql:"VARCHAR"`
	AssetGroupName        string `json:"assetGroup"                column:"tracking_histories.asset_group_name"                  sql:"VARCHAR"`
	AssetGroupId          int    `json:"assetGroupId"                column:"tracking_histories.asset_group_id"                  sql:"BIGINT NOT NULL"`

	DamageImage     string            `json:"damageImage"     column:"tracking_histories.damage_image"      sql:"VARCHAR"`
	PreviousStatus  AssetStatus       `json:"previousStatus"  column:"tracking_histories.previous_status"   sql:"asset_status_enum NOT NULL"`
	TransactionID   string            `json:"transactionId"   column:"tracking_histories.transaction_id"    sql:"VARCHAR NOT NULL"`
	Status          TrackingStatus    `json:"status"          column:"tracking_histories.status"            sql:"tracking_status_enum NOT NULL"`
	Condition       TrackingCondition `json:"condition"       column:"tracking_histories.condition"         sql:"tracking_condition_enum NOT NULL"`
	Notes           string            `json:"notes"           column:"tracking_histories.notes"             sql:"VARCHAR"`
	Address         string            `json:"address"         column:"tracking_histories.address"           sql:"VARCHAR"`
	Coordinate      string            `json:"coordinate"      column:"tracking_histories.coordinate"        sql:"VARCHAR"`
	DownTimeSeconds float64           `json:"downTimeSeconds" column:"tracking_histories.down_time_seconds" sql:"FLOAT"`
	ReportSource    AssetReportSource `json:"reportSource"    column:"tracking_histories.report_source"     sql:"asset_report_source_enum NOT NULL"`
	ApprovalRound   int               `json:"approvalRound"   column:"tracking_histories.approval_round"    sql:"INTEGER NOT NULL DEFAULT 0"`

	FoundAt    time.Time `json:"foundAt" column:"tracking_histories.found_at"`
	FoundBy    int       `json:"foundBy" column:"tracking_histories.found_by_id"`
	FoundGroup int       `json:"found_group" column:"tracking_histories.found_group_id"`

	CreatedAt time.Time `json:"createdAt"       column:"tracking_histories.created_at"        sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt"       column:"tracking_histories.updated_at"        sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *TrackingHistory) TableName() string {
	return "tracking_histories"
}

func (m *TrackingHistory) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)
}
