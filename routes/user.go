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

	//http://localhost:3001/api/v1/users/login
	routerGroup.POST("/login", usercontroller.Login)

	//http://localhost:3001/api/v1/users/1    ส่ง params
	routerGroup.GET("/:id", usercontroller.GetById)

	//http://localhost:3001/api/v1/users/search?fullname=j    ส่ง query
	routerGroup.GET("/search", usercontroller.SearchByFullname)
}
