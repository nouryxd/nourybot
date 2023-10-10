package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound                = errors.New("record not found")
	ErrChannelRecordAlreadyExists    = errors.New("channel already in database")
	ErrEditConflict                  = errors.New("edit conflict")
	ErrCommandRecordAlreadyExists    = errors.New("command already exists")
	ErrLastFMUserRecordAlreadyExists = errors.New("lastfm connection already set")
	ErrUserAlreadyExists             = errors.New("user already in database")
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
		Insert(login, twitchId string) error
		Get(login string) (*User, error)
		Check(twitchId string) (*User, error)
		SetLevel(login string, level int) error
		GetLevel(twitchId string) (int, error)
		SetLocation(twitchId, location string) error
		GetLocation(twitchId string) (string, error)
		SetLastFM(login, lastfmUser string) error
		GetLastFM(login string) (string, error)
		Delete(login string) error
	}
	Commands interface {
		Get(name string) (*Command, error)
		Insert(command *Command) error
		Update(command *Command) error
		SetLevel(name string, level int) error
		SetCategory(name, category string) error
		SetHelp(name, helptext string) error
		Delete(name string) error
	}
	Timers interface {
		Get(name string) (*Timer, error)
		GetIdentifier(name string) (string, error)
		Insert(timer *Timer) error
		Update(timer *Timer) error
		GetAll() ([]*Timer, error)
		Delete(name string) error
	}
	Uploads interface {
		Insert(twitchLogin, twitchID, twitchMessage, twitchChannel, filehoster, downloadURL, identifier string)
		UpdateUploadURL(identifier, uploadURL string)
	}
	CommandsLogs interface {
		Insert(twitchLogin, twitchId, twitchChannel, twitchMessage, commandName string, uLvl int, identifier, rawMsg string)
	}
	SentMessagesLogs interface {
		Insert(twitchChannel, twitchMessage, identifier string)
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Channels:         ChannelModel{DB: db},
		Users:            UserModel{DB: db},
		Commands:         CommandModel{DB: db},
		Timers:           TimerModel{DB: db},
		Uploads:          UploadModel{DB: db},
		CommandsLogs:     CommandsLogModel{DB: db},
		SentMessagesLogs: SentMessagesLogModel{DB: db},
	}
}
