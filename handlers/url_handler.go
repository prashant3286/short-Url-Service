package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/prashant3286/short-url-service/models"
	"github.com/prashant3286/short-url-service/repository"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var url models.URL
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil || url.LongURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURL, err := repository.ShortenURL(&url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := models.ShortenedURLResponse{ShortURL: shortURL}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	shortURL := r.URL.Path[1:]
	longURL, err := repository.GetLongURL(shortURL)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
