package helper

import (
	"context"
	"errors"
	"fmt"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/middleware"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"github.com/gofiber/fiber/v2"
)

// Group filter can be "$in":[] or "$eq":"" or just empty bson.M{}
func CheckSystemRolePermissionSQL(ctx *fiber.Ctx, permissionKey new_model.PermissionType, resourceTypes []new_model.SystemRoleAttribute, groupFilter sql_query.SQLCondition) *entity.HttpError {
	// Get userId from Locals in Fiber context
	userId, ok := ctx.Locals(middleware.UserKey).(string)
	if !ok {
		return entity.Unauthorized("User not authenticated")
	}

	permissionKeyStr := string(permissionKey)
	companyCode := ctx.Locals(middleware.CompanyCodeKey).(string)

	fixedAssetDbName := db.FixedAssetDBName.CompanyDBName(companyCode)
	service := service.MakePostgreSqlService(fixedAssetDbName, db.UserSystemRoleTableName)

	// Convert resourceTypes to a slice of strings
	resourceTypeStrs := make([]string, len(resourceTypes))
	for i, resourceType := range resourceTypes {
		resourceTypeStrs[i] = string(resourceType)

		filter := map[string]sql_query.SQLCondition{
			"user_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    userId,
			},
			"role_attribute": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    resourceType,
			},
			fmt.Sprintf("permission_%s", permissionKeyStr): {
				Operator: sql_query.SQLOperatorEqual,
				Value:    true,
			},
			"group_id": groupFilter,
			"is_active": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    true,
			},
		}

		// Check for matching documents where the resource type is in the given array
		count, err := service.CountWithFilter(ctx.Context(), filter)
		if err != nil {
			return entity.InternalServerError(err.Error())
		}

		if count < 1 {
			fmt.Printf("FILTER: %v\n", filter)
			return entity.Unauthorized("You do not have the required permissions")
		}
	}

	return nil
}

// For checking system role group permisson, usually for editing or deleting data
// Can handle multiple data with different groups
func CheckGroupPermissionSQL(
	service service.PostgreSqlService,
	tableName string,
	loggedInUser *string,
	dataFilter map[string]sql_query.SQLCondition,
	groupIdField string,
	permissions []new_model.PermissionType,
	systemRoleTypes []new_model.SystemRoleAttribute,
) error {
	filter := map[string]sql_query.SQLCondition{
		"usr.user_id": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    loggedInUser,
		},
		"usr.role_attribute": {
			Operator: sql_query.SQLOperatorIn,
			Value:    systemRoleTypes,
		},
		"usr.is_active": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    true,
		},
	}

	builder := sql_query.NewSQLSelectBuilder[any](tableName)

	for _, systemRoleType := range systemRoleTypes {
		for _, permission := range permissions {
			filterOr := make(map[string]sql_query.SQLCondition, 2)
			filterOr["usr.role_attribute"] = sql_query.SQLCondition{Operator: sql_query.SQLOperatorEqual, Value: systemRoleType}
			filterOr[fmt.Sprintf("usr.permission_%s", string(permission))] = sql_query.SQLCondition{Operator: sql_query.SQLOperatorEqual, Value: true}
			builder.WhereOr(filterOr)
		}
	}

	query, args, _ := builder.
		Select(tableName+".id", "usr.id as \"userSystemRoleId\"").
		LeftJoin(fmt.Sprintf("%s as usr", db.UserSystemRoleTableName), fmt.Sprintf("%s = usr.group_id", groupIdField), filter).
		Where(dataFilter).
		Build()

	permissionCheck := make([]map[string]any, 0)
	err := service.SelectMany(&permissionCheck, context.Background(), query, args...)
	if err != nil || len(permissionCheck) < 1 {
		return errors.New("Data not found")
	}

	for _, data := range permissionCheck {
		systemRoles := data["userSystemRoleId"]
		if systemRoles == nil {
			// If ONE OF data have no permission, throw error
			return errors.New("You have no permission")
		}
	}

	return nil
}

// For checking system role group permisson, usually for editing or deleting data
// Can handle multiple data with different groups
func CheckTransactionRoleMultiGroupPermissionsSQL(
	service service.PostgreSqlService, // service for data (assets, transactions, etc)
	tableName string,
	loggedInUser *string,
	dataFilter map[string]sql_query.SQLCondition,
	groupIdField string,
	transactionType []new_model.TransactionRoleAttribute,
	roleType []new_model.TransactionRoleType,
) error {
	filter := map[string]sql_query.SQLCondition{
		"utr.user_id": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    *loggedInUser,
		},
		"utr.role_type": {
			Operator: sql_query.SQLOperatorIn,
			Value:    roleType,
		},
		"utr.role_attribute": {
			Operator: sql_query.SQLOperatorIn,
			Value:    transactionType,
		},
		"utr.is_active": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    true,
		},
	}

	builder := sql_query.NewSQLSelectBuilder[any](tableName)

	query, args, _ := builder.
		Select(tableName+".id", "utr.id as \"userTransactionRoleId\"").
		LeftJoin(fmt.Sprintf("%s as utr", db.UserTransactionRoleTableName), fmt.Sprintf("%s = utr.group_id", groupIdField), filter).
		Where(dataFilter).
		Build()

	permissionCheck := make([]map[string]any, 0)
	err := service.SelectMany(&permissionCheck, context.Background(), query, args...)
	if err != nil {
		fmt.Printf("ERR: %v\n", err.Error())
		return errors.New("Data not found")
	}
	if len(permissionCheck) < 1 {
		fmt.Printf("PERMISSION: %v\n", len(permissionCheck))
		return errors.New("Data not found")
	}

	for _, data := range permissionCheck {
		transactionRoles := data["userTransactionRoleId"]
		if transactionRoles == nil {
			// If ONE OF data have no permission, throw error
			return errors.New("You have no permission")
		}
	}

	return nil
}

// For checking transaction role permisson
// Group filter is sql_query.SQLCondition so it can be bson.M{"$eq":group} or bson.M{"$in":groups}
// CANNOT HANDLE INPUT ARRAY WITH POTENTIALLY DIFFERENT GROUPS FOR EACH ITEM
// (e.g Have role in Group 1, but no role in Group 2: It will still pass)
func CheckTransactionRolePermissionSQL(
	ctx *fiber.Ctx,
	groupIdFilter sql_query.SQLCondition,
	transactionType []new_model.TransactionRoleAttribute,
	roleType []new_model.TransactionRoleType,
) error {
	// Get userId from Locals in Fiber context
	userId, ok := ctx.Locals(middleware.UserKey).(string)
	if !ok {
		return entity.Unauthorized("User not authenticated")
	}

	companyCode := ctx.Locals(middleware.CompanyCodeKey).(string)

	fixedAssetDbName := db.FixedAssetDBName.CompanyDBName(companyCode)
	service := service.MakePostgreSqlService(fixedAssetDbName, db.UserTransactionRoleTableName)

	filter := map[string]sql_query.SQLCondition{
		"user_id": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    userId,
		},
		"role_type": {
			Operator: sql_query.SQLOperatorIn,
			Value:    roleType,
		},
		"role_attribute": {
			Operator: sql_query.SQLOperatorIn,
			Value:    transactionType,
		},
		"group_id": groupIdFilter,
	}

	count, err := service.CountWithFilter(ctx.Context(), filter)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	if count < 1 {
		return entity.Unauthorized("You do not have the required permissions")
	}

	return nil
}

func CheckTotalControlPermissionSQL(ctx *fiber.Ctx) *entity.HttpError {
	// Get userId from Locals in Fiber context
	userId, ok := ctx.Locals(middleware.UserKey).(string)
	if !ok {
		return entity.Unauthorized("User not authenticated")
	}

	companyCode := ctx.Locals(middleware.CompanyCodeKey).(string)

	fixedAssetDbName := db.FixedAssetDBName.CompanyDBName(companyCode)
	service := service.MakePostgreSqlService(fixedAssetDbName, db.UserTableName)

	filter := map[string]sql_query.SQLCondition{
		"id": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    userId,
		},
		"total_control_is_active": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    true,
		},
		"total_control_is_assigned": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    true,
		},
	}

	count, err := service.CountWithFilter(ctx.Context(), filter)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	if count < 1 {
		return entity.Unauthorized("You do not have the required permissions")
	}

	return nil
}
