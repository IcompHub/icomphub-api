package middlewares

import (
	"errors"
	"strings"

	"icomphub-api/auth"
	"icomphub-api/codes"
	"icomphub-api/handlers"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			handlers.AbortUnauthorized(ctx, codes.AuthMissingToken, errors.New("missing or invalid token"))
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := auth.GetTokenData(tokenStr)

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
