package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/entity"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/service"

	"github.com/gofiber/fiber/v3"
)

// keyedCacheFake is a Valkey stand-in keyed per key: the client gate and the product gate share one
// cache, so their entries must not collide.
type keyedCacheFake struct {
	service.ValkeyServiceInterface
	data map[string]string
	sets int
}

func newKeyedCache() *keyedCacheFake { return &keyedCacheFake{data: map[string]string{}} }

func (f *keyedCacheFake) Get(_ context.Context, key string) (string, error) {
	v, ok := f.data[key]
	if !ok {
		return "", errors.New("miss")
	}
	return v, nil
}

func (f *keyedCacheFake) Set(_ context.Context, key string, value any, _ time.Duration) error {
	f.sets++
	f.data[key] = fmt.Sprint(value)
	return nil
}

func activeClient(context.Context, string) (string, error) { return "Active", nil }

func subApp(cache service.ValkeyServiceInterface, fetchEntity EntityStatusFetcher, fetchSub ProductSubscriptionFetcher) *fiber.App {
	app := fiber.New()
	app.Put("/", requireActiveProductWrite(cache, fetchEntity, fetchSub, "FAMS"), func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	return app
}

type subResponse struct {
	code   int
	errMsg string
	reason string
}

func subRequest(t *testing.T, app *fiber.App) subResponse {
	t.Helper()
	req := httptest.NewRequest(fiber.MethodPut, "/", nil)
	req.Header.Set(constant.HEADER_CLIENT_ID, "acme")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	var body struct {
		Error string `json:"error"`
		Data  struct {
			Reason string `json:"reason"`
		} `json:"data"`
	}
	if len(raw) > 0 && raw[0] == '{' {
		if err := json.Unmarshal(raw, &body); err != nil {
			t.Fatalf("decode body %q: %v", raw, err)
		}
	}
	return subResponse{resp.StatusCode, body.Error, body.Data.Reason}
}

func fixedSub(status string) ProductSubscriptionFetcher {
	return func(context.Context, string, string) (string, error) { return status, nil }
}

func TestRequireActiveProductWrite_ActivePasses(t *testing.T) {
	cache := newKeyedCache()
	var gotClient, gotProduct string
	app := subApp(cache, activeClient, func(_ context.Context, clientID, product string) (string, error) {
		gotClient, gotProduct = clientID, product
		return "active", nil
	})
	if r := subRequest(t, app); r.code != fiber.StatusOK {
		t.Fatalf("status = %d, want 200", r.code)
	}
	if gotClient != "acme" || gotProduct != "FAMS" {
		t.Errorf("fetch = %q/%q, want acme/FAMS", gotClient, gotProduct)
	}
	if got := cache.data[ProductSubscriptionCacheKey("acme", "FAMS")]; got != "active" {
		t.Errorf("product cache = %q, want active", got)
	}
}

func TestRequireActiveProductWrite_ReadOnlyStatusesRefused(t *testing.T) {
	for _, status := range []string{"grace", "expired", "archive", "purge"} {
		t.Run(status, func(t *testing.T) {
			app := subApp(newKeyedCache(), activeClient, fixedSub(status))
			r := subRequest(t, app)
			if r.code != fiber.StatusForbidden || r.errMsg != ErrCodeEntityInactive || r.reason != ReasonEntityInactive {
				t.Errorf("got %+v, want 403 %s reason %s", r, ErrCodeEntityInactive, ReasonEntityInactive)
			}
		})
	}
}

func TestRequireActiveProductWrite_NoCurrentTermRefused(t *testing.T) {
	cache := newKeyedCache()
	app := subApp(cache, activeClient, fixedSub(""))
	r := subRequest(t, app)
	if r.code != fiber.StatusForbidden || r.errMsg != ErrCodeSubscriptionNotFound {
		t.Fatalf("got %+v, want 403 %s", r, ErrCodeSubscriptionNotFound)
	}
	if got := cache.data[ProductSubscriptionCacheKey("acme", "FAMS")]; got != productSubscriptionNone {
		t.Errorf("cache = %q, want %q", got, productSubscriptionNone)
	}
	// the cached "none" must still refuse without another fetch
	app = subApp(cache, activeClient, func(context.Context, string, string) (string, error) {
		t.Error("fetch must not run on a cache hit")
		return "active", nil
	})
	if r := subRequest(t, app); r.code != fiber.StatusForbidden || r.errMsg != ErrCodeSubscriptionNotFound {
		t.Errorf("cached none: got %+v", r)
	}
}

func TestRequireActiveProductWrite_CacheHitSkipsFetch(t *testing.T) {
	cache := newKeyedCache()
	cache.data[EntityStatusCacheKey("acme")] = "Active"
	cache.data[ProductSubscriptionCacheKey("acme", "FAMS")] = "expired"
	app := subApp(cache, activeClient, func(context.Context, string, string) (string, error) {
		t.Error("fetch must not run on a cache hit")
		return "active", nil
	})
	if r := subRequest(t, app); r.code != fiber.StatusForbidden || r.reason != ReasonEntityInactive {
		t.Fatalf("got %+v, want 403 from the cached status", r)
	}
}

func TestRequireActiveProductWrite_ClientGateStillApplies(t *testing.T) {
	app := subApp(newKeyedCache(),
		func(context.Context, string) (string, error) { return "Suspended", nil },
		func(context.Context, string, string) (string, error) {
			t.Error("product lookup must not run for a suspended client")
			return "active", nil
		})
	if r := subRequest(t, app); r.code != fiber.StatusForbidden || r.errMsg != ErrCodeEntityInactive {
		t.Fatalf("got %+v, want 403 %s", r, ErrCodeEntityInactive)
	}
}

func TestRequireActiveProductWrite_FailsClosed(t *testing.T) {
	cases := []struct {
		name     string
		err      error
		wantCode int
		wantErr  string
	}{
		{"ems 404", &entity.HttpError{Code: 404, Message: "Client not found"}, fiber.StatusForbidden, ErrCodeEntityUnknown},
		{"ems 5xx", &entity.HttpError{Code: 500, Message: "boom"}, fiber.StatusServiceUnavailable, ""},
		{"transport", errors.New("connection refused"), fiber.StatusServiceUnavailable, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			app := subApp(newKeyedCache(), activeClient,
				func(context.Context, string, string) (string, error) { return "", tc.err })
			r := subRequest(t, app)
			if r.code != tc.wantCode || r.errMsg != tc.wantErr {
				t.Errorf("got %+v, want %d %q", r, tc.wantCode, tc.wantErr)
			}
		})
	}
}

func TestRequireActiveProductWrite_MissingClientIDUnauthorized(t *testing.T) {
	app := subApp(newKeyedCache(), activeClient, fixedSub("active"))
	resp, err := app.Test(httptest.NewRequest(fiber.MethodPut, "/", nil))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != fiber.StatusUnauthorized {
		t.Errorf("status = %d, want 401", resp.StatusCode)
	}
}

func TestProductSubscriptionCacheKey_PerProduct(t *testing.T) {
	if ProductSubscriptionCacheKey("acme", "fams") != "entity:subscription:acme:FAMS" {
		t.Errorf("key = %q", ProductSubscriptionCacheKey("acme", "fams"))
	}
	if ProductSubscriptionCacheKey("acme", "FAMS") == ProductSubscriptionCacheKey("acme", "SAMS") {
		t.Error("products must not share a cache key")
	}
}

func TestCurrentTermStatus_PreferActiveOverGraceOverExpired(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"expired", "active", "renewed"}, "active"},
		{[]string{"expired", "grace"}, "grace"},
		{[]string{"renewed", "Expired"}, "expired"},
		{[]string{"archive", "purge"}, "archive"},
		{[]string{"pending", "renewed", "canceled"}, ""},
		{nil, ""},
	}
	for _, tc := range cases {
		if got := currentTermStatus(tc.in); got != tc.want {
			t.Errorf("currentTermStatus(%v) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestFetchProductSubscriptionFromEMS(t *testing.T) {
	var gotPath, gotToken, gotClient string
	body := `{"data":{"subscriptions":[
		{"productCode":"SAMS","status":"active"},
		{"productCode":"FAMS","status":"renewed"},
		{"productCode":"fams","status":"expired"}]}}`
	status := http.StatusOK
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath, gotToken, gotClient = r.URL.Path, r.Header.Get(constant.HEADER_SERVICE_TOKEN), r.Header.Get(constant.HEADER_CLIENT_ID)
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	defer srv.Close()

	fetch := fetchProductSubscriptionFromEMS(srv.Client(), srv.URL+"/", "tok")
	got, err := fetch(context.Background(), "acme", "FAMS")
	if err != nil || got != "expired" {
		t.Fatalf("got %q, %v; want expired (SAMS active must not leak in)", got, err)
	}
	if gotPath != "/internal/subscriptions" || gotToken != "tok" || gotClient != "acme" {
		t.Errorf("request = %q token %q client %q", gotPath, gotToken, gotClient)
	}
	if got, _ := fetch(context.Background(), "acme", "SIEM"); got != "" {
		t.Errorf("SIEM = %q, want no term", got)
	}

	status = http.StatusNotFound
	_, err = fetch(context.Background(), "acme", "FAMS")
	var httpErr *entity.HttpError
	if !errors.As(err, &httpErr) || httpErr.Code != 404 {
		t.Errorf("err = %v, want HttpError 404", err)
	}
}
