package models

import (
	"time"

	"github.com/google/uuid"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
type User struct {
	Base
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email" gorm:"uniqueIndex"`
	Password    string `json:"password"`
}

type CreateUserInput struct {
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	DisplayName string `json:"display_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
