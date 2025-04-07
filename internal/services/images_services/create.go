package images_services

import (
	model "contents-api/internal/models"
	"fmt"
)

func (cis *ImageService) Create(imageContent model.Images) (model.Images, error) {
	imgInsert, err := cis.repository.Create(imageContent)
	if err != nil {
		fmt.Println(err)
		return model.Images{}, nil
	}

	return imgInsert, nil
}
