package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/prathameshlendghar/URL-Shortner/pkg/hash"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

var ErrInvalidCredentials = fmt.Errorf("Invalid Password")

var jwtSecret = []byte("super-secret-48-hour-sprint-key")

func (s *Service) RegisterUser(ctx context.Context, req RegisterUserReq) error {
	var hashedPassword string

	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("Failed to hash password: %s", err)
	}

	user := User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
		IsActive:     true,
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	err = s.repo.RegisterUser(ctx, user)
	return err
}

func (s *Service) LoginUser(ctx context.Context, email string, password string) (TokenResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return TokenResponse{}, err
	}

	if valid := hash.ComparePasswordHash(user.PasswordHash, password); !valid {
		return TokenResponse{}, ErrInvalidCredentials
	}

	tokenExpiration := time.Now().Add(24 * time.Hour)
	// 3. Passwords match! Generate the JWT.
	claims := CustomClaims{
		UserID: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenExpiration), // Token valid for 1 day
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token using the HS256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return TokenResponse{}, err
	}

	resp := TokenResponse{
		Token:     signedTokenString,
		ExpiresAt: tokenExpiration,
	}

	return resp, nil
}
