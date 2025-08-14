package helper

import (
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

type MockRolePermissionSQLHelper struct {
	mock.Mock
}

func (m *MockRolePermissionSQLHelper) CheckSystemRolePermissionSQL(
	ctx *fiber.Ctx,
	permissionKey new_model.PermissionType,
	resourceTypes []new_model.SystemRoleAttribute,
	groupFilter sql_query.SQLCondition,
) *entity.HttpError {
	args := m.Mock.Called(ctx, permissionKey, resourceTypes, groupFilter)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockRolePermissionSQLHelper) CheckGroupPermissionSQL(
	service service.PostgreSqlService,
	tableName string,
	loggedInUser *string,
	dataFilter map[string]sql_query.SQLCondition,
	groupIdField string,
	permissions []new_model.PermissionType,
	systemRoleTypes []new_model.SystemRoleAttribute,
) error {
	args := m.Mock.Called(
		service,
		tableName,
		loggedInUser,
		dataFilter,
		groupIdField,
		permissions,
		systemRoleTypes,
	)
	return args.Error(0)
}

func (m *MockRolePermissionSQLHelper) CheckTransactionRoleMultiGroupPermissionsSQL(
	service service.PostgreSqlService, // service for data (assets, transactions, etc)
	tableName string,
	loggedInUser *string,
	dataFilter map[string]sql_query.SQLCondition,
	groupIdField string,
	transactionType []new_model.TransactionRoleAttribute,
	roleType []new_model.TransactionRoleType,
) error {
	args := m.Mock.Called(
		service,
		tableName,
		loggedInUser,
		dataFilter,
		groupIdField,
		transactionType,
		roleType,
	)
	return args.Error(0)
}

func (m *MockRolePermissionSQLHelper) CheckTransactionRolePermissionSQL(
	ctx *fiber.Ctx,
	groupIdFilter sql_query.SQLCondition,
	transactionType []new_model.TransactionRoleAttribute,
	roleType []new_model.TransactionRoleType,
) error {
	args := m.Mock.Called(ctx, groupIdFilter, transactionType, roleType)
	return args.Error(0)
}

func (m *MockRolePermissionSQLHelper) CheckTotalControlPermissionSQL(ctx *fiber.Ctx) *entity.HttpError {
	args := m.Mock.Called(ctx)
	return args.Get(0).(*entity.HttpError)
}
