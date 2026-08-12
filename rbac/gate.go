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
//	borrowings.Post("/list", rbac.Require(rbac.TransactionRole(tx, Manager, Monitoring)), h.List)
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
				return entity.Forbidden("You do not have permission to do this").SendResponse(c)
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

// SystemRole passes when the caller holds `verb` on ANY of `keys`. A route may name several — V1
// took a slice of attributes for the same reason. Total control passes everything.
func SystemRole(verb middleware.PermissionVerb, keys ...enum.SystemRoleKey) Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		for _, key := range keys {
			if roles.HasSystemPermission(key, verb) {
				return true, nil
			}
		}
		return false, nil
	}
}

// TransactionRole passes when the caller holds ANY of `types` on `transaction` in at least one
// group (Manager / Monitoring / Staff — a route may accept several). Total control passes.
//
// Approver is deliberately NOT satisfiable here: it is absent from the snapshot on purpose
// (ADR 005), because approving is authorized by the approval ROW that names you, not by a role.
// Gate approver routes with Assignment — that is what survives a role revocation.
func TransactionRole(
	transaction enum.FixedAssetTransactionRoleTransactionKey,
	types ...enum.TransactionRolesType,
) Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		if roles.IsTotalControl() {
			return true, nil
		}
		byType, held := roles.TransactionRole[transaction]
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

// SeeAll passes for a total-control or read-only caller — the shortcut that lets someone see every
// group without holding a per-group role. Compose it with Any when a READ route must stay open to
// read-only users:
//
//	rbac.Require(rbac.Any(rbac.SeeAll(), rbac.TransactionRole(tx, Manager, Monitoring)))
//
// Deliberately separate from TransactionRole: read-only is not a licence to WRITE, so a write route
// must not pick it up implicitly.
func SeeAll() Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		return roles.IsTotalControl() || roles.IsReadOnly(), nil
	}
}

// TotalControl passes for a total-control admin only — stricter than SeeAll (which also admits
// read-only). Use it where holding the grant, not just seeing everything, is the requirement.
func TotalControl() Rule {
	return func(_ fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		return roles.IsTotalControl(), nil
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
