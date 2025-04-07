package images_services

import model "contents-api/internal/models"

func (fibis *ImageService) FindByID(imgID int) (model.Images, error) {
	return fibis.repository.FindByID(imgID)
}
