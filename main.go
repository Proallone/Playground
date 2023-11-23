package main

import (
	"example/web-service-gin/db"
	"example/web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.ConnectDatabase()

	routes.UserRoutes(router)

	router.Run()
}
