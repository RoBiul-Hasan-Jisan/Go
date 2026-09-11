package main

/*
net/http: Go's standard library ships a production-capable HTTP server and
client - no framework required for many real projects.

This file defines a small JSON API server AND shows how to call it (and any
other API) as a client. It's written so you can read it top-to-bottom;
to actually run the server, uncomment main()'s ListenAndServe call.
*/

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var todos = []Todo{
	{ID: 1, Text: "Learn Go basics", Done: true},
	{ID: 2, Text: "Master goroutines", Done: false},
}

// Handler: list todos as JSON
func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)

	case http.MethodPost:
		var t Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}
		t.ID = len(todos) + 1
		todos = append(todos, t)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// Middleware example: wraps a handler to log every request
func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s %s took %v", r.Method, r.URL.Path, time.Since(start))
	}
}

func setupServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/todos", withLogging(todosHandler))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	return mux
}

// ---------- Client side: calling an HTTP API ----------
func fetchURL(url string) (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading body failed: %w", err)
	}
	return string(body), nil
}

func main() {
	mux := setupServer()

	// To actually run the server, uncomment the line below and visit:
	//   http://localhost:8080/todos   (GET/POST)
	//   http://localhost:8080/health
	// go http.ListenAndServe(":8080", mux)
	_ = mux

	fmt.Println("Server routes configured: /todos, /health")
	fmt.Println("(ListenAndServe left commented out in this demo file)")

	// Example of the client side hitting a public API (network access required)
	// body, err := fetchURL("https://api.github.com/RoBiul-Hasan-Jisan/")
	// if err != nil {
	// 	fmt.Println("fetch error:", err)
	// 	return
	// }
	// fmt.Println("Response snippet:", body[:200])
}
