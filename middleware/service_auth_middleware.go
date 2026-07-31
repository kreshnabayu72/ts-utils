package middleware

import (
	"crypto/subtle"
	"log/slog"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/helper"

	"github.com/gofiber/fiber/v3"
)

// RequireServiceToken authenticates a SERVICE, not a user. Mount it on every /internal route.
//
// The /internal endpoints are the ones that answer without a user's authorization behind them: iam
// hands out any user's whole RBAC snapshot, master-asset-registry hands out any asset, and the
// tenant is chosen by the X-Client-ID header. Keeping them off the public gateway keeps the internet
// out; it does not keep the CLUSTER out. Anything able to reach the port could otherwise read any
// user's permissions in any tenant — and that snapshot is precisely what every other service trusts
// to make its own authorization decisions. It is the one door that must not be open.
//
// A shared secret is the floor, not the ceiling: it authenticates the caller as "some service we
// gave the secret to", nothing finer. mTLS or a signed service identity is the real answer, and this
// is deliberately a single choke point so replacing it later touches one file.
//
// Fails closed, and closed means closed: an unset secret refuses every request rather than waving
// them all through. A misconfigured deployment must break loudly, not silently unlock the door.
func RequireServiceToken(secret string) fiber.Handler {
	return func(c fiber.Ctx) error {
		if secret == "" {
			slog.ErrorContext(c.Context(),
				"service_auth: no service token configured; refusing internal request",
				"path", c.Path())
			return helper.SendResponse(c, fiber.StatusServiceUnavailable, nil, "Service unavailable", nil)
		}

		// Constant-time: a byte-by-byte comparison leaks the secret one character at a time to
		// anyone willing to time the responses.
		presented := c.Get(constant.HEADER_SERVICE_TOKEN)
		if subtle.ConstantTimeCompare([]byte(presented), []byte(secret)) != 1 {
			slog.WarnContext(c.Context(),
				"service_auth: rejected internal request", "path", c.Path(), "ip", c.IP())
			return helper.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized", nil)
		}

		return c.Next()
	}
}
