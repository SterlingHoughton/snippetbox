package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/new", NewSnippet)

	// Create a file server which serves files out of the "./ui/static" directory.
	// Path given to http.Dir function is relative to the project repository root.
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Use the mux.Handle() function to register the file server as the
	// handler for all URL paths that start with "/static/". For matching
	// paths, strip the "/static" prefix before the request reaches the 
	// file server.

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
