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
		HTMLDir:   *htmlDir,
		StaticDir: *staticDir,
	}

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}
