package images_repositories

import model "contents-api/internal/models"

func (gibir *ImageRepository) GetImageByIDRepository(imgID int) ([]model.Images, error) {
	query := `SELECT * FROM tb_imagens WHERE id = ?`

	row, err := gibir.connection.Query(query, imgID)
	if err != nil {
		return []model.Images{}, err
	}
	defer row.Close()

	var imageList []model.Images
	var imageObj model.Images

	for row.Next() {
		err := row.Scan(
			&imageObj.ID,
			&imageObj.Name,
			&imageObj.Description,
			&imageObj.Url,
			&imageObj.Type,
			&imageObj.Status,
		)
		if err != nil {
			return []model.Images{}, err
		}

		imageList = append(imageList, imageObj)
	}

	return imageList, nil
}
