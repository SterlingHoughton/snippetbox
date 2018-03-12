package main

import (
	"log"
	"net/http"
	"runtime/debug"
)

// ServerError writes an error message and stack trace to the log, then
// sends a generic 500 Internal Server Error response to the user.
func (app *App) ServerError(w http.ResponseWriter, err error) {
	log.Printf("%s\n%s", err.Error(), debug.Stack())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// ClientError sends a specific status code and corresponding description
// to the user.
func (app *App) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// NotFound Helper. Convient wrapper around ClientError for 404's
func (app *App) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
