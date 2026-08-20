package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/db"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/entity"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/helper"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/service"

	"github.com/gofiber/fiber/v3"
)

// RolesFetcher resolves the caller's RBAC snapshot on a cache miss. Injectable so the middleware
// can be tested without a live iam, and so iam itself can supply a fetch that reads its own
// database directly instead of calling its own HTTP endpoint (see LoadUserRolesFrom).
type RolesFetcher func(ctx context.Context, userID, clientID string) (UserAccessCache, error)

// LoadUserRoles returns middleware that puts the caller's fixed-asset RBAC snapshot in
// Locals(UserRolesKey).
//
// The hot path is a direct read of the shared Valkey cache — no service hop. iam is the only
// service that touches the role tables and the only writer of the cache; when this middleware
// misses, it asks iam, which rebuilds from its own database and SETs the key, warming it for the
// next request. So a consumer reads the cache directly and hops to iam only on a miss — never
// opening a connection to iam's database.
//
// iamBaseURL is iam's internal address, e.g. "http://ts-iam:5005". valkeyOptions is optional; with
// none, the shared Valkey connection is taken from the VALKEY_* environment (see be-core-utils).
func LoadUserRoles(iamBaseURL, serviceToken string, valkeyOptions ...db.ValkeyServiceOptions) fiber.Handler {
	client := &http.Client{Timeout: 5 * time.Second}
	cache := service.MakeValkeyService(valkeyOptions...)
	return loadUserRoles(cache, fetchRolesFromIAM(client, iamBaseURL, serviceToken))
}

// LoadUserRolesFrom is LoadUserRoles for the one service that owns the role tables: iam. Same
// cache-hit path, same Locals contract, same fail-closed behavior — only the cache-miss fallback
// differs. iam supplies a fetch that reads its own repository directly; every other service uses
// LoadUserRoles, whose fetch asks iam over HTTP, because they have no database to read.
//
// Calling iam's own HTTP endpoint from inside iam would be a self-referential network hop to reach
// code already in the same binary — this skips that hop while keeping the exact same middleware
// shape every route in every service (including iam's) is written against.
func LoadUserRolesFrom(fetch RolesFetcher, valkeyOptions ...db.ValkeyServiceOptions) fiber.Handler {
	cache := service.MakeValkeyService(valkeyOptions...)
	return loadUserRoles(cache, fetch)
}

func loadUserRoles(cache service.ValkeyServiceInterface, fetch RolesFetcher) fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Get(constant.HEADER_USER_ID)
		clientID := c.Get(constant.HEADER_CLIENT_ID)
		if userID == "" || clientID == "" {
			return helper.SendResponse(c, fiber.StatusUnauthorized, nil, "You are not authenticated!", nil)
		}

		// Hot path: read the shared cache directly. On a hit the consumer never touches iam.
		if raw, err := cache.Get(c.Context(), UserRolesCacheKey(userID)); err == nil {
			var roles UserAccessCache
			if jsonErr := json.Unmarshal([]byte(raw), &roles); jsonErr == nil {
				c.Locals(UserRolesKey, roles)
				return c.Next()
			}
			// A corrupt entry is treated as a miss and refetched below, not served.
			slog.WarnContext(
				c.Context(),
				"user_roles: cache entry unparseable, refetching from iam",
				"userId",
				userID,
			)
		}

		// Miss, corrupt entry, or Valkey unreachable: ask iam, which rebuilds from its DB and SETs
		// the cache for next time. Fail closed if it cannot answer — authorization needs the snapshot.
		roles, err := fetch(c.Context(), userID, clientID)
		if err != nil {
			slog.ErrorContext(c.Context(), "user_roles: fetch from iam failed", "userId", userID, "err", err)
			if httpErr, ok := err.(*entity.HttpError); ok && httpErr.Code >= 400 && httpErr.Code < 500 {
				return helper.SendResponse(c, fiber.StatusUnauthorized, nil, "You are not authenticated!", nil)
			}
			return helper.SendResponse(c, fiber.StatusServiceUnavailable, nil, "Authorization service unavailable", nil)
		}

		c.Locals(UserRolesKey, roles)
		return c.Next()
	}
}

// fetchRolesFromIAM calls iam's internal snapshot endpoint. iam serves it from its cache and
// rebuilds on a miss, so the consumer neither knows nor cares which happened.
func fetchRolesFromIAM(client *http.Client, baseURL, serviceToken string) RolesFetcher {
	return func(ctx context.Context, userID, clientID string) (UserAccessCache, error) {
		var roles UserAccessCache

		url := fmt.Sprintf("%s/internal/users/%s/roles", baseURL, userID)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return roles, err
		}
		req.Header.Set(constant.HEADER_USER_ID, userID)
		req.Header.Set(constant.HEADER_CLIENT_ID, clientID)
		// This IS an internal call: iam's snapshot endpoint is service-authenticated.
		req.Header.Set(constant.HEADER_SERVICE_TOKEN, serviceToken)

		resp, err := client.Do(req)
		if err != nil {
			return roles, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return roles, err
		}
		if resp.StatusCode != http.StatusOK {
			return roles, &entity.HttpError{Code: resp.StatusCode, Message: fmt.Sprintf("iam returned %d: %s", resp.StatusCode, string(body))}
		}

		// iam wraps its payload in the standard { status, message, data } envelope.
		var envelope struct {
			Data UserAccessCache `json:"data"`
		}
		if err := json.Unmarshal(body, &envelope); err != nil {
			return roles, err
		}
		return envelope.Data, nil
	}
}

// validClientID locks the tenant charset. The client id becomes a PostgreSQL SCHEMA NAME
// (interpolated into SET LOCAL search_path), so it must be pinned before it leaves the handler.
var validClientID = regexp.MustCompile(`^[a-zA-Z0-9_]{1,50}$`)

// ClientID is the tenant, taken from the X-Client-ID header the gateway sets.
//
// It is the ONE place a service decides which tenant it is serving. Never take the tenant from a
// query param or a body field: that lets a caller point a request at somebody else's schema and
// read data from a tenant they have nothing to do with. Header only, validated, fails closed.
func ClientID(c fiber.Ctx) (string, error) {
	id := c.Get(constant.HEADER_CLIENT_ID)
	if !validClientID.MatchString(id) {
		return "", entity.BadRequest("Invalid client id")
	}
	return id, nil
}

// CallerID is the authenticated caller's user id, taken from the X-User-ID header the gateway
// sets and a client cannot forge.
//
// It is the ONE place a service reads the caller's identity. Never take it from a body field or a
// path parameter: whose request this is decides what they may do, and letting the caller name
// themselves lets anyone act as anyone. Fails closed — no usable header means not authenticated.
func CallerID(c fiber.Ctx) (int64, error) {
	id, err := strconv.ParseInt(c.Get(constant.HEADER_USER_ID), 10, 64)
	if err != nil {
		return 0, entity.Unauthorized("You are not authenticated!")
	}
	return id, nil
}

// GetUserRoles returns the RBAC snapshot placed in Locals by LoadUserRoles.
func GetUserRoles(c fiber.Ctx) (UserAccessCache, bool) {
	roles, ok := c.Locals(UserRolesKey).(UserAccessCache)
	return roles, ok
}

// RequestIdentity bundles the three request-scoped facts a handler most often needs together:
// which tenant, whose RBAC snapshot, and who is calling — the snapshot LoadUserRoles placed in
// Locals, plus X-Client-ID and X-User-ID. It returns the raw snapshot, not a derived scope —
// decisions are asked directly of it (HasGroupPermissionFixedAsset, GroupIDsWithTransactionRole, ...), so
// nothing gets lost in a reduction step.
//
// Fails closed — no snapshot (middleware not mounted, or iam unreachable), no usable X-User-ID, or
// no usable X-Client-ID means not authenticated.
//
// Use it only where the handler actually needs all three; a handler that needs just the tenant, or
// is gated by role middleware without touching the caller's own id, should keep calling ClientID /
// CallerID / GetUserRoles directly rather than over-fetch.
func RequestIdentity(c fiber.Ctx) (clientID string, roles UserAccessCache, callerID int64, err error) {
	clientID, err = ClientID(c)
	if err != nil {
		return "", UserAccessCache{}, 0, err
	}
	roles, ok := GetUserRoles(c)
	if !ok {
		return "", UserAccessCache{}, 0, entity.Unauthorized("You are not authenticated!")
	}
	callerID, err = CallerID(c)
	if err != nil {
		return "", UserAccessCache{}, 0, err
	}
	return clientID, roles, callerID, nil
}
