package model

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"

type GlobalSettingsUserSystemRole struct {
	ID               int                               `json:"_id"              column:"fixed_asset_user_system_roles.id"`
	UserId           int                               `json:"userId"           column:"fixed_asset_user_system_roles.user_id"`
	IsActive         bool                              `json:"isActive"         column:"fixed_asset_user_system_roles.is_active"`
	PermissionView   bool                              `json:"permissionView"   column:"fixed_asset_user_system_roles.permission_view"`
	PermissionCreate bool                              `json:"permissionCreate" column:"fixed_asset_user_system_roles.permission_create"`
	PermissionUpdate bool                              `json:"permissionUpdate" column:"fixed_asset_user_system_roles.permission_update"`
	PermissionDelete bool                              `json:"permissionDelete" column:"fixed_asset_user_system_roles.permission_delete"`
	Name             enum.GlobalSettingsSystemRoleName `json:"name"             column:"fixed_asset_user_system_roles.name"`
	Key              enum.GlobalSettingsSystemRoleKey  `json:"key"              column:"fixed_asset_user_system_roles.key"`
}
