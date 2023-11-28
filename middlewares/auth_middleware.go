package middlewares

import (
	"example/web-service-gin/utils"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		extractedToken := strings.Split(token, "Bearer ")
		if isAuthenticated(extractedToken[1]) {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access!"})
	}
}

func isAuthenticated(token string) bool {
	userID, err := utils.ValidateToken(token, os.Getenv("JWT_TOKEN"))
	fmt.Print(userID)
	return err == nil
}
