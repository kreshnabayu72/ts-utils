package helper

import (
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"github.com/gofiber/fiber/v2"
)

type RolePermissionHelperSQL interface {
	CheckSystemRolePermissionSQL(
		ctx *fiber.Ctx,
		permissionKey new_model.PermissionType,
		resourceTypes []new_model.SystemRoleAttribute,
		groupFilter sql_query.SQLCondition,
	) *entity.HttpError

	CheckGroupPermissionSQL(
		service service.PostgreSqlService,
		tableName string,
		loggedInUser *string,
		dataFilter map[string]sql_query.SQLCondition,
		groupIdField string,
		permissions []new_model.PermissionType,
		systemRoleTypes []new_model.SystemRoleAttribute,
	) error

	CheckTransactionRoleMultiGroupPermissionsSQL(
		service service.PostgreSqlService, // service for data (assets, transactions, etc)
		tableName string,
		loggedInUser *string,
		dataFilter map[string]sql_query.SQLCondition,
		groupIdField string,
		transactionType []new_model.TransactionRoleAttribute,
		roleType []new_model.TransactionRoleType,
	) error

	CheckTransactionRolePermissionSQL(
		ctx *fiber.Ctx,
		groupIdFilter sql_query.SQLCondition,
		transactionType []new_model.TransactionRoleAttribute,
		roleType []new_model.TransactionRoleType,
	) error

	CheckTotalControlPermissionSQL(ctx *fiber.Ctx) *entity.HttpError
}

type RolePermissionSQLHelperImpl struct{}

func MakeRolePermisisonSQLHelperImpl() RolePermissionHelperSQL {
	return &RolePermissionSQLHelperImpl{}
}

func (r *RolePermissionSQLHelperImpl) CheckSystemRolePermissionSQL(
	ctx *fiber.Ctx,
	permissionKey new_model.PermissionType,
	resourceTypes []new_model.SystemRoleAttribute,
	groupFilter sql_query.SQLCondition,
) *entity.HttpError {
	return CheckSystemRolePermissionSQL(ctx, permissionKey, resourceTypes, groupFilter)
}

func (r *RolePermissionSQLHelperImpl) CheckGroupPermissionSQL(
	service service.PostgreSqlService,
	tableName string,
	loggedInUser *string,
	dataFilter map[string]sql_query.SQLCondition,
	groupIdField string,
	permissions []new_model.PermissionType,
	systemRoleTypes []new_model.SystemRoleAttribute,
) error {
	return CheckGroupPermissionSQL(
		service,
		tableName,
		loggedInUser,
		dataFilter,
		groupIdField,
		permissions,
		systemRoleTypes,
	)
}

func (r *RolePermissionSQLHelperImpl) CheckTransactionRoleMultiGroupPermissionsSQL(
	service service.PostgreSqlService,
	tableName string,
	loggedInUser *string,
	dataFilter map[string]sql_query.SQLCondition,
	groupIdField string,
	transactionType []new_model.TransactionRoleAttribute,
	roleType []new_model.TransactionRoleType,
) error {
	return CheckTransactionRoleMultiGroupPermissionsSQL(
		service,
		tableName,
		loggedInUser,
		dataFilter,
		groupIdField,
		transactionType,
		roleType,
	)
}

func (r *RolePermissionSQLHelperImpl) CheckTransactionRolePermissionSQL(
	ctx *fiber.Ctx,
	groupIdFilter sql_query.SQLCondition,
	transactionType []new_model.TransactionRoleAttribute,
	roleType []new_model.TransactionRoleType,
) error {
	return CheckTransactionRolePermissionSQL(ctx, groupIdFilter, transactionType, roleType)
}

func (r *RolePermissionSQLHelperImpl) CheckTotalControlPermissionSQL(ctx *fiber.Ctx) *entity.HttpError {
	return CheckTotalControlPermissionSQL(ctx)
}
