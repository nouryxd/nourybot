package main

import (
	"database/sql"
	"errors"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/data"
	"go.uber.org/zap"
)

var (
	ErrUserLevelNotInteger   = errors.New("user level must be a number")
	ErrRecordNotFound        = errors.New("user not found in the database")
	ErrUserInsufficientLevel = errors.New("user has insufficient level")
)

type config struct {
	twitchUsername string
	twitchOauth    string
	environment    string
	db             struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type Application struct {
	TwitchClient *twitch.Client
	Logger       *zap.SugaredLogger
	Db           *sql.DB
	Models       data.Models
}
