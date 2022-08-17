package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// func (app *application) readIDParam(r *http.Request) (int64, error) {
// 	params := httprouter.ParamsFromContext(r.Context())
//
// 	// Use `ByName()` function to get the value of the "id" parameter from the slice.
// 	// The value returned by `ByName()` is always a string so we try to convert it to
// 	// base64 with a bit size of 64.
// 	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
// 	if err != nil || id < 1 {
// 		return 0, errors.New("invalid id parameter")
// 	}
//
// 	return id, nil
// }

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	// Limit the size of the requst body to 1MB.
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains malformed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains invalid JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains invalid JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains invalid JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body contains no data")

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		case err.Error() == "htto: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)

		case errors.As(err, &invalidUnmarshalError):
			app.Logger.Panic(err)
		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}
	return nil
}

func (app *application) readCommandNameParam(r *http.Request) (string, error) {
	params := httprouter.ParamsFromContext(r.Context())

	// Use `ByName()` function to get the value of the "id" parameter from the slice.
	// The value returned by `ByName()` is always a string so we try to convert it to
	// base64 with a bit size of 64.
	name := params.ByName("name")

	return name, nil
}

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
