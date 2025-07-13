package routes

import (
	"icomphub-api/controllers"
	"icomphub-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterMemberRoutes(rg *gin.RouterGroup, memberController *controllers.MemberController) {
	memberRoute := rg.Group("/members")

	memberRoute.GET("/:id", memberController.Find)
	memberRoute.POST("/", middlewares.AuthMiddleware(), memberController.Create)
	memberRoute.PUT("/:id", middlewares.AuthMiddleware(), memberController.Update)
	memberRoute.DELETE("/:id", middlewares.AuthMiddleware(), memberController.Delete)
}
