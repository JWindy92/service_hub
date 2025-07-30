package services

import (
	"errors"
	"log"
	"log/slog"

	"github.com/JWindy92/service_hub/service_hub/internal/common"
	"github.com/JWindy92/service_hub/service_hub/internal/models"
	"github.com/JWindy92/service_hub/service_hub/internal/utils"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(conn *gorm.DB) *UserService {
	return &UserService{DB: conn}
}

func (s *UserService) CreateUser(user *models.User) error {
	utils.PrettyPrint(user)
	err := s.DB.Where("email = ?", user.Email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// user does NOT exist
		return s.DB.Create(user).Error
	} else if err != nil {
		// unexpected error
		return err
	} else {
		return common.ErrUserExists
	}
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := s.DB.Preload("Profile").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	slog.Info("Attempting to find user", slog.String("email", email))
	var user models.User
	if err := s.DB.Preload("Profile").Where("email = ?", email).First(&user).Error; err != nil {
		log.Println("No user found")
		return nil, err
	}
	// if err := s.DB.Preload("Reviews").Where("email = ?", email).First(&user).Error; err != nil {
	// 	log.Println("No user found")
	// 	return nil, err
	// }
	slog.Info("User found")
	return &user, nil
}

func (s *UserService) UpdateUserProfile(userId int, prof *models.UpdateProfileRequest) error {
	slog.Info("Updating user profile", slog.Int("userId", userId))
	return s.DB.Model(&models.UserProfile{}).Where("user_id = ?", userId).Updates(&prof).Error
}
