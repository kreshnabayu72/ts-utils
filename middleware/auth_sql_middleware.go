package middleware

import (
	"fmt"
	"os"
	"strings"

	"gitea.qwertysystem.net/BETS/ts-utils/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Middleware to validate JWT
func ValidateJWTSQL() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var jwtKey = []byte(os.Getenv("JWT_KEY"))
		tokenStr := ctx.Get("Authorization", "NOT_FOUND")
		if tokenStr == "NOT_FOUND" {
			return response.SendResponse(ctx, fiber.StatusBadRequest, nil, "Unauthorized")
		}

		// Extract the actual token by removing "Bearer " prefix
		tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)

		// Parse and validate the JWT
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm used
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			return response.SendResponse(ctx, fiber.StatusBadRequest, nil, "Unauthorized")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["id"] == nil || claims["companyCode"] == nil {
				return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
			}

			companyCode := claims["companyCode"].(string)
			userId, _ := claims["id"].(float64)
			session, _ := claims["session"].(float64)
			company, _ := claims["company"].(float64)

			ctx.Locals(UserKey, int(userId))
			ctx.Locals(SessionKey, int(session))
			ctx.Locals(CompanyCodeKey, companyCode)
			ctx.Locals(Company, int(company))

		} else {
			return response.SendResponse(ctx, fiber.StatusBadRequest, nil, "Unauthorized")
		}

		return ctx.Next()
	}
}
