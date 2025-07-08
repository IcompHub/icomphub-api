package middlewares

import (
	"errors"
	"strings"

	"icomphub-api/auth"
	"icomphub-api/codes"
	"icomphub-api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			handlers.AbortUnauthorized(ctx, codes.AuthMissingToken, errors.New("missing or invalid token"))
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenStr, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return auth.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			handlers.AbortUnauthorized(ctx, codes.AuthExpiredToken, errors.New("invalid or expired token"))
			return
		}

		claims := token.Claims.(*auth.Claims)

		ctx.Set("userID", claims.UserID)
		ctx.Set("systemRole", claims.SystemRole)
		ctx.Next()
	}
}
