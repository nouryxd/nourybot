package data

import (
	"database/sql"
)

type SentMessagesLog struct {
	ID                 int    `json:"id"`
	TwitchChannel      string `json:"twitch_channel,omitempty"`
	TwitchMessage      string `json:"twitch_message,omitempty"`
	ContextCommandName string `json:"context_command_name"`
	ContextUsername    string `json:"context_user"`
	ContextMessage     string `json:"context_message"`
	ContextUserID      string `json:"context_user_id"`
	Identifier         string `json:"identifier,omitempty"`
	ContextRawMsg      string `json:"context_raw"`
}

type SentMessagesLogModel struct {
	DB *sql.DB
}

// Insert adds the supplied values to the database sent_messages_logs.
func (s SentMessagesLogModel) Insert(twitchChannel, twitchMessage, ctxCommandName, ctxUser, ctxUserID, ctxMsg, identifier, ctxRaw string) {
	query := `
	INSERT into sent_messages_logs(twitch_channel, twitch_message, context_command_name, context_username, context_user_id, context_message, identifier, context_raw)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id;
	`

	args := []interface{}{twitchChannel, twitchMessage, ctxCommandName, ctxUser, ctxUserID, ctxMsg, identifier, ctxRaw}

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
