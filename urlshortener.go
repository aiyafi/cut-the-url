package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURLLength = 6
	urlExpiration = 24 * time.Hour // Set expiration to 24 hours
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func generateShortURL() string {
	shortURL := make([]byte, shortURLLength)
	for i := range shortURL {
		shortURL[i] = characters[rand.Intn(len(characters))]
	}
	return string(shortURL)
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL()
	
	// Ensure the generated shortURL is unique
	for {
		_, exists := getURL(shortURL)
		if !exists {
			break
		}
		shortURL = generateShortURL()
	}
	
	saveURL(originalURL, shortURL, urlExpiration)
	fmt.Fprintf(w, "Shortened URL: http://localhost:8080/%s", shortURL)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	url, exists := getURL(shortURL)
	
	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url.LongURL, http.StatusFound)
}
