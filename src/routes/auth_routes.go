package routes

import (
	"icomphub-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, authCtrl *controllers.AuthController) {
	authRoute := rg.Group("/auth")
	authRoute.POST("/login", authCtrl.Login)
}
