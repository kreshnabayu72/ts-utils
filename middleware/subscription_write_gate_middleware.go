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

// ProductSubscriptionFetcher resolves the status of a client's CURRENT term of one product on a cache
// miss. It returns "" when the client has no current term of that product. Injectable so the middleware
// can be tested without a live enterprise-management-service, and so EMS itself can read its own
// repository (see RequireActiveProductWriteFrom).
type ProductSubscriptionFetcher func(ctx context.Context, clientID, productCode string) (status string, err error)

const (
	// productSubscriptionTTL matches the client gate: short enough that a renewal or expiry is seen
	// quickly, and InvalidateProductSubscription drops it explicitly on a transition.
	productSubscriptionTTL = 30 * time.Second

	// productSubscriptionNone is what the cache holds for "no current term"; an empty value would be
	// indistinguishable from a miss.
	productSubscriptionNone = "none"

	subscriptionStatusActive  = "active"
	subscriptionStatusGrace   = "grace"
	subscriptionStatusExpired = "expired"
	subscriptionStatusArchive = "archive"
	subscriptionStatusPurge   = "purge"

	// ErrCodeSubscriptionNotFound: the client has no current term of the product the route belongs to.
	ErrCodeSubscriptionNotFound = "SUBSCRIPTION_NOT_FOUND"
)

// ProductSubscriptionCacheKey is the Valkey key holding one client's cached current-term status for one product.
func ProductSubscriptionCacheKey(clientID, productCode string) string {
	return fmt.Sprintf("entity:subscription:%s:%s", clientID, normalizeProduct(productCode))
}

func normalizeProduct(productCode string) string {
	return strings.ToUpper(strings.TrimSpace(productCode))
}

// RequireActiveProductWrite is the single middleware a product service mounts on its WRITE routes. It
// applies the client gate of RequireActiveEntity (Suspended/Expired client) and then the per-product
// subscription gate: a product whose current term is expired or grace is open read-only (session
// AUTHSES-LOGIC-03d, EC-59), so its writes are refused with the entity-isolation §18.8 403. The guard
// is per product subscription, not per client: FAMS expired refuses FAMS writes while SAMS active
// still writes. Mount after authentication and RBAC, on write routes only; reads stay available.
//
// productCode is the EMS product code of the mounting service (e.g. "FAMS", "SAMS"); the service owns
// exactly one product, so it is passed explicitly rather than read from a request header. emsBaseURL
// is EMS's internal address; valkeyOptions is optional (see RequireActiveEntity).
func RequireActiveProductWrite(emsBaseURL, serviceToken, productCode string, valkeyOptions ...db.ValkeyServiceOptions) fiber.Handler {
	client := &http.Client{Timeout: 5 * time.Second}
	return RequireActiveProductWriteFrom(
		fetchEntityStatusFromEMS(client, emsBaseURL, serviceToken),
		fetchProductSubscriptionFromEMS(client, emsBaseURL, serviceToken),
		productCode,
		valkeyOptions...,
	)
}

// RequireActiveProductWriteFrom is RequireActiveProductWrite with injected fetchers, for EMS (which
// reads its own repository) and for tests.
func RequireActiveProductWriteFrom(fetchEntity EntityStatusFetcher, fetchSub ProductSubscriptionFetcher, productCode string, valkeyOptions ...db.ValkeyServiceOptions) fiber.Handler {
	cache := service.MakeValkeyService(valkeyOptions...)
	return requireActiveProductWrite(cache, fetchEntity, fetchSub, productCode)
}

// RequireActiveProduct is the per-product subscription gate alone, for a route that must not also run
// the client gate. Services normally mount RequireActiveProductWrite instead.
func RequireActiveProduct(emsBaseURL, serviceToken, productCode string, valkeyOptions ...db.ValkeyServiceOptions) fiber.Handler {
	client := &http.Client{Timeout: 5 * time.Second}
	return RequireActiveProductFrom(fetchProductSubscriptionFromEMS(client, emsBaseURL, serviceToken), productCode, valkeyOptions...)
}

// RequireActiveProductFrom is RequireActiveProduct with an injected fetcher.
func RequireActiveProductFrom(fetchSub ProductSubscriptionFetcher, productCode string, valkeyOptions ...db.ValkeyServiceOptions) fiber.Handler {
	cache := service.MakeValkeyService(valkeyOptions...)
	return requireActiveProduct(cache, fetchSub, productCode)
}

func requireActiveProductWrite(cache service.ValkeyServiceInterface, fetchEntity EntityStatusFetcher, fetchSub ProductSubscriptionFetcher, productCode string) fiber.Handler {
	return func(c fiber.Ctx) error {
		if handled, err := checkActiveEntity(c, cache, fetchEntity); handled {
			return err
		}
		if handled, err := checkActiveProduct(c, cache, fetchSub, productCode); handled {
			return err
		}
		return c.Next()
	}
}

func requireActiveProduct(cache service.ValkeyServiceInterface, fetchSub ProductSubscriptionFetcher, productCode string) fiber.Handler {
	return func(c fiber.Ctx) error {
		if handled, err := checkActiveProduct(c, cache, fetchSub, productCode); handled {
			return err
		}
		return c.Next()
	}
}

// checkActiveProduct passes only an active current term. expired and grace (read-only) and archive and
// purge (data no longer live) answer the §18.8 403 with the same reason and error code as the client
// gate, so the apps' single lockdown toast applies. No current term at all (or only pending, renewed,
// canceled... terms) is a refusal too, with ErrCodeSubscriptionNotFound. handled is true once a response
// was written.
func checkActiveProduct(c fiber.Ctx, cache service.ValkeyServiceInterface, fetchSub ProductSubscriptionFetcher, productCode string) (handled bool, err error) {
	clientID, err := ClientID(c)
	if err != nil {
		return true, helper.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized", nil)
	}

	status, err := cachedProductSubscription(c.Context(), cache, fetchSub, clientID, productCode)
	if err != nil {
		slog.ErrorContext(c.Context(), "subscription_write_gate: lookup failed", "clientId", clientID, "product", productCode, "err", err)
		return true, sendLookupFailure(c, err)
	}

	switch strings.ToLower(status) {
	case subscriptionStatusActive:
		return false, nil
	case subscriptionStatusGrace, subscriptionStatusExpired, subscriptionStatusArchive, subscriptionStatusPurge:
		return true, sendEntityInactive(c)
	default:
		return true, helper.SendResponse(c, fiber.StatusForbidden, nil, "Forbidden", ErrCodeSubscriptionNotFound)
	}
}

// cachedProductSubscription reads the shared cache first and falls back to fetch on any miss or
// transport error. "no current term" is cached as productSubscriptionNone.
func cachedProductSubscription(ctx context.Context, cache service.ValkeyServiceInterface, fetch ProductSubscriptionFetcher, clientID, productCode string) (string, error) {
	key := ProductSubscriptionCacheKey(clientID, productCode)
	if raw, err := cache.Get(ctx, key); err == nil && raw != "" {
		if raw == productSubscriptionNone {
			return "", nil
		}
		return raw, nil
	}

	status, err := fetch(ctx, clientID, productCode)
	if err != nil {
		return "", err
	}

	cached := status
	if cached == "" {
		cached = productSubscriptionNone
	}
	if err := cache.Set(ctx, key, cached, productSubscriptionTTL); err != nil {
		slog.WarnContext(ctx, "subscription_write_gate: cache set failed", "clientId", clientID, "product", productCode, "err", err)
	}
	return status, nil
}

// InvalidateProductSubscription drops a client's cached current-term status for the given products
// (all products when none is named is not supported: name each one). Call it from every subscription
// transition (activate, renew, expire, grace, archive).
func InvalidateProductSubscription(ctx context.Context, cache service.ValkeyServiceInterface, clientID string, productCodes ...string) error {
	if cache == nil || clientID == "" || len(productCodes) == 0 {
		return nil
	}
	keys := make([]string, 0, len(productCodes))
	for _, p := range productCodes {
		keys = append(keys, ProductSubscriptionCacheKey(clientID, p))
	}
	return cache.Del(ctx, keys...)
}

// currentTermStatus picks the status of the client's current term of a product from every term EMS
// lists: active over grace over expired over archive over purge, mirroring EMS's quota pick. Terms
// in any other status (pending, renewed, canceled, ...) are not current. "" means no current term.
func currentTermStatus(statuses []string) string {
	order := []string{subscriptionStatusActive, subscriptionStatusGrace, subscriptionStatusExpired, subscriptionStatusArchive, subscriptionStatusPurge}
	for _, want := range order {
		for _, st := range statuses {
			if strings.EqualFold(strings.TrimSpace(st), want) {
				return want
			}
		}
	}
	return ""
}

// fetchProductSubscriptionFromEMS calls EMS's service-authenticated GET /internal/subscriptions, which
// lists every term of the client (keyed off X-Client-ID, like every internal read), and picks the
// current term of productCode.
func fetchProductSubscriptionFromEMS(client *http.Client, baseURL, serviceToken string) ProductSubscriptionFetcher {
	baseURL = strings.TrimRight(baseURL, "/")
	return func(ctx context.Context, clientID, productCode string) (string, error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"/internal/subscriptions", nil)
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

		var envelope struct {
			Data struct {
				Subscriptions []struct {
					ProductCode string `json:"productCode"`
					Status      string `json:"status"`
				} `json:"subscriptions"`
			} `json:"data"`
		}
		if err := json.Unmarshal(body, &envelope); err != nil {
			return "", err
		}

		want := normalizeProduct(productCode)
		var statuses []string
		for _, sub := range envelope.Data.Subscriptions {
			if normalizeProduct(sub.ProductCode) == want {
				statuses = append(statuses, sub.Status)
			}
		}
		return currentTermStatus(statuses), nil
	}
}
