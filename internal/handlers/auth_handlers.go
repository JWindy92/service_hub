package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/JWindy92/service_hub/service_hub/internal/common"
	"github.com/JWindy92/service_hub/service_hub/internal/models"
	"github.com/JWindy92/service_hub/service_hub/internal/services"
	"github.com/JWindy92/service_hub/service_hub/internal/utils"
	"github.com/gin-gonic/gin"
)

type PasswordHashInterface interface {
	Hash(string) (string, error)
	CheckHash(pw string, hash string) bool
}

/* -------------------------------------------------------------------------- */
/*                    Basic Auth Passthrough Implementation                   */
/* -------------------------------------------------------------------------- */

type UserAuthInterface interface {
	GetUserByEmail(string) (*models.User, error)
	CreateUser(*models.User) error
}

// Basic implementation of the routes.AuthInterface to allow for dev work without full authentication logic
// ! Not for production
type AuthPassthroughHandler struct {
	UserService UserAuthInterface
	Hasher      PasswordHashInterface
}

func NewAuthPassthroughHandler(db common.DBInterface, hasher PasswordHashInterface) *AuthPassthroughHandler {
	return &AuthPassthroughHandler{
		UserService: services.NewUserService(db.ConnectDB()),
		Hasher:      hasher,
	}
}

func (a *AuthPassthroughHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	utils.PrettyPrint(req)

	user, err := a.UserService.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("unexpected error: %v", err)})
	}

	if a.Hasher.CheckHash(req.Password, user.PasswordHash) {
		resp := models.LoginSuccess{
			Message: "login successful",
			Token:   "1234-567-891011",
			User: models.UserLoginResponse{
				ID: user.ID,
			},
		}
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
	}
}

func (a *AuthPassthroughHandler) SignUp(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// utils.PrettyPrint(req)
	hash, err := a.Hasher.Hash(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error performing password hash"})
		return
	}
	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hash, // replace with actual hash function
		Profile: &models.UserProfile{
			IsComplete: false,
		},
	}

	err = a.UserService.CreateUser(&user)
	if err != nil {
		if errors.Is(err, common.ErrUserExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
			return
		}
	}

	// utils.PrettyPrint(user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
