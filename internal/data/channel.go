package data

import (
	"database/sql"
	"errors"
	"time"
)

type Channel struct {
	ID       int       `json:"id"`
	AddedAt  time.Time `json:"-"`
	Login    string    `json:"login"`
	TwitchID string    `json:"twitchid"`
}

type ChannelModel struct {
	DB *sql.DB
}

func (c ChannelModel) Insert(channel *Channel) error {
	query := `
	INSERT INTO channels(login, twitchid)
	VALUES ($1, $2)
	ON CONFLICT (login)
	DO NOTHING
	RETURNING id, added_at;
	`

	args := []interface{}{channel.Login, channel.TwitchID}

	// Execute the query returning the number of affected rows.
	result, err := c.DB.Exec(query, args...)
	if err != nil {
		return err
	}

	// Check how many rows were affected.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordAlreadyExists
	}

	return nil
}

func (c ChannelModel) Get(login string) (*Channel, error) {
	query := `
	SELECT id, added_at, login, twitchid
	FROM channels
	WHERE login = $1`

	var channel Channel

	err := c.DB.QueryRow(query, login).Scan(
		&channel.ID,
		&channel.AddedAt,
		&channel.Login,
		&channel.TwitchID,
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

func (c ChannelModel) Delete(login string) error {
	// Prepare the statement.
	query := `
	DELETE FROM channels
	WHERE login = $1`

	// Execute the query returning the number of affected rows.
	result, err := c.DB.Exec(query, login)
	if err != nil {
		return err
	}

	// Check how many rows were affected.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// We want atleast 1, if it is 0 the entry did not exist.
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
