package entity

import "time"

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []Shorten
}

type Shorten struct {
	ID          int       `json:"id"`
	OriginalUrl string    `json:"original_url"`
	ShortenUrl  string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
}
