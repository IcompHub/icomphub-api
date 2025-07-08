package middlewares

import (
	"errors"
	"slices"

	"icomphub-api/codes"
	"icomphub-api/handlers"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleVal, exists := ctx.Get("systemRole")
		if !exists {
			handlers.AbortForbidden(ctx, codes.RoleAccessDenied, errors.New("access denied"))
			return
		}
		userRole := roleVal.(string)
		if slices.Contains(allowedRoles, userRole) {
			ctx.Next()
			return
		}
		handlers.AbortForbidden(ctx, codes.RoleInsufficient, errors.New("insufficient role"))
	}
}
