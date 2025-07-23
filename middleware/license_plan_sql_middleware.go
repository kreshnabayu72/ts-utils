package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/dto"
	"gitea.qwertysystem.net/BETS/ts-utils/response"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

// GetLicensePlanSQL is a Fiber handler that retrieves the license plan information
// for a specific company from the PostgreSQL database. It checks if the request is
// in test mode and skips processing if it is. The function expects the company ID
// to be stored in the context and uses it to query the database for the corresponding
// client data.
//
// The function performs the following steps:
//  1. Checks if the request is in test mode. If so, it calls the next handler in the chain.
//  2. Retrieves the company ID from the context.
//  3. Creates a PostgreSQL service instance for querying the database.
//  4. Constructs a SQL SELECT query to fetch the client data based on the company ID.
//  5. Executes the query and populates a `new_model.Client` struct with the result.
//  6. If the query execution is successful, it saves the fixed asset plan and supply asset
//     plan in the context for further processing by later handlers.
//
// Returns:
// - A Fiber error if any step fails, with an appropriate HTTP status code and error message.
// - Calls `ctx.Next()` to pass control to the next handler if successful.
func GetLicensePlanSQL() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var err error

		isTest := ctx.Locals(IsTest)
		if isTest != nil {
			return ctx.Next()
		}

		company := ctx.Locals(Company).(string)

		companyService := service.MakePostgreSqlService(db.AdminDBName, db.ClientTableName)

		query, args, err := sql_query.NewSQLSelectBuilder[dto.ClientPlan](db.ClientTableName).
			Where(map[string]sql_query.SQLCondition{
				"id": {
					Operator: sql_query.SQLOperatorEqual,
					Value:    company,
					IsRef:    false,
				},
			}).
			SetLimit(1).
			Build()
		if err != nil {
			return response.SendResponseWithError(ctx, fiber.StatusBadRequest, nil, "Invalid JWT", err.Error())
		}

		companyData := dto.ClientPlan{}
		err = companyService.SelectOne(
			&companyData,
			ctx.UserContext(),
			query,
			args...,
		)
		if err != nil {
			fmt.Println("Jwt query: " + query)
			fmt.Printf("Jwt args: %v\n", args)
			return response.SendResponseWithError(ctx, fiber.StatusBadRequest, nil, "Invalid JWT", err.Error())
		}

		// Save plan in context.
		ctx.Locals(FixedAssetPlan, companyData.PlanFixedAsset)
		ctx.Locals(SupplyPlan, companyData.PlanSupplyAsset)

		return ctx.Next()
	}
}
