package main

import (
	"html/template"
	"path/filepath"
	"snippetbox.garonazarian.net/internal/models"
	"time"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Init template.FuncMap obj and store it in global vars (string-keyed map which acts as a lookup b/w
// names of our custom template func & func themselves
var functions = template.FuncMap{
	"humanDate": humanDate,
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

		// create an empty template set, use Funcs method to register FuncMap and then parse it
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
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
