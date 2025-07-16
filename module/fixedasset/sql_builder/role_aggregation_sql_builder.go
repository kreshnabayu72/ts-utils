package sql_builder

import (
	"fmt"

	fixedasset_model "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

type TransactionRoleRestriction struct {
	ManageType []fixedasset_model.TransactionManageType `json:"manageType" bson:"manageType"`
}

type SystemRoleRestriction struct {
	RoleType    fixedasset_model.RoleType   `json:"roleType"    bson:"roleType"`
	Permissions fixedasset_model.Permission `json:"permissions" bson:"permissions"`
}

type RoleRestrictionOptions struct {
	TransactionRole *TransactionRoleRestriction `json:"transactionRole" bson:"transactionRole"`
	SystemRole      *SystemRoleRestriction      `json:"systemRole"      bson:"systemRole"`
}

type UserSystemRole struct {
	ID       string                    `json:"id" column:"usr.id::text"`
	GroupID  string                    `json:"group_id" column:"usr.group_id"`
	RoleType fixedasset_model.RoleType `json:"role_attribute" column:"usr.role_attribute"`
	Create   bool                      `json:"create" column:"usr.permission_create"`
	View     bool                      `json:"view" column:"usr.permission_view"`
	Update   bool                      `json:"update" column:"usr.permission_update"`
	Delete   bool                      `json:"delete" column:"usr.permission_delete"`
}

type UserTransactionRole struct {
	ID                        string                                     `json:"id" column:"utr.id::text"`
	GroupID                   string                                     `json:"group_id" column:"utr.group_id"`
	RoleType                  fixedasset_model.TransactionRoleType       `json:"role_type" column:"utr.role_type"`
	TransactionGroupAttribute fixedasset_model.TransactionGroupAttribute `json:"role_attribute" column:"utr.role_attribute"`
}

func generateUserSystemRolesBuilder(
	loggedInUserId int,
	groupColumn string,
	isFieldArray bool,
	roleType fixedasset_model.RoleType,
	isGlobalRole ...bool,
) sql_query.SQLSelectChainBuilder {
	builder := sql_query.NewSQLSelectBuilder[UserSystemRole]("user_system_roles usr").
		Where(map[string]sql_query.SQLCondition{
			"usr\".\"user_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    loggedInUserId,
			},
			"usr\".\"is_active": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    true,
			},
			"usr\".\"role_attribute": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    roleType,
			},
		})

	if len(isGlobalRole) < 1 && isGlobalRole[0] && groupColumn != "" {
		if !isFieldArray {
			builder.Where(map[string]sql_query.SQLCondition{
				"usr.group_id": {
					Operator: sql_query.SQLOperatorEqual,
					Value:    groupColumn,
					IsRef:    true,
				},
			})

			return builder
		}

		builder.Where(map[string]sql_query.SQLCondition{
			"usr.group_id": {
				Operator: sql_query.SQLOperatorAny,
				Value:    groupColumn,
				IsRef:    true,
			},
		})
	}

	return builder
}

func generateUserTransactionRolesBuilder(
	loggedInUserId int,
	transactionGroupAttribute fixedasset_model.TransactionGroupAttribute,
) sql_query.SQLSelectChainBuilder {
	return sql_query.NewSQLSelectBuilder[UserTransactionRole]("user_transaction_roles utr").
		Where(map[string]sql_query.SQLCondition{
			"utr\".\"user_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    loggedInUserId,
			},
			"utr\".\"is_active": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    true,
			},
			"utr\".\"role_type": {
				Operator: sql_query.SQLOperatorIn,
				Value:    []string{"Monitoring", "Manager"},
			},
			"utr\".\"role_attribute": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    transactionGroupAttribute,
			},
		})
}

func getSystemRolePermission(permission string) string {
	return fmt.Sprintf("usr.%s", permission)
}

func GetHasSystemRoleSection(builder *sql_query.SelectBuilder, roleType fixedasset_model.RoleType) {
	rt, ok := fixedasset_model.RoleTypeMap[string(roleType)]
	if !ok {
		return
	}

	builder.SelectJSONAggregateFunc("hasSystemRole", func(b *sql_query.SelectBuilder) {
		b.SelectJSONAggregate(string(rt),
			map[string]string{
				"create": getSystemRolePermission("create"),
				"view":   getSystemRolePermission("view"),
				"update": getSystemRolePermission("update"),
				"delete": getSystemRolePermission("delete"),
			}, fmt.Sprintf("usr.role_attribute = $%d", len(builder.Args)+1), false,
		)
	})
	// builder.SelectJSONAggregateFunc("hasSystemRole", func(b *sql_query.SelectBuilder) {
	// 	// To make json below to be children of this alias
	// 	b.AggregationChild = true

	// 	b.SelectJSONAggregateFunc("userAndRole", func(sb *sql_query.SelectBuilder) {
	// 		sb.SelectJSONAggregate(map[string]string{
	// 			"create": getSystemRolePermission("create"),
	// 			"view":   getSystemRolePermission("view"),
	// 			"update": getSystemRolePermission("update"),
	// 			"delete": getSystemRolePermission("delete"),
	// 		}, false).
	// 			Where(map[string]sql_query.SQLCondition{
	// 				"usr\".\"role_attribute": {
	// 					Operator: sql_query.SQLOperatorEqual,
	// 					Value:    "userAndRole",
	// 				},
	// 			})
	// 	})
	// 	b.SelectJSONAggregateFunc("openAPI", func(sb *sql_query.SelectBuilder) {
	// 		sb.SelectJSONAggregate(map[string]string{
	// 			"create": getSystemRolePermission("create"),
	// 			"view":   getSystemRolePermission("view"),
	// 			"update": getSystemRolePermission("update"),
	// 			"delete": getSystemRolePermission("delete"),
	// 		}, false).
	// 			Where(map[string]sql_query.SQLCondition{
	// 				"usr\".\"role_attribute": {
	// 					Operator: sql_query.SQLOperatorEqual,
	// 					Value:    "openAPI",
	// 				},
	// 			})
	// 	})
	// })
}

func SystemRolePermissionAggregation(
	builder *sql_query.SelectBuilder,
	loggedInUserId int,
	groupColumn string,
	isFieldArray bool,
	roleType fixedasset_model.RoleType,
	isGlobalRole ...bool,
) {
	var globalRole bool
	if len(isGlobalRole) > 0 {
		globalRole = isGlobalRole[0]
	} else {
		globalRole = false // default
	}

	cteBuilder := generateUserSystemRolesBuilder(loggedInUserId, groupColumn, isFieldArray, roleType, globalRole)
	builder.WithCTEBuilder("user_system_roles_cte", cteBuilder)

	GetHasSystemRoleSection(builder, roleType)

	builder.LeftJoin("user_system_roles_cte usr", "true")
}

func GetHasTransactionRole(builder *sql_query.SelectBuilder, transactionGroupAttribute fixedasset_model.TransactionGroupAttribute) {
	builder.SelectBoolOr(fmt.Sprintf("utr.role_type = 'Manager' AND utr.role_attribute = $%d", len(builder.Args)+1), "isManager", transactionGroupAttribute)
	builder.SelectBoolOr(fmt.Sprintf("utr.role_type = 'Monitoring' AND utr.role_attribute = $%d", len(builder.Args)+1), "isMonitoring", transactionGroupAttribute)
}

func TransactionRolePermissionAggregation(
	builder *sql_query.SelectBuilder,
	loggedInUserId int,
	groupColumn string,
	transactionGroupAttribute fixedasset_model.TransactionGroupAttribute,
	filterOptions RoleRestrictionOptions,
	filterOutEmptyRole ...bool,
) {
	cteBuilder := generateUserTransactionRolesBuilder(loggedInUserId, transactionGroupAttribute)
	builder.WithCTEBuilder("user_transaction_roles_cte", cteBuilder)

	GetHasTransactionRole(builder, transactionGroupAttribute)

	builder.LeftJoin("user_transaction_roles_cte utr", fmt.Sprintf("utr.group_id = %s", groupColumn))
	builder.Where(map[string]sql_query.SQLCondition{
		"utr\".\"group_id": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    groupColumn,
			IsRef:    true,
		},
	})

	if len(filterOutEmptyRole) > 0 {
		builder.HavingClauses = append(builder.HavingClauses, "COUNT(utr.id) > 0")
	}
}
