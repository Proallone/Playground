package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// uuid  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
