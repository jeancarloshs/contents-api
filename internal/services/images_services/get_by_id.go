package images_services

import model "contents-api/internal/models"

func (gibis *ImageService) GetImageByIDService(imgID int) ([]model.Images, error) {
	return gibis.repository.GetImageByIDRepository(imgID)
}
