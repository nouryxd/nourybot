package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Channel struct {
	ID       int
	Username string
	TwitchID string
	Added    time.Time
}
