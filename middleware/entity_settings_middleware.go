package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/entity"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/helper"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/sync/errgroup"
)

// EntityAndSettings is the tenant's entity and settings, the aggregate of iam's GET /me/entity and
// GET /me/settings payloads.
type EntityAndSettings struct {
	Entity          EntityInfo      `json:"entity"`
	ActiveProducts  []string        `json:"activeProducts"`
	GeneralSettings GeneralSettings `json:"generalSettings"`
	SystemLimits    SystemLimits    `json:"systemLimits"`
}

type EntityInfo struct {
	ID           string `json:"_id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	ClientStatus string `json:"clientStatus"`
}

type GeneralSettings struct {
	CurrencyCode          string `json:"currencyCode"`
	CurrencySymbol        string `json:"currencySymbol"`
	CurrencyDecimalPlaces int16  `json:"currencyDecimalPlaces"`
	Timezone              string `json:"timezone"`
	DateFormat            string `json:"dateFormat"` // e.g. "dd/MM/yy", "DDD, MMM dd, yyyy"
	TimeFormat            string `json:"timeFormat"` // "HH:mm" or "hh:mm a"
	SessionIdleTimeout    int    `json:"sessionIdleTimeout"`
}

type SystemLimits struct {
	BulkTransactionRowLimit int `json:"bulkTransactionRowLimit"`
	SelectionLimit          int `json:"selectionLimit"`
	RegisterAssetBulkLimit  int `json:"registerAssetBulkLimit"`
	HierarchyDepthLimit     int `json:"hierarchyDepthLimit"`
	MasterDataCountLimit    int `json:"masterDataCountLimit"`
}

// settingsLayout maps General Settings format tokens to Go layout tokens. Longer tokens come first
// because strings.Replacer tries its pairs in argument order at each position. DDD is the
// weekday ("Mon"), as the General Settings options use it.
var settingsLayout = strings.NewReplacer(
	"yyyy", "2006", "yy", "06",
	"MMM", "Jan", "MM", "01", "M", "1",
	"DDD", "Mon", "dd", "02", "d", "2",
	"HH", "15", "hh", "03", "mm", "04", "a", "PM",
)

// DateLayout is DateFormat as a Go time layout.
func (s GeneralSettings) DateLayout() string { return settingsLayout.Replace(s.DateFormat) }

// TimeLayout is TimeFormat as a Go time layout.
func (s GeneralSettings) TimeLayout() string { return settingsLayout.Replace(s.TimeFormat) }

// Location is the tenant time zone. The service binary must embed tzdata (import _ "time/tzdata")
// when its image has no zoneinfo.
func (s GeneralSettings) Location() (*time.Location, error) { return time.LoadLocation(s.Timezone) }

// EntityAndSettingsFetcher resolves one tenant's entity and settings. Injectable so EMS, which owns
// the data, can read its own services instead of calling its own HTTP endpoints.
type EntityAndSettingsFetcher func(ctx context.Context, clientID string) (EntityAndSettings, error)

// LoadEntityAndSettings returns middleware that puts the caller tenant's EntityAndSettings in
// Locals(EntitySettingsKey), fetched from enterprise-management-service's internal endpoints (the
// same ones iam aggregates for /me/entity and /me/settings). emsBaseURL is EMS's internal address,
// e.g. "http://ts-enterprise-management:5002".
//
// shortcut: no cache, every request pays four EMS reads; mount it only on routes that need the
// data (exports, renders), and add a Valkey layer like LoadUserRoles once a hot route needs it.
func LoadEntityAndSettings(emsBaseURL, serviceToken string) fiber.Handler {
	client := &http.Client{Timeout: 5 * time.Second}
	return LoadEntityAndSettingsFrom(fetchEntityAndSettingsFromEMS(client, emsBaseURL, serviceToken))
}

// LoadEntityAndSettingsFrom is LoadEntityAndSettings with a caller-supplied fetch, for EMS itself.
// Fails closed: a 4xx from fetch (unknown tenant) is 403, anything else is 503.
func LoadEntityAndSettingsFrom(fetch EntityAndSettingsFetcher) fiber.Handler {
	return func(c fiber.Ctx) error {
		clientID, err := ClientID(c)
		if err != nil {
			return helper.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized", nil)
		}

		data, err := fetch(c.Context(), clientID)
		if err != nil {
			slog.ErrorContext(c.Context(), "entity_settings: fetch failed", "clientId", clientID, "err", err)
			if httpErr, ok := err.(*entity.HttpError); ok && httpErr.Code >= 400 && httpErr.Code < 500 {
				return helper.SendResponse(c, fiber.StatusForbidden, nil, "Forbidden", ErrCodeEntityUnknown)
			}
			return helper.SendResponse(c, fiber.StatusServiceUnavailable, nil, "Settings service unavailable", nil)
		}

		c.Locals(EntitySettingsKey, data)
		return c.Next()
	}
}

// RequestEntityAndSettings returns the EntityAndSettings LoadEntityAndSettings placed in Locals.
// Fails closed: no data (middleware not mounted) means the request cannot be served.
func RequestEntityAndSettings(c fiber.Ctx) (EntityAndSettings, error) {
	data, ok := c.Locals(EntitySettingsKey).(EntityAndSettings)
	if !ok {
		return EntityAndSettings{}, entity.InternalServerError("Entity and settings not loaded")
	}
	return data, nil
}

// fetchEntityAndSettingsFromEMS reads the four EMS internal endpoints in parallel. Any failure
// rejects the whole result rather than returning partial settings.
func fetchEntityAndSettingsFromEMS(client *http.Client, baseURL, serviceToken string) EntityAndSettingsFetcher {
	return func(ctx context.Context, clientID string) (EntityAndSettings, error) {
		var (
			out       EntityAndSettings
			clientRow struct {
				ID          string `json:"_id"`
				CompanyName string `json:"companyName"`
				ClientCode  string `json:"clientCode"`
				Status      string `json:"status"`
			}
			subs struct {
				ActiveProducts []string `json:"activeProducts"`
			}
		)
		get := func(path string, dst any) func() error {
			return func() error { return getEMSData(ctx, client, baseURL+path, serviceToken, clientID, dst) }
		}

		g, _ := errgroup.WithContext(ctx)
		g.Go(get("/internal/client", &clientRow))
		g.Go(get("/internal/subscriptions", &subs))
		g.Go(get("/internal/general-settings", &out.GeneralSettings))
		g.Go(get("/internal/lifecycle-settings/system-limits", &out.SystemLimits))
		if err := g.Wait(); err != nil {
			return EntityAndSettings{}, err
		}

		out.Entity = EntityInfo{ID: clientRow.ID, Name: clientRow.CompanyName, Code: clientRow.ClientCode, ClientStatus: clientRow.Status}
		out.ActiveProducts = subs.ActiveProducts
		if out.ActiveProducts == nil {
			out.ActiveProducts = []string{}
		}
		return out, nil
	}
}

// getEMSData GETs one service-authenticated EMS route and decodes the envelope's data into dst.
func getEMSData(ctx context.Context, client *http.Client, url, serviceToken, clientID string, dst any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set(constant.HEADER_SERVICE_TOKEN, serviceToken)
	req.Header.Set(constant.HEADER_CLIENT_ID, clientID)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return &entity.HttpError{
			Code:    resp.StatusCode,
			Message: fmt.Sprintf("enterprise-management returned %d: %s", resp.StatusCode, string(body)),
		}
	}

	var envelope struct {
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		return err
	}
	return json.Unmarshal(envelope.Data, dst)
}
