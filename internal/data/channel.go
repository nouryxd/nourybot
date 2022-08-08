package data

import (
	"database/sql"
	"errors"
	"time"
)

type Channel struct {
	ID       int       `json:"id"`
	JoinedAt time.Time `json:"-"`
	Login    string    `json:"login"`
	TwitchID string    `json:"twitchid"`
	Announce bool      `json:"announce"`
}

type ChannelModel struct {
	DB *sql.DB
}

func (c ChannelModel) Insert(channel *Channel) error {
	query := `
	INSERT INTO channels (login, twitchid, announce)
	VALUES ($1, $2, $3)
	RETURNING id, created_at`

	args := []interface{}{channel.Login, channel.TwitchID, channel.Announce}

	return c.DB.QueryRow(query, args...).Scan(&channel.ID, &channel.JoinedAt)
}

func (c ChannelModel) Get(login string) (*Channel, error) {
	query := `
	SELECT id, joined_at, login, twitchid, announce
	FROM channels
	WHERE login = $1`

	var channel Channel

	err := c.DB.QueryRow(query, login).Scan(
		&channel.ID,
		&channel.JoinedAt,
		&channel.Login,
		&channel.TwitchID,
		&channel.Announce,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &channel, nil
}
