package vod_content

import model "contents-api/internal/models"

func (cvs *VodContentService) GetVodContentByIDService(vodID int) (model.VodContent, error) {
	return cvs.repository.GetVodContentByIDRepository(vodID)
}
