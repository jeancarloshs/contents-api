package images_services

import model "contents-api/internal/models"

func (gais *ImageService) GetAllImagesService() ([]model.Images, error) {
	return gais.repository.GetAllImageRepository()
}
