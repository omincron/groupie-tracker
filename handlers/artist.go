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
		return
	}

	artistsFull, err := services.GetCachedData()
	if err != nil {
		renderTemplate(w, "error", TemplateData{Title: "Error"})
		return
	}

	var artistFull models.ArtistFull
	for _, artist := range artistsFull {
		if artist.Artist.ID == id { // Changed from artist.ID to artist.Artist.ID
			artistFull = artist
			break
		}
	}

	if artistFull.Artist.ID == 0 { // Changed from artistFull.ID to artistFull.Artist.ID
		http.NotFound(w, r)
		return
	}

	data := TemplateData{
		Title:  artistFull.Artist.Name + " - Groopie Tracker", // Changed from artistFull.Name
		Artist: artistFull,
	}

	renderTemplate(w, "artist", data)
}
