package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/betsv2/tsv2-go-utils/types"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Middleware to validate JWT
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {

	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Extract the actual token by removing "Bearer " prefix
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse and validate the JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm used
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["id"] == nil || claims["companyCode"] == nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			companyCode := claims["companyCode"].(string)
			userId, _ := primitive.ObjectIDFromHex(claims["id"].(string))

			userKey := types.ContextKey("user")
			companyCodeKey := types.ContextKey("companyCode")

			ctx := r.Context()
			ctx = context.WithValue(ctx, userKey, userId)
			ctx = context.WithValue(ctx, companyCodeKey, companyCode)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	}
}
