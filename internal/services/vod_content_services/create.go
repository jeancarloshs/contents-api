package vod_content

import (
	model "contents-api/internal/models"
)

func (cvs *VodContentService) Create(vodContent model.VodContent) (model.VodContent, error) {
	vod, err := cvs.repository.Create(vodContent)
	if err != nil {
		return model.VodContent{}, err
	}
	return vod, nil
}
