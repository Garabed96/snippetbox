package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)                  // Log to Stdout
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) // Log to Stderr

	// Initialize a new instance of app struct, containing dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	// Use http.NewServeMux() function initialize a new multiplexer.
	// Then, register func home as the handler/controller for the "/"
	// Local scope ServerMux (SECURE) vs Global scope DefaultServeMux (INSECURE)

	//INSECURE GLOBAL SCOPE
	//http.HandleFunc("/", home)                    // Subtree path matching
	//http.HandleFunc("/snippet/view", snippetView) // Fixed path matching
	//http.HandleFunc("/snippet/create", snippetCreate)
	//
	//Use http.ListenAndServer to start a new web server (arg1, arg2):
	//arg1 := TCP network address where server will listen for requests
	//arg2 := servermux we created.
	//
	//log.Fatal() log the error message and exit the program.
	//(Any error returned by http.ListenAndServe() is always non-nil)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
