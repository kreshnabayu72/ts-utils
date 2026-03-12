package model

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

type FixedAssetGroupApprovalType struct {
	ID             int                                          `json:"_id"            column:"fixed_asset_group_approval_types.id"`
	Level          int                                          `json:"level"          column:"fixed_asset_group_approval_types.level"`
	TransactionKey enum.FixedAssetTransactionRoleTransactionKey `json:"transactionKey" column:"fixed_asset_group_approval_types.transaction_key"`
	GroupId        *int                                         `json:"groupId"        column:"fixed_asset_group_approval_types.group_id"`
	ApprovalType   enum.ApprovalType                            `json:"approvalType"   column:"fixed_asset_group_approval_types.approval_type"`
	IsDeleted      bool                                         `json:"isDeleted"      column:"fixed_asset_group_approval_types.is_deleted"`
	DeletedAt      *time.Time                                   `json:"deletedAt"      column:"fixed_asset_group_approval_types.deleted_at"`
	CreatedAt      time.Time                                    `json:"createdAt"      column:"fixed_asset_group_approval_types.created_at"`
	UpdatedAt      time.Time                                    `json:"updatedAt"      column:"fixed_asset_group_approval_types.updated_at"`
}
func (r *FixedAssetGroupApprovalType) GetColumns() []string {
	return []string{
		"id",
		"level",
		"transaction_key",
		"group_id",
		"approval_type",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}
