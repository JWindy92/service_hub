package models

import "gorm.io/gorm"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User struct defines the user model.
type User struct {
	gorm.Model
	// ID           uint         `gorm:"primaryKey" json:"id"`
	Name         string       `json:"name,omitempty"`
	Email        string       `gorm:"unique" json:"email"`
	PasswordHash string       `json:"password_hash"`
	Profile      *UserProfile `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// UserProfile struct defines the user's profile model.
type UserProfile struct {
	gorm.Model
	// ID         uint   `gorm:"primaryKey" json:"id"`
	IsComplete bool   `gorm:"default:false"` // Default to false
	Address    string `json:"Address"`
	UserID     uint   // Foreign key to User
}
