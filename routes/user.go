package routes

import "github.com/gin-gonic/gin"

func InitUserRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users") //localhost/api/v1/users/

	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "users",
		})
	})
}
