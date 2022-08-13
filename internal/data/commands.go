package data

import (
	"database/sql"
	"errors"
)

type Command struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Text  string `json:"text"`
	Level int    `json:"level"`
}

type CommandModel struct {
	DB *sql.DB
}

// Get tries to find a command in the database with the provided name.
func (c CommandModel) Get(name string) (*Command, error) {
	query := `
	SELECT id, name, text, level
	FROM commands
	WHERE name = $1`

	var command Command

	err := c.DB.QueryRow(query, name).Scan(
		&command.ID,
		&command.Name,
		&command.Text,
		&command.Level,
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

// SetLevel queries the database for an entry with the provided name,
// if there is one it updates the entrys level with the provided level.
func (c CommandModel) SetLevel(name string, level int) error {
	query := `
	UPDATE commands
	SET level = $2
	WHERE name = $1`

	result, err := c.DB.Exec(query, name, level)
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

// Insert adds a command into the database.
func (c CommandModel) Insert(name, text string) error {
	perms := 0
	query := `
	INSERT into commands(name, text, level)
	VALUES ($1, $2, $3)
	ON CONFLICT (name)
	DO NOTHING
	RETURNING id;
	`

	args := []interface{}{name, text, perms}

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

// Delete takes in a command name and queries the database for an entry with
// the same name and tries to delete that entry.
func (c CommandModel) Delete(name string) error {
	// Prepare the statement.
	query := `
	DELETE FROM commands
	WHERE name = $1`

	// Execute the query returning the number of affected rows.
	result, err := c.DB.Exec(query, name)
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
