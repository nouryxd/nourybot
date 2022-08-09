package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound      = errors.New("record not found")
	ErrRecordAlreadyExists = errors.New("channel already in database")
	ErrUserAlreadyExists   = errors.New("user already in database")
)

// struct Models wraps the models, making them callable
// as app.models.Channels.Get(login)
type Models struct {
	Channels interface {
		Insert(channel *Channel) error
		Get(login string) (*Channel, error)
		GetAll() ([]*Channel, error)
		Delete(login string) error
	}
	Users interface {
		Insert(user *User) error
		Get(login string) (*User, error)
		SetLevel(login string, level int) error
		Delete(login string) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Channels: ChannelModel{DB: db},
		Users:    UserModel{DB: db},
	}
}
