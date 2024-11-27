package handlers

import (
	"groopie_local/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type TemplateData struct {
	Title       string
	Artists     []models.ArtistFull
	Artist      models.ArtistFull // Add this to store a single artist
	SearchQuery string
}

func renderTemplate(w http.ResponseWriter, name string, data TemplateData) {
	tmplPath := filepath.Join("templates", name+".html")
	layoutPath := filepath.Join("templates", "layout.html")

	// Parse the templates
	tmpl, err := template.ParseFiles(layoutPath, tmplPath)
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return // Prevent further execution
	}

	// Execute the layout template
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return // Prevent further execution
	}
}
