package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

// struct Models wraps the models, making them callable
// as app.models.Channels.Get(login)
type Models struct {
	Channels interface {
		Insert(channel *Channel) error
		Get(login string) (*Channel, error)
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Channels: ChannelModel{DB: db},
	}
}
