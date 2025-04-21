package routes

import (
	usercontroller "example.com/gin-backend-api/controllers/user"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users") //http://localhost:3001/api/v1/users

	//http://localhost:3001/api/v1/users
	routerGroup.GET("/", usercontroller.GetAll)

	//http://localhost:3001/api/v1/users/register
	routerGroup.POST("/register", usercontroller.Register)
}
