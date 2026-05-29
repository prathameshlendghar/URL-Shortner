package auth

import (
	"context"
	"fmt"

	"github.com/prathameshlendghar/URL-Shortner/pkg/hash"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

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
