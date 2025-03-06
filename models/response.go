package models

type ShortUrlResp struct {
	LongUrl   string `json long-url`
	ShortUrl  string `json short-url`
	CreatedAt string `json created-at`
	ExpiresAt string `json expires-at`
	Tag       string `json tag`
}
