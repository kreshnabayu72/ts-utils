package model

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

type FixedAssetUserTransactionRole struct {
	ID             int                                          `json:"_id"            column:"fixed_asset_user_transaction_roles.id"`
	UserId         int                                          `json:"userId"         column:"fixed_asset_user_transaction_roles.user_id"`
	IsActive       bool                                         `json:"isActive"       column:"fixed_asset_user_transaction_roles.is_active"`
	Type           enum.TransactionRolesType                    `json:"type"           column:"fixed_asset_user_transaction_roles.type"`
	Transaction    enum.FixedAssetTransactionRoleTransaction    `json:"transaction"    column:"fixed_asset_user_transaction_roles.transaction"`
	TransactionKey enum.FixedAssetTransactionRoleTransactionKey `json:"transactionKey" column:"fixed_asset_user_transaction_roles.transaction_key"`
	GroupId        int                                          `json:"groupId"        column:"fixed_asset_user_transaction_roles.group_id"`
	ApprovalLevel  int                                          `json:"approvalLevel"  column:"fixed_asset_user_transaction_roles.approval_level"`
	ApprovalType   enum.ApprovalType                            `json:"approvalType"   column:"fixed_asset_user_transaction_roles.approval_type"`
	CreatedAt      time.Time                                    `json:"createdAt"      column:"fixed_asset_user_transaction_roles.created_at"`
	UpdatedAt      time.Time                                    `json:"updatedAt"      column:"fixed_asset_user_transaction_roles.updated_at"`
}
func (r *FixedAssetUserTransactionRole) GetColumns() []string {
	return []string{
		"id",
		"user_id",
		"is_active",
		"type",
		"transaction",
		"transaction_key",
		"group_id",
		"approval_level",
		"approval_type",
		"created_at",
		"updated_at",
	}
}
