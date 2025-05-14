package helper

import (
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RolePermissionHelper interface {
	CheckSystemRolePermission(
		ctx *fiber.Ctx,
		permissionKey model.PermissionType,
		resourceTypes []model.RoleType,
		groupFilter primitive.M,
	) *entity.HttpError

	CheckGroupPermission(
		service service.Service,
		loggedInUser *primitive.ObjectID,
		dataFilter primitive.M,
		groupIdField string,
		permissions []model.PermissionType,
		systemRoleTypes []model.RoleType,
	) error

	CheckTransactionRoleMultiGroupPermissions(
		service service.Service, // service for data (assets, transactions, etc)
		loggedInUser *primitive.ObjectID,
		dataFilter primitive.M,
		groupIdField string,
		transactionType []model.TransactionGroupAttribute,
		roleType []model.TransactionRoleType,
	) error

	CheckTransactionRolePermission(
		ctx *fiber.Ctx,
		groupIdFilter primitive.M,
		transactionType []model.TransactionGroupAttribute,
		roleType []model.TransactionRoleType,
	) error

	CheckTotalControlPermission(ctx *fiber.Ctx) *entity.HttpError
}

type RolePermissionHelperImpl struct{}

func MakeRolePermisisonHelperImpl() *RolePermissionHelperImpl {
	return &RolePermissionHelperImpl{}
}

func (r *RolePermissionHelperImpl) CheckSystemRolePermission(
	ctx *fiber.Ctx,
	permissionKey model.PermissionType,
	resourceTypes []model.RoleType,
	groupFilter primitive.M,
) *entity.HttpError {
	return CheckSystemRolePermission(ctx, permissionKey, resourceTypes, groupFilter)
}

func (r *RolePermissionHelperImpl) CheckGroupPermission(
	service service.Service,
	loggedInUser *primitive.ObjectID,
	dataFilter primitive.M,
	groupIdField string,
	permissions []model.PermissionType,
	systemRoleTypes []model.RoleType,
) error {
	return CheckGroupPermission(
		service,
		loggedInUser,
		dataFilter,
		groupIdField,
		permissions,
		systemRoleTypes,
	)
}

func (r *RolePermissionHelperImpl) CheckTransactionRoleMultiGroupPermissions(
	service service.Service,
	loggedInUser *primitive.ObjectID,
	dataFilter primitive.M,
	groupIdField string,
	transactionType []model.TransactionGroupAttribute,
	roleType []model.TransactionRoleType,
) error {
	return CheckTransactionRoleMultiGroupPermissions(
		service,
		loggedInUser,
		dataFilter,
		groupIdField,
		transactionType,
		roleType,
	)
}

func (r *RolePermissionHelperImpl) CheckTransactionRolePermission(
	ctx *fiber.Ctx,
	groupIdFilter primitive.M,
	transactionType []model.TransactionGroupAttribute,
	roleType []model.TransactionRoleType,
) error {
	return CheckTransactionRolePermission(ctx, groupIdFilter, transactionType, roleType)
}

func (r *RolePermissionHelperImpl) CheckTotalControlPermission(ctx *fiber.Ctx) *entity.HttpError {
	return CheckTotalControlPermission(ctx)
}
