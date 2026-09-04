package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/service"
	"github.com/gofiber/fiber/v3"
)

// fakeCache is a Valkey stand-in: it embeds the real interface (so it satisfies it) and overrides
// only Get, the one method the middleware reads.
type fakeCache struct {
	service.ValkeyServiceInterface
	get func(ctx context.Context, key string) (string, error)
}

func (f fakeCache) Get(ctx context.Context, key string) (string, error) { return f.get(ctx, key) }

// missCache always misses, so the middleware falls through to the iam fetch.
func missCache() fakeCache {
	return fakeCache{get: func(context.Context, string) (string, error) {
		return "", errors.New("cache miss")
	}}
}

func newRolesFixture() UserAccessCache {
	var r UserAccessCache
	r.ID = 42
	r.FullName = "Elaina"
	r.ClientID = "acme"
	r.ModuleAccess = map[enum.ModuleAccess]ModuleFlags{"fixedAsset": {IsTotalControl: true}}
	return r
}

// captureRoles is a terminal handler that records what the middleware placed in Locals.
func captureRoles(got *UserAccessCache, found *bool) fiber.Handler {
	return func(c fiber.Ctx) error {
		r, ok := GetUserRoles(c)
		*got, *found = r, ok
		return c.SendStatus(fiber.StatusOK)
	}
}

func TestLoadUserRoles_PutsSnapshotInLocals(t *testing.T) {
	fixture := newRolesFixture()

	var gotUserID, gotClientID string
	fetch := func(_ context.Context, userID, clientID string) (UserAccessCache, error) {
		gotUserID, gotClientID = userID, clientID
		return fixture, nil
	}

	var got UserAccessCache
	var found bool
	app := fiber.New()
	app.Use(loadUserRoles(missCache(), fetch))
	app.Get("/", captureRoles(&got, &found))

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set(constant.HEADER_USER_ID, "42")
	req.Header.Set(constant.HEADER_CLIENT_ID, "acme")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	if gotUserID != "42" || gotClientID != "acme" {
		t.Errorf("fetch called with %q/%q, want 42/acme", gotUserID, gotClientID)
	}
	if !found || got.ID != 42 || got.FullName != "Elaina" {
		t.Errorf("Locals roles = %+v found=%v, want fixture", got, found)
	}
}

func TestLoadUserRoles_CacheHitSkipsIAM(t *testing.T) {
	fixture := newRolesFixture()
	payload, _ := json.Marshal(fixture)

	hit := fakeCache{get: func(_ context.Context, key string) (string, error) {
		if key != UserRolesCacheKey("42") {
			t.Errorf("cache key = %q, want %q", key, UserRolesCacheKey("42"))
		}
		return string(payload), nil
	}}
	fetch := func(context.Context, string, string) (UserAccessCache, error) {
		t.Error("iam must not be called on a cache hit")
		return UserAccessCache{}, nil
	}

	var got UserAccessCache
	var found bool
	app := fiber.New()
	app.Use(loadUserRoles(hit, fetch))
	app.Get("/", captureRoles(&got, &found))

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set(constant.HEADER_USER_ID, "42")
	req.Header.Set(constant.HEADER_CLIENT_ID, "acme")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	if !found || got.ID != 42 {
		t.Errorf("Locals roles = %+v found=%v, want fixture from cache", got, found)
	}
}

func TestLoadUserRoles_FailsClosedWhenIAMUnavailable(t *testing.T) {
	fetch := func(context.Context, string, string) (UserAccessCache, error) {
		return UserAccessCache{}, errors.New("connection refused")
	}

	reached := false
	app := fiber.New()
	app.Use(loadUserRoles(missCache(), fetch))
	app.Get("/", func(c fiber.Ctx) error {
		reached = true
		return c.SendStatus(fiber.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set(constant.HEADER_USER_ID, "42")
	req.Header.Set(constant.HEADER_CLIENT_ID, "acme")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	if resp.StatusCode != fiber.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503 (authorization must fail closed)", resp.StatusCode)
	}
	if reached {
		t.Error("handler ran without an RBAC snapshot; the request must not get through unscoped")
	}
}

func TestLoadUserRoles_MissingHeaders_Unauthorized(t *testing.T) {
	fetch := func(context.Context, string, string) (UserAccessCache, error) {
		t.Error("fetch must not be called without identity headers")
		return UserAccessCache{}, nil
	}
	app := fiber.New()
	app.Use(loadUserRoles(missCache(), fetch))
	app.Get("/", func(c fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil)) // no headers
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	if resp.StatusCode != fiber.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", resp.StatusCode)
	}
}

func TestClientID(t *testing.T) {
	cases := []struct {
		name    string
		header  string
		wantID  string
		wantErr bool
	}{
		{name: "valid simple", header: "acme", wantID: "acme"},
		{name: "valid alphanumeric with underscore", header: "tenant_123", wantID: "tenant_123"},
		{name: "empty header", header: "", wantErr: true},
		{name: "invalid characters with hyphen", header: "acme-corp", wantErr: true},
		{name: "invalid characters with space", header: "acme corp", wantErr: true},
		{name: "exceeds length limit", header: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New()
			var gotID string
			var gotErr error
			app.Get("/test", func(c fiber.Ctx) error {
				gotID, gotErr = ClientID(c)
				return c.SendStatus(fiber.StatusOK)
			})

			req := httptest.NewRequest("GET", "/test", nil)
			if tc.header != "" {
				req.Header.Set(constant.HEADER_CLIENT_ID, tc.header)
			}

			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("app.Test: %v", err)
			}
			defer resp.Body.Close()

			if tc.wantErr {
				if gotErr == nil {
					t.Fatalf("ClientID() expected error, got id %q", gotID)
				}
				return
			}
			if gotErr != nil {
				t.Fatalf("ClientID() unexpected error: %v", gotErr)
			}
			if gotID != tc.wantID {
				t.Errorf("ClientID() = %q, want %q", gotID, tc.wantID)
			}
		})
	}
}

func TestEntityID(t *testing.T) {
	cases := []struct {
		name    string
		header  string
		wantID  string
		wantErr bool
	}{
		{name: "valid simple", header: "acme", wantID: "acme"},
		{name: "valid alphanumeric with underscore", header: "entity_123", wantID: "entity_123"},
		{name: "empty header", header: "", wantErr: true},
		{name: "invalid characters with hyphen", header: "entity-corp", wantErr: true},
		{name: "invalid characters with special char", header: "entity@123", wantErr: true},
		{name: "exceeds length limit", header: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New()
			var gotID string
			var gotErr error
			app.Get("/test", func(c fiber.Ctx) error {
				gotID, gotErr = EntityID(c)
				return c.SendStatus(fiber.StatusOK)
			})

			req := httptest.NewRequest("GET", "/test", nil)
			if tc.header != "" {
				req.Header.Set(constant.HEADER_ENTITY_ID, tc.header)
			}

			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("app.Test: %v", err)
			}
			defer resp.Body.Close()

			if tc.wantErr {
				if gotErr == nil {
					t.Fatalf("EntityID() expected error, got id %q", gotID)
				}
				return
			}
			if gotErr != nil {
				t.Fatalf("EntityID() unexpected error: %v", gotErr)
			}
			if gotID != tc.wantID {
				t.Errorf("EntityID() = %q, want %q", gotID, tc.wantID)
			}
		})
	}
}
