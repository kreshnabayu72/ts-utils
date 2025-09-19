package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type ServiceActivity struct {
	ID int `json:"_id"                  column:"service_activities.id"                     sql:"BIGINT PRIMARY KEY"`

	AssetBrandName        string `json:"assetBrandName"         column:"service_activities.asset_brand_name"         sql:"VARCHAR NOT NULL"`
	AssetNameWithSequence string `json:"assetNameWithSequence"         column:"service_activities.asset_name_with_sequence"         sql:"VARCHAR NOT NULL"`
	AssetModelName        string `json:"assetModelName"         column:"service_activities.asset_model_name"         sql:"VARCHAR NOT NULL"`
	AssetGroupName        string `json:"assetGroupName"         column:"service_activities.asset_group_name"         sql:"VARCHAR NOT NULL"`
	Status                string `json:"status"               column:"service_activities.status"                 sql:"service_activity_status_enum"`

	// Repair / Maintenance
	Type string `json:"type"      column:"service_activities.type"       sql:"VARCHAR NOT NULL"`

	// Refer to repair, repair history / routine,routine history
	TransactionID *string `json:"transactionId"        column:"service_activities.transaction_id"         sql:"VARCHAR"`

	ServiceCenterID   *int   `json:"serviceCenterId"      column:"service_activities.service_center_id"      sql:"BIGINT"`
	ServiceCenterName string `json:"serviceCenterName"      column:"service_activities.service_center_name"      sql:"VARCHAR"`

	DownTime        *string `json:"downTime"             column:"service_activities.down_time"              sql:"VARCHAR"`
	DownTimeSeconds *int    `json:"downTimeSeconds"      column:"service_activities.down_time_seconds"      sql:"BIGINT"`

	DeletedAt *time.Time `json:"deletedAt"            column:"service_activities.deleted_at"             sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt time.Time  `json:"createdAt"            column:"service_activities.created_at"             sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time  `json:"updatedAt"            column:"service_activities.updated_at"             sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (s *ServiceActivity) TableName() string {
	return "service_activities"
}

func (s *ServiceActivity) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(s)
}
