package new_model

import "time"

type AssetName struct {
	ID   int    `json:"id"         column:"id"`
	Name string `json:"name"       column:"name"`
	Code string `json:"code"       column:"code"`

	// ─────── Add on ───────
	AddOnRepairTicketing bool `json:"addOnRepairTicketing" column:"add_on_repair_ticketing"`
	AddOnMaintenance     bool `json:"addOnMaintenance"     column:"add_on_maintenance"`
	AddOnTracking        bool `json:"addOnTracking"        column:"add_on_tracking"`
	AddOnAssetControl    bool `json:"addOnAssetControl"    column:"add_on_asset_control"`
	AddOnAudit           bool `json:"addOnAudit"           column:"add_on_audit"`

	// ─────── References ───────
	CategoryId          int `json:"categoryId"          column:"category_id"`
	MeasurementId       int `json:"measurementId"       column:"measurement_id"`
	DepreciationGroupId int `json:"depreciationGroupId" column:"depreciation_group_id"`
	AccountCodeId       int `json:"accountCodeId"       column:"account_code_id"`

	// ─────── Transactions policy ───────
	TransactionsBorrow      bool `json:"transactionsBorrow"      column:"transactions_borrow"`
	TransactionsAssign      bool `json:"transactionsAssign"      column:"transactions_assign"`
	TransactionsTransfer    bool `json:"transactionsTransfer"    column:"transactions_transfer"`
	TransactionsSelfService bool `json:"transactionsSelfService" column:"transactions_self_service"`

	TagType TagType `json:"tagType" column:"tag_type"`

	IsDeleted            bool       `json:"isDeleted"            column:"is_deleted"`
	DeleteFromDatabaseAt *time.Time `json:"deleteFromDatabaseAt" column:"delete_from_database_at"`
	CreatedAt            time.Time  `json:"createdAt"            column:"created_at"`
	UpdatedAt            time.Time  `json:"updatedAt"            column:"updated_at"`
}
