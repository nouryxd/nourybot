package data

import (
	"database/sql"
	"errors"
	"time"
)

type Upload struct {
	ID            int       `json:"id"`
	AddedAt       time.Time `json:"-"`
	TwitchLogin   string    `json:"twitchlogin"`
	TwitchID      string    `json:"twitchid"`
	TwitchChannel string    `json:"twitchchannel"`
	TwitchMessage string    `json:"twitchmessage"`
	Filehoster    string    `json:"filehoster"`
	DownloadURL   string    `json:"downloadurl"`
	UploadURL     string    `json:"uploadurl"`
	Identifier    string    `json:"identifier"`
}

type UploadModel struct {
	DB *sql.DB
}

// Insert takes in a channel struct and inserts it into the database.
func (u UploadModel) Insert(twitchLogin, twitchID, twitchChannel, twitchMessage, filehoster, downloadURL, identifier string) {
	query := `
	INSERT INTO uploads(twitchlogin, twitchid, twitchchannel, twitchmessage,  filehoster, downloadurl, uploadurl, identifier)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, added_at, identifier;
	`

	args := []interface{}{
		twitchLogin,
		twitchID,
		twitchChannel,
		twitchMessage,
		filehoster,
		downloadURL,
		"undefined",
		identifier,
	}

	// Execute the query returning the number of affected rows.
	result, err := u.DB.Exec(query, args...)
	if err != nil {
		return
	}

	// Check how many rows were affected.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected == 0 {
		return
	}
}

func (u UploadModel) UpdateUploadURL(identifier, uploadURL string) {
	var id string
	query := `
	UPDATE uploads
	SET uploadurl = $2
	WHERE identifier = $1
	RETURNING id`

	args := []interface{}{
		identifier,
		uploadURL,
	}

	err := u.DB.QueryRow(query, args...).Scan(id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return
		default:
			return
		}
	}
}
