package routes

import "github.com/gin-gonic/gin"

func Route(route *gin.Engine) {

	route.GET("/users", getAllUser)

}
