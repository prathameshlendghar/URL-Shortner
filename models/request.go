package models

import "time"

type NewUrlReq struct {
	LongUrl     string `json:"url"`
	ExpireAfter int32  `json:"expiry"`
	Tag         string `json:"tag"`
}

type GetInfoReq struct {
	ShortUrl string `json:"shorturl"`
}

//TODO: Extend time for URL

type UpdateReq struct {
	ShortUrl    string `json:"shorturl"`
	LongUrl     string `json:"longurl"`
	ExpireAfter int32  `json:"expiry"`
	Tag         string `json:"tag"`
}

type UpdateReqDB struct {
	LongUrl  *string
	ExpireAt *time.Time
	Tag      *string
}

type ShortUrlDB struct {
	Id        int64
	LongUrl   string
	ShortUrl  string
	CreatedAt time.Time
	ExpireAt  time.Time
	Tag       string
}
