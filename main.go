package main

import (
	"log"
	"os"

	"example.com/gin-backend-api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := SetupRouter()
	router.Run(":" + os.Getenv("GO_PORT"))
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	apiV1 := router.Group("/api/v1") //localhost:3000/api/v1/...
	routes.InitHomeRoutes(apiV1)
	routes.InitUserRoutes(apiV1)
	return router
}
