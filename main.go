package main

import (
	"groopie_local/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	// Set the port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register route handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artist/", handlers.ArtistHandler)
	http.HandleFunc("/search", handlers.SearchHandler)

	// Start the server
	log.Printf("Server started on :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
