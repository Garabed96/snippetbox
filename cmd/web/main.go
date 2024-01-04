package main

import (
	"log"
	"net/http"
)

func main() {
	// Use http.NewServeMux() function initialize a new multiplexer.
	// Then, register func home as the handler/controller for the "/"
	// Local scope ServerMux (SECURE) vs Global scope DefaultServeMux (INSECURE)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)                    // Subtree path matching
	mux.HandleFunc("/snippet/view", snippetView) // Fixed path matching
	mux.HandleFunc("/snippet/create", snippetCreate)
	// INSECURE GLOBAL SCOPE
	//http.HandleFunc("/", home)                    // Subtree path matching
	//http.HandleFunc("/snippet/view", snippetView) // Fixed path matching
	//http.HandleFunc("/snippet/create", snippetCreate)

	// Use http.ListenAndServer to start a new web server (arg1, arg2):
	// arg1 := TCP network address where server will listen for requests
	// arg2 := servermux we created.

	// log.Fatal() log the error message and exit the program.
	// (Any error returned by http.ListenAndServe() is always non-nil)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
