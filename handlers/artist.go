package handlers

import (
	"groopie_local/models"
	"groopie_local/services"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		log.Printf("Invalid artist ID: %s", idStr)
		http.NotFound(w, r)
		return
	}

	artistsFull, err := services.GetCachedData()
	if err != nil {
		log.Printf("Error fetching cached data: %v", err)
		renderTemplate(w, "error", TemplateData{Title: "Error"})
		return
	}

	artistMap := make(map[int]models.ArtistFull)
	for _, artist := range artistsFull {
		artistMap[artist.Artist.ID] = artist
	}

	artistFull, found := artistMap[id]
	if !found {
		log.Printf("Artist with ID %d not found", id)
		http.NotFound(w, r)
		return
	}

	data := TemplateData{
		Title:  artistFull.Artist.Name + " - Groopie Tracker",
		Artist: artistFull,
	}

	renderTemplate(w, "artist", data)
}
