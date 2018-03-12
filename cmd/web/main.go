package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define command-line flags for network address and location of static files
	// directory.
	addr := flag.String("addr", ":4000", "HTTP network address")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")

	// Parse command line flags.
	flag.Parse()

	app := &App{
		HTMLDir: *htmlDir,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/new", app.NewSnippet)

	// Create a file server which serves files out of the "./ui/static" directory.
	// Path given to http.Dir function is relative to the project repository root.
	// Derefrence pointer passed from flag.String().
	fileServer := http.FileServer(http.Dir(*staticDir))

	// Use the mux.Handle() function to register the file server as the
	// handler for all URL paths that start with "/static/". For matching
	// paths, strip the "/static" prefix before the request reaches the
	// file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
