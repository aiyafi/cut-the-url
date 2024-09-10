package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    // Define routes
    http.HandleFunc("/shorten", shortenURLHandler) // URL shortening route
    http.HandleFunc("/", redirectHandler)          // URL redirect route

    fmt.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server on port 8080
}
