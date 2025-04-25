package routes

import (
	"example.com/gin-backend-api/configs"
	"github.com/gin-gonic/gin"
)

func InitHomeRoutes(rg *gin.RouterGroup) {
	// connect db
	configs.Connection()

	routerGroup := rg.Group("/")

	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API VERSION": "1.0.0",
		})
	})
}
