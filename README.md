# URL Shortener

A simple URL shortener service built in Go. This project demonstrates basic backend development skills, including HTTP handling, in-memory data storage, and URL redirection.

## Features

- Shorten long URLs to a 6-character code
- Redirect short URLs to their original long URLs
- In-memory storage with expiration (24 hours)
- Basic concurrency handling

## Getting Started

1. Ensure you have Go installed on your system.
2. Clone this repository.
3. Run the server:

```bash
go run main.go
```
4. The server will start on `http://localhost:8080`.

## Usage

### Shorten a URL

Send a GET request to `/shorten` with a `url` query parameter:

```bash
curl "http://localhost:8080/shorten?url=https://www.example.com"
```

### Redirect to Original URL

Send a GET request to the short URL:

```bash
curl "http://localhost:8080/short-url"
```


## Project Structure

- `main.go`: Entry point and server setup
- `urlshortener.go`: Core URL shortening and redirection logic
- `db.go`: In-memory storage implementation

## Learning Outcomes

This project demonstrates:

- HTTP handling in Go
- Basic data structures and algorithms
- Concurrency with mutexes
- URL parsing and generation
- Simple persistence layer simulation
