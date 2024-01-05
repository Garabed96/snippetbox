package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r) // Use the http.NotFound() function to send a 404 response
		return
	}
	// Create a slice of paths to the two files. Note Base Layout must be first
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/home.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// ***** TLDR: Handler is a Controller from MVC, basically controls LOGIC and HTTP requests *****
// Define a home handler function which writes a byte slice containing
// "Hello from Mom" as the response body.

// TODO: VERY IMPORTANT TO UNDERSTAND THESE FUNDAMENTALS
// First param; assembles HTTP response and sends it to client
// Second param; pointer to a struct which holds info about current request
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet"))
}
