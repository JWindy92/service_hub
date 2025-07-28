package models

type LoginSuccess struct {
	Message string       `json:"message"`
	Token   string       `json:"token"`
	User    UserResponse `json:"user"`
}

type UserResponse struct {
	ID      uint        `json:"id"`
	Name    string      `json:"name"`
	Email   string      `json:"email"`
	Profile UserProfile `json:"profile"`
}
