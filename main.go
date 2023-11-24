package main

import (
	"example/web-service-gin/db"
	"example/web-service-gin/middlewares"
	"example/web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.Logger())

	db.ConnectDatabase()

	routes.UserRoutes(router)

	router.Run()
}
