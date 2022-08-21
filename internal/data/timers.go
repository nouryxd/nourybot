package data

import (
	"database/sql"
	"errors"
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
