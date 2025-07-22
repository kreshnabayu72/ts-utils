package middleware

import (
	"gitea.qwertysystem.net/BETS/ts-utils/middleware"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/helper"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	"gitea.qwertysystem.net/BETS/ts-utils/response"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"github.com/gofiber/fiber/v2"
)

// Not filtered by group
func GlobalSystemRoleRestrictionSQL(permission model.PermissionType, resourceTypes []model.RoleType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		isTest := c.Locals(middleware.IsTest)
		if isTest != nil {
			return c.Next()
		}

		// Check permissions based on the provided permissionKey and resourceTypes
		err := helper.CheckSystemRolePermissionSQL(c, permission, resourceTypes, sql_query.SQLCondition{
			Operator: sql_query.SQLOperatorIsNotNull,
		})
		if err != nil {
			return response.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized: "+err.Message)
		}

		// Continue to the next handler if permission is granted
		return c.Next()
	}
}

// Check user data, totalControl.isAssigned & totalControl.isActive
func TotalControlRoleRestrictionSQL() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isTest := c.Locals(middleware.IsTest)
		if isTest != nil {
			return c.Next()
		}

		// Check permissions based on the provided permissionKey and resourceTypes
		err := helper.CheckTotalControlPermissionSQL(c)
		if err != nil {
			return response.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized: "+err.Message)
		}

		// Continue to the next handler if permission is granted
		return c.Next()
	}
}

// Not filtered by group
func TransactionRoleRestrictionSQL(roleTypes []model.TransactionRoleType, transactionAttributes []model.TransactionGroupAttribute) fiber.Handler {
	return func(c *fiber.Ctx) error {
		isTest := c.Locals(middleware.IsTest)
		if isTest != nil {
			return c.Next()
		}

		// Check permissions based on the provided permissionKey and resourceTypes
		err := helper.CheckTransactionRolePermissionSQL(c, sql_query.SQLCondition{
			Operator: sql_query.SQLOperatorIsNotNull,
		}, transactionAttributes, roleTypes)
		if err != nil {
			return response.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized: "+err.Error())
		}

		// Continue to the next handler if permission is granted
		return c.Next()
	}
}

// Not filtered by group
func AssetListRoleRestrictionSQL() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isTest := c.Locals(middleware.IsTest)
		if isTest != nil {
			return c.Next()
		}

		// Check permissions based on the provided permissionKey and resourceTypes
		err := helper.CheckTransactionRolePermissionSQL(c, sql_query.SQLCondition{
			Operator: sql_query.SQLOperatorIsNotNull,
		},
			[]model.TransactionGroupAttribute{model.TGABorrowingRole, model.TGADisposalRole, model.TGATransferRole, model.TGAAuditRole, model.TGAMaintenanceRole, model.TGARepairRole, model.TGATrackingRole, model.TGADamagedRole, model.TGAMissingRole},
			[]model.TransactionRoleType{model.TRTManager, model.TRTMonitoring, model.TRTStaff})

		// Check permissions based on the provided permissionKey and resourceTypes
		systErr := helper.CheckSystemRolePermissionSQL(c, model.PermissionView,
			[]model.RoleType{model.RTTag, model.RTPurchaseInformation, model.RTRegisterAsset, model.RTPurchaseDocument, model.RTDepreciationGroup, model.RTDepreciationMethod},
			sql_query.SQLCondition{
				Operator: sql_query.SQLOperatorIsNotNull,
			})
		if systErr != nil && err != nil {
			return response.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized: "+systErr.Message)
		}

		// Continue to the next handler if permission is granted
		return c.Next()
	}
}
