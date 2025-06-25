package channels_repositories

import (
	model "contents-api/internal/models"
	"database/sql"
)

type ChannelsRepository struct {
	connection *sql.DB
}

func NewChannels(connection *sql.DB) ChannelsRepository {
	return ChannelsRepository{
		connection: connection,
	}
}

type Repository interface {
	FindAll() ([]model.Channels, error)
}
