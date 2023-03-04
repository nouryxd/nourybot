package data

import (
	"database/sql"
	"errors"
)

type LastFMUser struct {
	ID          int    `json:"id"`
	TwitchLogin string `json:"twitch_login"`
	TwitchID    string `json:"twitch_id"`
	LastFMUser  string `json:"lastfm_username"`
}

type LastFMUserModel struct {
	DB *sql.DB
}

func (l LastFMUserModel) Get(login string) (*LastFMUser, error) {
	query := `
  SELECT id, twitch_login, twitch_id, lastfm_username
  FROM lastfm_users
  WHERE twitch_login = $1`

	var lastfm LastFMUser

	err := l.DB.QueryRow(query, login).Scan(
		&lastfm.ID,
		&lastfm.TwitchLogin,
		&lastfm.TwitchID,
		&lastfm.LastFMUser,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &lastfm, nil
}

func (l LastFMUserModel) Insert(lastfm *LastFMUser) error {
	query := `
	INSERT into lastfm_users(twitch_login, twitch_id, lastfm_username)
	VALUES ($1, $2, $3)
	ON CONFLICT (twitch_id)
	DO NOTHING
	RETURNING id;
	`

	args := []interface{}{lastfm.TwitchLogin, lastfm.TwitchID, lastfm.LastFMUser}

	result, err := l.DB.Exec(query, args...)
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
