// Package query_builder builds the SQL that materializes a user's fixed-asset RBAC
// snapshot (middleware.UserAccessCache) from the iam database. It is run once at
// login (and on cache-miss) and its result is cached in Valkey; downstream services
// read the cache, never this query.
package query_builder

import (
	"fmt"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/db"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/sql_query"
)

// AuthUserRolesQueryBuilder builds a single-row query producing the JSON columns that
// unmarshal into middleware.UserAccessCache. All role tables live in the iam DB,
// so this is a plain single-database query (unlike the cross-service read path).
//
// Shape parity with BETS-V1 is deliberate:
//   - moduleAccess.fixedAsset -> {isTotalControl, isReadOnly} from the users row.
//   - systemRole.globalRole   -> per system-role key, booleans (category = 'Global').
//   - systemRole.groupRoles   -> per system-role key, group_id lists (category = 'Group').
//   - transactionRole          -> per transaction, per role type, group_id lists.
func AuthUserRolesQueryBuilder(loggedInUserID string, clientID string) (string, []any, error) {
	rolesCTE := authUserRolesMainCTE(loggedInUserID, clientID)

	return sql_query.NewSQLSelectBuilder[any]("roles_cte").
		Select(
			`roles_cte.id as "_id"`,
			`roles_cte.full_name as "fullName"`,
			`roles_cte.email as "email"`,
			`roles_cte.client_id as "clientId"`,
		).
		SelectJSONAggregate(
			"moduleAccess",
			map[string]string{
				string(enum.ModuleAccessFixedAsset): "roles_cte.fixed_asset",
			},
			"",
			false,
		).
		SelectJSONAggregate(
			"systemRole",
			map[string]string{
				"globalRole": "roles_cte.global_role",
				"groupRoles": "roles_cte.group_role",
			},
			"",
			false,
		).
		Select(
			`roles_cte.transaction_role as "transactionRole"`,
		).
		WithCTEBuilder("roles_cte", rolesCTE.(*sql_query.SelectBuilder).SQLEloquentQuery).
		Build()
}

// authUserRolesGroupRoleCTE aggregates, per system-role key, the group_ids the user is
// allowed to act on for each CRUD permission (category = 'Group').
//
// user_system_roles is shared by every product, so the product filter is not optional:
// role keys collide across products (importData, tag, userAndRole), and without it a
// Global Settings grant would be read as the fixed-asset permission of the same name.
func authUserRolesGroupRoleCTE(loggedInUserID string) sql_query.Selector {
	return sql_query.NewSQLSelectBuilder[any](db.UserSystemRoleTableName, "fasr").
		Select(
			"fasr.key",
			`jsonb_build_object(
				'create', COALESCE(jsonb_agg(fasr.fixed_asset_group_id::text) FILTER (WHERE fasr.permission_create), '[]'::jsonb),
				'view',   COALESCE(jsonb_agg(fasr.fixed_asset_group_id::text) FILTER (WHERE fasr.permission_view),   '[]'::jsonb),
				'edit',   COALESCE(jsonb_agg(fasr.fixed_asset_group_id::text) FILTER (WHERE fasr.permission_update), '[]'::jsonb),
				'delete', COALESCE(jsonb_agg(fasr.fixed_asset_group_id::text) FILTER (WHERE fasr.permission_delete), '[]'::jsonb)
			) AS permissions`,
		).
		Where(sql_query.WhereQuery{
			"fasr.user_id":   {Operator: sql_query.SQLOperatorEqual, Value: loggedInUserID},
			"fasr.is_active": {Operator: sql_query.SQLOperatorEqual, Value: true},
			"fasr.product":   {Operator: sql_query.SQLOperatorEqual, Value: string(enum.ModuleAccessFixedAsset)},
			"fasr.category":  {Operator: sql_query.SQLOperatorEqual, Value: string(enum.SystemRolesCategoryGroup)},
		}).
		GroupBy("fasr.key")
}

// authUserRolesTransactionRoleCTE produces, per transaction, a map of role type -> group_ids
// for the user's active transaction roles (excluding Approver, matching BETS-V1).
func authUserRolesTransactionRoleCTE(loggedInUserID string) sql_query.Selector {
	innerCTE := sql_query.NewSQLSelectBuilder[any](db.FixedAssetUserTransactionRoleTableName, "fatr").
		Select(
			"fatr.transaction",
			"fatr.type",
			"jsonb_agg(fatr.group_id::text) AS group_ids",
		).
		Where(sql_query.WhereQuery{
			"fatr.user_id": {Operator: sql_query.SQLOperatorEqual, Value: loggedInUserID},
			"fatr.type": {
				Operator: sql_query.SQLOperatorNotEqual,
				Value:    string(enum.TransactionRolesTypeApprover),
			},
			"fatr.is_active": {Operator: sql_query.SQLOperatorEqual, Value: true},
		}).
		GroupBy("fatr.transaction, fatr.type")

	return sql_query.NewSQLSelectBuilder[any]("transaction_role_inner_cte tri").
		Select(
			"tri.transaction",
			"jsonb_object_agg(tri.type, tri.group_ids) AS types",
		).
		GroupBy("tri.transaction").
		WithCTEBuilder("transaction_role_inner_cte", innerCTE.(*sql_query.SelectBuilder).SQLEloquentQuery)
}

// authUserRolesMainCTE joins the user row with the global system-role rows and the two
// pre-aggregated CTEs (group + transaction), producing one row per user.
func authUserRolesMainCTE(loggedInUserID string, clientID string) sql_query.Selector {
	groupRoleCTE := authUserRolesGroupRoleCTE(loggedInUserID)
	txCTE := authUserRolesTransactionRoleCTE(loggedInUserID)

	return sql_query.NewSQLSelectBuilder[any](db.UserTableName, "u").
		Select(
			"u.id",
			"u.full_name",
			"u.email",
			fmt.Sprintf("'%s' as client_id", clientID),
		).
		SelectJSONAggregate(
			"fixed_asset",
			map[string]string{
				"isTotalControl": "u.is_total_control_fixed_asset",
				"isReadOnly":     "u.is_read_only_fixed_asset",
			},
			"",
			false,
		).
		SelectJSONObjectAggregation(
			"global_role",
			"fasr_global.key",
			map[string]string{
				"create": "fasr_global.permission_create",
				"view":   "fasr_global.permission_view",
				"edit":   "fasr_global.permission_update",
				"delete": "fasr_global.permission_delete",
			},
			// FILTER out the NULL key produced by the LEFT JOIN when the user has no
			// global roles; without it jsonb_object_agg errors on a NULL object key.
			"fasr_global.key IS NOT NULL",
		).
		Select(
			`COALESCE(jsonb_object_agg(group_role_cte.key, group_role_cte.permissions) FILTER (WHERE group_role_cte.key IS NOT NULL), '{}'::jsonb) AS group_role`,
			`COALESCE(jsonb_object_agg(transaction_role_cte.transaction, transaction_role_cte.types) FILTER (WHERE transaction_role_cte.transaction IS NOT NULL), '{}'::jsonb) AS transaction_role`,
		).
		LeftJoin(db.UserSystemRoleTableName+" fasr_global", "fasr_global.user_id = u.id", sql_query.WhereQuery{
			"fasr_global.category": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    string(enum.SystemRolesCategoryGlobal),
			},
			// Same reason as the group CTE: without the product filter, a Global
			// Settings grant lands in the fixed-asset globalRole map.
			"fasr_global.product": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    string(enum.ModuleAccessFixedAsset),
			},
			"fasr_global.is_active": {Operator: sql_query.SQLOperatorEqual, Value: true},
		}).
		LeftJoin("group_role_cte", "group_role_cte.key IS NOT NULL").
		LeftJoin("transaction_role_cte", "transaction_role_cte.transaction IS NOT NULL").
		Where(sql_query.WhereQuery{
			"u.id": {Operator: sql_query.SQLOperatorEqual, Value: loggedInUserID},
		}).
		GroupBy("u.id").
		WithCTEBuilder("group_role_cte", groupRoleCTE.(*sql_query.SelectBuilder).SQLEloquentQuery).
		WithCTEBuilder("transaction_role_cte", txCTE.(*sql_query.SelectBuilder).SQLEloquentQuery)
}
