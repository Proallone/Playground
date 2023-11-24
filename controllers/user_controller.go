package controllers

import (
	"example/web-service-gin/db"
	"example/web-service-gin/models"
	"example/web-service-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FindUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"res": users})
}

func CreateUser(c *gin.Context) {

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

func RegisterUser(c *gin.Context) {
	// Validate input
	var input models.CreateUserInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := utils.HashPassword(input.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"An error occured": err})
		return
	}
	// Register new user
	user := models.User{Name: input.Name, Surname: input.Surname, DisplayName: input.DisplayName, Email: input.Email, Password: hash}
	result := db.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"An error occured: ": result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"New user created with ID": user.Base.ID})
}

func LoginUser(c *gin.Context) {
	var User models.User

	if err := c.BindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe dane logowania"})
		return
	}

	if err := db.DB.Table("users").Select("password").Where("email = ?", User.Email).First(&User).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"User not found": err})
		return
	}

	c.JSON(http.StatusOK, "D")

}
