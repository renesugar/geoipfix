package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	// Use default options
	handler = cors.Default().Handler(handler)
	http.ListenAndServe(":8080", handler)
}