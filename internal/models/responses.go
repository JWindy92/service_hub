package models

type LoginSuccess struct {
	Message string            `json:"message"`
	Token   string            `json:"token"`
	User    UserLoginResponse `json:"user"`
}

type UserLoginResponse struct {
	ID uint `json:"id"`
	// Name    string      `json:"name"`
	// Email   string      `json:"email"`
	// Profile UserProfile `json:"profile"` //TODO: probably unnecessary to return with successful login
}

type UserPublicDataResponse struct {
	ID      uint         `json:"id"` //? Should the "Public" response be purely data needed to present to the user?
	Name    string       `json:"name,omitempty"`
	Email   string       `gorm:"unique" json:"email"`
	Profile *UserProfile `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UpdateProfileResponse struct {
	ID     uint `json:"id"`
	UserId uint `json:"user_id"`
}
