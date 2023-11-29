package main

import (
	"example/web-service-gin/db"
	"example/web-service-gin/docs"
	"example/web-service-gin/middlewares"
	"example/web-service-gin/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title First Go Gin API project
// @version 1.0
// @description This is a sample server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	router.Use(middlewares.LogMiddleware())
	router.Use(middlewares.ErrorHandlerMiddleware())

	db.ConnectDatabases()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.UserRoutes(router)

	router.Run()
}
