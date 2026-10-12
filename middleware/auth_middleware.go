// Package middleware provides Fiber middleware for authentication,
// service-to-service auth, RBAC snapshot loading, and the shared
// request-identity helpers (ClientID, CallerID, GetUserRoles).
package middleware

import (
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/helper"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Get(constant.HEADER_USER_ID)
		clientID := c.Get(constant.HEADER_CLIENT_ID)

		if userID == "" || clientID == "" {
			return helper.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized", nil)
		}

		return c.Next()
	}
}
