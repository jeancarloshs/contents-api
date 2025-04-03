package vod_content

import (
	model "contents-api/internal/models"
)

func (cvs *VodContentService) CreateVodContentService(vodContent model.VodContent) (model.VodContent, error) {
	vod, err := cvs.repository.CreateVodContentRepository(vodContent)
	if err != nil {
		return model.VodContent{}, err
	}
	return vod, nil
}
