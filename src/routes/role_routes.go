package routes

import (
	"icomphub-api/controllers"
	"icomphub-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(rg *gin.RouterGroup, roleController *controllers.RoleController) {
	roleRoute := rg.Group("/roles")

	roleRoute.GET("/:id", roleController.Find)
	roleRoute.POST("/", middlewares.AuthMiddleware(), roleController.Create)
	roleRoute.PUT("/:id", middlewares.AuthMiddleware(), roleController.Update)
	roleRoute.DELETE("/:id", middlewares.AuthMiddleware(), roleController.Delete)
}
