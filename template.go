package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// TemplateCache to store parsed templates
var TemplateCache = map[string]*template.Template{}

// LoadTemplates loads all templates from the "templates" directory
func LoadTemplates() {
	pages, err := filepath.Glob("templates/*.html")
	if err != nil {
		log.Fatal("Error loading templates:", err)
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl, err := template.ParseFiles(page)
		if err != nil {
			log.Fatal("Error parsing template:", err)
		}
		TemplateCache[name] = tmpl
	}
}

// RenderTemplate renders a template with data
func RenderTemplate(w http.ResponseWriter, name string, data any) {
	tmpl, ok := TemplateCache[name]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// HomeHandler handles the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title": "Home Page",
		"Body":  "Welcome to the Go Web Template!",
	}
	RenderTemplate(w, "homepage.html", data)
}

// AboutHandler handles the about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "About Page",
		"Body":  "This is an advanced Go web template example.",
	}
	RenderTemplate(w, "about.html", data)
}

func main() {
	// Load templates at startup
	LoadTemplates()

	// Define routes
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)

	// Start the server
	log.Println("Starting server on :3030...")
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
