package server

import "time"

type CreateUrl struct {
	ActualUrl  string  `json:"actual_url"  validate:"required"`
	ShortAlias *string `json:"short_alias"`
	IsActive   *bool   `json:"is_active"`
}

type RedirectUrlResponse struct {
	ShortAlias string `json:"short_alias"`
}

type UrlResponse struct {
	UserId     int       `json:"user_id"`
	ActualUrl  string    `json:"actual_url"`
	ShortAlias string    `json:"short_alias"`
	Clicks     int       `json:"clicks"`
	CreatedAt  time.Time `json:"created_at"`
	IsActive   bool      `json:"is_active"`
}
