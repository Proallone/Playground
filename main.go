package main

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)

	r.Run()
}
