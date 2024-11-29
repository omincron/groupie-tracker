package handlers

import (
	"groopie_local/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const templateDir = "templates"

type TemplateData struct {
	Title       string
	Artists     []models.ArtistFull
	Artist      models.ArtistFull
	SearchQuery string
	Message     string // Add this field
}

func renderTemplate(w http.ResponseWriter, name string, data TemplateData) {
	tmplPath := filepath.Join(templateDir, name+".html")
	layoutPath := filepath.Join(templateDir, "layout.html")

	tmpl, err := template.ParseFiles(layoutPath, tmplPath)
	if err != nil {
		log.Printf("Error parsing template (%s): %v", name, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		log.Printf("Error executing template (%s): %v", name, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
