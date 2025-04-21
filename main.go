package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API VERSION": "1.0.0",
		})
	})
	router.Run(":3001") // listen and serve on 0.0.0.0:8080
}
