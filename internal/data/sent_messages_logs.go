package data

import (
	"database/sql"
)

type SentMessagesLog struct {
	ID            int    `json:"id"`
	TwitchChannel string `json:"twitch_channel,omitempty"`
	TwitchMessage string `json:"twitch_message,omitempty"`
	Identifier    string `json:"identifier,omitempty"`
}

type SentMessagesLogModel struct {
	DB *sql.DB
}

// Get tries to find a command in the database with the provided name.
func (s SentMessagesLogModel) Insert(twitchChannel, twitchMessage, identifier string) {
	query := `
	INSERT into sent_messages_logs(twitch_channel, twitch_message, identifier)
	VALUES ($1, $2, $3)
	RETURNING id;
	`

	args := []interface{}{twitchChannel, twitchMessage, identifier}

	result, err := s.DB.Exec(query, args...)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected == 0 {
		return
	}
}
