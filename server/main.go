package main

import (
	"example/web-service-gin/db"
	"example/web-service-gin/docs"
	logger "example/web-service-gin/log"
	"example/web-service-gin/middlewares"
	"example/web-service-gin/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	err := godotenv.Load()
	// log.Log.InitLogrus()

	if err != nil {
		log.Fatal("Error occured. No .env file found. Please check the path.")
	}

	logger.Logger.Info("Logger initialized")
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"
	router.Use(middlewares.LogMiddleware())
	router.Use(middlewares.ErrorHandlerMiddleware())

	db.ConnectDatabases()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.UserRoutes(router)

	router.Run()
	fmt.Println("Server is running!")
}
