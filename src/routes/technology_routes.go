package routes

import (
	"icomphub-api/controllers"
	"icomphub-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterTechnologyRoutes(rg *gin.RouterGroup, technologyController *controllers.TechnologyController) {
	technologyRoute := rg.Group("/technologies")
	technologyRoute.GET("/", technologyController.GetAll)
	technologyRoute.GET("/:id", technologyController.Find)
	technologyRoute.POST("/", technologyController.Create)
	technologyRoute.PUT("/:id", technologyController.Update)
	technologyRoute.DELETE("/:id", technologyController.Delete)
	technologyRoute.PUT("/image/:id", middlewares.AuthMiddleware(), technologyController.UpdateImage)
	technologyRoute.DELETE("/image/:id", middlewares.AuthMiddleware(), technologyController.DeleteImage)
	technologyRoute.GET("/image/:id", technologyController.GetImage)
}
