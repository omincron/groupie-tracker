package handlers

import (
	"groopie_local/models"
	"groopie_local/services"
	"net/http"
	"strconv"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return // Stop further execution
	}

	artistsFull, err := services.GetCachedData()
	if err != nil {
		renderTemplate(w, "error", TemplateData{Title: "Error"})
		return // Stop further execution
	}

	var artistFull models.ArtistFull
	for _, artist := range artistsFull {
		if artist.ID == id {
			artistFull = artist
			break
		}
	}

	if artistFull.ID == 0 {
		http.NotFound(w, r)
		return // Stop further execution
	}

	// Pass the individual artist to the template
	data := TemplateData{
		Title:  artistFull.Artist.Name + " - Groopie Tracker",
		Artist: artistFull, // Populate the Artist field
	}

	renderTemplate(w, "artist", data)
}
