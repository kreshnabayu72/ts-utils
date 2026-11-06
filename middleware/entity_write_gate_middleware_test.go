package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/entity"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/service"

	"github.com/gofiber/fiber/v3"
)

// statusCacheFake is a stateful Valkey stand-in: the gate's hit path, warming, and invalidation
// all observe the same value, so a test can prove the cache changed rather than merely called.
type statusCacheFake struct {
	service.ValkeyServiceInterface
	value    string
	getErr   error
	setCalls int
	delCalls []string
}

func (f *statusCacheFake) Get(context.Context, string) (string, error) {
	if f.getErr != nil {
		return "", f.getErr
	}
	return f.value, nil
}

func (f *statusCacheFake) Set(_ context.Context, _ string, value any, _ time.Duration) error {
	f.setCalls++
	f.value = fmt.Sprint(value)
	return nil
}

func (f *statusCacheFake) Del(_ context.Context, keys ...string) error {
	f.delCalls = append(f.delCalls, keys...)
	f.value = ""
	return nil
}

func gateApp(cache service.ValkeyServiceInterface, fetch EntityStatusFetcher) *fiber.App {
	app := fiber.New()
	app.Put("/", requireActiveEntity(cache, fetch), func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	return app
}

func gateRequest(t *testing.T, app *fiber.App) (int, string) {
	code, message, _ := gateRequestErr(t, app)
	return code, message
}

func gateRequestErr(t *testing.T, app *fiber.App) (int, string, string) {
	t.Helper()
	req := httptest.NewRequest(fiber.MethodPut, "/", nil)
	req.Header.Set(constant.HEADER_CLIENT_ID, "acme")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	var body struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}
	if len(raw) > 0 && raw[0] == '{' {
		if err := json.Unmarshal(raw, &body); err != nil {
			t.Fatalf("decode body %q: %v", raw, err)
		}
	}
	return resp.StatusCode, body.Message, body.Error
}

func TestRequireActiveEntity_ActiveWritePasses(t *testing.T) {
	cache := &statusCacheFake{getErr: errors.New("miss")}
	var fetched string
	app := gateApp(cache, func(_ context.Context, clientID string) (string, error) {
		fetched = clientID
		return "Active", nil
	})

	status, _ := gateRequest(t, app)
	if status != fiber.StatusOK {
		t.Fatalf("status = %d, want 200", status)
	}
	if fetched != "acme" {
		t.Errorf("fetch clientID = %q, want acme", fetched)
	}
	if cache.setCalls != 1 || cache.value != "Active" {
		t.Errorf("cache = %q after %d sets, want Active cached once", cache.value, cache.setCalls)
	}
}

func TestRequireActiveEntity_StatusBlocksWrite(t *testing.T) {
	for _, status := range []string{"Suspended", "suspended", "Expired"} {
		t.Run(status, func(t *testing.T) {
			cache := &statusCacheFake{getErr: errors.New("miss")}
			app := gateApp(cache, func(context.Context, string) (string, error) {
				return status, nil
			})

			code, message, errCode := gateRequestErr(t, app)
			if code != fiber.StatusForbidden {
				t.Fatalf("status = %d, want 403", code)
			}
			if message != "Forbidden" || errCode != ErrCodeEntityInactive {
				t.Errorf("body = %q/%q, want Forbidden/%s", message, errCode, ErrCodeEntityInactive)
			}
		})
	}
}

func TestRequireActiveEntity_CacheHitSkipsFetch(t *testing.T) {
	cache := &statusCacheFake{value: "Suspended"}
	app := gateApp(cache, func(context.Context, string) (string, error) {
		t.Error("fetch must not run on a cache hit")
		return "Active", nil
	})

	code, _ := gateRequest(t, app)
	if code != fiber.StatusForbidden {
		t.Fatalf("status = %d, want 403 from the cached status", code)
	}
}

func TestRequireActiveEntity_FailsClosedWhenStatusUnavailable(t *testing.T) {
	cache := &statusCacheFake{getErr: errors.New("miss")}
	app := gateApp(cache, func(context.Context, string) (string, error) {
		return "", errors.New("connection refused")
	})

	code, message := gateRequest(t, app)
	if code != fiber.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503", code)
	}
	if message != "Authorization service unavailable" {
		t.Errorf("message = %q, want Authorization service unavailable", message)
	}
}

func TestRequireActiveEntity_UnknownTenantIsRefused(t *testing.T) {
	cache := &statusCacheFake{getErr: errors.New("miss")}
	app := gateApp(cache, func(context.Context, string) (string, error) {
		return "", entity.NotFound("Client not found")
	})

	code, message := gateRequest(t, app)
	if code != fiber.StatusForbidden {
		t.Fatalf("status = %d, want 403", code)
	}
	if message != "Forbidden" {
		t.Errorf("message = %q, want Forbidden", message)
	}
}

func TestRequireActiveEntity_EmptyStatusFailsClosed(t *testing.T) {
	cache := &statusCacheFake{getErr: errors.New("miss")}
	app := gateApp(cache, func(context.Context, string) (string, error) {
		return "", nil
	})

	code, _ := gateRequest(t, app)
	if code != fiber.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503", code)
	}
	if cache.setCalls != 0 {
		t.Errorf("cache set %d times, want 0 for an empty status", cache.setCalls)
	}
}

func TestInvalidateEntityStatus_DropsCachedStatus(t *testing.T) {
	cache := &statusCacheFake{value: "Suspended"}
	fetchActive := false
	app := gateApp(cache, func(context.Context, string) (string, error) {
		fetchActive = true
		return "Active", nil
	})

	if err := InvalidateEntityStatus(context.Background(), cache, "acme"); err != nil {
		t.Fatalf("InvalidateEntityStatus: %v", err)
	}
	if len(cache.delCalls) != 2 || cache.delCalls[0] != EntityStatusCacheKey("acme") || cache.delCalls[1] != MeEntityCacheKey("acme") {
		t.Fatalf("deleted keys = %v, want [%q %q]", cache.delCalls, EntityStatusCacheKey("acme"), MeEntityCacheKey("acme"))
	}

	code, _ := gateRequest(t, app)
	if code != fiber.StatusOK {
		t.Fatalf("status = %d, want 200 after invalidation", code)
	}
	if !fetchActive {
		t.Error("fetch did not run after invalidation")
	}
}

// The 403 for a Suspended or Expired tenant carries data.reason, the value the apps key the
// lockdown toast on, next to the error code.
func TestRequireActiveEntity_InactiveCarriesReason(t *testing.T) {
	cache := &statusCacheFake{getErr: errors.New("miss")}
	app := gateApp(cache, func(context.Context, string) (string, error) { return "Suspended", nil })

	req := httptest.NewRequest(fiber.MethodPut, "/", nil)
	req.Header.Set(constant.HEADER_CLIENT_ID, "acme")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	defer resp.Body.Close()
	var body struct {
		Error string `json:"error"`
		Data  struct {
			Reason string `json:"reason"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if body.Data.Reason != ReasonEntityInactive || body.Error != ErrCodeEntityInactive {
		t.Errorf("body = %+v, want data.reason %q and error %q", body, ReasonEntityInactive, ErrCodeEntityInactive)
	}
}
