package main

import (
	"errors"
	"fmt"
	"net/http"
	"snippetbox.garonazarian.net/internal/models"
	"strconv"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/static/file.zip")
}

// We have refactored our home handler func to a method on the application struct
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // Use the http.NotFound() function to send a 404 response
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// We're using the new render helper here
	app.render(w, http.StatusOK, "home.tmpl", &templateData{
		Snippets: snippets,
	})
}

// ***** TLDR: Handler is a Controller from MVC, basically controls LOGIC and HTTP requests *****
// Define a home handler function which writes a byte slice containing
// "Hello from Mom" as the response body.

// TODO: VERY IMPORTANT TO UNDERSTAND THESE FUNDAMENTALS
// First param; assembles HTTP response and sends it to client
// Second param; pointer to a struct which holds info about current request
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// We're using the new render helper here too
	app.render(w, http.StatusOK, "view.tmpl", &templateData{
		Snippet: snippet,
	})
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
