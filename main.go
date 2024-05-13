package main

import (
	"crud_basic/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	routes.Route(server)
	server.Run(":8080")
}
