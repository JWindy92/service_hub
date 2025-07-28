package models

import "gorm.io/gorm"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `json:"name,omitempty"`
	Email        string `gorm:"unique" json:"email"`
	PasswordHash string `json:"password_hash"` // omit from JSON responses for security
}
