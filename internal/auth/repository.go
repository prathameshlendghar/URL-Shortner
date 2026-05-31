package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

const dbUniqueKeyViolationCode = "23505"

var ErrDuplicateKey = fmt.Errorf("DB unique key violation")

func (r *Repository) RegisterUser(ctx context.Context, userReq User) error {
	query := `INSERT INTO users (email, password_hash, is_active) VALUES ($1, $2, $3)`

	_, err := r.db.ExecContext(ctx, query, userReq.Email, userReq.PasswordHash, userReq.IsActive)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == dbUniqueKeyViolationCode {
			return ErrDuplicateKey
		}
		return err
	}

	return nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	var user User

	query := `SELECT id, email, password_hash FROM users WHERE email = $1 AND is_active = true`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.Email, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrInvalidCredentials
		}
		return User{}, err
	}

	return user, nil
}
