package main

import (
	"fmt"
	"net/http"
)

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// The logError() method is a generic helper for logging an error message. Later in the
// book we'll upgrade this to use structured logging, and record additional information
// about the request including the HTTP method and URL.
func (app *application) logError(r *http.Request, err error) {
	app.Logger.Errorw("Error",
		"Request URI", r.RequestURI,
		"error", err,
	)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {

	// Write the response using the writeJSON() helper. If this happens to return an
	// error then log it, and fall back to sending the client an empty response with a
	// 500 Internal Server Error status code.
	fmt.Fprintf(w, "Error: %d", status)
}
