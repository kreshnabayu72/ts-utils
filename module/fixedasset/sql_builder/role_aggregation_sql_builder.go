package sql_builder

import (
	"context"
	"fmt"
	"log"

	fixedasset_model "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
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
	ID       string                    `json:"_id" column:"usr.id::text"`
	GroupID  string                    `json:"group_id" column:"usr.group_id"`
	RoleType fixedasset_model.RoleType `json:"role_attribute" column:"usr.role_attribute"`
	Create   bool                      `json:"create" column:"usr.permission_create"`
	View     bool                      `json:"view" column:"usr.permission_view"`
	Update   bool                      `json:"update" column:"usr.permission_update"`
	Delete   bool                      `json:"delete" column:"usr.permission_delete"`
}

type UserTransactionRole struct {
	ID                        string                                     `json:"_id" column:"utr.id::text"`
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

	cteBuilder := generateUserSystemRolesBuilder(
		loggedInUserId,
		groupColumn,
		isFieldArray,
		roleType,
		globalRole,
	)
	builder.WithCTEBuilder("user_system_roles_cte", cteBuilder.(*sql_query.SelectBuilder).SQLEloquentQuery)

	GetHasSystemRoleSection(builder, roleType)

	builder.LeftJoin("user_system_roles_cte usr", "true")
}

func GetHasTransactionRole(
	builder *sql_query.SelectBuilder,
	transactionGroupAttribute fixedasset_model.TransactionGroupAttribute,
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
	builder.WithCTEBuilder("user_transaction_roles_cte", cteBuilder.(*sql_query.SelectBuilder).SQLEloquentQuery)

	GetHasTransactionRole(builder, transactionGroupAttribute)

	builder.LeftJoin(
		"user_transaction_roles_cte utr",
		fmt.Sprintf("utr.group_id = %s", groupColumn),
	)
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

func getTransactionRolePermission(
	transactionGroupAttribute new_model.TransactionRoleAttribute,
	roleType new_model.TransactionRoleType,
) string {
	return fmt.Sprintf(
		"BOOL_OR(utr.role_type = '%s' AND utr.role_attribute = '%s')",
		roleType,
		transactionGroupAttribute,
	)
}

func transactionRoleCTEBuilder(
	loggedInUserId int,
) sql_query.SQLSelectChainBuilder {
	return sql_query.NewSQLSelectBuilder[UserTransactionRole](
		"user_transaction_roles utr",
	).
		Where(map[string]sql_query.SQLCondition{
			"utr\".\"user_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    loggedInUserId,
			},
			"utr\".\"is_active": {
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
							new_model.TransactionRoleAttributeBorrowing,
							new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							new_model.TransactionRoleAttributeBorrowing,
							new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							new_model.TransactionRoleAttributeBorrowing,
							new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							new_model.TransactionRoleAttributeBorrowing,
							new_model.TransactionRoleTypeApproval,
						),
					},
					"true",
					false,
				)
				b.SelectJSONAggregate(
					"transfer",
					map[string]string{
						"manager": getTransactionRolePermission(
							new_model.TransactionRoleAttributeTransfer,
							new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							new_model.TransactionRoleAttributeTransfer,
							new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							new_model.TransactionRoleAttributeTransfer,
							new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							new_model.TransactionRoleAttributeTransfer,
							new_model.TransactionRoleTypeApproval,
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"disposal",
					map[string]string{
						"manager": getTransactionRolePermission(
							new_model.TransactionRoleAttributeDisposal,
							new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							new_model.TransactionRoleAttributeDisposal,
							new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							new_model.TransactionRoleAttributeDisposal,
							new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							new_model.TransactionRoleAttributeDisposal,
							new_model.TransactionRoleTypeApproval,
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"tracking",
					map[string]string{
						"manager": getTransactionRolePermission(
							new_model.TransactionRoleAttributeTracking,
							new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							new_model.TransactionRoleAttributeTracking,
							new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							new_model.TransactionRoleAttributeTracking,
							new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							new_model.TransactionRoleAttributeTracking,
							new_model.TransactionRoleTypeApproval,
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"repairTikecting",
					map[string]string{
						"manager": getTransactionRolePermission(
							new_model.TransactionRoleAttributeRepair,
							new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							new_model.TransactionRoleAttributeRepair,
							new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							new_model.TransactionRoleAttributeRepair,
							new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							new_model.TransactionRoleAttributeRepair,
							new_model.TransactionRoleTypeApproval,
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"maintenance",
					map[string]string{
						"manager": getTransactionRolePermission(
							new_model.TransactionRoleAttributeMaintenance,
							new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							new_model.TransactionRoleAttributeMaintenance,
							new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							new_model.TransactionRoleAttributeMaintenance,
							new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							new_model.TransactionRoleAttributeMaintenance,
							new_model.TransactionRoleTypeApproval,
						),
					},
					"",
					false,
				)
				b.SelectJSONAggregate(
					"audit",
					map[string]string{
						"manager": getTransactionRolePermission(
							new_model.TransactionRoleAttributeAudit,
							new_model.TransactionRoleTypeManager,
						),
						"monitoring": getTransactionRolePermission(
							new_model.TransactionRoleAttributeAudit,
							new_model.TransactionRoleTypeMonitoring,
						),
						"staff": getTransactionRolePermission(
							new_model.TransactionRoleAttributeAudit,
							new_model.TransactionRoleTypeStaff,
						),
						"approver": getTransactionRolePermission(
							new_model.TransactionRoleAttributeAudit,
							new_model.TransactionRoleTypeApproval,
						),
					},
					"",
					false,
				)
			},
		).
		GroupBy("utr.group_id")
}

func getSystemRolePermissions(
	roleType new_model.SystemRoleAttribute,
	permission string,
) string {
	return fmt.Sprintf(
		"BOOL_OR(usr.role_attribute = '%s' AND usr.permission_%s = true)",
		roleType,
		permission,
	)
}

func systemRoleCTEBuilder(
	loggedInUserId int,
) sql_query.SQLSelectChainBuilder {
	return sql_query.NewSQLSelectBuilder[UserSystemRole]("user_system_roles usr").
		Where(map[string]sql_query.SQLCondition{
			"usr\".\"user_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    loggedInUserId,
			},
			"usr\".\"is_active": {
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
							new_model.SystemRoleAttributeImportData,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributeImportData,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributeImportData,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributeImportData,
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
							new_model.SystemRoleAttributeAssetAttribute,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributeAssetAttribute,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributeAssetAttribute,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributeAssetAttribute,
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
							new_model.SystemRoleAttributeAssetPolicies,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributeAssetPolicies,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributeAssetPolicies,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributeAssetPolicies,
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
							new_model.SystemRoleAttributeDepreciationGroup,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributeDepreciationGroup,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributeDepreciationGroup,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributeDepreciationGroup,
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
							new_model.SystemRoleAttributeDepreciationMethod,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributeDepreciationMethod,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributeDepreciationMethod,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributeDepreciationMethod,
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
							new_model.SystemRoleAttributeDocumentDeletion,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributeDocumentDeletion,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributeDocumentDeletion,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributeDocumentDeletion,
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
							new_model.SystemRoleAttributeGroup,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributeGroup,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributeGroup,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributeGroup,
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
							new_model.SystemRoleAttributeIoTReader,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributeIoTReader,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributeIoTReader,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributeIoTReader,
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
							new_model.SystemRoleAttributePurchaseDocument,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributePurchaseDocument,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributePurchaseDocument,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributePurchaseDocument,
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
							new_model.SystemRoleAttributePurchaseInfo,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributePurchaseInfo,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributePurchaseInfo,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributePurchaseInfo,
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
							new_model.SystemRoleAttributeUserAndRole,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributeUserAndRole,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributeUserAndRole,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributeUserAndRole,
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
							new_model.SystemRoleAttributeTag,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributeTag,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributeTag,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributeTag,
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
							new_model.SystemRoleAttributeRegisterAsset,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributeRegisterAsset,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributeRegisterAsset,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributeRegisterAsset,
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
							new_model.SystemRoleAttributeTagType,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributeTagType,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributeTagType,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributeTagType,
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
							new_model.SystemRoleAttributeOpenAPI,
							"create",
						),
						"view": getSystemRolePermissions(
							new_model.SystemRoleAttributeOpenAPI,
							"view",
						),
						"update": getSystemRolePermissions(
							new_model.SystemRoleAttributeOpenAPI,
							"update",
						),
						"delete": getSystemRolePermissions(
							new_model.SystemRoleAttributeOpenAPI,
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

func RolePermissionAggregation(
	builder *sql_query.SelectBuilder,
	loggedInUserId int,
	groupColumn string,
) {
	transactionRoleBuilder := transactionRoleCTEBuilder(loggedInUserId)
	builder.WithCTEBuilder("user_transaction_roles_cte", transactionRoleBuilder.(*sql_query.SelectBuilder).SQLEloquentQuery)
	builder.LeftJoin(
		"user_transaction_roles_cte utr",
		fmt.Sprintf("utr.group_id = %s", groupColumn),
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

func FindUserSystemRoles[T any](
	v *[]T,
	service service.PostgreSqlService,
	ctx context.Context,
	loggedInUserId int,
	groupId ...int,
) error {
	var err error

	isGlobalRole := true
	if len(groupId) > 0 {
		isGlobalRole = false
	}

	sytemRoleBuilder := sql_query.NewSQLSelectBuilder[T]("user_system_roles")

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

	if !isGlobalRole {
		filter["group_id"] = sql_query.SQLCondition{
			Operator: sql_query.SQLOperatorEqual,
			Value:    groupId[0],
		}
	}

	queryString, args, _ := sytemRoleBuilder.
		Where(filter).
		Build()

	err = service.SelectMany(v, ctx, queryString, args...)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func FindUserTransactionRoles[T any](
	v *[]T,
	service service.PostgreSqlService,
	ctx context.Context,
	loggedInUserId int,
	groupId int,
) error {
	var err error

	transactionRoleBuilder := sql_query.NewSQLSelectBuilder[T]("user_transaction_roles")

	filter := map[string]sql_query.SQLCondition{
		"user_id": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    loggedInUserId,
		},
		"is_active": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    true,
		},
		"group_id": {
			Operator: sql_query.SQLOperatorEqual,
			Value:    groupId,
		},
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

type GroupFieldAndAttributeName struct {
	Field         string // origin_group.id / destination_group.id
	AttributeName string // Origin / Destination (For isDestinationManager)
}

func TransactionRoleMultipleGroupAggregation(
	builder *sql_query.SelectBuilder,
	loggedInUserId int,
	groupFieldsAndAttribute []GroupFieldAndAttributeName,
	transactionGroupAttribute new_model.TransactionRoleAttribute,
	roleTypes []new_model.TransactionRoleType,
) {
	builder.
		LeftJoin(
			"user_transaction_roles utr",
			fmt.Sprintf(
				"utr.user_id = %d AND utr.is_active = true AND utr.role_attribute = %s",
				loggedInUserId,
				transactionGroupAttribute,
			),
		)

	if len(roleTypes) > 0 {
		builder.Where(map[string]sql_query.SQLCondition{
			"utr\".\"role_type": {
				Operator: sql_query.SQLOperatorIn,
				Value:    roleTypes,
				IsRef:    true,
			},
		})
	}

	for _, groupField := range groupFieldsAndAttribute {
		builder.SelectBoolOr(
			fmt.Sprintf(
				"utr.group_id = '%s' AND utr.role_type = '%s'",
				groupField.Field,
				new_model.TransactionRoleTypeManager,
			),
			fmt.Sprintf("is%sManager", groupField.AttributeName),
		)
		builder.SelectBoolOr(
			fmt.Sprintf(
				"utr.group_id = '%s' AND utr.role_type = '%s'",
				groupField.Field,
				new_model.TransactionRoleTypeMonitoring,
			),
			fmt.Sprintf("is%sMonitoring", groupField.AttributeName),
		)
	}
}

func transactionRoleMultiPermissionAggregationCTE(
	loggedInUserId int,
) sql_query.SQLSelectChainBuilder {
	return sql_query.NewSQLSelectBuilder[UserTransactionRole](
		"user_transaction_roles utr",
	).
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
				Operator: sql_query.SQLOperatorIn,
				Value: []new_model.TransactionRoleAttribute{
					new_model.TransactionRoleAttributeBorrowing,
					new_model.TransactionRoleAttributeTransfer,
					new_model.TransactionRoleAttributeDisposal,
				},
			},
		}).
		Select("utr.group_id").
		SelectJSONAggregate(
			"manager",
			map[string]string{
				"borrowing": fmt.Sprintf(
					"BOOL_OR(utr.role_type = '%s' AND utr.role_attribute = '%s')",
					new_model.TransactionRoleTypeManager,
					new_model.TransactionRoleAttributeBorrowing,
				),
				"transfer": fmt.Sprintf(
					"BOOL_OR(utr.role_type = '%s' AND utr.role_attribute = '%s')",
					new_model.TransactionRoleTypeManager,
					new_model.TransactionRoleAttributeTransfer,
				),
				"disposal": fmt.Sprintf(
					"BOOL_OR(utr.role_type = '%s' AND utr.role_attribute = '%s')",
					new_model.TransactionRoleTypeManager,
					new_model.TransactionRoleAttributeDisposal,
				),
			},
			"",
			false,
		).
		SelectJSONAggregate(
			"monitoring",
			map[string]string{
				"borrowing": fmt.Sprintf(
					"BOOL_OR(utr.role_type = '%s' AND utr.role_attribute = '%s')",
					new_model.TransactionRoleTypeMonitoring,
					new_model.TransactionRoleAttributeBorrowing,
				),
				"transfer": fmt.Sprintf(
					"BOOL_OR(utr.role_type = '%s' AND utr.role_attribute = '%s')",
					new_model.TransactionRoleTypeMonitoring,
					new_model.TransactionRoleAttributeTransfer,
				),
				"disposal": fmt.Sprintf(
					"BOOL_OR(utr.role_type = '%s' AND utr.role_attribute = '%s')",
					new_model.TransactionRoleTypeMonitoring,
					new_model.TransactionRoleAttributeDisposal,
				),
			},
			"",
			false,
		).
		GroupBy("utr.group_id")
}

func TransactionRoleMultiPermissionAggregation(
	builder *sql_query.SelectBuilder,
	loggedInUserId int,
	groupColumn string,
) {
	cteBuilder := transactionRoleMultiPermissionAggregationCTE(loggedInUserId)
	builder.WithCTEBuilder("user_transaction_roles_cte", cteBuilder.(*sql_query.SelectBuilder).SQLEloquentQuery)
	builder.LeftJoin(
		"user_transaction_roles_cte utr",
		fmt.Sprintf("utr.group_id = %s", groupColumn),
	)
	builder.Select(
		"\"utr\".\"manager\" AS manager",
		"\"utr\".\"monitoring\" AS monitoring",
	)
}

func RoleRestrictionAggregation(
	builder *sql_query.SelectBuilder,
	loggedInUserId int,
	groupColumn string,
	filterOptions RoleRestrictionOptions,
) {
	if filterOptions.TransactionRole != nil {
		joinCondition := fmt.Sprintf(
			"utr.user_id = %d AND utr.is_active = true",
			loggedInUserId,
		)
		if groupColumn != "" {
			joinCondition += fmt.Sprintf(" AND utr.group_id = %s", groupColumn)
		}

		builder.
			Join(
				"user_transaction_roles utr",
				joinCondition,
			).
			Where(map[string]sql_query.SQLCondition{
				"utr\".\"role_type": {
					Operator: sql_query.SQLOperatorIn,
					Value:    filterOptions.TransactionRole.ManageType,
					IsRef:    true,
				},
			})
	}

	if filterOptions.SystemRole != nil {
		joinCondition := fmt.Sprintf(
			"usr.user_id = %d AND usr.is_active = true AND usr.role_name = '%s'",
			loggedInUserId,
			filterOptions.SystemRole.RoleType,
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
					"usr\".\"permission_create": {
						Operator: sql_query.SQLOperatorEqual,
						Value:    filterOptions.SystemRole.Permissions.Create,
						IsRef:    true,
					},
					"usr\".\"permission_view": {
						Operator: sql_query.SQLOperatorEqual,
						Value:    filterOptions.SystemRole.Permissions.View,
						IsRef:    true,
					},
					"usr\".\"permission_update": {
						Operator: sql_query.SQLOperatorEqual,
						Value:    filterOptions.SystemRole.Permissions.Update,
						IsRef:    true,
					},
					"usr\".\"permission_delete": {
						Operator: sql_query.SQLOperatorEqual,
						Value:    filterOptions.SystemRole.Permissions.Delete,
						IsRef:    true,
					},
				},
			)
	}
}
