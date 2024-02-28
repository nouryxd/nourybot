package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Command struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Channel     string `json:"channel"`
	Text        string `json:"text,omitempty"`
	Level       int    `json:"level,omitempty"`
	Description string `json:"description,omitempty"`
}

type CommandModel struct {
	DB *sql.DB
}

// Get tries to find a command in the database with the provided name and channel.
func (c CommandModel) Get(name, channel string) (*Command, error) {
	query := `
	SELECT *
	FROM commands
	WHERE name = $1 AND channel = $2`

	var command Command

	err := c.DB.QueryRow(query, name, channel).Scan(
		&command.ID,
		&command.Name,
		&command.Channel,
		&command.Text,
		&command.Level,
		&command.Description,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &command, nil
}

// GetAll returns  a slice of all channels in the database.
func (c CommandModel) GetAll() ([]*Command, error) {
	query := `
	SELECT *
	FROM commands
	ORDER BY id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Use QueryContext() the context and query. This returns a
	// sql.Rows resultset containing our channels.
	rows, err := c.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	commands := []*Command{}

	// Iterate over the resultset.
	for rows.Next() {
		// Initialize an empty Channel struct where we put on
		// a single channel value.
		var command Command

		// Scan the values onto the channel struct
		err := rows.Scan(
			&command.ID,
			&command.Name,
			&command.Channel,
			&command.Text,
			&command.Level,
			&command.Description,
		)
		if err != nil {
			return nil, err
		}
		// Add the single movie struct onto the slice.
		commands = append(commands, &command)
	}

	// When rows.Next() finishes call rows.Err() to retrieve any errors.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return commands, nil
}

// GetAll returns  a slice of all channels in the database.
func (c CommandModel) GetAllChannel(channel string) ([]*Command, error) {
	query := `
	SELECT id, name, channel, text, level, description
	FROM commands
	WHERE channel = $1
	ORDER BY name`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Use QueryContext() the context and query. This returns a
	// sql.Rows resultset containing our channels.
	rows, err := c.DB.QueryContext(ctx, query, channel)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Initialize an empty slice to hold the data.
	commands := []*Command{}

	// Iterate over the resultset.
	for rows.Next() {
		// Initialize an empty Channel struct where we put on
		// a single channel value.
		var command Command

		// Scan the values onto the channel struct
		err := rows.Scan(
			&command.ID,
			&command.Name,
			&command.Channel,
			&command.Text,
			&command.Level,
			&command.Description,
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
		commands = append(commands, &command)
	}

	// When rows.Next() finishes call rows.Err() to retrieve any errors.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return commands, nil
}

// Insert adds a command into the database.
func (c CommandModel) Insert(command *Command) error {
	query := `
	INSERT into commands(name, channel, text, level, description)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;
	`

	args := []interface{}{command.Name, command.Channel, command.Text, command.Level, command.Description}

	result, err := c.DB.Exec(query, args...)
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

// Update updates a command with the new supplied values.
func (c CommandModel) Update(command *Command) error {
	query := `
	UPDATE commands
	SET text = $3
	WHERE name = $1 AND channel = $2
	RETURNING id`

	args := []interface{}{
		command.Name,
		command.Channel,
		command.Text,
	}

	err := c.DB.QueryRow(query, args...).Scan(&command.ID)
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

// SetLevel queries the database for an entry with the provided name,
// if there is one it updates the entries level with the newly provided level.
func (c CommandModel) SetLevel(name, channel string, level int) error {
	query := `
	UPDATE commands
	SET level = $3
	WHERE name = $1 AND channel = $2`

	result, err := c.DB.Exec(query, name, channel, level)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

// SetHelp sets the help text for a given name of a command in the database.
func (c CommandModel) SetDescription(name, channel string, description string) error {
	query := `
	UPDATE commands
	SET description = $3
	WHERE name = $1 AND channel = $2`

	result, err := c.DB.Exec(query, name, channel, description)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

// Delete takes in a command name and queries the database for an entry with
// the same name and tries to delete that entry.
func (c CommandModel) Delete(name, channel string) error {
	// Prepare the statement.
	query := `
	DELETE FROM commands
	WHERE name = $1 AND channel = $2`

	// Execute the query returning the number of affected rows.
	result, err := c.DB.Exec(query, name, channel)
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
