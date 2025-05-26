package routes

import (
	"icomphub-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTechnologyRoutes(rg *gin.RouterGroup, technologyController *controllers.TechnologyController) {
	technologyRoute := rg.Group("/technologies")
	technologyRoute.GET("/", technologyController.GetAll)
	technologyRoute.GET("/:id", technologyController.Find)
	technologyRoute.POST("/", technologyController.Create)
	technologyRoute.PUT("/:id", technologyController.Update)
	technologyRoute.DELETE("/:id", technologyController.Delete)
}
