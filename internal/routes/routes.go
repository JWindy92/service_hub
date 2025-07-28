package routes

import (
	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
}

type UserInterface interface {
	GetUserByID(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	// CreateUser(c *gin.Context)
}

type ReviewInterface interface {
	GetReviewByID(c *gin.Context)
	CreateReview(c *gin.Context)
}

type LocationInterface interface {
	GetLocationByID(c *gin.Context)
	CreateLocation(c *gin.Context)
}

func RegisterRoutes(
	r *gin.Engine,
	auth AuthInterface,
	// users UserInterface,
	// reviews ReviewInterface,
	// locs LocationInterface,
) {
	r.POST("/signup", auth.SignUp)
	r.POST("/login", auth.Login)

	// r.GET("/users/:id", users.GetUserByID)
	// r.GET("/users", users.GetUserByEmail) //TODO: should make less ambiguous. Maybe implement a broader search function
	// r.POST("/users", users.CreateUser) //TODO: not sure if this is unnecessary yet, but /login performs the creation of users
}
