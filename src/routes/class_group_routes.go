package routes

import (
	"icomphub-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterClassGroupRoutes(rg *gin.RouterGroup, classGroupController *controllers.ClassGroupController) {
	classGroupRoute := rg.Group("/class_groups")
	classGroupRoute.GET("/", classGroupController.GetAll)
	classGroupRoute.GET("/:id", classGroupController.Find)
	classGroupRoute.POST("/", classGroupController.Create)
	classGroupRoute.PUT("/:id", classGroupController.Update)
	classGroupRoute.DELETE("/:id", classGroupController.Delete)
}
