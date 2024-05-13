package main

import (
	"crud_basic/db"
	"crud_basic/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.Route(server)

	server.Run(":8080")
}
