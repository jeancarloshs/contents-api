package images_repositories

import (
	model "contents-api/internal/models"
	"database/sql"
	"fmt"
)

func (gibir *ImageRepository) GetImageByIDRepository(imgID int) (model.Images, error) {
	query := `SELECT * FROM tb_imagens WHERE id = ?`

	row := gibir.connection.QueryRow(query, imgID)

	var imageObj model.Images
	err := row.Scan(
		&imageObj.ID,
		&imageObj.Name,
		&imageObj.Description,
		&imageObj.Url,
		&imageObj.Type,
		&imageObj.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Images{}, fmt.Errorf("imagem com ID %d não encontrada", imgID)
		}
		return model.Images{}, err
	}

	return imageObj, nil
}
