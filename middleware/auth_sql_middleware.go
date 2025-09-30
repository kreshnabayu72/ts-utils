package middleware

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gitea.qwertysystem.net/BETS/ts-utils/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Middleware to validate JWT
func ValidateJWTSQL() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		jwtKey := []byte(os.Getenv("JWT_KEY"))
		tokenStr := ctx.Get("Authorization", "NOT_FOUND")
		if tokenStr == "NOT_FOUND" {
			log.Println("Error: no auth header")
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
			log.Printf("Error parsing token: %v\n", err.Error())
			return response.SendResponse(ctx, fiber.StatusBadRequest, nil, "Unauthorized")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["id"] == nil || claims["companyCode"] == nil {
				return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
			}

			companyCode := claims["companyCode"].(string)
			userId, _ := claims["id"].(string)
			session, _ := claims["session"].(string)
			company, _ := claims["company"].(string)

			ctx.Locals(UserKey, userId)
			ctx.Locals(SessionKey, session)
			ctx.Locals(CompanyCodeKey, companyCode)
			ctx.Locals(Company, company)

		} else {
			log.Println("invalid token")
			return response.SendResponse(ctx, fiber.StatusBadRequest, nil, "Unauthorized")
		}

		return ctx.Next()
	}
}
