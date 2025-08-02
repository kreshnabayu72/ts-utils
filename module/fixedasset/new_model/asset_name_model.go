package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AssetName struct {
	ID   int    `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	Name string `json:"name" column:"name" sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(name) > 0)"`
	Code string `json:"code" column:"code" sql:"VARCHAR NULL DEFAULT NULL"`

	// ─────── Add on ───────
	AddOnRepairTicketing bool `json:"addonRepairTicketing" column:"addon_repair_ticketing" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnMaintenance     bool `json:"addonMaintenance" column:"addon_maintenance" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnTracking        bool `json:"addonTracking" column:"addon_tracking" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnAssetControl    bool `json:"addonAssetControl" column:"addon_asset_control" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnAudit           bool `json:"addonAudit" column:"addon_audit" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`

	// ─────── References ───────
	CategoryId          int  `json:"categoryId" column:"category_id" sql:"BIGINT"`
	MeasurementId       int  `json:"measurementId" column:"measurement_id" sql:"BIGINT"`
	DepreciationGroupId *int `json:"depreciationGroupId" column:"depreciation_group_id" sql:"BIGINT NULL"`
	AccountCodeId       *int `json:"accountCodeId" column:"account_code_id" sql:"BIGINT NULL"`

	// ─────── Transactions policy ───────
	TransactionsBorrow      bool `json:"transactionsBorrow" column:"transactions_borrow" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TransactionsAssign      bool `json:"transactionsAssign" column:"transactions_assign" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TransactionsTransfer    bool `json:"transactionsTransfer" column:"transactions_transfer" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TransactionsSelfService bool `json:"transactionsSelfService" column:"transactions_self_service" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`

	TagType TagType `json:"tagType" column:"tag_type" sql:"tag_type_enum DEFAULT 'Non TAG'"`

	IsDeleted bool       `json:"isDeleted" column:"is_deleted" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeletedAt *time.Time `json:"deletedAt" column:"deleted_at" sql:"TIMESTAMP NULL DEFAULT NULL"`
	CreatedAt time.Time  `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time  `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
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
