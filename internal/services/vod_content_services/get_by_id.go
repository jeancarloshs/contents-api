package vod_content

import model "contents-api/internal/models"

func (fbyis *VodContentService) FindByID(vodID int) (model.VodContent, error) {
	return fbyis.repository.FindByID(vodID)
}
