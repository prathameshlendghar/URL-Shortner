package url

import (
	"context"

	base62 "github.com/prathameshlendghar/URL-Shortner/pkg/base62"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateNewShortUrl(ctx context.Context, req CreateUrl, userId int) (UrlResponse, error) {
	urlReq := Url{
		ActualUrl: req.ActualUrl,
		IsActive:  true,
		UserId:    userId,
		Clicks:    0,
	}
	if req.IsActive != nil {
		urlReq.IsActive = *req.IsActive
	}
	if req.ShortAlias != nil {
		urlReq.ShortAlias = *req.ShortAlias
	} else {
		shortAlias, err := base62.GenerateRandomAlias()
		if err != nil {
			return UrlResponse{}, err
		}
		urlReq.ShortAlias = shortAlias
	}

	resp, err := s.repo.CreateNewShortUrl(ctx, urlReq)
	if err != nil {
		return UrlResponse{}, err
	}

	return UrlResponse{
		ActualUrl:  resp.ActualUrl,
		ShortAlias: resp.ShortAlias,
		Clicks:     resp.Clicks,
		CreatedAt:  resp.CreatedAt,
		IsActive:   resp.IsActive,
	}, nil
}

func (s *Service) GetActualUrl(ctx context.Context, shortAlias string) (string, error) {
	actualUrl, err := s.repo.GetActualUrl(ctx, shortAlias)
	if err != nil {
		return "", err
	}
	return actualUrl, err
}
