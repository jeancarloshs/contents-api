package channels_service

import (
	model "contents-api/internal/models"
	channels_repositories "contents-api/internal/repository/channels_repositories"
)

type Service interface {
	FindAll() ([]model.Channels, error)
}

type ChannelsService struct {
	repository channels_repositories.Repository
}

func NewChannelsService(repo channels_repositories.Repository) Service {
	return &ChannelsService{
		repository: repo,
	}
}
