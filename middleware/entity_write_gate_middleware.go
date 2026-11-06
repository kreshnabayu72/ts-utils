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
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/db"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/entity"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/helper"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/service"

	"github.com/gofiber/fiber/v3"
)

// EntityStatusFetcher resolves a tenant's lifecycle status on a cache miss. Injectable so the
// middleware can be tested without a live enterprise-management-service, and so EMS itself can
// read clients.status from its own repository instead of calling its own HTTP endpoint (see
// RequireActiveEntityFrom).
type EntityStatusFetcher func(ctx context.Context, clientID string) (status string, err error)

const (
	// entityStatusTTL keeps a suspend immediately visible without paying an EMS lookup on every
	// write. Status changes invalidate the key explicitly; the TTL is only the backstop.
	entityStatusTTL = 30 * time.Second

	entityStatusSuspended = "Suspended"
	entityStatusExpired   = "Expired"
)

// EntityStatusCacheKey is the Valkey key holding one tenant's cached lifecycle status.
func EntityStatusCacheKey(clientID string) string {
	return fmt.Sprintf("entity:status:%s", clientID)
}

// RequireActiveEntity blocks writes from a suspended or expired tenant. Mount it on WRITE routes
// only, after authentication and RBAC: reads stay available (entity-isolation §18.8), and routes
// that reactivate a subscription must stay exempt when they are added.
//
// The hot path is a direct read of the shared Valkey cache. On a miss it asks
// enterprise-management-service's internal client endpoint, whose data already carries `status`.
// emsBaseURL is EMS's internal address, e.g. "http://ts-enterprise-management:5002".
// valkeyOptions is optional; with none, the shared Valkey connection is taken from the VALKEY_*
// environment (see be-core-utils).
func RequireActiveEntity(emsBaseURL, serviceToken string, valkeyOptions ...db.ValkeyServiceOptions) fiber.Handler {
	client := &http.Client{Timeout: 5 * time.Second}
	return RequireActiveEntityFrom(fetchEntityStatusFromEMS(client, emsBaseURL, serviceToken), valkeyOptions...)
}

// RequireActiveEntityFrom is RequireActiveEntity for the service that owns the clients table: EMS.
// Same cache-hit path, same status contract, same fail-closed behavior — only the cache-miss
// fallback differs. EMS supplies a fetch that reads its own repository directly; every other
// service uses RequireActiveEntity, whose fetch asks EMS over HTTP.
func RequireActiveEntityFrom(fetch EntityStatusFetcher, valkeyOptions ...db.ValkeyServiceOptions) fiber.Handler {
	cache := service.MakeValkeyService(valkeyOptions...)
	return requireActiveEntity(cache, fetch)
}

// Error codes sent in the "error" field of a 403, so the app can tell the cases apart.
const (
	// ErrCodeEntityInactive: the tenant is Suspended or Expired (entity-isolation §18.8 toast).
	ErrCodeEntityInactive = "ENTITY_INACTIVE"
	// ErrCodeEntityUnknown: the tenant lookup refused the client ID.
	ErrCodeEntityUnknown = "ENTITY_UNKNOWN"
	// ErrCodePermissionDenied: the user's roles do not allow the action.
	ErrCodePermissionDenied = "PERMISSION_DENIED"

	// ReasonEntityInactive is data.reason of that 403, the value the apps key the lockdown toast on.
	ReasonEntityInactive = "client_suspended_or_expired"
)

func requireActiveEntity(cache service.ValkeyServiceInterface, fetch EntityStatusFetcher) fiber.Handler {
	return func(c fiber.Ctx) error {
		clientID, err := ClientID(c)
		if err != nil {
			return helper.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized", nil)
		}

		status, err := cachedEntityStatus(c.Context(), cache, fetch, clientID)
		if err != nil {
			slog.ErrorContext(c.Context(), "entity_write_gate: status lookup failed", "clientId", clientID, "err", err)
			// A tenant whose state cannot be established is not a tenant we write for. A 4xx from
			// the fetcher means the tenant itself is unknown, which is a refusal, not an outage.
			if httpErr, ok := err.(*entity.HttpError); ok && httpErr.Code >= 400 && httpErr.Code < 500 {
				return helper.SendResponse(c, fiber.StatusForbidden, nil, "Forbidden", ErrCodeEntityUnknown)
			}
			return helper.SendResponse(c, fiber.StatusServiceUnavailable, nil, "Authorization service unavailable", nil)
		}

		// Expired is a status value the fetcher reports (EMS derives it from subscription state).
		if strings.EqualFold(status, entityStatusSuspended) || strings.EqualFold(status, entityStatusExpired) {
			// data.reason is what the apps read for the suspended/expired lockdown toast (the same reason iam's
			// own module gates answer, LLD §2 / PRD EC-27); the error code stays for callers that switch on it.
			return helper.SendResponse(c, fiber.StatusForbidden, map[string]string{"reason": ReasonEntityInactive}, "Forbidden", ErrCodeEntityInactive)
		}

		return c.Next()
	}
}

// cachedEntityStatus reads the shared cache first and falls back to fetch on any miss or transport
// error. An empty fetched status is an error, never a cacheable "active".
func cachedEntityStatus(
	ctx context.Context,
	cache service.ValkeyServiceInterface,
	fetch EntityStatusFetcher,
	clientID string,
) (string, error) {
	if raw, err := cache.Get(ctx, EntityStatusCacheKey(clientID)); err == nil && raw != "" {
		return raw, nil
	}

	status, err := fetch(ctx, clientID)
	if err != nil {
		return "", err
	}
	if status == "" {
		return "", fmt.Errorf("entity status fetch returned an empty status")
	}

	if err := cache.Set(ctx, EntityStatusCacheKey(clientID), status, entityStatusTTL); err != nil {
		// A cache write failure must not fail the request: the status was established, and the
		// next write simply pays another lookup.
		slog.WarnContext(ctx, "entity_write_gate: cache set failed", "clientId", clientID, "err", err)
	}
	return status, nil
}

// InvalidateEntityStatus drops a tenant's cached status so the next write observes a change made
// just now instead of waiting out the TTL. Call it from every status transition.
func InvalidateEntityStatus(ctx context.Context, cache service.ValkeyServiceInterface, clientID string) error {
	if cache == nil || clientID == "" {
		return nil
	}
	// iam caches GET /v2/me/entity (clientStatus included) under me-entity:<tenantCode> in the same
	// Valkey; without this a status change kept answering the old clientStatus for up to a day.
	return cache.Del(ctx, EntityStatusCacheKey(clientID), MeEntityCacheKey(clientID))
}

// MeEntityCacheKey is iam's cache key for a tenant's GET /v2/me/entity answer.
func MeEntityCacheKey(clientID string) string { return "me-entity:" + clientID }

// fetchEntityStatusFromEMS calls EMS's internal client endpoint. The route is service-authenticated
// and keys the tenant off X-Client-ID, the same header every other internal read uses.
func fetchEntityStatusFromEMS(client *http.Client, baseURL, serviceToken string) EntityStatusFetcher {
	return func(ctx context.Context, clientID string) (string, error) {
		url := fmt.Sprintf("%s/internal/client", baseURL)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Set(constant.HEADER_SERVICE_TOKEN, serviceToken)
		req.Header.Set(constant.HEADER_CLIENT_ID, clientID)

		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		if resp.StatusCode != http.StatusOK {
			return "", &entity.HttpError{
				Code:    resp.StatusCode,
				Message: fmt.Sprintf("enterprise-management returned %d: %s", resp.StatusCode, string(body)),
			}
		}

		// EMS wraps its payload in the standard { status, message, data } envelope.
		var envelope struct {
			Data struct {
				Status string `json:"status"`
			} `json:"data"`
		}
		if err := json.Unmarshal(body, &envelope); err != nil {
			return "", err
		}
		return envelope.Data.Status, nil
	}
}
