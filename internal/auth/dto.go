package auth

type RegisterUserReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	IsActive *bool   `json:"is_active"`
}