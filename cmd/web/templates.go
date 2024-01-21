package main

import (
	"html/template"
	"path/filepath"
	"snippetbox.garonazarian.net/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Creating in-memory map
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			page,
		}

		// Parse files into template set
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		// Add template set to the map, using name of the page
		cache[name] = ts
	}
	// return map
	return cache, nil
}
