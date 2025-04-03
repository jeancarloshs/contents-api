package images_repositories

import (
	model "contents-api/internal/models"
	"fmt"
)

func (uir *ImageRepository) CreateImageRepository(imageContent model.Images) (model.Images, error) {

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
