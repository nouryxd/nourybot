package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Timer struct {
	ID         int    `json:"id" redis:"timer-id"`
	Name       string `json:"name" redis:"timer-name"`
	Identifier string `json:"identifier"`
	Text       string `json:"text" redis:"timer-text"`
	Channel    string `json:"channel" redis:"timer-channel"`
	Repeat     string `json:"repeat" redis:"timer-repeat"`
}

type TimerModel struct {
	DB *sql.DB
}

// Get tries to find a timer with the supplied name.
func (t TimerModel) Get(name string) (*Timer, error) {
	query := `
	SELECT id, name, identifier, text, channel, repeat
	FROM timers
	WHERE name = $1
	`

	var timer Timer

	err := t.DB.QueryRow(query, name).Scan(
		&timer.ID,
		&timer.Name,
		&timer.Identifier,
		&timer.Text,
		&timer.Channel,
		&timer.Repeat,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &timer, nil
}

// GetIdentifier returns the internal identifier for a supplied timer name.
func (t TimerModel) GetIdentifier(name string) (string, error) {
	query := `
	SELECT id, name, identifier, text, channel, repeat
	FROM timers
	WHERE name = $1
	`

	var timer Timer

	err := t.DB.QueryRow(query, name).Scan(
		&timer.ID,
		&timer.Name,
		&timer.Identifier,
		&timer.Text,
		&timer.Channel,
		&timer.Repeat,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return "", ErrRecordNotFound
		default:
			return "", err
		}
	}

	return timer.Identifier, nil
}

// Insert adds a new timer into the database.
func (t TimerModel) Insert(timer *Timer) error {
	query := `
	INSERT into timers(name, identifier, text, channel, repeat)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;
	`

	args := []interface{}{timer.Name, timer.Identifier, timer.Text, timer.Channel, timer.Repeat}

	result, err := t.DB.Exec(query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrCommandRecordAlreadyExists
	}

	return nil
}

// GetAll returns a pointer to a slice of all timers in the database.
func (t TimerModel) GetAll() ([]*Timer, error) {
	query := `
	SELECT id, name, identifier, text, channel, repeat
	FROM timers
	ORDER BY id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Use QueryContext() the context and query. This returns a
	// sql.Rows resultset containing our channels.
	rows, err := t.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	timers := []*Timer{}

	// Iterate over the resultset.
	for rows.Next() {
		// Initialize an empty Channel struct where we put on
		// a single channel value.
		var timer Timer

		// Scan the values onto the channel struct
		err := rows.Scan(
			&timer.ID,
			&timer.Name,
			&timer.Identifier,
			&timer.Text,
			&timer.Channel,
			&timer.Repeat,
		)
		if err != nil {
			return nil, err
		}
		// Add the single movie struct onto the slice.
		timers = append(timers, &timer)
	}

	// When rows.Next() finishes call rows.Err() to retrieve any errors.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return timers, nil
}

// GetChanneltimer returns a pointer to a slice of all timers of a given channel in the database.
func (t TimerModel) GetChannelTimer(channel string) ([]*Timer, error) {
	query := `
	SELECT id, name, identifier, text, channel, repeat
	FROM timers
	WHERE channel = $1
	ORDER BY name`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Use QueryContext() the context and query. This returns a
	// sql.Rows resultset containing our channels.
	rows, err := t.DB.QueryContext(ctx, query, channel)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	timers := []*Timer{}

	// Iterate over the resultset.
	for rows.Next() {
		// Initialize an empty Channel struct where we put on
		// a single channel value.
		var timer Timer

		err := rows.Scan(
			&timer.ID,
			&timer.Name,
			&timer.Identifier,
			&timer.Text,
			&timer.Channel,
			&timer.Repeat,
		)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, ErrRecordNotFound
			default:
				return nil, err
			}
		}
		// Add the single movie struct onto the slice.
		timers = append(timers, &timer)
	}

	// When rows.Next() finishes call rows.Err() to retrieve any errors.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return timers, nil
}

func (t TimerModel) Update(timer *Timer) error {
	query := `
	UPDATE timers
	SET text = $2, channel = $3, repeat = $4
	WHERE name = $1
	RETURNING id`

	args := []interface{}{
		timer.Name,
		timer.Text,
		timer.Channel,
		timer.Repeat,
	}

	err := t.DB.QueryRow(query, args...).Scan(&timer.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

// Delete takes in a command name and queries the database for an entry with
// the name and tries to delete that entry.
func (t TimerModel) Delete(identifier string) error {
	// Prepare the statement.
	query := `
	DELETE FROM timers
	WHERE identifier = $1`

	// Execute the query returning the number of affected rows.
	result, err := t.DB.Exec(query, identifier)
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
