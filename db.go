package main

import (
	"sync"
	"time"
)

type URL struct {
    LongURL     string
    ShortURL    string
    Expiration  time.Time
    AccessCount int
}

var (
    urlStore = make(map[string]URL)
    mu       sync.Mutex
)

func saveURL(longURL, shortURL string, expiration time.Duration) {
    mu.Lock()
    defer mu.Unlock()
    urlStore[shortURL] = URL{LongURL: longURL, ShortURL: shortURL, Expiration: time.Now().Add(expiration)}
}

func getURL(shortURL string) (URL, bool) {
    mu.Lock()
    defer mu.Unlock()
    url, exists := urlStore[shortURL]
    if exists && time.Now().Before(url.Expiration) {
        url.AccessCount++
        urlStore[shortURL] = url // Update access count
        return url, true
    }
    return URL{}, false
}
