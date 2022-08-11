package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound             = errors.New("record not found")
	ErrChannelRecordAlreadyExists = errors.New("channel already in database")
	ErrCommandRecordAlreadyExists = errors.New("command already exists")
	ErrUserAlreadyExists          = errors.New("user already in database")
)

// struct Models wraps the models, making them callable
// as app.models.Channels.Get(login)
type Models struct {
	Channels interface {
		Insert(channel *Channel) error
		Get(login string) (*Channel, error)
		GetAll() ([]*Channel, error)
		GetJoinable() ([]string, error)
		Delete(login string) error
	}
	Users interface {
		Insert(user *User) error
		Get(login string) (*User, error)
		SetLevel(login string, level int) error
		Delete(login string) error
	}
	Commands interface {
		Get(name string) (*Command, error)
		Insert(name, text string) error
		SetLevel(name string, level int) error
		Delete(name string) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Channels: ChannelModel{DB: db},
		Users:    UserModel{DB: db},
		Commands: CommandModel{DB: db},
	}
}
