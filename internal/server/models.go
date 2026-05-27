package server

import "time"

type User struct {
	Id           int       `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
	IsActive     bool      `db:"is_active"`
}

type Url struct {
	Id         int       `db:"id"`
	UserId     int       `db:"user_id"`
	ActualUrl  string    `db:"actual_url"`
	ShortAlias string    `db:"short_alias"`
	Clicks     int       `db:"clicks"`
	CreatedAt  time.Time `db:"created_at"`
	IsActive   bool      `db:"is_active"`
}
