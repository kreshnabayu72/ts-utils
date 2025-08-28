package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	fixed_asset_migration "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model/migration"
)

type UserSystemRole struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	IsActive         bool `json:"isActive" column:"is_active" sql:"BOOLEAN DEFAULT TRUE"`
	PermissionView   bool `json:"permissionView" column:"permission_view" sql:"BOOLEAN DEFAULT TRUE"`
	PermissionCreate bool `json:"permissionCreate" column:"permission_create" sql:"BOOLEAN DEFAULT TRUE"`
	PermissionUpdate bool `json:"permissionUpdate" column:"permission_update" sql:"BOOLEAN DEFAULT TRUE"`
	PermissionDelete bool `json:"permissionDelete" column:"permission_delete" sql:"BOOLEAN DEFAULT TRUE"`

	GroupID *int `json:"groupId" column:"group_id" sql:"BIGINT"`
	UserID  int  `json:"userId" column:"user_id" sql:"BIGINT NOT NULL"`

	RoleType      SystemRoleType      `json:"roleType" column:"role_type" sql:"system_role_type_enum NOT NULL"`
	RoleName      SystemRoleName      `json:"roleName" column:"role_name" sql:"system_role_name_enum NOT NULL"`
	RoleAttribute SystemRoleAttribute `json:"roleAttribute" column:"role_attribute" sql:"system_role_attribute_enum NOT NULL"`

	CreatedAt *time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NULL"`
	UpdatedAt *time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NULL"`
}

func (u *UserSystemRole) TableName() string {
	return "user_system_roles"
}

func (u *UserSystemRole) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(u)

	builder.AddForeignKey(
		"fk_user_system_roles_user_id",
		"user_id",
		"users",
		"id",
		"CASCADE",
	)

	builder.AddForeignKey(
		"fk_user_system_roles_group_id",
		"group_id",
		"groups",
		"id",
		"CASCADE",
	)

	builder.AddIndex(
		"user_system_roles_unique_index",
		"group_id,user_id,role_type",
		true,
		"",
	)

	// ─────── Insert default data ───────
	// NOTE: reading from sql file has been tried and doesnt work in server, so use string instead
	builder.SetPostQuery(fixed_asset_migration.UserSystemRolePostQuery)
}
