package model

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

// UserSystemRole is one system-role grant, for any product. Fixed Asset and Global
// Settings share this table because they share a grain — one row per user × product ×
// role key × group — against the same tenant users.id. They differ only in which
// columns are populated, and that is a column problem, not a table problem.
//
// Product is the discriminator and is never optional: role keys collide across products
// ("importData", "tag", "userAndRole" exist in more than one), so a query matching on
// Key without also constraining Product reads another product's grants. The unique key
// is (user_id, product, key, <scope fk>).
//
// Scope is an exclusive arc — one nullable FK per group-scoped product, with a DB CHECK
// that at most one is set (Karwin, SQL Antipatterns ch. 7). Fixed Asset is the only
// scoped product today; Global Settings has no groups, so its rows are always category
// Global with every scope FK NULL. Adding a scoped product costs one column here plus
// one term in the CHECK. Once a third one lands, promote the arc to a shared `groups`
// supertype and collapse these back to a single FK.
//
// Product is typed enum.ModuleAccess, the iam-side vocabulary ("Fixed Asset"), NOT
// analytics_reporting_service/model.Product ("fams") — same concept, two value sets.
// Consolidate before a third caller picks a side.
type UserSystemRole struct {
	ID                int                      `json:"_id"               column:"user_system_roles.id"`
	UserId            int                      `json:"userId"            column:"user_system_roles.user_id"`
	Product           enum.ModuleAccess        `json:"product"           column:"user_system_roles.product"`
	Category          enum.SystemRolesCategory `json:"category"          column:"user_system_roles.category"`
	Key               enum.SystemRoleKey       `json:"key"               column:"user_system_roles.key"`
	Name              enum.SystemRoleName      `json:"name"              column:"user_system_roles.name"`
	IsActive          bool                     `json:"isActive"          column:"user_system_roles.is_active"`
	PermissionView    bool                     `json:"permissionView"    column:"user_system_roles.permission_view"`
	PermissionCreate  bool                     `json:"permissionCreate"  column:"user_system_roles.permission_create"`
	PermissionUpdate  bool                     `json:"permissionUpdate"  column:"user_system_roles.permission_update"`
	PermissionDelete  bool                     `json:"permissionDelete"  column:"user_system_roles.permission_delete"`
	FixedAssetGroupId *int                     `json:"fixedAssetGroupId" column:"user_system_roles.fixed_asset_group_id"`
	ModifiedBy        int                      `json:"modifiedBy"            column:"user_system_roles.modified_by"`

	CreatedAt time.Time `json:"createdAt"     column:"user_system_roles.created_at"`
	UpdatedAt time.Time `json:"updatedAt"     column:"user_system_roles.updated_at"`
}
