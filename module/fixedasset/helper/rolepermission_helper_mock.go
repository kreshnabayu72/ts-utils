package helper

import (
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockRolePermissionHelper struct {
	mock.Mock
}

func (m *MockRolePermissionHelper) CheckSystemRolePermission(
	ctx *fiber.Ctx,
	permissionKey model.PermissionType,
	resourceTypes []model.RoleType,
	groupFilter primitive.M,
) *entity.HttpError {
	args := m.Mock.Called(ctx, permissionKey, resourceTypes, groupFilter)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockRolePermissionHelper) CheckGroupPermission(
	service service.Service,
	loggedInUser *primitive.ObjectID,
	dataFilter primitive.M,
	groupIdField string,
	permissions []model.PermissionType,
	systemRoleTypes []model.RoleType,
) error {
	args := m.Mock.Called(
		service,
		loggedInUser,
		dataFilter,
		groupIdField,
		permissions,
		systemRoleTypes,
	)
	return args.Error(0)
}

func (m *MockRolePermissionHelper) CheckTransactionRoleMultiGroupPermissions(
	service service.Service,
	loggedInUser *primitive.ObjectID,
	dataFilter primitive.M,
	groupIdField string,
	transactionType []model.TransactionGroupAttribute,
	roleType []model.TransactionRoleType,
) error {
	args := m.Mock.Called(
		service,
		loggedInUser,
		dataFilter,
		groupIdField,
		transactionType,
		roleType,
	)
	return args.Error(0)
}

func (m *MockRolePermissionHelper) CheckTransactionRolePermission(
	ctx *fiber.Ctx,
	groupIdFilter primitive.M,
	transactionType []model.TransactionGroupAttribute,
	roleType []model.TransactionRoleType,
) error {
	args := m.Mock.Called(ctx, groupIdFilter, transactionType, roleType)
	return args.Error(0)
}

func (m *MockRolePermissionHelper) CheckTotalControlPermission(ctx *fiber.Ctx) *entity.HttpError {
	args := m.Mock.Called(ctx)
	return args.Get(0).(*entity.HttpError)
}
