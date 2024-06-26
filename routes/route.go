package routes

import "github.com/gin-gonic/gin"

func Route(route *gin.Engine) {

	route.POST("/users", createUser)

	route.GET("/users", getUsers)

	route.PUT("/users/:name", updateUser)

}
