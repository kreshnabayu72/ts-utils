package middleware

import (
	"fmt"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/middleware"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func checkPermission(ctx *fiber.Ctx, permissionKey model.PermissionType, resourceTypes []model.RoleType) *entity.HttpError {
	// Get userId from Locals in Fiber context
	userId, ok := ctx.Locals(middleware.UserKey).(primitive.ObjectID)
	if !ok {
		return entity.Unauthorized("User not authenticated")
	}

	permissionKeyStr := string(permissionKey)
	companyCode := ctx.Locals(middleware.CompanyCodeKey).(string)

	db.ConnectToDB(db.FixedAssetDBName.CompanyDBName(companyCode))
	coll := db.GetCollection(db.UserSystemRoleModelName)

	// Convert resourceTypes to a slice of strings
	resourceTypeStrs := make([]string, len(resourceTypes))
	for i, resourceType := range resourceTypes {
		resourceTypeStrs[i] = string(resourceType)

		// Check for matching documents where the resource type is in the given array
		count, err := coll.CountDocuments(ctx.Context(), bson.M{
			"user._id": userId,
			fmt.Sprintf("managePermission.%s.%s", resourceType, permissionKeyStr): true,
			"type": resourceType,
		})
		if err != nil {
			return entity.InternalServerError(err.Error())
		}

		if count < 1 {
			return entity.Unauthorized("You do not have the required permissions")
		}
	}

	return nil
}

// Group system role need to filter by group
func GlobalSystemRoleRestriction(permission model.PermissionType, resourceTypes []model.RoleType) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// Check permissions based on the provided permissionKey and resourceTypes
		err := checkPermission(c, permission, resourceTypes)
		if err != nil {
			return fiber.NewError(err.Code, err.Message)
		}

		// Continue to the next handler if permission is granted
		return c.Next()
	}
}
