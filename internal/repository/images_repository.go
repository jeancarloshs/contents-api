package repository

import (
	model "contents-api/internal/models"
	"database/sql"
	"fmt"
)

type ImageRepository struct {
	connection *sql.DB
}

func GetImageRepository(connection *sql.DB) ImageRepository {
	return ImageRepository{
		connection: connection,
	}
}

func (gair *ImageRepository) GetAllImageRepository() ([]model.Images, error) {
	query := `SELECT * FROM tb_imagens`
	rows, err := gair.connection.Query(query)
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
		return model.Images{}, err
	}

	return imageObj, nil
}

func (uir *ImageRepository) UploadImageRepository(imageContent model.Images) (model.Images, error) {

	insertIMG, err := uir.connection.Prepare(`
		INSERT INTO tb_imagens (nome, descricao, url, tipo, status) VALUES (?, ?, ?, ?, ?)
	`)

	if err != nil {
		fmt.Println(err)
		return model.Images{}, err
	}
	defer insertIMG.Close()

	_, err = insertIMG.Exec(
		imageContent.Name,
		imageContent.Description,
		imageContent.Url,
		imageContent.Type,
		imageContent.Status,
	)

	if err != nil {
		fmt.Println("Erro ao inserir Imagem:", err)
		return model.Images{}, err
	}

	insertIMG.Close()
	return imageContent, nil
}

func (uir *ImageRepository) UpdateImageRepository(imgContent model.Images) (model.Images, error) {
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
