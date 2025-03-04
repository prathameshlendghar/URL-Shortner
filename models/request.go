package models

type NewUrlReq struct {
	LongUrl   string `json:"url"`
	ExpireAt  string `json:"expiry,omitempty"`
	CreatedAt string `json:"created,omitempty"`
	Tag       string `json:"tag,omitempty"`
}

//TODO: Extend time for URL
