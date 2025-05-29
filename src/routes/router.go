package routes

import (
	"icomphub-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userController *controllers.UserController,
	techlonogyController *controllers.TechnologyController,
	classGroupController *controllers.ClassGroupController,
	projectController *controllers.ProjectController,
) *gin.Engine {
	r := gin.Default()

	group := r.Group("/")
	RegisterUserRoutes(group, userController)
	RegisterTechnologyRoutes(group, techlonogyController)
	RegisterClassGroupRoutes(group, classGroupController)
	RegisterProjectRoutes(group, projectController)

	return r
}
