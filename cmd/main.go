package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	apiPort := os.Getenv("INTERNAL_API_PORT")

	if apiPort == "" {
		log.Fatal("Error while loading INTERNAL_API_PORT from .env file")
	}

	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "IcompHub API is running! hehe",
		})
	})

	server.Run(":" + apiPort)
}