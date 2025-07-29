package handlers

import (
	"net/http"
	"strconv"

	"github.com/JWindy92/service_hub/service_hub/internal/common"
	"github.com/JWindy92/service_hub/service_hub/internal/models"
	"github.com/JWindy92/service_hub/service_hub/internal/services"
	"github.com/JWindy92/service_hub/service_hub/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	GetUserByID(string) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	UpdateUserProfile(int, *models.UpdateProfileRequest) error
	CreateUser(*models.User) error //TODO: dont need here
}

type UserHandler struct {
	Service UserServiceInterface
}

func NewDefaultUserHandler(db common.DBInterface) *UserHandler {
	return &UserHandler{
		Service: services.NewUserService(db.ConnectDB()),
	}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.Service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	resp := models.UserPublicDataResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Profile: user.Profile,
	}

	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) UpdateUserProfile(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	var req models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	utils.PrettyPrint(req)

	err = h.Service.UpdateUserProfile(userId, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error updating user profile"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "profile updated",
		"profile": models.UpdateProfileResponse{
			UserId: uint(userId),
		},
	})
}

// // TODO: Check for SQL injection (does gorm handle it natively?)
// func (h *UserHandler) GetUserByEmail(c *gin.Context) {
// 	email := c.Query("email") // or use Param() if you're using path params

// 	if email == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
// 		return
// 	}

// 	user, err := h.Service.GetUserByEmail(email)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }
