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

	"context"
	"errors"

	"gitea.qwertysystem.net/BETS/ts-utils/pipeline"
	utilpipeline "gitea.qwertysystem.net/BETS/ts-utils/pipeline"
	"gitea.qwertysystem.net/BETS/ts-utils/service"

	"go.mongodb.org/mongo-driver/mongo/options"
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

		if count < 1 {
			return entity.Unauthorized("You do not have the required permissions")
		}
	}

	return nil
}

// For checking system role group permisson, usually for editing or deleting data
func CheckGroupPermission(service service.Service,
	loggedInUser *primitive.ObjectID,
	dataFilter primitive.M,
	groupIdField string,
	permissions []model.PermissionType,
	systemRoleTypes []model.RoleType,
) error {
	filter := bson.M{
		"user._id": loggedInUser,
		"type":     bson.M{"$in": systemRoleTypes},
	}

	for _, systemRoleType := range systemRoleTypes {
		for _, permission := range permissions {
			filter[fmt.Sprintf("managePermission.%s.%s", systemRoleType, string(permission))] = true
		}
	}

	pipeline := utilpipeline.NewPipelineBuilder().
		Match(dataFilter).
		Lookup(bson.M{
			"from":         db.UserSystemRoleModelName,
			"localField":   groupIdField,
			"foreignField": "group._id",
			"as":           "systemRoles",
			"pipeline": pipeline.NewPipelineBuilder().
				Match(filter).
				Build(),
		}).Unwind(bson.M{"path": "$systemRoles", "preserveNullAndEmptyArrays": false}).
		Build()

	var permissionCheck []map[string]any
	err := service.Aggregate(
		&permissionCheck,
		context.Background(),
		pipeline,
		&options.AggregateOptions{
			Collation: &options.Collation{
				Locale:   "en",
				Strength: 3,
			},
		},
	)
	if err != nil {
		return errors.New("You have no permission")
	}

	return nil
}
