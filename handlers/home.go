package handlers

import (
	"groopie_local/models"
	"groopie_local/services"
	"net/http"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return // Stop further execution
	}

	artistsFull, err := services.GetCachedData()
	if err != nil {
		renderTemplate(w, "error", TemplateData{Title: "Error"})
		return // Stop further execution
	}

	searchQuery := r.URL.Query().Get("search")
	if searchQuery != "" {
		var filtered []models.ArtistFull
		for _, artist := range artistsFull {
			if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(searchQuery)) {
				filtered = append(filtered, artist)
			}
		}
		artistsFull = filtered
	}

	data := TemplateData{
		Title:       "Home - Groopie Tracker",
		Artists:     artistsFull,
		SearchQuery: searchQuery,
	}

	renderTemplate(w, "home", data)
}
