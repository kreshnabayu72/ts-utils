package middleware

import (
	"gitea.qwertysystem.net/BETS/ts-utils/middleware"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/helper"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	"gitea.qwertysystem.net/BETS/ts-utils/response"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Not filtered by group
func GlobalSystemRoleRestriction(permission model.PermissionType, resourceTypes []model.RoleType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		isTest := c.Locals(middleware.IsTest)
		if isTest != nil {
			return c.Next()
		}

		// Check permissions based on the provided permissionKey and resourceTypes
		err := helper.CheckSystemRolePermission(c, permission, resourceTypes, bson.M{"$ne": nil})
		if err != nil {
			return response.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized: "+err.Message)
		}

		// Continue to the next handler if permission is granted
		return c.Next()
	}
}

// Check user data, totalControl.isAssigned & totalControl.isActive
func TotalControlRoleRestriction() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isTest := c.Locals(middleware.IsTest)
		if isTest != nil {
			return c.Next()
		}

		// Check permissions based on the provided permissionKey and resourceTypes
		err := helper.CheckTotalControlPermission(c)
		if err != nil {
			return response.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized: "+err.Message)
		}

		// Continue to the next handler if permission is granted
		return c.Next()
	}
}

// Not filtered by group
func TransactionRoleRestriction(roleTypes []model.TransactionRoleType, transactionAttributes []model.TransactionGroupAttribute) fiber.Handler {
	return func(c *fiber.Ctx) error {
		isTest := c.Locals(middleware.IsTest)
		if isTest != nil {
			return c.Next()
		}

		// Check permissions based on the provided permissionKey and resourceTypes
		err := helper.CheckTransactionRolePermission(c, bson.M{"$ne": nil}, transactionAttributes, roleTypes)
		if err != nil {
			return response.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized: "+err.Error())
		}

		// Continue to the next handler if permission is granted
		return c.Next()
	}
}

// Not filtered by group
func AssetListRoleRestriction() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isTest := c.Locals(middleware.IsTest)
		if isTest != nil {
			return c.Next()
		}

		// Check permissions based on the provided permissionKey and resourceTypes
		err := helper.CheckTransactionRolePermission(c, bson.M{"$ne": nil},
			[]model.TransactionGroupAttribute{model.TGABorrowingRole, model.TGADisposalRole, model.TGATransferRole, model.TGAAuditRole, model.TGAMaintenanceRole, model.TGARepairRole, model.TGATrackingRole, model.TGADamagedRole, model.TGAMissingRole},
			[]model.TransactionRoleType{model.TRTManager, model.TRTMonitoring, model.TRTStaff})
		if err != nil {
			return response.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized: "+err.Error())
		}

		// Check permissions based on the provided permissionKey and resourceTypes
		systErr := helper.CheckSystemRolePermission(c, model.PermissionView,
			[]model.RoleType{model.RTTag, model.RTPurchaseInformation, model.RTRegisterAsset, model.RTPurchaseDocument, model.RTDepreciationGroup, model.RTDepreciationMethod},
			bson.M{"$ne": nil})
		if systErr != nil {
			return response.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized: "+systErr.Message)
		}

		// Continue to the next handler if permission is granted
		return c.Next()
	}
}
