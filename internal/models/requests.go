package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ? Good practice? Or just use UserProfile?
type UpdateProfileRequest struct {
	FName   string `json:"first_name"`
	LName   string `json:"last_name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
