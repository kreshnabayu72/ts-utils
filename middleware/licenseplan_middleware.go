package middleware

import (
	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	"gitea.qwertysystem.net/BETS/ts-utils/response"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
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
		//get company id
		company := ctx.Locals(Company).(primitive.ObjectID)

		//get the plan attribute
		companyService := service.MakeService(db.AdminDBName, db.ClientModelName)

		var companyData *model.Company
		err := companyService.FindOne(&companyData, ctx.Context(), bson.M{"_id": company})
		if err != nil {
			return response.SendResponse(ctx, fiber.ErrBadRequest.Code, map[string]interface{}{"message": "Invalid JWT", "status": 400}, "Invalid JWT")
		} //Throw error if no plan

		if companyData.Plan == nil {
			return response.SendResponse(ctx, fiber.ErrBadRequest.Code, map[string]interface{}{"message": "Invalid License Plan", "status": 400}, "Invalid License Plan")
		}

		//Save plan in context
		ctx.Locals(FixedAssetPlan, companyData.Plan.FixedAsset)

		return ctx.Next()
	}
}
