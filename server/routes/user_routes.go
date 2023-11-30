package routes

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	prefix := router.Group("/api")

	user := prefix.Group("/user", middlewares.AuthMiddleware())
	{
		user.GET("/", controllers.FindUsers)
		user.PATCH("/:id", controllers.UpdateUser)
		user.GET("/:id", controllers.FindUserByID)
	}

	auth := prefix.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.LoginUser)
		auth.POST("/logout", controllers.LogoutUser)
	}
}
