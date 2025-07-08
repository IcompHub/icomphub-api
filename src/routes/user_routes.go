package routes

import (
	"icomphub-api/controllers"
	"icomphub-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, userController *controllers.UserController) {
	userRoute := rg.Group("/users")
	userRoute.GET("/", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("admin"), userController.GetAll)
	userRoute.GET("/:id", userController.Find)
	userRoute.POST("/", userController.Create)
	userRoute.PUT("/:id", userController.Update)
	userRoute.DELETE("/:id", userController.Delete)
}
