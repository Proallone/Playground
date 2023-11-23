package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("/", controllers.FindUsers)
		user.POST("/", controllers.CreateUser)
		user.PATCH("/:id", controllers.UpdateUser)
		user.GET("/:id", controllers.FindUserByID)
	}
}
