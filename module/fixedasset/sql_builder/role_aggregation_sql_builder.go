package sql_builder

import (
	"context"
	"fmt"
	"log"
	"strings"
	"unicode"

	fixedasset_new_model "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

type TransactionRoleRestriction struct {
	ManageType []fixedasset_new_model.TransactionRoleType `json:"manageType"`
}

type SystemRoleRestriction struct {
	RoleAttribute fixedasset_new_model.SystemRoleAttribute `json:"roleAttribute"`
	Permissions   fixedasset_new_model.Permission          `json:"permissions"`
}

type RoleRestrictionOptions struct {
	TransactionRole *TransactionRoleRestriction `json:"transactionRole"`
	SystemRole      *SystemRoleRestriction      `json:"systemRole"`
}

type UserSystemRole struct {
	ID       string                                   `json:"_id"            column:"usr.id::text"`
	GroupID  string                                   `json:"group_id"       column:"usr.group_id"`
	RoleType fixedasset_new_model.SystemRoleAttribute `json:"role_attribute" column:"usr.role_attribute"`
	Create   bool                                     `json:"create"         column:"usr.permission_create"`
	View     bool                                     `json:"view"           column:"usr.permission_view"`
	Update   bool                                     `json:"update"         column:"usr.permission_update"`
	Delete   bool                                     `json:"delete"         column:"usr.permission_delete"`
}

type UserTransactionRole struct {
	ID                        string                                        `json:"_id"            column:"utr.id::text"`
	GroupID                   string                                        `json:"group_id"       column:"utr.group_id"`
	RoleType                  fixedasset_new_model.TransactionRoleType      `json:"role_type"      column:"utr.role_type"`
	TransactionGroupAttribute fixedasset_new_model.TransactionRoleAttribute `json:"role_attribute" column:"utr.role_attribute"`
}

func ToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 { // Don't add underscore before the first character if it's uppercase
				result.WriteByte('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// generateUserSystemRolesBuilder creates a query builder to find all active system roles
// for a specific user and role attribute, either globally or within a specific group.
//
// This helper is the core logic for checking system-level permissions. It constructs a
// SQL query that filters 'user_system_roles' by the user's ID, checks for an active
// status, and matches a given role attribute (e.g., 'assetAttribute').
//
// A key feature is the conditional `WhereOr` clause:
//
//   - It checks if the role is 'Global'. If so, it bypasses the group check,
//     granting permission to the user regardless of the resource's group.
//   - Otherwise, it ensures the user's role is specifically tied to the
//     group of the main resource being queried.
//
// The function returns a fully configured `SQLSelectChainBuilder` ready to be
// executed or used as a Common Table Expression (CTE).
func generateUserSystemRolesBuilder(
	loggedInUserId int,
	roleAttribute fixedasset_new_model.SystemRoleAttribute,
) sql_query.SQLSelectChainBuilder {
	builder := sql_query.NewSQLSelectBuilder[any]("user_system_roles usr").
		Where(
			map[string]sql_query.SQLCondition{
				"usr.user_id": {
					Operator: sql_query.SQLOperatorEqual,
					Value:    loggedInUserId,
				},
				"usr.is_active": {
					Operator: sql_query.SQLOperatorEqual,
					Value:    true,
				},
				"usr.role_attribute": {
					Operator: sql_query.SQLOperatorEqual,
					Value:    roleAttribute,
				},
			},
		)

	builder.Select(
		"usr.group_id",
		"bool_or(usr.permission_create) AS \"create\"",
		"bool_or(usr.permission_view)   AS \"view\"",
		"bool_or(usr.permission_update) AS \"update\"",
		"bool_or(usr.permission_delete) AS \"delete\"",
	)

	builder.GroupBy("usr.group_id")

	return builder
}

// generateUserTransactionRolesBuilder creates a query builder to find all active transaction
// roles for a specific user, transaction attribute, and a list of role types.
//
// This is the transaction-specific equivalent of `generateUserSystemRolesBuilder`. It's used
// to build the SQL logic for checking permissions for actions like transfers or disposals.
// The query filters for active roles belonging to the specified user, matches a
// transaction attribute (e.g., 'transferRole'), and checks against a provided list of
// role types (e.g., ['Manager', 'Staff']).
//
// It returns a `SQLSelectChainBuilder` that's typically used to get a list of group IDs
// where the user has the specified transaction roles. This list can then be
// used to restrict a main query to only those groups.
func generateUserTransactionRolesBuilder(
	loggedInUserId int,
	transactionGroupAttribute fixedasset_new_model.TransactionRoleAttribute,
	roleType []fixedasset_new_model.TransactionRoleType,
) sql_query.SQLSelectChainBuilder {
	return sql_query.NewSQLSelectBuilder[any]("user_transaction_roles utr").
		Select("utr.group_id").
		Where(map[string]sql_query.SQLCondition{
			"utr.user_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    loggedInUserId,
			},
			"utr.is_active": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    true,
			},
			"utr.role_type": {
				Operator: sql_query.SQLOperatorIn,
				Value:    roleType,
			},
			"utr.role_attribute": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    transactionGroupAttribute,
			},
		})
}

// SystemRolePermissionAggregation aggregates system-level role permissions by creating
// Common Table Expressions (CTEs) for each specified role attribute and performing a LEFT JOIN.
//
// This function dynamically generates a CTE for each `fixedasset_new_model.SystemRoleAttribute`
// provided in `roleAttributes`. Each CTE checks if the `loggedInUserId` has the specified
// system role and then joins it back to the main query builder. This approach efficiently
// denormalizes the user's permissions into the main query's result set.
//
// The join condition for the CTE is now conditional:
// - For group-specific roles, it joins on a matching ID from the main table (`%s.id = %s.id`).
// - For global roles, it performs a LEFT JOIN on `TRUE`, making the permissions available
// to the main query without a group-based filter.
//
// Parameters:
//   - builder: The main `sql_query.SelectBuilder` to which the CTEs and joins will be added.
//   - loggedInUserId: The ID of the currently authenticated user.
//   - groupColumn: The column name used to join the CTE with the main query (e.g., "id" from the main table).
//   - mainTable: The name of the main table being queried.
//   - aliasTable: The alias of the main table in the query.
//   - roleAttributes: A slice of `fixedasset_new_model.SystemRoleAttribute` to check for.
func SystemRolePermissionAggregation(
	builder *sql_query.SelectBuilder,
	loggedInUserId int,
	groupColumn string,
	mainTable string,
	aliasTable string,
	roleAttributes []fixedasset_new_model.SystemRoleAttribute,
) {
	for _, roleAttribute := range roleAttributes {
		cteName := fmt.Sprintf("has_system_role_%s", ToSnakeCase(string(roleAttribute)))

		cteBuilder := generateUserSystemRolesBuilder(
			loggedInUserId,
			roleAttribute,
		)

		builder.WithCTEBuilder(
			cteName,
			cteBuilder.(*sql_query.SelectBuilder).SQLEloquentQuery,
		)

		roleType := fixedasset_new_model.GetSystemRoleTypeFromAttribute(roleAttribute)
		if roleType == fixedasset_new_model.SystemRoleTypeGroup {
			builder.LeftJoin(cteName, fmt.Sprintf("%s.group_id = %s.group_id", aliasTable, cteName))
		} else {
			// dont need to filter by group.id if its global role
			builder.LeftJoin(cteName, "TRUE")
		}
	}
}

// GetHasTransactionRole adds boolean columns to a SQL query builder to check for a user's
// transaction-specific roles within a group. It's a specialized helper that appends
// a series of BOOL_OR expressions to the builder's SELECT clause.
//
// These expressions create new columns—"isManager", "isMonitoring", and "isStaff"—which will be
// TRUE if the user holds the corresponding role for the specified transaction attribute, and FALSE otherwise.
// This function is typically used within a subquery or Common Table Expression (CTE) to
// summarize a user's permissions cleanly before joining them to a main query.
func GetHasTransactionRole(
	builder *sql_query.SelectBuilder,
	transactionGroupAttribute fixedasset_new_model.TransactionRoleAttribute,
) {
	builder.SelectBoolOr(
		fmt.Sprintf("utr.role_type = 'Manager' AND utr.role_attribute = $%d", len(builder.Args)+1),
		"isManager",
		transactionGroupAttribute,
	)
	builder.SelectBoolOr(
		fmt.Sprintf(
			"utr.role_type = 'Monitoring' AND utr.role_attribute = $%d",
			len(builder.Args)+1,
		),
		"isMonitoring",
		transactionGroupAttribute,
	)
	builder.SelectBoolOr(
		fmt.Sprintf(
			"utr.role_type = 'Staff' AND utr.role_attribute = $%d",
			len(builder.Args)+1,
		),
		"isStaff",
		transactionGroupAttribute,
	)
	builder.SelectBoolOr(
		fmt.Sprintf(
			"utr.role_type = 'Approval' AND utr.role_attribute = $%d",
			len(builder.Args)+1,
		),
		"isApprover",
		transactionGroupAttribute,
	)
}

// TransactionRolePermissionAggregation constructs and attaches a Common Table Expression (CTE)
// containing a user's transaction-based permissions to a main SQL query. This function
// streamlines the process of filtering or enriching a query with permission data.
//
// The process is as follows:
//  1. It first calls `generateUserTransactionRolesBuilder` to create a subquery that
//     identifies all groups where the user has roles for a specific transaction type.
//  2. It then uses `GetHasTransactionRole` to add the "isManager" and "isMonitoring"
//     boolean columns to this subquery, summarizing the user's permissions per group.
//  3. The resulting subquery is registered as a CTE named "user_transaction_roles_cte".
//  4. Finally, it joins this CTE to the main query using `groupColumn` to link permissions
//     to the main dataset. The join type is determined by `filterOutEmptyRole`:
//     - LEFT JOIN (default): Keeps all records from the main query, with permission columns as
//     NULL where the user has no matching roles.
//     - INNER JOIN (if `filterOutEmptyRole` is true): Only includes records where the user
//     has at least one of the specified roles, effectively filtering the result set.
func TransactionRolePermissionAggregation(
	builder *sql_query.SelectBuilder,
	loggedInUserId int,
	groupColumn string,
	transactionGroupAttribute fixedasset_new_model.TransactionRoleAttribute,
	filterOptions RoleRestrictionOptions,
	filterOutEmptyRole ...bool,
) {
	cteBuilder := generateUserTransactionRolesBuilder(loggedInUserId, transactionGroupAttribute, filterOptions.TransactionRole.ManageType)
	GetHasTransactionRole(cteBuilder.(*sql_query.SelectBuilder), transactionGroupAttribute)
	cteBuilder.GroupBy("utr.group_id")
	builder.WithCTEBuilder("user_transaction_roles_cte", cteBuilder.(*sql_query.SelectBuilder).SQLEloquentQuery)
	builder.Select(
		`utr."isManager"`,
		`utr."isMonitoring"`,
		`utr."isStaff"`,
		`utr."isApprover"`,
	)

	joinCteKeyword := builder.LeftJoin
	if len(filterOutEmptyRole) > 0 && filterOutEmptyRole[0] {
		joinCteKeyword = builder.Join
	}

	builder.Join("groups gp", "gp.id = "+groupColumn)
	joinCteKeyword(
		"user_transaction_roles_cte utr",
		"utr.group_id = gp.ancestor_id",
	)
}

// getTransactionRolePermission generates a SQL fragment using BOOL_OR to check for a specific
// transaction role attribute and type. This is a helper function used within a CTE builder
// to aggregate permissions efficiently. The output is a string like:
// `BOOL_OR(utr.role_type = 'Manager' AND utr.role_attribute = 'Borrowing')`.
//
// Parameters:
//   - transactionGroupAttribute: The transaction attribute (e.g., "Borrowing", "Transfer").
//   - roleType: The type of role (e.g., "Manager", "Staff", "Approver").
//
// Returns:
//   - A string containing the SQL fragment.
func getTransactionRolePermission(
	transactionGroupAttribute fixedasset_new_model.TransactionRoleAttribute,
	roleType fixedasset_new_model.TransactionRoleType,
) string {
	return fmt.Sprintf(
		"BOOL_OR(utr.role_type = '%s' AND utr.role_attribute = '%s')",
		roleType,
		transactionGroupAttribute,
	)
}

// transactionRoleCTEBuilder creates a Common Table Expression (CTE) builder that aggregates a user's
// transaction-level permissions into a single JSON object. This provides an efficient way to fetch all
// transaction roles and their permissions for a user in a single database round-trip.
//
// The function builds a subquery on the `user_transaction_roles` table, filtered by `loggedInUserId`
// and `is_active`. It then uses a series of `SelectJSONAggregate` calls to create a nested JSON
// structure. The outer JSON object, aliased as "transactionRoles", contains keys for each transaction
// attribute (e.g., "borrowing", "transfer"). Each of these keys points to an inner JSON object that
// holds boolean values for the user's `manager`, `monitoring`, `staff`, and `approver` permissions
// for that specific attribute. The results are grouped by `group_id` to provide permissions on a
// per-group basis.
//
// Parameters:
//   - loggedInUserId: The ID of the currently authenticated user.
//
// Returns:
//   - A `sql_query.SQLSelectChainBuilder` object that can be used to construct the CTE.
func transactionRoleCTEBuilder(
	loggedInUserId int,
) sql_query.SQLSelectChainBuilder {
	return sql_query.NewSQLSelectBuilder[any](
		"user_transaction_roles utr",
	).
		Where(map[string]sql_query.SQLCondition{
			"utr.user_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    loggedInUserId,
			},
			"utr.is_active": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    true,
			},
		}).
		Select("utr.group_id").
		SelectJSONAggregateFunc("transactionRoles",
			func(b *sql_query.SelectBuilder) {
				b.SelectJSONAggregate(
					"borrowing",
					map[string]string{
						"manager": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeBorrowing,
							fixedasset_new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeBorrowing,
							fixedasset_new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeBorrowing,
							fixedasset_new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeBorrowing,
							fixedasset_new_model.TransactionRoleTypeApproval,
						),
					},
					"true",
					false,
				)
				b.SelectJSONAggregate(
					"transfer",
					map[string]string{
						"manager": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeTransfer,
							fixedasset_new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeTransfer,
							fixedasset_new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeTransfer,
							fixedasset_new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeTransfer,
							fixedasset_new_model.TransactionRoleTypeApproval,
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"disposal",
					map[string]string{
						"manager": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeDisposal,
							fixedasset_new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeDisposal,
							fixedasset_new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeDisposal,
							fixedasset_new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeDisposal,
							fixedasset_new_model.TransactionRoleTypeApproval,
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"tracking",
					map[string]string{
						"manager": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeTracking,
							fixedasset_new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeTracking,
							fixedasset_new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeTracking,
							fixedasset_new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeTracking,
							fixedasset_new_model.TransactionRoleTypeApproval,
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"repairTikecting",
					map[string]string{
						"manager": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeRepair,
							fixedasset_new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeRepair,
							fixedasset_new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeRepair,
							fixedasset_new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeRepair,
							fixedasset_new_model.TransactionRoleTypeApproval,
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"maintenance",
					map[string]string{
						"manager": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeMaintenance,
							fixedasset_new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeMaintenance,
							fixedasset_new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeMaintenance,
							fixedasset_new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeMaintenance,
							fixedasset_new_model.TransactionRoleTypeApproval,
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"audit",
					map[string]string{
						"manager": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeAudit,
							fixedasset_new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeAudit,
							fixedasset_new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeAudit,
							fixedasset_new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							fixedasset_new_model.TransactionRoleAttributeAudit,
							fixedasset_new_model.TransactionRoleTypeApproval,
						),
					},
					"",
					false,
				)
			},
		).
		GroupBy("utr.group_id")
}

// getSystemRolePermissions generates a SQL fragment using BOOL_OR to check for a specific
// system role permission (e.g., 'create', 'view') and role attribute (e.g., 'importData').
// This helper function is used to create a concise and readable SQL expression for
// aggregating permissions. The output is a string like:
// `BOOL_OR(usr.role_attribute = 'importData' AND usr.permission_create = true)`
// This makes the CTE builder code cleaner and easier to maintain.
func getSystemRolePermissions(
	roleType fixedasset_new_model.SystemRoleAttribute,
	permission string,
) string {
	return fmt.Sprintf(
		"BOOL_OR(usr.role_attribute = '%s' AND usr.permission_%s = true)",
		roleType,
		permission,
	)
}

// systemRoleCTEBuilder creates a Common Table Expression (CTE) builder that aggregates a user's
// system-level permissions into a single JSON object. This is an efficient way to fetch all
// system roles and their permissions for a user in a single database round-trip.
//
// The function builds a subquery on the `user_system_roles` table, filtered by `loggedInUserId`
// and `is_active`. It then uses a series of `SelectJSONAggregate` calls to create a nested
// JSON structure. The outer JSON object, aliased as "systemRoles", contains keys for each
// system attribute (e.g., "importData", "assetAttribute"). Each of these keys points to an
// inner JSON object that holds boolean values for the user's `create`, `view`, `update`, and `delete`
// permissions for that specific attribute. The results are grouped by `group_id` to provide
// permissions on a per-group basis.
func systemRoleCTEBuilder(
	loggedInUserId int,
) sql_query.SQLSelectChainBuilder {
	return sql_query.NewSQLSelectBuilder[UserSystemRole]("user_system_roles usr").
		Where(map[string]sql_query.SQLCondition{
			"usr.user_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    loggedInUserId,
			},
			"usr.is_active": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    true,
			},
		}).
		Select("usr.group_id").
		SelectJSONAggregateFunc("systemRoles",
			func(b *sql_query.SelectBuilder) {
				b.SelectJSONAggregate(
					"importData",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeImportData,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeImportData,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeImportData,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeImportData,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"assetAttribute",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeAssetAttribute,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeAssetAttribute,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeAssetAttribute,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeAssetAttribute,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"assetPolicies",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeAssetPolicies,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeAssetPolicies,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeAssetPolicies,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeAssetPolicies,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"depreciationGroup",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeDepreciationGroup,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeDepreciationGroup,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeDepreciationGroup,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeDepreciationGroup,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"depreciationMethod",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeDepreciationMethod,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeDepreciationMethod,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeDepreciationMethod,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeDepreciationMethod,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"documentDeletion",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeDocumentDeletion,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeDocumentDeletion,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeDocumentDeletion,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeDocumentDeletion,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"group",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeGroup,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeGroup,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeGroup,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeGroup,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"iotReader",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeIoTReader,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeIoTReader,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeIoTReader,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeIoTReader,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"purchaseDocument",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributePurchaseDocument,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributePurchaseDocument,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributePurchaseDocument,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributePurchaseDocument,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"purchaseInformation",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributePurchaseInfo,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributePurchaseInfo,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributePurchaseInfo,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributePurchaseInfo,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"userAndRole",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeUserAndRole,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeUserAndRole,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeUserAndRole,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeUserAndRole,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"tag",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeTag,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeTag,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeTag,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeTag,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"registerAsset",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeRegisterAsset,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeRegisterAsset,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeRegisterAsset,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeRegisterAsset,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"tagType",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeTagType,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeTagType,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeTagType,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeTagType,
							"delete",
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"openApi",
					map[string]string{
						"create": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeOpenAPI,
							"create",
						),
						"view": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeOpenAPI,
							"view",
						),
						"update": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeOpenAPI,
							"update",
						),
						"delete": getSystemRolePermissions(
							fixedasset_new_model.SystemRoleAttributeOpenAPI,
							"delete",
						),
					},
					"",
					false,
				)
			},
		).
		GroupBy("usr.group_id")
}

// RolePermissionAggregation is a comprehensive function that aggregates both system and
// transaction-level permissions for a given user and attaches them to a main query as
// a single, unified JSON object. This is a powerful, high-level helper for fetching all
// permission data relevant to a specific user and group.
//
// The function operates in three main steps:
//
//  1. Transaction Role CTE: It calls `transactionRoleCTEBuilder` (assumed to be a helper function)
//     to create a CTE named "user_transaction_roles_cte". This CTE summarizes the user's
//     transaction roles. It then LEFT JOINs this CTE to the main query using `groupColumn`.
//
//  2. System Role CTE: It calls `systemRoleCTEBuilder` to create a CTE named
//     "user_system_roles_cte". This CTE summarizes the user's system roles and their
//     respective permissions. It then LEFT JOINs this CTE to the main query using `groupColumn`.
//
//  3. Final Aggregation: It uses `SelectJSONAggregate` on the main query to combine the
//     "systemRoles" and "transactionRoles" JSON objects from the two CTEs into a single
//     "permissions" JSON object. This final output provides a complete picture of the user's
//     permissions, ready for use in the application logic.
func RolePermissionAggregation(
	builder *sql_query.SelectBuilder,
	loggedInUserId int,
	groupColumn string,
) {
	transactionRoleBuilder := transactionRoleCTEBuilder(loggedInUserId)
	builder.WithCTEBuilder("user_transaction_roles_cte", transactionRoleBuilder.(*sql_query.SelectBuilder).SQLEloquentQuery)
	builder.Join("groups gp", "gp.id = "+groupColumn)
	builder.LeftJoin(
		"user_transaction_roles_cte utr",
		"utr.group_id = gp.ancestor_id",
	)

	systemRolesBuilder := systemRoleCTEBuilder(loggedInUserId)
	builder.WithCTEBuilder("user_system_roles_cte", systemRolesBuilder.(*sql_query.SelectBuilder).SQLEloquentQuery)
	builder.LeftJoin(
		"user_system_roles_cte usr",
		fmt.Sprintf("usr.group_id = %s", groupColumn),
	)

	builder.SelectJSONAggregate(
		"permissions",
		map[string]string{
			"transactionRoles": "utr.transactionRoles",
			"systemRoles":      "usr.systemRoles",
		},
		"",
		false,
	)
}

// SystemRoleFilter is a struct for filtering user system roles.
type SystemRoleFilter struct {
	RoleAttribute    fixedasset_new_model.SystemRoleAttribute `json:"roleAttribute" column:"role_attribute"` // Enum: system_role_attribute_enum
	PermissionView   bool                                     `json:"permissionView" column:"permission_view"`
	PermissionCreate bool                                     `json:"permissionCreate" column:"permission_create"`
	PermissionUpdate bool                                     `json:"permissionUpdate" column:"permission_update"`
	PermissionDelete bool                                     `json:"permissionDelete" column:"permission_delete"`
	GroupID          *int                                     `json:"groupId" column:"group_id"`
}

// FindUserSystemRoles is a generic function to find user system roles based on a logged-in user ID
// and a set of filters. It dynamically builds a SQL query and executes it against the database.
func FindUserSystemRoles[T any](
	v *[]T,
	service service.PostgreSqlService,
	ctx context.Context,
	loggedInUserId int,
	roleFilter ...SystemRoleFilter,
) error {
	var err error

	systemRolesBuilder := sql_query.NewSQLSelectBuilder[T]("user_system_roles")

	filter := map[string]sql_query.SQLCondition{
		"user_id": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    loggedInUserId,
		},
		"is_active": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    true,
		},
	}

	orFilter := []map[string]sql_query.SQLCondition{}
	for _, each := range roleFilter {
		filterPart := map[string]sql_query.SQLCondition{}
		if each.RoleAttribute != "" {
			filterPart["role_attribute"] = sql_query.SQLCondition{Operator: sql_query.SQLOperatorEqual, Value: each.RoleAttribute}
		}
		if each.GroupID != nil {
			filterPart["group_id"] = sql_query.SQLCondition{Operator: sql_query.SQLOperatorEqual, Value: each.GroupID}
		}
		if each.PermissionCreate {
			filterPart["permission_create"] = sql_query.SQLCondition{Operator: sql_query.SQLOperatorEqual, Value: each.PermissionCreate}
		}
		if each.PermissionView {
			filterPart["permission_view"] = sql_query.SQLCondition{Operator: sql_query.SQLOperatorEqual, Value: each.PermissionView}
		}
		if each.PermissionUpdate {
			filterPart["permission_update"] = sql_query.SQLCondition{Operator: sql_query.SQLOperatorEqual, Value: each.PermissionUpdate}
		}
		if each.PermissionDelete {
			filterPart["permission_delete"] = sql_query.SQLCondition{Operator: sql_query.SQLOperatorEqual, Value: each.PermissionDelete}
		}
		orFilter = append(orFilter, filterPart)
	}
	if len(orFilter) > 0 {
		systemRolesBuilder.WhereOr(orFilter...)
	}

	queryString, args, _ := systemRolesBuilder.
		Where(filter).
		Build()

	err = service.SelectMany(v, ctx, queryString, args...)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// TransactionRoleFilter is a struct for filtering user transaction roles.
type TransactionRoleFilter struct {
	RoleAttribute fixedasset_new_model.TransactionRoleAttribute `json:"roleAttribute" column:"role_attribute"`
	RoleType      fixedasset_new_model.TransactionRoleType      `json:"roleType" column:"role_type"`

	// Use a pointer to allow for a nil value, indicating no group-based filtering
	GroupID *int `json:"groupId" column:"group_id"`
}

// FindUserTransactionRoles is a generic function to find user transaction roles.
// This function finds a user's transaction roles by building a SQL query filtered by the logged-in user's ID and a set of role filters.
func FindUserTransactionRoles[T any](
	v *[]T,
	service service.PostgreSqlService,
	ctx context.Context,
	loggedInUserId int,
	roleFilter ...TransactionRoleFilter,
) error {
	var err error

	transactionRoleBuilder := sql_query.NewSQLSelectBuilder[T]("user_transaction_roles")
	transactionRoleBuilder.LeftJoin("groups gp", "gp.ancestor_id = user_transaction_roles.group_id")

	filter := map[string]sql_query.SQLCondition{
		"user_id": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    loggedInUserId,
		},
		"is_active": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    true,
		},
	}

	orFilter := []map[string]sql_query.SQLCondition{}
	for _, each := range roleFilter {
		filterPart := map[string]sql_query.SQLCondition{}
		if each.RoleAttribute != "" {
			filterPart["role_attribute"] = sql_query.SQLCondition{Operator: sql_query.SQLOperatorEqual, Value: each.RoleAttribute}
		}
		if each.RoleType != "" {
			filterPart["role_type"] = sql_query.SQLCondition{Operator: sql_query.SQLOperatorEqual, Value: each.RoleType}
		}
		if each.GroupID != nil {
			filterPart["gp.id"] = sql_query.SQLCondition{Operator: sql_query.SQLOperatorEqual, Value: each.GroupID}
		}
		orFilter = append(orFilter, filterPart)
	}
	if len(orFilter) > 0 {
		transactionRoleBuilder.WhereOr(orFilter...)
	}

	queryString, args, _ := transactionRoleBuilder.
		Where(filter).
		Build()

	err = service.SelectMany(v, ctx, queryString, args...)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// GroupFieldAndAttributeName is a struct that defines a SQL column and its corresponding
// attribute name for transaction role checks.
//
// Field:         The name of the column representing a group ID in the main table (e.g., "origin_group.id").
// AttributeName: A human-readable name used to generate the output column alias (e.g., "Origin" will result in "isOriginManager").
type GroupFieldAndAttributeName struct {
	Field         string // origin_group.id / destination_group.id
	AttributeName string // Origin / Destination (For isDestinationManager)
}

// TransactionRoleMultipleGroupAggregation dynamically adds a JOIN and SELECT clauses to a query
// builder to check for transaction roles across multiple groups. This function is useful for
// queries that involve assets moving between groups, such as transfers.
//
// It performs a LEFT JOIN on the `user_transaction_roles` table, filtered by user ID,
// active status, and transaction attribute. It then iterates through the provided
// `groupFieldsAndAttribute` slice to add `BOOL_OR` expressions for each group field.
// This results in boolean columns like "isOriginManager", "isDestinationManager", etc.,
// which indicate if the user holds a specific role in each of the groups.
func TransactionRoleMultipleGroupAggregation(
	builder *sql_query.SelectBuilder,
	loggedInUserId int,
	groupFieldsAndAttribute []GroupFieldAndAttributeName,
	transactionGroupAttribute fixedasset_new_model.TransactionRoleAttribute,
	roleTypes []fixedasset_new_model.TransactionRoleType,
) {
	builder.
		LeftJoin(
			"user_transaction_roles utr",
			fmt.Sprintf(
				"utr.user_id = $%d AND utr.is_active = true AND utr.role_attribute = $%d",
				builder.GetCurrentArgIndex()+1,
				builder.GetCurrentArgIndex()+2,
			),
		)
	builder.AddArgs(loggedInUserId, transactionGroupAttribute)

	if len(roleTypes) > 0 {
		builder.Where(map[string]sql_query.SQLCondition{
			"utr.role_type": {
				Operator: sql_query.SQLOperatorIn,
				Value:    roleTypes,
				IsRef:    true,
			},
		})
	}

	for idx, groupField := range groupFieldsAndAttribute {
		builder.Join(fmt.Sprintf("groups gp%d", idx), fmt.Sprintf("gp%d.id = %s", idx, groupField.Field))
		builder.SelectBoolOr(
			fmt.Sprintf(
				"utr.group_id = gp%d.ancestor_id AND utr.role_type = $%d",
				idx,
				builder.GetCurrentArgIndex()+1,
			),
			fmt.Sprintf("is%sManager", groupField.AttributeName),
			fixedasset_new_model.TransactionRoleTypeManager,
		)
		builder.SelectBoolOr(
			fmt.Sprintf(
				"utr.group_id = gp%d.ancestor_id AND utr.role_type = $%d",
				idx,
				builder.GetCurrentArgIndex()+1,
			),
			fmt.Sprintf("is%sMonitoring", groupField.AttributeName),
			fixedasset_new_model.TransactionRoleTypeMonitoring,
		)
	}
}

// transactionRoleMultiPermissionAggregationCTE builds a Common Table Expression (CTE)
// that aggregates all relevant transaction permissions for a given user across multiple
// transaction types (Borrowing, Transfer, Disposal).
//
// This internal helper creates a subquery that selects the user's active 'Manager' and
// 'Monitoring' roles, then uses `BOOL_OR` to produce boolean columns for each role-transaction
// combination (e.g., "manager_borrowing", "monitoring_transfer"). The result is grouped by
// `group_id`, providing a concise summary of the user's permissions within each group.
// The returned builder can then be used to join this permission data to a main query.
func transactionRoleMultiPermissionAggregationCTE(
	loggedInUserId int,
) sql_query.SQLSelectChainBuilder {
	builder := sql_query.NewSQLSelectBuilder[any](
		"user_transaction_roles utr",
	).
		Where(map[string]sql_query.SQLCondition{
			"utr.user_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    loggedInUserId,
			},
			"utr.is_active": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    true,
			},
			"utr.role_type": {
				Operator: sql_query.SQLOperatorIn,
				Value:    []string{"Monitoring", "Manager"},
			},
			"utr.role_attribute": {
				Operator: sql_query.SQLOperatorIn,
				Value: []fixedasset_new_model.TransactionRoleAttribute{
					fixedasset_new_model.TransactionRoleAttributeBorrowing,
					fixedasset_new_model.TransactionRoleAttributeTransfer,
					fixedasset_new_model.TransactionRoleAttributeDisposal,
				},
			},
		})

	builder.Select(
		"utr.group_id",
		fmt.Sprintf(
			`BOOL_OR(utr.role_type = $%d AND utr.role_attribute = $%d) AS "manager_borrowing"`,
			builder.GetCurrentArgIndex()+1,
			builder.GetCurrentArgIndex()+2,
		),
		fmt.Sprintf(
			`BOOL_OR(utr.role_type = $%d AND utr.role_attribute = $%d) AS "manager_transfer"`,
			builder.GetCurrentArgIndex()+3,
			builder.GetCurrentArgIndex()+4,
		),
		fmt.Sprintf(
			`BOOL_OR(utr.role_type = $%d AND utr.role_attribute = $%d) AS "manager_disposal"`,
			builder.GetCurrentArgIndex()+5,
			builder.GetCurrentArgIndex()+6,
		),
		fmt.Sprintf(
			`BOOL_OR(utr.role_type = $%d AND utr.role_attribute = $%d) AS "monitoring_borrowing"`,
			builder.GetCurrentArgIndex()+7,
			builder.GetCurrentArgIndex()+8,
		),
		fmt.Sprintf(
			`BOOL_OR(utr.role_type = $%d AND utr.role_attribute = $%d) AS "monitoring_transfer"`,
			builder.GetCurrentArgIndex()+9,
			builder.GetCurrentArgIndex()+10,
		),
		fmt.Sprintf(
			`BOOL_OR(utr.role_type = $%d AND utr.role_attribute = $%d) AS "monitoring_disposal"`,
			builder.GetCurrentArgIndex()+11,
			builder.GetCurrentArgIndex()+12,
		),
	)

	builder.AddArgs(
		fixedasset_new_model.TransactionRoleTypeManager,
		fixedasset_new_model.TransactionRoleAttributeBorrowing,
		fixedasset_new_model.TransactionRoleTypeManager,
		fixedasset_new_model.TransactionRoleAttributeTransfer,
		fixedasset_new_model.TransactionRoleTypeManager,
		fixedasset_new_model.TransactionRoleAttributeDisposal,
		fixedasset_new_model.TransactionRoleTypeMonitoring,
		fixedasset_new_model.TransactionRoleAttributeBorrowing,
		fixedasset_new_model.TransactionRoleTypeMonitoring,
		fixedasset_new_model.TransactionRoleAttributeTransfer,
		fixedasset_new_model.TransactionRoleTypeMonitoring,
		fixedasset_new_model.TransactionRoleAttributeDisposal,
	)

	return builder.GroupBy("utr.group_id")
}

// TransactionRoleMultiPermissionAggregation constructs a `user_transaction_roles_cte`
// and attaches it to the main query using a LEFT JOIN. This is the top-level
// function for integrating multi-transaction permission data into a query.
//
// It first calls `transactionRoleMultiPermissionAggregationCTE` to get a builder for the
// permission-summary CTE. It then registers this CTE with the main query builder and
// performs a LEFT JOIN on the provided `groupColumn`, making the aggregated permission
// data available for filtering or selection in the main query.
func TransactionRoleMultiPermissionAggregation(
	builder *sql_query.SelectBuilder,
	loggedInUserId int,
	groupColumn string,
) {
	cteBuilder := transactionRoleMultiPermissionAggregationCTE(loggedInUserId)

	builder.WithCTEBuilder("user_transaction_roles_cte", cteBuilder.(*sql_query.SelectBuilder).SQLEloquentQuery)
	builder.Join("groups gp", "gp.id = "+groupColumn)
	builder.LeftJoin(
		"user_transaction_roles_cte utr",
		"utr.group_id = gp.ancestor_id",
	)
}

// RoleRestrictionAggregation restricts query results based on a user's system and transaction roles.
// It dynamically adds JOIN and WHERE clauses to implement row-level security. This function
// serves as a flexible, top-level helper for implementing row-level security.
//
// It checks the `RoleRestrictionOptions` to determine which type of role-based filtering to apply:
//
// 1. Transaction Role Filtering:
// If `filterOptions.TransactionRole` is not nil, the function adds an INNER JOIN to the
// `user_transaction_roles` table. This join is conditioned on the `loggedInUserId`,
// an active role, and optionally a `groupColumn`. It then adds a WHERE clause to filter
// for specific transaction role types (e.g., 'Manager', 'Monitoring') defined in
// `filterOptions.TransactionRole.ManageType`.
//
// 2. System Role Filtering:
// If `filterOptions.SystemRole` is not nil, the function adds an INNER JOIN to the
// `user_system_roles` table. This join is conditioned on the `loggedInUserId`, an active
// role, a specific `role_attribute`, and optionally a `groupColumn`. It then adds a
// WHERE_OR clause to filter for records where the user has at least one of the required
// permissions (create, view, update, delete) as specified in `filterOptions.SystemRole.Permissions`.
//
// The function can combine both transaction and system role filtering if both options are
// provided, creating a comprehensive permission-based restriction on the query.
func RoleRestrictionAggregation(
	builder *sql_query.SelectBuilder,
	loggedInUserId int,
	groupColumn string,
	filterOptions RoleRestrictionOptions,
) {
	if filterOptions.TransactionRole != nil {
		joinCondition := fmt.Sprintf(
			"utr.user_id = $%d AND utr.is_active = true",
			builder.GetCurrentArgIndex()+1,
		)
		builder.AddArgs(loggedInUserId)
		if groupColumn != "" {
			joinCondition += fmt.Sprintf(" AND utr.group_id = %s", groupColumn)
		}

		builder.
			Join(
				"user_transaction_roles utr",
				joinCondition,
			).
			Where(map[string]sql_query.SQLCondition{
				"utr.role_type": {
					Operator: sql_query.SQLOperatorIn,
					Value:    filterOptions.TransactionRole.ManageType,
				},
			})
	}

	if filterOptions.SystemRole != nil {
		joinCondition := fmt.Sprintf(
			"usr.user_id = $%d AND usr.is_active = true AND usr.role_attribute = $%d",
			builder.GetCurrentArgIndex()+1,
			builder.GetCurrentArgIndex()+2,
		)
		builder.AddArgs(
			loggedInUserId,
			filterOptions.SystemRole.RoleAttribute,
		)
		if groupColumn != "" {
			joinCondition += fmt.Sprintf(" AND usr.group_id = %s", groupColumn)
		}

		builder.
			Join(
				"user_system_roles usr",
				joinCondition,
			).
			WhereOr(
				map[string]sql_query.SQLCondition{
					"usr.permission_create": {
						Operator: sql_query.SQLOperatorEqual,
						Value:    filterOptions.SystemRole.Permissions.Create,
					},
					"usr.permission_view": {
						Operator: sql_query.SQLOperatorEqual,
						Value:    filterOptions.SystemRole.Permissions.View,
					},
					"usr.permission_update": {
						Operator: sql_query.SQLOperatorEqual,
						Value:    filterOptions.SystemRole.Permissions.Update,
					},
					"usr.permission_delete": {
						Operator: sql_query.SQLOperatorEqual,
						Value:    filterOptions.SystemRole.Permissions.Delete,
					},
				},
			)
	}
}
