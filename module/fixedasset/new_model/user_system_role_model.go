package new_model

import (
	"time"
)

type UserSystemRole struct {
	ID int `json:"_id" column:"id"`

	IsActive         bool `json:"isActive" column:"is_active"`
	UserID           int  `json:"userId" column:"user_id"`
	PermissionView   bool `json:"permissionView" column:"permission_view"`
	PermissionCreate bool `json:"permissionCreate" column:"permission_create"`
	PermissionUpdate bool `json:"permissionUpdate" column:"permission_update"`
	PermissionDelete bool `json:"permissionDelete" column:"permission_delete"`
	GroupID          int  `json:"groupId" column:"group_id"`

	RoleName      SystemRoleName      `json:"roleName" column:"role_name"`           // Enum: system_role_name_enum
	RoleAttribute SystemRoleAttribute `json:"roleAttribute" column:"role_attribute"` // Enum: system_role_attribute_enum
	RoleType      SystemRoleType      `json:"roleType" column:"role_type"`

	CreatedAt *time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" column:"updated_at"`
}
