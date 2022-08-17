package main

import (
	"errors"
	"fmt"
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

func (app *application) createCommandHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Text     string `json:"text"`
		Category string `json:"category"`
		Level    int    `json:"level"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	command := &data.Command{
		Name:     input.Name,
		Text:     input.Text,
		Category: input.Category,
		Level:    input.Level,
	}

	err = app.Models.Commands.Insert(command)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/commands/%s", command.Name))

	err = app.writeJSON(w, http.StatusCreated, envelope{"command": command}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteCommandHandler(w http.ResponseWriter, r *http.Request) {
	name, err := app.readCommandNameParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.Models.Commands.Delete(name)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": fmt.Sprintf("command %s deleted", name)}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type envelope map[string]interface{}
