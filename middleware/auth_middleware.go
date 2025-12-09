package middleware

import (
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/helper"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Get("X-User-ID")
		clientID := c.Get("X-Client-ID")

		if userID == "" || clientID == "" {
			return helper.SendResponse(c, fiber.StatusUnauthorized, nil, "You are not authenticated!", nil)
		}

		return c.Next()
	}
}
