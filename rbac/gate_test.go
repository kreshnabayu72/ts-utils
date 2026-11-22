package rbac

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/middleware"

	"github.com/gofiber/fiber/v3"
)

const borrowTx = enum.FixedAssetTransactionRoleTransactionKeyBorrowRole

// runGate mounts the gate behind a stub that injects roles into Locals, and reports the status the
// request ended with plus whether the handler was reached.
func runGate(t *testing.T, roles *middleware.UserAccessCache, rules ...Rule) (int, bool) {
	t.Helper()

	reached := false
	app := fiber.New()
	app.Use(func(c fiber.Ctx) error {
		if roles != nil {
			c.Locals(middleware.UserRolesKey, *roles)
		}
		return c.Next()
	})
	app.Get("/", Require(rules...), func(c fiber.Ctx) error {
		reached = true
		return c.SendStatus(fiber.StatusOK)
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	return resp.StatusCode, reached
}

func rolesWithTx(types map[enum.TransactionRolesType][]string) *middleware.UserAccessCache {
	var r middleware.UserAccessCache
	r.FixedAssetTransactionRole = map[enum.FixedAssetTransactionRoleTransactionKey]map[enum.TransactionRolesType][]string{
		borrowTx: types,
	}
	return &r
}

func TestRequire_FailsClosedWithoutSnapshot(t *testing.T) {
	// LoadUserRoles not mounted / iam unreachable: no snapshot in Locals.
	status, reached := runGate(t, nil, FixedAssetTransactionRole(borrowTx, enum.TransactionRolesTypeManager))
	if status != fiber.StatusUnauthorized {
		t.Errorf("status = %d, want 401 — a gate with no snapshot must fail closed", status)
	}
	if reached {
		t.Error("handler ran without an RBAC snapshot")
	}
}

func TestFixedAssetTransactionRole(t *testing.T) {
	manager := rolesWithTx(map[enum.TransactionRolesType][]string{
		enum.TransactionRolesTypeManager: {"7"},
	})
	status, reached := runGate(t, manager, FixedAssetTransactionRole(borrowTx, enum.TransactionRolesTypeManager))
	if status != fiber.StatusOK || !reached {
		t.Errorf("manager should pass: status=%d reached=%v", status, reached)
	}

	// Holds Monitoring, route wants Manager or Staff -> rejected.
	monitoring := rolesWithTx(map[enum.TransactionRolesType][]string{
		enum.TransactionRolesTypeMonitoring: {"7"},
	})
	status, reached = runGate(t, monitoring,
		FixedAssetTransactionRole(borrowTx, enum.TransactionRolesTypeManager, enum.TransactionRolesTypeStaff))
	if status != fiber.StatusForbidden || reached {
		t.Errorf("monitoring must not pass a Manager/Staff gate: status=%d reached=%v", status, reached)
	}

	// The route accepting several types accepts Monitoring too.
	status, _ = runGate(t, monitoring,
		FixedAssetTransactionRole(borrowTx, enum.TransactionRolesTypeManager, enum.TransactionRolesTypeMonitoring))
	if status != fiber.StatusOK {
		t.Errorf("multi-type gate should accept monitoring: status=%d", status)
	}

	// An empty group list is not a grant.
	empty := rolesWithTx(map[enum.TransactionRolesType][]string{
		enum.TransactionRolesTypeManager: {},
	})
	status, _ = runGate(t, empty, FixedAssetTransactionRole(borrowTx, enum.TransactionRolesTypeManager))
	if status != fiber.StatusForbidden {
		t.Errorf("empty group list must not grant: status=%d", status)
	}
}

func TestTransactionRole_TotalControlPasses(t *testing.T) {
	var tc middleware.UserAccessCache
	tc.ModuleAccess = map[enum.ModuleAccess]middleware.ModuleFlags{
		enum.ModuleAccessFixedAsset: {IsTotalControl: true},
	}
	status, _ := runGate(t, &tc, FixedAssetTransactionRole(borrowTx, enum.TransactionRolesTypeManager))
	if status != fiber.StatusOK {
		t.Errorf("total control should pass any transaction gate: status=%d", status)
	}
}

// Read-only must NOT satisfy a transaction-role gate on its own — read-only is not a licence to
// write. It passes only where the route explicitly composes SeeAll.
func TestSeeAll_ReadOnlyOnlyPassesWhenComposed(t *testing.T) {
	var ro middleware.UserAccessCache
	ro.ModuleAccess = map[enum.ModuleAccess]middleware.ModuleFlags{
		enum.ModuleAccessFixedAsset: {IsReadOnly: true},
	}

	status, _ := runGate(t, &ro, FixedAssetTransactionRole(borrowTx, enum.TransactionRolesTypeManager))
	if status != fiber.StatusForbidden {
		t.Errorf("read-only must not pass a bare transaction gate: status=%d", status)
	}

	status, _ = runGate(t, &ro, Any(FixedAssetSeeAll(), FixedAssetTransactionRole(borrowTx, enum.TransactionRolesTypeManager)))
	if status != fiber.StatusOK {
		t.Errorf("read-only should pass a read route composed with SeeAll: status=%d", status)
	}
}

// Group visibility is checked in the SERVICE (Scope.CanSeeGroup), not by a route gate: it depends on
// the request body, and binding assets to their real group needs the asset owner's data. See the
// note in gate.go for why no body-reading rule lives here.

// ModuleAccess rules — MBAC (Module-Based Access Control)

func TestModuleAccess(t *testing.T) {
	// User with total control on fixedAsset should pass ModuleAccess(fixedAsset).
	var tc middleware.UserAccessCache
	tc.ModuleAccess = map[enum.ModuleAccess]middleware.ModuleFlags{
		enum.ModuleAccessFixedAsset: {IsTotalControl: true},
	}
	status, reached := runGate(t, &tc, ModuleAccess(enum.ModuleAccessFixedAsset))
	if status != fiber.StatusOK || !reached {
		t.Errorf("total control should pass ModuleAccess: status=%d reached=%v", status, reached)
	}

	// User with read-only on fixedAsset should also pass ModuleAccess(fixedAsset).
	var ro middleware.UserAccessCache
	ro.ModuleAccess = map[enum.ModuleAccess]middleware.ModuleFlags{
		enum.ModuleAccessFixedAsset: {IsReadOnly: true},
	}
	status, reached = runGate(t, &ro, ModuleAccess(enum.ModuleAccessFixedAsset))
	if status != fiber.StatusOK || !reached {
		t.Errorf("read-only should pass ModuleAccess: status=%d reached=%v", status, reached)
	}

	// User with no access to fixedAsset must not pass.
	var empty middleware.UserAccessCache
	status, reached = runGate(t, &empty, ModuleAccess(enum.ModuleAccessFixedAsset))
	if status != fiber.StatusForbidden || reached {
		t.Errorf("no module access must be refused: status=%d reached=%v", status, reached)
	}
}

// GlobalSettings module access — user creation gate example.

func TestModuleAccess_GlobalSettings(t *testing.T) {
	var tc middleware.UserAccessCache
	tc.ModuleAccess = map[enum.ModuleAccess]middleware.ModuleFlags{
		enum.ModuleAccessGlobalSettings: {IsTotalControl: true},
	}
	status, reached := runGate(t, &tc, ModuleAccess(enum.ModuleAccessGlobalSettings))
	if status != fiber.StatusOK || !reached {
		t.Errorf("global-settings total control should pass: status=%d reached=%v", status, reached)
	}

	// User without global-settings access must not pass.
	var onlyFixed middleware.UserAccessCache
	onlyFixed.ModuleAccess = map[enum.ModuleAccess]middleware.ModuleFlags{
		enum.ModuleAccessFixedAsset: {IsTotalControl: true},
	}
	status, reached = runGate(t, &onlyFixed, ModuleAccess(enum.ModuleAccessGlobalSettings))
	if status != fiber.StatusForbidden || reached {
		t.Errorf("fixed-asset access must not satisfy global-settings gate: status=%d reached=%v", status, reached)
	}
}

func TestRequire_DeniedBodyCarriesPermissionCode(t *testing.T) {
	roles := rolesWithTx(map[enum.TransactionRolesType][]string{enum.TransactionRolesTypeMonitoring: {"7"}})
	app := fiber.New()
	app.Use(func(c fiber.Ctx) error {
		c.Locals(middleware.UserRolesKey, *roles)
		return c.Next()
	})
	app.Get("/", Require(FixedAssetTransactionRole(borrowTx, enum.TransactionRolesTypeManager)), func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	defer resp.Body.Close()
	var body struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if resp.StatusCode != fiber.StatusForbidden || body.Message != "Forbidden" || body.Error != middleware.ErrCodePermissionDenied {
		t.Errorf("got %d %q/%q, want 403 Forbidden/%s", resp.StatusCode, body.Message, body.Error, middleware.ErrCodePermissionDenied)
	}
}
