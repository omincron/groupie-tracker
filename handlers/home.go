package handlers

import (
	"groopie_local/models"
	"groopie_local/services"
	"log"
	"net/http"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Fetch cached data
	artistsFull, err := services.GetCachedData()
	if err != nil {
		log.Printf("Error fetching cached data: %v", err)
		renderTemplate(w, "error", TemplateData{
			Title:   "Error",
			Message: "Unable to load data. Please try again later.",
		})
		return
	}

	// Handle search query
	searchQuery := r.URL.Query().Get("search")
	var filtered []models.ArtistFull
	if searchQuery != "" {
		queryLower := strings.ToLower(searchQuery)
		for _, artist := range artistsFull {
			if strings.Contains(strings.ToLower(artist.Artist.Name), queryLower) {
				filtered = append(filtered, artist)
			}
		}

		if len(filtered) == 0 {
			log.Printf("No results found for search query: %s", searchQuery)
		}
	} else {
		filtered = artistsFull
	}

	// Prepare template data
	data := TemplateData{
		Title:       "Home - Groopie Tracker",
		Artists:     filtered,
		SearchQuery: searchQuery,
	}

	renderTemplate(w, "home", data)
}
