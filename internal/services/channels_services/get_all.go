package channels_service

import (
	model "contents-api/internal/models"
)

func (crr *ChannelsService) FindAll() ([]model.Channels, error) {
	return crr.repository.FindAll()
}
