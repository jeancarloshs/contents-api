package images_services

import model "contents-api/internal/models"

func (gais *ImageService) FindAll() ([]model.Images, error) {
	return gais.repository.FindAll()
}
