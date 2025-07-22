package middleware

import (
	"gitea.qwertysystem.net/BETS/ts-utils/middleware"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/helper"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/response"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"github.com/gofiber/fiber/v2"
)

// Not filtered by group
func GlobalSystemRoleRestrictionSQL(permission new_model.PermissionType, resourceTypes []new_model.SystemRoleAttribute) fiber.Handler {
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
func TransactionRoleRestrictionSQL(roleTypes []new_model.TransactionRoleType, transactionAttributes []new_model.TransactionRoleAttribute) fiber.Handler {
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
			[]new_model.TransactionRoleAttribute{new_model.TransactionRoleAttributeBorrowing, new_model.TransactionRoleAttributeDisposal, new_model.TransactionRoleAttributeTransfer, new_model.TransactionRoleAttributeAudit, new_model.TransactionRoleAttributeMaintenance, new_model.TransactionRoleAttributeRepair, new_model.TransactionRoleAttributeTracking, new_model.TransactionRoleAttributeDamaged, new_model.TransactionRoleAttributeMissing},
			[]new_model.TransactionRoleType{new_model.TransactionRoleTypeManager, new_model.TransactionRoleTypeMonitoring, new_model.TransactionRoleTypeStaff})

		// Check permissions based on the provided permissionKey and resourceTypes
		systErr := helper.CheckSystemRolePermissionSQL(c, new_model.PermissionView,
			[]new_model.SystemRoleAttribute{new_model.SystemRoleAttributeTag, new_model.SystemRoleAttributePurchaseInfo, new_model.SystemRoleAttributeRegisterAsset, new_model.SystemRoleAttributePurchaseDocument, new_model.SystemRoleAttributeDepreciationGroup, new_model.SystemRoleAttributeDepreciationMethod},
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
