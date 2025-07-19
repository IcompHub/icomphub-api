package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"

	"icomphub-api/controllers"
	"icomphub-api/db"
	"icomphub-api/docs"
	"icomphub-api/files"
	"icomphub-api/repositories"
	"icomphub-api/routes"
	"icomphub-api/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	isInProd := os.Getenv("IN_PRODUCTION")

	if isInProd != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error while loading .env file")
		}
	}

	apiPort := os.Getenv("INTERNAL_API_PORT")

	if apiPort == "" {
		log.Fatal("Error while loading INTERNAL_API_PORT from .env file")
	}

	swaggerApiUrl := os.Getenv("SWAGGER_API_URL")

	if swaggerApiUrl == "" {
		log.Fatal("Error while loading SWAGGER_API_URL from .env file")
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	dbConnection, err := db.ConnectDB(host, port, user, password, dbname)
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.Title = "IcompHub API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = swaggerApiUrl

	swaggerURL := ginSwagger.URL("/swagger/doc.json")

	uploadService := files.NewLocalFileUploadService("./static/uploads")

	userRepository := repositories.NewUserRepository(dbConnection)
	userService := services.NewUserService(userRepository, uploadService)
	userController := controllers.NewUserController(userService)

	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)

	technologyRepository := repositories.NewTechnologyRepository(dbConnection)
	technologyService := services.NewTechnologyService(technologyRepository, uploadService)
	technologyController := controllers.NewTechnologyController(technologyService)

	classGroupRepository := repositories.NewClassGroupRepository(dbConnection)
	classGroupService := services.NewClassGroupService(classGroupRepository)
	classGroupController := controllers.NewClassGroupController(classGroupService)

	projectImageRepository := repositories.NewProjectImageRepository(dbConnection)

	roleRepository := repositories.NewRoleRepository(dbConnection)
	roleService := services.NewRoleService(roleRepository)
	roleController := controllers.NewRoleController(roleService)

	memberRepository := repositories.NewMemberRepository(dbConnection)
	memberService := services.NewMemberService(memberRepository, roleService)
	memberController := controllers.NewMemberController(memberService)

	projectRepository := repositories.NewProjectRepository(dbConnection)
	projectService := services.NewProjectService(projectRepository, uploadService, projectImageRepository, technologyService, memberService)
	projectController := controllers.NewProjectController(projectService)

	router := routes.SetupRouter(userController, technologyController, classGroupController, projectController, authController, roleController, memberController)

	// Need to be updated in the future to allow only authorized clients in production
	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
	}

	router.Use(cors.New(config))

	router.GET("/swagger/*any", func(ctx *gin.Context) {
		if ctx.Param("any") == "" || ctx.Param("any") == "/" {
			ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
			return
		}

		ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL)(ctx)
	})

	error := router.Run(":" + apiPort)

	if error != nil {
		log.Fatal(error.Error())
	}
}
