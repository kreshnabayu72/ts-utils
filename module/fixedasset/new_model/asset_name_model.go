package new_model

import (
	"time"
)

type AssetName struct {
	ID                      int        `json:"_id"                     column:"asset_names.id"`
	Name                    string     `json:"name"                    column:"asset_names.name"`
	Code                    *string    `json:"code"                    column:"asset_names.code"`
	AddOnRepairTicketing    bool       `json:"addonRepairTicketing"    column:"asset_names.addon_repair_ticketing"`
	AddOnMaintenance        bool       `json:"addonMaintenance"        column:"asset_names.addon_maintenance"`
	AddOnTracking           bool       `json:"addonTracking"           column:"asset_names.addon_tracking"`
	AddOnAssetControl       bool       `json:"addonAssetControl"       column:"asset_names.addon_asset_control"`
	AddOnAudit              bool       `json:"addonAudit"              column:"asset_names.addon_audit"`
	CategoryId              int        `json:"categoryId"              column:"asset_names.category_id"`
	MeasurementId           int        `json:"measurementId"           column:"asset_names.measurement_id"`
	DepreciationGroupId     *int       `json:"depreciationGroupId"     column:"asset_names.depreciation_group_id"`
	AccountCodeId           *int       `json:"accountCodeId"           column:"asset_names.account_code_id"`
	TransactionsBorrow      bool       `json:"transactionsBorrow"      column:"asset_names.transactions_borrow"`
	TransactionsAssign      bool       `json:"transactionsAssign"      column:"asset_names.transactions_assign"`
	TransactionsTransfer    bool       `json:"transactionsTransfer"    column:"asset_names.transactions_transfer"`
	TransactionsSelfService bool       `json:"transactionsSelfService" column:"asset_names.transactions_self_service"`
	TagType                 *TagType   `json:"tagType"                 column:"asset_names.tag_type"`
	IsDeleted               bool       `json:"isDeleted"               column:"asset_names.is_deleted"`
	DeletedAt               *time.Time `json:"deletedAt"               column:"asset_names.deleted_at"`
	CreatedAt               time.Time  `json:"createdAt"               column:"asset_names.created_at"`
	UpdatedAt               time.Time  `json:"updatedAt"               column:"asset_names.updated_at"`
}
