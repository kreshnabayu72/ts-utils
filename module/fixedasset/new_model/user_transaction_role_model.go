package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	fixed_asset_migration "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model/migration"
)

type UserTransactionRole struct {
	ID            int                      `json:"_id"           column:"user_transaction_roles.id"             sql:"BIGINT PRIMARY KEY"`
	IsActive      bool                     `json:"isActive"      column:"user_transaction_roles.is_active"      sql:"BOOLEAN DEFAULT TRUE"`
	UserID        int                      `json:"userId"        column:"user_transaction_roles.user_id"        sql:"BIGINT NOT NULL"`
	RoleAttribute TransactionRoleAttribute `json:"roleAttribute" column:"user_transaction_roles.role_attribute" sql:"transaction_role_attribute_enum NOT NULL"`
	RoleName      TransactionRoleName      `json:"roleName"      column:"user_transaction_roles.role_name"      sql:"transaction_role_name_enum NOT NULL"`
	RoleType      TransactionRoleType      `json:"roleType"      column:"user_transaction_roles.role_type"      sql:"transaction_role_type_enum NOT NULL"`
	GroupID       int                      `json:"groupId"       column:"user_transaction_roles.group_id"       sql:"BIGINT NOT NULL"`
	ApprovalLevel *int                     `json:"approvalLevel" column:"user_transaction_roles.approval_level" sql:"INTEGER"`
	ApprovalType  ApprovalType             `json:"approvalType"  column:"user_transaction_roles.approval_type"  sql:"approval_type_enum NOT NULL"`
	CreatedAt     *time.Time               `json:"createdAt"     column:"user_transaction_roles.created_at"     sql:"TIMESTAMPTZ NULL"`
	UpdatedAt     *time.Time               `json:"updatedAt"     column:"user_transaction_roles.updated_at"     sql:"TIMESTAMPTZ NULL"`
}

func (u *UserTransactionRole) TableName() string {
	return "user_transaction_roles"
}

func (u *UserTransactionRole) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(u)

	builder.AddForeignKey(
		"fk_user_transaction_roles_user_id",
		"user_id",
		"users",
		"id",
		"CASCADE",
	)

	builder.AddForeignKey(
		"fk_user_transaction_roles_group_id",
		"group_id",
		"groups",
		"id",
		"CASCADE",
	)

	builder.AddIndex(
		"user_transaction_roles_unique_index",
		"group_id,user_id,role_type,role_attribute",
		true,
		"",
	)

	// ─────── Insert default data ───────
	// NOTE: reading from sql file has been tried and doesnt work in server, so use string instead
	builder.SetPostQuery(fixed_asset_migration.UserTransactionRolePostQuery)
}
