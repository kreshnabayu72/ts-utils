package new_model

import "time"

type AssetName struct {
	ID   int    `json:"_id"         column:"id"`
	Name string `json:"name"       column:"name"`
	Code string `json:"code"       column:"code"`

	// ─────── Add on ───────
	AddOnRepairTicketing bool `json:"addonRepairTicketing" column:"addon_repair_ticketing"`
	AddOnMaintenance     bool `json:"addonMaintenance"     column:"addon_maintenance"`
	AddOnTracking        bool `json:"addonTracking"        column:"addon_tracking"`
	AddOnAssetControl    bool `json:"addonAssetControl"    column:"addon_asset_control"`
	AddOnAudit           bool `json:"addonAudit"           column:"addon_audit"`

	// ─────── References ───────
	CategoryId          int  `json:"categoryId"          column:"category_id"`
	MeasurementId       int  `json:"measurementId"       column:"measurement_id"`
	DepreciationGroupId *int `json:"depreciationGroupId" column:"depreciation_group_id"`
	AccountCodeId       *int `json:"accountCodeId"       column:"account_code_id"`

	// ─────── Transactions policy ───────
	TransactionsBorrow      bool `json:"transactionsBorrow"      column:"transactions_borrow"`
	TransactionsAssign      bool `json:"transactionsAssign"      column:"transactions_assign"`
	TransactionsTransfer    bool `json:"transactionsTransfer"    column:"transactions_transfer"`
	TransactionsSelfService bool `json:"transactionsSelfService" column:"transactions_self_service"`

	TagType TagType `json:"tagType" column:"tag_type"`

	IsDeleted bool       `json:"isDeleted"            column:"is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"deleted_at"`
	CreatedAt time.Time  `json:"createdAt"            column:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt"            column:"updated_at"`
}
