package rbac

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/middleware"

	"github.com/gofiber/fiber/v3"
)

const approverKey = enum.SystemRoleKeyManageHardware

// runApprover stubs iam with `body`/`status`, runs the rules, and returns the gate status and iam call count.
func runApprover(t *testing.T, roles middleware.UserAccessCache, status int, body string, rules ...Rule) (int, int32) {
	t.Helper()
	var calls int32
	iam := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		if r.URL.Path != "/internal/users/7/approver-scope" || r.URL.Query().Get("roleKey") != string(approverKey) ||
			r.Header.Get("X-Service-Token") != "tok" || r.Header.Get("X-Client-ID") != "acme" {
			t.Errorf("bad iam request: %s %v", r.URL, r.Header)
		}
		w.WriteHeader(status)
		fmt.Fprint(w, body)
	}))
	defer iam.Close()

	app := fiber.New()
	app.Use(func(c fiber.Ctx) error {
		c.Locals(middleware.UserRolesKey, roles)
		c.Locals(middleware.IAMEndpointKey, middleware.IAMEndpoint{
			Client: &http.Client{Timeout: 2 * time.Second}, BaseURL: iam.URL, ServiceToken: "tok",
		})
		return c.Next()
	})
	app.Get("/", Require(rules...), func(c fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-User-ID", "7")
	req.Header.Set("X-Client-ID", "acme")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	return resp.StatusCode, atomic.LoadInt32(&calls)
}

func scope(g, a string, global bool) string {
	return fmt.Sprintf(`{"data":{"groupManagerGroupIds":[%s],"additionalApproverGroupIds":[%s],"globalAdditionalApprover":%t}}`, g, a, global)
}

func TestAdminConsoleApprover(t *testing.T) {
	var none middleware.UserAccessCache
	tests := []struct {
		name     string
		withMgr  bool
		body     string
		wantCode int
	}{
		{"global approver", false, scope("", "", true), 200},
		{"group approver", false, scope("", `"g1"`, false), 200},
		{"manager without flag", false, scope(`"g1"`, "", false), 403},
		{"manager with flag", true, scope(`"g1"`, "", false), 200},
		{"none", true, scope("", "", false), 403},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, calls := runApprover(t, none, 200, tt.body, AdminConsoleApprover(approverKey, tt.withMgr))
			if code != tt.wantCode || calls != 1 {
				t.Errorf("code=%d calls=%d, want %d/1", code, calls, tt.wantCode)
			}
		})
	}
}

func TestAdminConsoleApprover_TotalControlSkipsIAM(t *testing.T) {
	var roles middleware.UserAccessCache
	roles.ModuleAccess = map[enum.ModuleAccess]middleware.ModuleFlags{enum.ModuleAccessAdminConsole: {IsTotalControl: true}}
	code, calls := runApprover(t, roles, 500, "", AdminConsoleApprover(approverKey, true))
	if code != 200 || calls != 0 {
		t.Errorf("code=%d calls=%d, want 200/0", code, calls)
	}
}

func TestAdminConsoleApprover_IAMErrorIs500(t *testing.T) {
	code, _ := runApprover(t, middleware.UserAccessCache{}, 500, "boom", AdminConsoleApprover(approverKey, true))
	if code != 500 {
		t.Errorf("code=%d, want 500", code)
	}
}

func TestAdminConsoleApprover_LocalsCacheSingleCall(t *testing.T) {
	code, calls := runApprover(t, middleware.UserAccessCache{}, 200, scope("", "", false),
		Any(AdminConsoleApprover(approverKey, false), AdminConsoleApprover(approverKey, true)))
	if code != 403 || calls != 1 {
		t.Errorf("code=%d calls=%d, want 403/1", code, calls)
	}
}
