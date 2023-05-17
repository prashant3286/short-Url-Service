package models

type URL struct {
	LongURL string `json:"longURL"`
}

type ShortenedURLResponse struct {
	ShortURL string `json:"shortURL"`
}
