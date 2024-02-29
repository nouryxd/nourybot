package main

import (
	"errors"
)

var (
	ErrGenericErrorMessage    = errors.New("something went wrong FeelsBadMan")
	ErrUserLevelNotInteger    = errors.New("user level must be a number")
	ErrCommandLevelNotInteger = errors.New("command level must be a number")
	ErrRecordNotFound         = errors.New("user not found in the database")
	ErrUserInsufficientLevel  = errors.New("user has insufficient level")
	ErrInternalServerError    = errors.New("internal server error")
	ErrDuringPasteUpload      = errors.New("could not upload paste")
)
