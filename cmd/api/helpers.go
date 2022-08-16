package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	// Use `ByName()` function to get the value of the "id" parameter from the slice.
	// The value returned by `ByName()` is always a string so we try to convert it to
	// base64 with a bit size of 64.
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) readCommandNameParam(r *http.Request) (string, error) {
	params := httprouter.ParamsFromContext(r.Context())
	app.Logger.Info(r.Context())

	// Use `ByName()` function to get the value of the "id" parameter from the slice.
	// The value returned by `ByName()` is always a string so we try to convert it to
	// base64 with a bit size of 64.
	name := params.ByName("name")

	return name, nil
}
