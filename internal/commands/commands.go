package commands

import "errors"

var (
	ErrInternalServerError     = errors.New("internal server error")
	ErrWeatherLocationNotFound = errors.New("location not found")
)
