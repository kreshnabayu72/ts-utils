package middleware

import (
	"gitea.qwertysystem.net/Wangsit-Developer/be-core-utils/helper"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Get(constant.HEADER_USER_ID)
		clientID := c.Get(constant.HEADER_CLIENT_ID)

		if userID == "" || clientID == "" {
			return helper.SendResponse(c, fiber.StatusUnauthorized, nil, "You are not authenticated!", nil)
		}

		return c.Next()
	}
}
