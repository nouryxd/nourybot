package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Timer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Text    string `json:"text"`
	Channel string `json:"channel"`
	Repeat  string `json:"repeat"`
}

type TimerModel struct {
	DB *sql.DB
}

func (t TimerModel) Get(name string) (*Timer, error) {
	query := `
	SELECT id, name, text, channel, repeat
	FROM timers
	WHERE name = $1
	`

	var timer Timer

	err := t.DB.QueryRow(query, name).Scan(
		&timer.ID,
		&timer.Name,
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

// Insert adds a command into the database.
func (t TimerModel) Insert(timer *Timer) error {
	query := `
	INSERT into timers(name, text, channel, repeat)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
	`

	args := []interface{}{timer.Name, timer.Text, timer.Channel, timer.Repeat}

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

// GetAll() returns a pointer to a slice of all channels (`[]*Channel`) in the database.
func (t TimerModel) GetAll() ([]*Timer, error) {
	query := `
	SELECT id, name, text, channel, repeat
	FROM timers
	ORDER BY id`

	// Create a context with 3 seconds timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Use QueryContext() the context and query. This returns a
	// sql.Rows resultset containing our channels.
	rows, err := t.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	// Need to defer a call to rows.Close() to ensure the resultset
	// is closed before GetAll() returns.
	defer rows.Close()

	// Initialize an empty slice to hold the data.
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
