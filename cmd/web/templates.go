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

		// Parse the base template file into a template set
		ts, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		// Call ParseGlob() *on this template set* to add any partials.
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		// Call ParseFiles() *on this template set* to add the page template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add template set to the map, using name of the page
		cache[name] = ts
	}
	// return map
	return cache, nil
}
