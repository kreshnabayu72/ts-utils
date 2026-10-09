// Package rbac provides the route-gate DSL: Rule predicates composed
// with Any (OR) and evaluated by Require (AND). Rules are pure
// predicates on the RBAC snapshot; Require is the middleware that
// evaluates them.
package rbac

import (
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/middleware"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/entity"

	"github.com/gofiber/fiber/v3"
)

// Rule reports whether the caller may proceed. It gets the request (for path/body params) and the
// caller's RBAC snapshot.
type Rule func(c fiber.Ctx, roles middleware.UserAccessCache) (bool, error)

// Require gates a route on rules. ALL must pass (AND); use Any for OR.
//
//	borrowings.Post("/list", rbac.Require(rbac.FixedAssetTransactionRole(tx, Manager, Monitoring)), h.List)
func Require(rules ...Rule) fiber.Handler {
	return func(c fiber.Ctx) error {
		roles, ok := middleware.GetUserRoles(c)
		if !ok {
			return entity.Unauthorized("You are not authenticated!").SendResponse(c)
		}
		for _, rule := range rules {
			passed, err := rule(c, roles)
			if err != nil {
				return entity.ToHttpError(err).SendResponse(c)
			}
			if !passed {
				return entity.Forbidden("Forbidden").SendResponse(c)
			}
		}
		return c.Next()
	}
}

// Any passes when ANY of the rules passes — the OR-combinator (V1's CombineOrRestriction).
//
// The detail endpoint is what needs it: manager OR borrower OR named approver may open one
// transaction, and no single rule expresses that.
func Any(rules ...Rule) Rule {
	return func(c fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		for _, rule := range rules {
			passed, err := rule(c, roles)
			if err != nil {
				return false, err
			}
			if passed {
				return true, nil
			}
		}
		return false, nil
	}
}

// AdminConsoleSystemRole passes when the caller holds `verb` on ANY of `keys` — either globally or
// in at least one group. A route may name several — V1 took a slice of attributes for the same
// reason. Total control passes everything.
func AdminConsoleSystemRole(verb middleware.PermissionVerb, keys ...enum.SystemRoleKey) Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		for _, key := range keys {
			if roles.HasSystemPermissionAdminConsole(key, verb) {
				return true, nil
			}
		}
		return false, nil
	}
}

// FixedAssetSystemRole passes when the caller holds `verb` on ANY of `keys` — either globally or
// in at least one group. A route may name several — V1 took a slice of attributes for the same
// reason. Total control passes everything.
func FixedAssetSystemRole(verb middleware.PermissionVerb, keys ...enum.SystemRoleKey) Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		for _, key := range keys {
			if roles.HasSystemPermissionFixedAsset(key, verb) {
				return true, nil
			}
		}
		return false, nil
	}
}

// GlobalSettingsSystemRole passes when the caller holds `verb` on ANY of `keys` in the
// global-settings module. A route may name several. Unlike FixedAssetSystemRole, this reads
// GlobalSettingsSystemRole, not FixedAssetSystemRole. Global-settings total control passes
// everything.
func GlobalSettingsSystemRole(verb middleware.PermissionVerb, keys ...enum.SystemRoleKey) Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		for _, key := range keys {
			if roles.HasGlobalSettingsPermission(key, verb) {
				return true, nil
			}
		}
		return false, nil
	}
}

// FixedAssetTransactionRole passes when the caller holds ANY of `types` on `transaction` in at least one
// group (Manager / Monitoring / Staff — a route may accept several). Total control passes.
//
// Approver is deliberately NOT satisfiable here: it is absent from the snapshot on purpose
// (ADR 005), because approving is authorized by the approval ROW that names you, not by a role.
// Gate approver routes with Assignment — that is what survives a role revocation.
func FixedAssetTransactionRole(
	transaction enum.FixedAssetTransactionRoleTransactionKey,
	types ...enum.TransactionRolesType,
) Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		if roles.IsTotalControl(enum.ModuleAccessFixedAsset) {
			return true, nil
		}
		byType, held := roles.FixedAssetTransactionRole[transaction]
		if !held {
			return false, nil
		}
		for _, t := range types {
			if len(byType[t]) > 0 {
				return true, nil
			}
		}
		return false, nil
	}
}

// AdminConsoleSeeAll passes for a total-control or read-only caller — the shortcut that lets someone see every
// group without holding a per-group role. Compose it with Any when a READ route must stay open to
// read-only users:
//
//	rbac.Require(rbac.Any(rbac.AdminConsoleSeeAll(), rbac.AdminConsoleTransactionRole(tx, Manager, Monitoring)))
//
// Deliberately separate from AdminConsoleTransactionRole: read-only is not a licence to WRITE, so a write route
// must not pick it up implicitly.
func AdminConsoleSeeAll() Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		return roles.IsTotalControl(enum.ModuleAccessAdminConsole) || roles.IsReadOnly(enum.ModuleAccessAdminConsole), nil
	}
}

// AdminConsoleTotalControl passes for a total-control admin only — stricter than AdminConsoleSeeAll
// (which also admits read-only). Use it where holding the grant, not just seeing everything, is the requirement.
func AdminConsoleTotalControl() Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		return roles.IsTotalControl(enum.ModuleAccessAdminConsole), nil
	}
}

// FixedAssetSeeAll passes for a total-control or read-only caller — the shortcut that lets someone see every
// group without holding a per-group role. Compose it with Any when a READ route must stay open to
// read-only users:
//
//	rbac.Require(rbac.Any(rbac.FixedAssetSeeAll(), rbac.FixedAssetTransactionRole(tx, Manager, Monitoring)))
//
// Deliberately separate from FixedAssetTransactionRole: read-only is not a licence to WRITE, so a write route
// must not pick it up implicitly.
func FixedAssetSeeAll() Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		return roles.IsTotalControl(enum.ModuleAccessFixedAsset) || roles.IsReadOnly(enum.ModuleAccessFixedAsset), nil
	}
}

// FixedAssetTotalControl passes for a total-control admin only — stricter than FixedAssetSeeAll
// (which also admits read-only). Use it where holding the grant, not just seeing everything, is the requirement.
func FixedAssetTotalControl() Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		return roles.IsTotalControl(enum.ModuleAccessFixedAsset), nil
	}
}

// GlobalSettingsSeeAll passes for a total-control or read-only caller of the global-settings
// module — the read-side counterpart to FixedAssetSeeAll.
func GlobalSettingsSeeAll() Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		return roles.IsTotalControl(enum.ModuleAccessGlobalSettings) || roles.IsReadOnly(enum.ModuleAccessGlobalSettings), nil
	}
}

// GlobalSettingsTotalControl passes for a global-settings total-control admin only — stricter than
// GlobalSettingsSeeAll, which also admits read-only.
func GlobalSettingsTotalControl() Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		return roles.IsTotalControl(enum.ModuleAccessGlobalSettings), nil
	}
}

// ModuleAccess passes when the caller has any access (total control or read-only) to the given
// module. Use it for module-level gates where the module itself must be reachable, regardless of
// the specific permission level:
//
//	rbac.Require(rbac.ModuleAccess(enum.ModuleAccessGlobalSettings))
func ModuleAccess(moduleKey enum.ModuleAccess) Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		_, ok := roles.ModuleAccess[moduleKey]
		return ok, nil
	}
}
