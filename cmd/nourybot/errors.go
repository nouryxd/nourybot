package main

import (
	"errors"
)

var (
	ErrUserLevelNotInteger    = errors.New("user level must be a number")
	ErrCommandLevelNotInteger = errors.New("command level must be a number")
	ErrRecordNotFound         = errors.New("user not found in the database")
	ErrUserInsufficientLevel  = errors.New("user has insufficient level")
)
