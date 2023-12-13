// db/postgres.go

package db

import (
	"example/web-service-gin/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToPG(connectionString string) (*gorm.DB, error) {
	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate postgres schema: %w", err)
	}

	fmt.Println("Server connected to postgres")
	return database, nil
}
