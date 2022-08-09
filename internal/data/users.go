package data

import (
	"database/sql"
	"errors"
	"time"
)

type User struct {
	ID       int       `json:"id"`
	AddedAt  time.Time `json:"-"`
	Login    string    `json:"login"`
	TwitchID string    `json:"twitchid"`
	Level    int       `json:"level"`
}

type UserModel struct {
	DB *sql.DB
}

func (u UserModel) Insert(user *User) error {
	query := `
	INSERT INTO users(login, twitchid, level)
	VALUES ($1, $2, $3)
	ON CONFLICT (login)
	DO NOTHING
	RETURNING id, added_at;
	`

	args := []interface{}{user.Login, user.TwitchID, user.Level}

	// Execute the query returning the number of affected rows.
	result, err := u.DB.Exec(query, args...)
	if err != nil {
		return err
	}

	// Check how many rows were affected.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrUserAlreadyExists
	}

	return nil
}

func (u UserModel) SetLevel(login string, level int) error {
	query := `
	UPDATE users
	SET level = $2
	WHERE login = $1`

	// err := u.DB.QueryRow(query, args...).Scan(&user)
	result, err := u.DB.Exec(query, login, level)
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

func (u UserModel) Get(login string) (*User, error) {
	query := `
	SELECT id, added_at, login, twitchid, level
	FROM users
	WHERE login = $1`

	var user User

	err := u.DB.QueryRow(query, login).Scan(
		&user.ID,
		&user.AddedAt,
		&user.Login,
		&user.TwitchID,
		&user.Level,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (u UserModel) Delete(login string) error {
	// Prepare the statement.
	query := `
	DELETE FROM users
	WHERE login = $1`

	// Execute the query returning the number of affected rows.
	result, err := u.DB.Exec(query, login)
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
