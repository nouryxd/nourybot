package data

import (
	"context"
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

// Get takes the login name for a channel and queries the database for an
// existing entry with that login value. If it exists it returns a
// pointer to a Channel.
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

// Insert takes in a channel struct and inserts it into the database.
func (c ChannelModel) Insert(login, id string) error {
	query := `
	INSERT INTO channels(login, twitchid)
	VALUES ($1, $2)
	ON CONFLICT (login)
	DO NOTHING
	RETURNING id, added_at;
	`

	args := []interface{}{login, id}

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
		return ErrChannelRecordAlreadyExists
	}

	return nil
}

// GetAll() returns a pointer to a slice of all channels (`[]*Channel`) in the database.
func (c ChannelModel) GetAll() ([]*Channel, error) {
	query := `
	SELECT id, added_at, login, twitchid
	FROM channels
	ORDER BY id`

	// Create a context with 3 seconds timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Use QueryContext() the context and query. This returns a
	// sql.Rows resultset containing our channels.
	rows, err := c.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	// Need to defer a call to rows.Close() to ensure the resultset
	// is closed before GetAll() returns.
	defer rows.Close()

	// Initialize an empty slice to hold the data.
	channels := []*Channel{}

	// Iterate over the resultset.
	for rows.Next() {
		// Initialize an empty Channel struct where we put on
		// a single channel value.
		var channel Channel

		// Scan the values onto the channel struct
		err := rows.Scan(
			&channel.ID,
			&channel.AddedAt,
			&channel.Login,
			&channel.TwitchID,
		)
		if err != nil {
			return nil, err
		}
		// Add the single movie struct onto the slice.
		channels = append(channels, &channel)
	}

	// When rows.Next() finishes call rows.Err() to retrieve any errors.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return channels, nil
}

// GetJoinable() returns a slice of channel names (Channel.Login) in the database.
func (c ChannelModel) GetJoinable() ([]string, error) {
	query := `
	SELECT login
	FROM channels
	ORDER BY id`

	// Create a context with 3 seconds timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Use QueryContext() the context and query. This returns a
	// sql.Rows resultset containing our channels.
	rows, err := c.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	// Need to defer a call to rows.Close() to ensure the resultset
	// is closed before GetAll() returns.
	defer rows.Close()

	// Initialize an empty slice to hold the data.
	channels := []string{}

	// Iterate over the resultset.
	for rows.Next() {
		// Initialize an empty Channel struct where we put on
		// a single channel value.
		var channel Channel

		// Scan the values onto the channel struct
		err := rows.Scan(
			&channel.Login,
		)
		if err != nil {
			return nil, err
		}
		// Add the single movie struct onto the slice.
		channels = append(channels, channel.Login)
	}

	// When rows.Next() finishes call rows.Err() to retrieve any errors.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return channels, nil
}

// Delete takes in a login name and queries the database and if there is an
// entry with that login name deletes the entry.
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
