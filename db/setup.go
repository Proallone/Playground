// models/setup.go

package db

import (
	"example/web-service-gin/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var dsn string

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}

	if os.Getenv("ENVIROMENT") == "docker" {
		dsn = os.Getenv("DOCKER_PG_CONN_STRING")
	} else {
		dsn = os.Getenv("LOCAL_PG_CONN_STRING")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	DB = database
}
