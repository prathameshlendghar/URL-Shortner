package auth

import "time"

type User struct {
	Id           int       `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
	IsActive     bool      `db:"is_active"`
}
