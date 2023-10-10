package data

import (
	"database/sql"
)

type CommandsLog struct {
	ID            int    `json:"id"`
	TwitchLogin   string `json:"twitch_login"`
	TwitchID      string `json:"twitch_id,omitempty"`
	TwitchChannel string `json:"twitch_channel,omitempty"`
	TwitchMessage string `json:"twitch_message,omitempty"`
	CommandName   string `json:"command_name,omitempty"`
	UserLevel     int    `json:"user_level,omitempty"`
	Identifier    string `json:"identifier,omitempty"`
	RawMessage    string `json:"raw_message,omitempty"`
}

type CommandsLogModel struct {
	DB *sql.DB
}

// Get tries to find a command in the database with the provided name.
func (c CommandsLogModel) Insert(twitchLogin, twitchId, twitchChannel, twitchMessage, commandName string, uLvl int, identifier, rawMsg string) {
	query := `
	INSERT into commands_logs(twitch_login, twitch_id, twitch_channel, twitch_message, command_name, user_level, identifier, raw_message)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id;
	`

	args := []interface{}{twitchLogin, twitchId, twitchChannel, twitchMessage, commandName, uLvl, identifier, rawMsg}

	result, err := c.DB.Exec(query, args...)
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
