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

type UserTransactionRole struct {
	ID                        string                                     `json:"id" column:"utr.id::text"`
	GroupID                   string                                     `json:"group_id" column:"utr.group_id"`
	RoleType                  fixedasset_model.RoleType                  `json:"role_type" column:"utr.role_type"`
	TransactionGroupAttribute fixedasset_model.TransactionGroupAttribute `json:"role_attribute" column:"utr.role_attribute"`
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
