package routes

import (
	"icomphub-api/controllers"
	"icomphub-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, userController *controllers.UserController) {
	userRoute := rg.Group("/users")
	userRoute.GET("/", userController.GetAll)
	userRoute.GET("/:id", userController.Find)
	userRoute.POST("/", userController.Create)
	userRoute.PUT("/:id", middlewares.AuthMiddleware(), userController.Update)
	userRoute.DELETE("/:id", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("admin"), userController.Delete)
	userRoute.POST("/profile-picture", middlewares.AuthMiddleware(), userController.UpdateProfilePicture)
	userRoute.DELETE("/profile-picture", middlewares.AuthMiddleware(), userController.DeleteProfilePicture)
	userRoute.GET("/profile-picture", middlewares.AuthMiddleware(), userController.GetProfilePicture)
	userRoute.GET("/profile-picture/:id", userController.GetProfilePictureById)
}
