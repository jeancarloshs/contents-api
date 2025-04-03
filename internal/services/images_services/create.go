package images_services

import (
	model "contents-api/internal/models"
	"fmt"
)

func (uis *ImageService) CreateImageService(imageContent model.Images) (model.Images, error) {
	imgInsert, err := uis.repository.CreateImageRepository(imageContent)
	if err != nil {
		fmt.Println(err)
		return model.Images{}, nil
	}

	return imgInsert, nil
}
