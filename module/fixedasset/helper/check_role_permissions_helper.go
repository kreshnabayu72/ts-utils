package helper

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

// Group filter can be "$in":[] or "$eq":"" or just empty bson.M{}
func CheckSystemRolePermission(ctx *fiber.Ctx, permissionKey model.PermissionType, resourceTypes []model.RoleType, groupFilter primitive.M) *entity.HttpError {
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
			"type":      resourceType,
			"group.key": groupFilter,
		})
		if err != nil {
			return entity.InternalServerError(err.Error())
		}

		fmt.Println(bson.M{
			"user._id": userId,
			fmt.Sprintf("managePermission.%s.%s", resourceType, permissionKeyStr): true,
			"type":      resourceType,
			"group.key": groupFilter,
		})

		count2, _ := coll.CountDocuments(ctx.Context(), bson.M{})
		fmt.Printf("COUNT 2 %v\n", count2)

		if count < 1 {
			return entity.Unauthorized("You do not have the required permissions")
		}
	}

	return nil
}
