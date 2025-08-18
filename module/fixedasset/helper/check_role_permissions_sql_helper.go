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

// CheckSystemRolePermissionSQL confirms if a user has a specific permission for a given group.
//
// This is a direct check for a single user, a single permission, and a single group. It's
// ideal for simple validation, like checking if the logged-in user can 'view' a dashboard.
// The function looks for an active system role entry that matches the user's ID, the specified
// permission key (e.g., "view"), and the role attribute (e.g., "registerAssets").
//
// If no matching permission is found, it returns an HTTP unauthorized error.
func CheckSystemRolePermissionSQL(ctx *fiber.Ctx, permissionKey new_model.PermissionType, resourceTypes []new_model.SystemRoleAttribute, groupFilter sql_query.SQLCondition) *entity.HttpError {
	// Get userId from Locals in Fiber context
	userId, ok := ctx.Locals(middleware.UserKey).(string)
	if !ok {
		return entity.Unauthorized("User not authenticated")
	}

	permissionKeyStr := string(permissionKey)
	companyCode := ctx.Locals(middleware.CompanyCodeKey).(string)

	fixedAssetDbName := db.FixedAssetDBName.CompanyDBName(companyCode)
	service := service.MakePostgreSqlService(fixedAssetDbName)

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
		count, err := service.CountWithFilter(ctx.Context(), db.UserSystemRoleTableName, filter)
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

// CheckGroupPermissionSQL verifies that a user has the required permissions for a set of records,
// even if those records belong to different groups.
//
// This helper is designed for bulk operations (like updating or deleting multiple items). It's
// critical for ensuring a user has permissions for *every single item* in the list, not just some.
//
// It works by joining the main data table with the user's system roles. For each record, it
// checks if there's a corresponding active permission. If even one record lacks a valid
// permission, the function fails, preventing partial or unauthorized updates.
//
// Parameters:
//   - service: PostgreSQL service used to execute the query.
//   - tableName: The name of the main table to check permissions against.
//   - loggedInUser: The ID of the currently logged-in user (pointer to string).
//   - dataFilter: Additional SQL filter conditions to apply to the query.
//   - groupIdField: The column in the main table that maps to the group ID.
//   - permissions: The set of permission types (e.g., view, edit, delete) the user must have.
//   - systemRoleTypes: The system role attributes that qualify the user for the action.
//
// Returns:
//   - error: Returns "Data not found" if no matching data exists, or
//     "You have no permission" if the user lacks access to at least one record.
//
// Example scenario:
//   - You have an "assets" table, where each asset belongs to a group via `assets.group_id`.
//   - A user tries to update several assets at once.
//   - `CheckGroupPermissionSQL` ensures the user has the correct role and update
//     permission for every asset's group before allowing the operation.
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

// CheckTransactionRoleMultiGroupPermissionsSQL verifies that a user has the required transaction
// roles for a set of records across multiple groups.
//
// This function is the transaction-specific version of `CheckGroupPermissionSQL`. It's used for
// actions like asset transfers or approvals, where a user's role might be limited to certain
// groups or transaction types.
//
// It performs a comprehensive check on all specified records, ensuring the user has the correct
// combination of `transactionType` and `roleType` for each item. If any item is found without
// a valid, active transaction role, the function returns an error, preventing the operation.
//
// Parameters:
//   - service: The PostgreSQL service used to execute the permission check query.
//   - tableName: Name of the target table (e.g., "assets") to check permissions against.
//   - loggedInUser: Pointer to the ID of the currently logged-in user.
//   - dataFilter: Additional SQL conditions applied to filter the target table rows
//     (e.g., filter by asset IDs).
//   - groupIdField: The name of the group ID column in the target table used for joining
//     with the `user_transaction_role` table.
//   - transactionType: List of allowed transaction role attributes
//     (e.g., TransactionRoleAttributeTransfer).
//   - roleType: List of allowed transaction role types
//     (e.g., TransactionRoleTypeManager).
//
// Returns:
//   - nil if the user has the required permission for all matching rows.
//   - error if the user is missing the required permission or no matching data is found.
//
// Example:
//
//	Logged in user should have role attribute of "transferRole" and role type of "Manager" to be able to process assets in `assetIds`.
//
//	err := util_fixedasset_helper.CheckTransactionRoleMultiGroupPermissionsSQL(
//		u.AssetService,
//		"assets",
//		&param.LoggedInUserId,
//		map[string]util_sql_query.SQLCondition{
//			"assets.id": {Operator: util_sql_query.SQLOperatorAny, Value: assetIds},
//		},
//		"assets.group_id",
//		[]util_fixedasset_new_model.TransactionRoleAttribute{
//			util_fixedasset_new_model.TransactionRoleAttributeTransfer,
//		},
//		[]util_fixedasset_new_model.TransactionRoleType{
//			util_fixedasset_new_model.TransactionRoleTypeManager,
//		},
//	)
//
//	if err != nil {
//		// Handle missing permission
//	}
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
		Join("groups g", "g.id = "+groupIdField).
		LeftJoin(fmt.Sprintf("%s as utr", db.UserTransactionRoleTableName), "g.ancestor_id = utr.group_id", filter).
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

// CheckTransactionRolePermissionSQL confirms if a user holds a specific transaction role
// within a designated group or set of groups.
//
// This is a straightforward permission check, much like `CheckSystemRolePermissionSQL`. It
// quickly verifies if the user has an active transaction role (e.g., "transfer manager")
// for the group(s) specified in the filter.
//
// **Important Note:** This helper should not be used for checking permissions on a list of
// items that might belong to different groups. For that, use `CheckTransactionRoleMultiGroupPermissionsSQL`
// to avoid false positives. This function will return `true` if the user has permission in
// *any* of the groups, which can be misleading for bulk actions.
//
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
	service := service.MakePostgreSqlService(fixedAssetDbName)

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
		"g.id": groupIdFilter,
	}

	query, args, _ := sql_query.NewSQLCountBuilder(db.UserTransactionRoleTableName).
		Join("groups g", "group_id = g.ancestor_id").
		Where(filter).
		Build()

	count, err := service.Count(ctx.Context(), query, args...)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	if count < 1 {
		return entity.Unauthorized("You do not have the required permissions")
	}

	return nil
}

// CheckTotalControlPermissionSQL confirms if a user has "total control" access.
//
// This is the highest level of permission check. It doesn't use groups or roles but directly
// queries the user's account to see if they are a designated "total control" user.
//
// This helper is for actions that bypass all other group-based permissions, such as
// system-wide configuration changes or data migration. It returns an unauthorized
// error if the user is not a total control.
func CheckTotalControlPermissionSQL(ctx *fiber.Ctx) *entity.HttpError {
	// Get userId from Locals in Fiber context
	userId, ok := ctx.Locals(middleware.UserKey).(string)
	if !ok {
		return entity.Unauthorized("User not authenticated")
	}

	companyCode := ctx.Locals(middleware.CompanyCodeKey).(string)

	fixedAssetDbName := db.FixedAssetDBName.CompanyDBName(companyCode)
	service := service.MakePostgreSqlService(fixedAssetDbName)

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

	count, err := service.CountWithFilter(ctx.Context(), db.UserTableName, filter)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	if count < 1 {
		return entity.Unauthorized("You do not have the required permissions")
	}

	return nil
}

// CheckApprovalExistance finds the top-level ancestor (level 1) of a given group and checks for an active "Approval" role.
//
// The function uses a recursive SQL query to traverse the group hierarchy upwards from the
// provided `groupId`. It identifies the highest ancestor as the group with a parent_id of 0.
// It then checks for an active "Approval" role that matches the specified `roleAttribute`
// on that single, top-level ancestor. It returns `true` if such a role is found, otherwise `false`.
//
// Example:
//
// The function is typically used to determine a transfer request's initial status.
// A common use case is to check for an approval role on the highest ancestor group,
// and set the transaction status accordingly:
//
//	hasApproval, err := helper.CheckApprovalExistance(u.FixedAssetService, ctx, groupId, util_fixedasset_new_model.TransactionRoleAttributeTransfer)
//	if err != nil { ... }
//
//	if hasApproval {
//		return "WaitingForApproval"
//	}
//
//	return "WaitingForHandover"
func CheckApprovalExistance(
	service service.PostgreSqlService,
	ctx context.Context,
	groupId int,
	roleAttribute new_model.TransactionRoleAttribute,
) (bool, error) {
	// Defined as anchor, starting point to recursively check the groups upwards
	groupCTE := sql_query.NewSQLSelectBuilder[any](db.GroupTableName).
		Select("id", "parent_id").
		Where(map[string]sql_query.SQLCondition{
			"id": {Operator: sql_query.SQLOperatorEqual, Value: groupId},
		})

	// Recursive step to climbs up, find the parent of current chain
	groupParentCTE := sql_query.NewSQLSelectBuilder[any](db.GroupTableName, "g").
		Select("g.id", "g.parent_id").
		Join("parent_groups pg", "g.id = pg.parent_id")

	// Write "WITH RECURSIVE" CTE to main query
	withRecursiveCTE := sql_query.NewSQLSelectBuilder[any](db.GroupTableName).
		UnionAll(
			groupCTE.(*sql_query.SelectBuilder).SQLEloquentQuery,
			groupParentCTE.(*sql_query.SelectBuilder).SQLEloquentQuery,
		)

	// Build final query, `select` basically just contain "hasApproval"
	query, args, _ := sql_query.NewSQLSelectBuilder[any]("parent_groups pg").
		Select(`COUNT(utr.id) > 0 as "hasApproval"`).
		WithRecursiveCTEBuilder("parent_groups", withRecursiveCTE.(*sql_query.SelectBuilder).SQLEloquentQuery).
		Join(db.UserTransactionRoleTableName+" utr", "utr.group_id = pg.id").
		Where(map[string]sql_query.SQLCondition{
			// Top level parent (level 1)
			"pg.parent_id":       {Operator: sql_query.SQLOperatorEqual, Value: 0},
			"utr.role_type":      {Operator: sql_query.SQLOperatorEqual, Value: "Approval"},
			"utr.role_attribute": {Operator: sql_query.SQLOperatorEqual, Value: roleAttribute},
			"utr.is_active":      {Operator: sql_query.SQLOperatorEqual, Value: true},
		}).
		Build()

	var result struct{ HasApproval bool }
	err := service.SelectOne(&result, ctx, query, args...)
	if err != nil {
		return false, err
	}
	fmt.Println("result", result)

	return result.HasApproval, nil
}
