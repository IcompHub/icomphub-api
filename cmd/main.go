package main

import (
	"icomphub-api/controllers"
	"icomphub-api/db"
	"icomphub-api/docs"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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

	server := gin.Default()

	server.GET("/swagger/*any", func(ctx *gin.Context) {
		if ctx.Param("any") == "" || ctx.Param("any") == "/" {
			ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
			return
		}

		ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL)(ctx)
	})

	UserController := controllers.NewUserController(dbConnection)
	server.GET("/users", UserController.GetAllUsers)

	server.Run(":" + apiPort)
}