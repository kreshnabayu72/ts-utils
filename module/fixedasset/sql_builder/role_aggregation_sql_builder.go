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

func generateUserSystemRolesBuilder(
	loggedInUserId int,
	groupColumn string,
	mainTable string,
	roleAttribute fixedasset_new_model.SystemRoleAttribute,
) sql_query.SQLSelectChainBuilder {
	builder := sql_query.NewSQLSelectBuilder[UserSystemRole]("user_system_roles usr").
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

	builder.WhereOr(
		map[string]sql_query.SQLCondition{
			// If the system role type is 'Global', skip checking for the group_id.
			"usr.role_type": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    fixedasset_new_model.SystemRoleTypeGlobal,
			},
			"usr.group_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    groupColumn,
				IsRef:    true,
			},
		},
	)

	builder.Select(
		fmt.Sprintf("%s.id AS id", mainTable),
		"usr.permission_create AS \"create\"",
		"usr.permission_view AS \"view\"",
		"usr.permission_update AS \"update\"",
		"usr.permission_delete AS \"delete\"",
	)

	builder.Join(mainTable, "true")

	return builder
}

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
			groupColumn,
			mainTable,
			roleAttribute,
		)

		builder.WithCTEBuilder(
			cteName,
			cteBuilder.(*sql_query.SelectBuilder).SQLEloquentQuery,
		)

		builder.LeftJoin(cteName, fmt.Sprintf("%s.id = %s.id", aliasTable, cteName))
	}
}

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
}

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
	builder.Select(`utr."isManager"`, `utr."isMonitoring"`)

	joinCteKeyword := builder.LeftJoin
	if len(filterOutEmptyRole) > 0 && filterOutEmptyRole[0] {
		joinCteKeyword = builder.Join
	}

	joinCteKeyword(
		"user_transaction_roles_cte utr",
		fmt.Sprintf("utr.group_id = %s", groupColumn),
	)
}

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

type SystemRoleFilter struct {
	RoleAttribute    fixedasset_new_model.SystemRoleAttribute `json:"roleAttribute" column:"role_attribute"` // Enum: system_role_attribute_enum
	PermissionView   bool                                     `json:"permissionView" column:"permission_view"`
	PermissionCreate bool                                     `json:"permissionCreate" column:"permission_create"`
	PermissionUpdate bool                                     `json:"permissionUpdate" column:"permission_update"`
	PermissionDelete bool                                     `json:"permissionDelete" column:"permission_delete"`
	GroupID          *int                                     `json:"groupId" column:"group_id"`
}

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

type TransactionRoleFilter struct {
	RoleAttribute fixedasset_new_model.TransactionRoleAttribute `json:"roleAttribute" column:"role_attribute"`
	RoleType      fixedasset_new_model.TransactionRoleType      `json:"roleType" column:"role_type"`
	GroupID       *int                                          `json:"groupId" column:"group_id"`
}

func FindUserTransactionRoles[T any](
	v *[]T,
	service service.PostgreSqlService,
	ctx context.Context,
	loggedInUserId int,
	roleFilter ...TransactionRoleFilter,
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
			filterPart["group_id"] = sql_query.SQLCondition{Operator: sql_query.SQLOperatorEqual, Value: each.GroupID}
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

type GroupFieldAndAttributeName struct {
	Field         string // origin_group.id / destination_group.id
	AttributeName string // Origin / Destination (For isDestinationManager)
}

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

	for _, groupField := range groupFieldsAndAttribute {
		builder.SelectBoolOr(
			fmt.Sprintf(
				"utr.group_id = %s AND utr.role_type = $%d",
				groupField.Field,
				builder.GetCurrentArgIndex()+1,
			),
			fmt.Sprintf("is%sManager", groupField.AttributeName),
			fixedasset_new_model.TransactionRoleTypeManager,
		)
		builder.SelectBoolOr(
			fmt.Sprintf(
				"utr.group_id = %s AND utr.role_type = $%d",
				groupField.Field,
				builder.GetCurrentArgIndex()+1,
			),
			fmt.Sprintf("is%sMonitoring", groupField.AttributeName),
			fixedasset_new_model.TransactionRoleTypeMonitoring,
		)
	}
}

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
}

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
						IsRef:    true,
					},
					"usr.permission_view": {
						Operator: sql_query.SQLOperatorEqual,
						Value:    filterOptions.SystemRole.Permissions.View,
						IsRef:    true,
					},
					"usr.permission_update": {
						Operator: sql_query.SQLOperatorEqual,
						Value:    filterOptions.SystemRole.Permissions.Update,
						IsRef:    true,
					},
					"usr.permission_delete": {
						Operator: sql_query.SQLOperatorEqual,
						Value:    filterOptions.SystemRole.Permissions.Delete,
						IsRef:    true,
					},
				},
			)
	}
}
