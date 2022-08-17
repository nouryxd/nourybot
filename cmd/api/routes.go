package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/commands/:name", app.showCommandHandler)
	router.HandlerFunc(http.MethodPost, "/v1/commands", app.createCommandHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/commands/:name", app.updateCommandHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/commands/:name", app.deleteCommandHandler)
	return router
}
