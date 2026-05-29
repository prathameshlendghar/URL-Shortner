package url

import "time"

type Url struct {
	Id         int       `db:"id"`
	UserId     int       `db:"user_id"`
	ActualUrl  string    `db:"actual_url"`
	ShortAlias string    `db:"short_alias"`
	Clicks     int       `db:"clicks"`
	CreatedAt  time.Time `db:"created_at"`
	IsActive   bool      `db:"is_active"`
}