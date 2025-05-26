package routes

import (
	"icomphub-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, userController *controllers.UserController) {
	userRoute := rg.Group("/users")
	userRoute.GET("/", userController.GetAll)
}
