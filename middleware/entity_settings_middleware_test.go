package middleware

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/entity"
	"github.com/gofiber/fiber/v3"
)

func TestGeneralSettingsLayouts(t *testing.T) {
	at := time.Date(2026, 8, 10, 20, 30, 0, 0, time.UTC)
	cases := map[string]string{
		"dd/MM/yy":          "10/08/26",
		"M/d/yy":            "8/10/26",
		"dd MMM yy":         "10 Aug 26",
		"DDD, MMM dd, yy":   "Mon, Aug 10, 26",
		"DDD, MMM dd, yyyy": "Mon, Aug 10, 2026",
	}
	for format, want := range cases {
		if got := at.Format(GeneralSettings{DateFormat: format}.DateLayout()); got != want {
			t.Errorf("DateFormat %q = %q, want %q", format, got, want)
		}
	}
	if got := at.Format(GeneralSettings{TimeFormat: "hh:mm a"}.TimeLayout()); got != "08:30 PM" {
		t.Errorf("hh:mm a = %q", got)
	}
	if got := at.Format(GeneralSettings{TimeFormat: "HH:mm"}.TimeLayout()); got != "20:30" {
		t.Errorf("HH:mm = %q", got)
	}

	loc, err := GeneralSettings{Timezone: "Asia/Jakarta"}.Location()
	if err != nil {
		t.Fatal(err)
	}
	if got := at.In(loc).Format(GeneralSettings{DateFormat: "dd MMM yy"}.DateLayout()); got != "11 Aug 26" {
		t.Errorf("Asia/Jakarta date = %q, want 11 Aug 26", got)
	}
}

func entitySettingsApp(fetch EntityAndSettingsFetcher) *fiber.App {
	app := fiber.New()
	app.Get("/", LoadEntityAndSettingsFrom(fetch), func(c fiber.Ctx) error {
		data, err := RequestEntityAndSettings(c)
		if err != nil {
			return err
		}
		return c.SendString(data.GeneralSettings.Timezone)
	})
	return app
}

func TestLoadEntityAndSettingsFrom(t *testing.T) {
	ok := func(context.Context, string) (EntityAndSettings, error) {
		return EntityAndSettings{GeneralSettings: GeneralSettings{Timezone: "Asia/Jakarta"}}, nil
	}
	cases := []struct {
		name     string
		clientID string
		fetch    EntityAndSettingsFetcher
		want     int
	}{
		{"loaded", "acme", ok, http.StatusOK},
		{"no client id", "", ok, http.StatusUnauthorized},
		{"unknown tenant", "acme", func(context.Context, string) (EntityAndSettings, error) {
			return EntityAndSettings{}, entity.NotFound("Not found")
		}, http.StatusForbidden},
		{"ems down", "acme", func(context.Context, string) (EntityAndSettings, error) {
			return EntityAndSettings{}, errors.New("dial tcp: refused")
		}, http.StatusServiceUnavailable},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(constant.HEADER_CLIENT_ID, tc.clientID)
			resp, err := entitySettingsApp(tc.fetch).Test(req)
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != tc.want {
				t.Errorf("status = %d, want %d", resp.StatusCode, tc.want)
			}
		})
	}
}

func TestRequestEntityAndSettingsNotLoaded(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		if _, err := RequestEntityAndSettings(c); err == nil {
			t.Error("expected error when middleware is not mounted")
		}
		return nil
	})
	if _, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil)); err != nil {
		t.Fatal(err)
	}
}

func TestFetchEntityAndSettingsFromEMS(t *testing.T) {
	bodies := map[string]string{
		"/internal/client":                           `{"data":{"_id":7,"companyName":"Acme","clientCode":"acme","status":"Active"}}`,
		"/internal/subscriptions":                    `{"data":{"activeProducts":["fixedAsset"]}}`,
		"/internal/general-settings":                 `{"data":{"timezone":"Asia/Jakarta","dateFormat":"dd MMM yy","timeFormat":"HH:mm"}}`,
		"/internal/lifecycle-settings/system-limits": `{"data":{"selectionLimit":100}}`,
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(constant.HEADER_SERVICE_TOKEN) != "tok" || r.Header.Get(constant.HEADER_CLIENT_ID) != "acme" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		body, ok := bodies[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = w.Write([]byte(body))
	}))
	defer srv.Close()

	got, err := fetchEntityAndSettingsFromEMS(srv.Client(), srv.URL, "tok")(context.Background(), "acme")
	if err != nil {
		t.Fatal(err)
	}
	if got.Entity.ID != 7 || got.Entity.Name != "Acme" || got.Entity.ClientStatus != "Active" ||
		len(got.ActiveProducts) != 1 || got.GeneralSettings.DateFormat != "dd MMM yy" || got.SystemLimits.SelectionLimit != 100 {
		t.Errorf("unexpected aggregate: %+v", got)
	}

	delete(bodies, "/internal/general-settings")
	_, err = fetchEntityAndSettingsFromEMS(srv.Client(), srv.URL, "tok")(context.Background(), "acme")
	var httpErr *entity.HttpError
	if !errors.As(err, &httpErr) || httpErr.Code != http.StatusNotFound {
		t.Errorf("err = %v, want HttpError 404", err)
	}
}
