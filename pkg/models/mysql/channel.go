package mysql

import (
	"database/sql"

	"github.com/lyx0/nourybot/pkg/models"
)

type ChannelModel struct {
	DB *sql.DB
}

func (m *ChannelModel) Insert(username string, twitchid string) (int, error) {
	return 0, nil
}

func (m *ChannelModel) Get(id int) (*models.Channel, error) {
	return nil, nil
}

func (m *ChannelModel) Latest() ([]*models.Channel, error) {
	return nil, nil
}
