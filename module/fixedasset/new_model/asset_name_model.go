package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AssetName struct {
	ID                      int        `json:"_id"                     column:"asset_names.id"                        sql:"BIGINT PRIMARY KEY"`
	Name                    string     `json:"name"                    column:"asset_names.name"                      sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(name) > 0)"`
	Code                    *string    `json:"code"                    column:"asset_names.code"                      sql:"VARCHAR NULL DEFAULT NULL"`
	AddOnRepairTicketing    bool       `json:"addonRepairTicketing"    column:"asset_names.addon_repair_ticketing"    sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnMaintenance        bool       `json:"addonMaintenance"        column:"asset_names.addon_maintenance"         sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnTracking           bool       `json:"addonTracking"           column:"asset_names.addon_tracking"            sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnAssetControl       bool       `json:"addonAssetControl"       column:"asset_names.addon_asset_control"       sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnAudit              bool       `json:"addonAudit"              column:"asset_names.addon_audit"               sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	CategoryId              int        `json:"categoryId"              column:"asset_names.category_id"               sql:"BIGINT"`
	MeasurementId           int        `json:"measurementId"           column:"asset_names.measurement_id"            sql:"BIGINT"`
	DepreciationGroupId     *int       `json:"depreciationGroupId"     column:"asset_names.depreciation_group_id"     sql:"BIGINT NULL"`
	AccountCodeId           *int       `json:"accountCodeId"           column:"asset_names.account_code_id"           sql:"BIGINT NULL"`
	TransactionsBorrow      bool       `json:"transactionsBorrow"      column:"asset_names.transactions_borrow"       sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TransactionsAssign      bool       `json:"transactionsAssign"      column:"asset_names.transactions_assign"       sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TransactionsTransfer    bool       `json:"transactionsTransfer"    column:"asset_names.transactions_transfer"     sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TransactionsSelfService bool       `json:"transactionsSelfService" column:"asset_names.transactions_self_service" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TagType                 TagType    `json:"tagType"                 column:"asset_names.tag_type"                  sql:"tag_type_enum DEFAULT 'Non TAG'"`
	IsDeleted               bool       `json:"isDeleted"               column:"asset_names.is_deleted"                sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeletedAt               *time.Time `json:"deletedAt"               column:"asset_names.deleted_at"                sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt               time.Time  `json:"createdAt"               column:"asset_names.created_at"                sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt               time.Time  `json:"updatedAt"               column:"asset_names.updated_at"                sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *AssetName) TableName() string {
	return "asset_names"
}

func (m *AssetName) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("asset_names_name_uindex", "name", true, "is_deleted = FALSE")
	builder.AddIndex(
		"asset_names_code_uindex",
		"code",
		true,
		"code IS NOT NULL AND is_deleted = FALSE",
	)
}
