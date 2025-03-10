package interceptor

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ValidateJWT(
	ctx context.Context,
	req any,
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "failed to get metadata")
	}

	authorization, ok := md["authorization"]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "authorization header not found")
	}

	tokenStr := strings.TrimPrefix(authorization[0], "Bearer ")

	// Parse and validate the JWT
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok = token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["id"] == nil || claims["companyCode"] == nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		}

		companyCode := claims["companyCode"].(string)
		userId, _ := claims["id"].(string)
		session, _ := claims["session"].(string)
		company, _ := claims["company"].(string)

		md.Append(string(UserKey), userId)
		md.Append(string(SessionKey), session)
		md.Append(string(CompanyCodeKey), companyCode)
		md.Append(string(Company), company)

		ctx = metadata.NewIncomingContext(ctx, md)
	} else {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	return handler(ctx, req)
}
