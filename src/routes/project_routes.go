package routes

import (
	"icomphub-api/controllers"
	"icomphub-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(rg *gin.RouterGroup, projectController *controllers.ProjectController) {
	projectRoute := rg.Group("/projects")
	projectRoute.GET("/", projectController.GetAll)
	projectRoute.GET("/:id", projectController.Find)
	projectRoute.POST("/", projectController.Create)
	projectRoute.PUT("/:id", projectController.Update)
	projectRoute.DELETE("/:id", projectController.Delete)
	projectRoute.POST("/thumbnail/:id", middlewares.AuthMiddleware(), projectController.UpdateThumbnail)
	projectRoute.DELETE("/thumbnail/:id", middlewares.AuthMiddleware(), projectController.DeleteThumbnail)
	projectRoute.GET("/thumbnail/:id", projectController.GetThumbnail)
}
