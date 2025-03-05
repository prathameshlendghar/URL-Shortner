package models

import "time"

type NewUrlReq struct {
	LongUrl     string `json:"url"`
	ExpireAfter int16  `json:"expiry,omitempty"`
	Tag         string `json:"tag,omitempty"`
}

//TODO: Extend time for URL

type ShortUrlDB struct {
	Id        int64
	LongUrl   string
	ShortUrl  string
	CreatedAt time.Time
	ExpireAt  time.Time
	Tag       string
}
