package images_repositories

import model "contents-api/internal/models"

func (uir *ImageRepository) Update(imgContent model.Images) (model.Images, error) {
	updateIMG, err := uir.connection.Prepare(`
		UPDATE tb_imagens SET  nome=?, descricao=?, url=?, tipo=?, status=? WHERE id=?
	`)
	if err != nil {
		return model.Images{}, err
	}
	defer updateIMG.Close()

	_, err = updateIMG.Exec(
		imgContent.Name,
		imgContent.Description,
		imgContent.Url,
		imgContent.Type,
		imgContent.Status,
		imgContent.ID,
	)
	if err != nil {
		return model.Images{}, err
	}

	updateIMG.Close()
	return imgContent, nil
}
