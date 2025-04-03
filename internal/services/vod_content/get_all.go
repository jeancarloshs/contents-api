package vod_content

import (
	model "contents-api/internal/models"
)

func (cvs *VodContentService) GetAllVodContentService() ([]model.VodContent, error) {
	return cvs.repository.GetAllVodContentRepository()
}
