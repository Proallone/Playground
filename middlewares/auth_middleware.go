package middlewares

import (
	"context"
	"example/web-service-gin/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization") // Pobranie tokenu JWT z nagłówka
		if isAuthenticated(token) {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access!"})
	}
}

func isAuthenticated(token string) bool {
	exists, err := db.REDIS.Exists(context.Background(), token).Result()
	if err != nil {
		return false
	}
	return exists == 1
}
