package main

import (
	"icomphub-api/controllers"
	"icomphub-api/docs"
	"log"
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

	docs.SwaggerInfo.Title = "IcompHub API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8016"

	server := gin.Default()

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	UserController := controllers.NewUserController()
	server.GET("/users", UserController.GetUsers)

	server.Run(":" + apiPort)
}