package middleware

import (
	"gitea.qwertysystem.net/BETS/ts-utils/middleware"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/helper"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/response"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"github.com/gofiber/fiber/v2"
)

func CombineOrRestriction(middlewares ...fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var lastErr error

		for _, mw := range middlewares {
			// Run each middleware
			err := mw(c)
			if err == nil {
				// At least one middleware passed then authorized
				return nil
			}
			lastErr = err
		}

		// All middlewares fails
		if lastErr != nil {
			return lastErr
		}
		return response.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized")
	}
}

// PermittedUserRestrictionSQL creates a middleware that restricts access based on a user's
// given access (either "Assigned" or "Borrower" based on given param value).
func PermittedUserRestrictionSQL(typeUser helper.PermittedUser) fiber.Handler {
	return func(c *fiber.Ctx) error {
		isTest := c.Locals(middleware.IsTest)
		if isTest != nil {
			return c.Next()
		}

		// Check permissions based on the provided permissionKey and resourceTypes
		err := helper.CheckPermittedUserPermissionSQL(c, typeUser)
		if err != nil {
			return response.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized: "+err.Message)
		}

		// Continue to the next handler if permission is granted
		return c.Next()
	}
}

// GlobalSystemRoleRestrictionSQL creates a middleware that restricts access based on a user's
// system role permissions, regardless of which group they belong to (which means global system role).
//
// This is used for routes where a user only needs to have a specific permission *somewhere*
// in the system, not tied to a particular group context. It checks for a user's
// permission (e.g., "edit") and role attribute (e.g., "assetAttribute") without
// filtering by a specific group ID. If the user has a matching permission in any group,
// they are granted access.
//
// Use this for global actions or routes that display information not specific to a single group.
func GlobalSystemRoleRestrictionSQL(permission new_model.PermissionType, resourceTypes []new_model.SystemRoleAttribute) fiber.Handler {
	return func(c *fiber.Ctx) error {
		isTest := c.Locals(middleware.IsTest)
		if isTest != nil {
			return c.Next()
		}

		// Check permissions based on the provided permissionKey and resourceTypes
		err := helper.CheckSystemRolePermissionSQL(c, permission, resourceTypes, sql_query.SQLCondition{
			Operator: sql_query.SQLOperatorIsNull,
		})
		if err != nil {
			return response.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized: "+err.Message)
		}

		// Continue to the next handler if permission is granted
		return c.Next()
	}
}

// Check user data, totalControl.isAssigned & totalControl.isActive
// TotalControlRoleRestrictionSQL creates a middleware that grants access only to users with
// "total control" administrative privileges.
//
// This is the highest level of access control. It's intended for routes that handle
// system-wide settings, sensitive data, or critical administrative functions that should
// bypass all other group-based permissions. It directly checks the user's account for
// the "total control" flags.
//
// Only use this for routes where a user must have full, unrestricted access.
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

// TransactionRoleRestrictionSQL creates a middleware that checks for a user's transaction
// role permissions, ignoring group filters.
//
// This is similar to `GlobalSystemRoleRestrictionSQL` but for transaction-specific roles. It's
// useful for routes where the user needs a certain transaction role (e.g., a "manager" role for
// "transfer" transactions) in any group they belong to. The check is passed if an active
// transaction role is found that matches the provided types and attributes.
//
// This helper assumes that the action itself is not group-specific.
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

// AssetListRoleRestrictionSQL creates a middleware to determine if a user has any permission
// to view assets.
//
// This middleware is a combination of both system and transaction role checks. It grants
// access if the user has **either** a view-related system role permission or any of the
// listed transaction-based permissions. This is designed for routes that display a list of
// assets, as a user might be authorized to see assets for a variety of reasons (e.g.,
// being a 'Manager' for 'borrowingRole' or having a 'view' system permission for 'registerAsset').
//
// The check succeeds if at least one of the two permission checks passes.
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
