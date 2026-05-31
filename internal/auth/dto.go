package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RegisterUserReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	IsActive *bool  `json:"is_active"`
}

type CustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenResponse struct {
	Token     string
	ExpiresAt time.Time
}
