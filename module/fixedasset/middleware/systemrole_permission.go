package middleware

import (
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/helper"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	"gitea.qwertysystem.net/BETS/ts-utils/response"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Group system role need to filter by group
func GlobalSystemRoleRestriction(permission model.PermissionType, resourceTypes []model.RoleType) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// Check permissions based on the provided permissionKey and resourceTypes
		err := helper.CheckSystemRolePermission(c, permission, resourceTypes, bson.M{})
		if err != nil {
			return response.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized: "+err.Message)
		}

		// Continue to the next handler if permission is granted
		return c.Next()
	}
}
