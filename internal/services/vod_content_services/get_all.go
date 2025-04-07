package vod_content

import (
	model "contents-api/internal/models"
)

func (fas *VodContentService) FindAll() ([]model.VodContent, error) {
	return fas.repository.FindAll()
}
