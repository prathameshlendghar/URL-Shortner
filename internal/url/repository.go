package url

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

var ErrUrlNotFound = fmt.Errorf("Url not found")
var ErrUrlNotActive = fmt.Errorf("Url is not accessible")

func (r *Repository) CreateNewShortUrl(ctx context.Context, req Url) (Url, error) {
	query := `INSERT INTO urls (user_id, actual_url, short_alias, clicks, is_active) 
				VALUES ($1, $2, $3, $4, $5)
				RETURNING created_at`

	err := r.db.QueryRowContext(ctx, query, req.UserId, req.ActualUrl, req.ShortAlias, req.Clicks, req.IsActive).Scan(&req.CreatedAt)
	if err != nil {
		return Url{}, err
	}

	return req, nil
}

func (r *Repository) GetActualUrl(ctx context.Context, shortAlias string) (string, error) {
	query := `UPDATE urls 
		    SET clicks = clicks + 1
    		WHERE short_alias = $1
    		RETURNING actual_url, is_active;`

	var actualUrl string
	var isActive bool

	err := r.db.QueryRowContext(ctx, query, shortAlias).Scan(&actualUrl, &isActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrUrlNotFound
		}
		return "", err
	}
	if !isActive {
		return "", ErrUrlNotActive
	}

	return actualUrl, nil
}
