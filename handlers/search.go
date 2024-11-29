package handlers

import (
	"encoding/json"
	"groopie_local/models"
	"groopie_local/services"
	"log"
	"net/http"
	"strings"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		log.Println("Search query parameter 'q' is missing")
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	artists, err := services.FetchArtists()
	if err != nil {
		log.Printf("Error fetching artist data: %v", err)
		http.Error(w, "Unable to fetch artists", http.StatusInternalServerError)
		return
	}

	var filtered []models.Artist
	queryLower := strings.ToLower(query)
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), queryLower) {
			filtered = append(filtered, artist)
		}
	}

	response := map[string]interface{}{
		"query":         query,
		"total_results": len(filtered),
		"artists":       filtered,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Unable to encode response", http.StatusInternalServerError)
	}
}
