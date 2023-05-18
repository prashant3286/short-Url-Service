package main

import (
	"log"
	"net/http"

	"github.com/prashant3286/short-url-service/handlers"
)

func main() {
	http.HandleFunc("/shorten", handlers.ShortenURL)
	http.HandleFunc("/", handlers.RedirectURL)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
