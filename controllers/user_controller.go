package controllers

import (
	"example/web-service-gin/db"
	"example/web-service-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"res": users})
}

func CreateUser(c *gin.Context) {
	// // Validate input
	// var input models.CreateUserInput
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// // Create user
	// user := models.User{Name: input.Name, Surname: input.Surname, DisplayName: input.DisplayName, Email: input.Email, Password: input.Password}
	// models.DB.Create(&user)

	// c.JSON(http.StatusOK, gin.H{"res": user})
	if gin.Mode() != gin.DebugMode {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var userData []models.User
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Użyj transakcji, aby zapewnić atomowość operacji
	tx := db.DB.Begin()
	for _, user := range userData {
		if err := tx.Create(&user).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to seed data"})
			return
		}
	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Dane zostały zaseedowane"})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	message := id + " is " + "good"
	c.JSON(http.StatusOK, gin.H{"success": message})
}

func FindUserByID(c *gin.Context) {
	var user models.User
	userID := c.Param("id")

	// Sprawdzenie poprawności ID
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Znalezienie użytkownika w bazie danych
	if err := db.DB.First(&user, "id = ?", parsedUserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
