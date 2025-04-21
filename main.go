package main

import (
	"log"
	"os"

	"example.com/gin-backend-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := SetupRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router.Run(":" + os.Getenv("GO_PORT"))
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	apiV1 := router.Group("/api/v1") //localhost:3000/api/v1/...
	routes.InitHomeRoutes(apiV1)
	routes.InitUserRoutes(apiV1)
	return router
}
