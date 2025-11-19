package middleware

import (
	"slices"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/db"
	common_model "gitea.qwertysystem.net/BETS/ts-utils/v2/module/common/model"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/fixedasset/model"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/response"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/service"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
	Company/Client data has "plan" attribute with type: {"fixedAsset":string,"supplyAsset":string} (will add more)

	1 API can lead to different use case, based on company plan

	e.g: Create transfer, in Basic plan it creates finished transfer data instantly.
		in Enterprise (& Standard?) it has approval process, handover, etc.
*/

func GetLicensePlan() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		isTest := ctx.Locals(IsTest)
		if isTest != nil {
			return ctx.Next()
		}

		company := ctx.Locals(Company).(primitive.ObjectID)

		// get the plan attribute
		companyService := service.MakeService(db.AdminDBName, db.ClientModelName)

		var companyData *model.Company
		err := companyService.FindOne(&companyData, ctx.Context(), bson.M{"_id": company})
		if err != nil {
			return response.SendResponse(ctx, fiber.StatusBadRequest, nil, "Invalid JWT")
		} // Throw error if no plan

		if companyData.Plan == nil {
			return response.SendResponse(ctx, fiber.StatusBadRequest, nil, "Invalid License Plan")
		}

		// Save plan in context
		ctx.Locals(FixedAssetPlan, companyData.Plan.FixedAsset)
		ctx.Locals(SupplyPlan, companyData.Plan.SupplyAsset)

		return ctx.Next()
	}
}

func ValidateLicensePlan(module common_model.AppModule) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		contextKey := map[common_model.AppModule]ContextKey{
			common_model.FixedAssetModule:  FixedAssetPlan,
			common_model.SupplyAssetModule: SupplyPlan,
		}

		_, ok := ctx.Locals(contextKey[module]).(model.PlanTypes)
		if !ok {
			return response.SendResponse(ctx, fiber.StatusUnauthorized, nil, "Invalid License Plan")
		}

		return ctx.Next()
	}
}

func LicensePlanGuard(module common_model.AppModule, allowedPlan []model.PlanTypes) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		contextKey := map[common_model.AppModule]ContextKey{
			common_model.FixedAssetModule:  FixedAssetPlan,
			common_model.SupplyAssetModule: SupplyPlan,
		}

		plan := ctx.Locals(contextKey[module]).(model.PlanTypes)

		isAllowed := slices.Contains(allowedPlan, plan)

		if !isAllowed {
			return response.SendResponse(
				ctx,
				fiber.StatusForbidden,
				nil,
				"This action is not allowed in your current plan",
			)
		}

		return ctx.Next()
	}
}
