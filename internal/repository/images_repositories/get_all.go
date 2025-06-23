package images_repositories

import (
	model "contents-api/internal/models"
)

func (fair *ImageRepository) FindAll() ([]model.Images, error) {
	query := `SELECT * FROM tb_imagens`
	rows, err := fair.connection.Query(query)
	if err != nil {
		return []model.Images{}, err
	}
	defer rows.Close()

	var imagesList []model.Images

	for rows.Next() {
		var imagesObj model.Images

		err = rows.Scan(
			&imagesObj.ID,
			&imagesObj.Name,
			&imagesObj.Description,
			&imagesObj.Url,
			&imagesObj.Type,
			&imagesObj.Status,
		)

		if err != nil {
			return []model.Images{}, err
		}

		imagesList = append(imagesList, imagesObj)
	}

	return imagesList, nil
}
