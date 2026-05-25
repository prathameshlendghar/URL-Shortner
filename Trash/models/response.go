package models

type ShortUrlResp struct {
	LongUrl   string `json:"longurl"`
	ShortUrl  string `json:"shorturl"`
	CreatedAt string `json:"createdat,omitempty"`
	ExpiresAt string `json:"expires,omitempty"`
	Tag       string `json:"tag,omitempty"`
}
