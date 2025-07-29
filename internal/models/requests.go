package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ? Good practice? Or just use UserProfile?
type UpdateProfileRequest struct {
	Address string `json:"address"`
}
