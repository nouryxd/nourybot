package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lyx0/nourybot/internal/data"
)

func (app *application) showCommandHandler(w http.ResponseWriter, r *http.Request) {
	name, err := app.readCommandNameParam(r)
	if err != nil {
		app.logError(r, err)
		return
	}

	// Get the data for a specific movie from our helper method,
	// then check if an error was returned, and which.
	command, err := app.Models.Commands.Get(name)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.logError(r, err)
			return
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	app.Logger.Infow("GET Command",
		"Command", command,
	)
	err = app.writeJSON(w, http.StatusOK, envelope{"command": command}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type envelope map[string]interface{}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// Encode the data into JSON and return any errors if there were any.
	// Use MarshalIndent instead of normal Marshal so it looks prettier on terminals.
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Append a newline to make it prettier on terminals.
	js = append(js, '\n')

	// Iterate over the header map and add each header to the
	// http.ResponseWriter header map.
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Set `Content-Type` to `application/json` because go
	// defaults to `text-plain; charset=utf8`.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
