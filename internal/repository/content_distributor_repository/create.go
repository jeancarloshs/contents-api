package content_distributor_repository

import model "contents-api/internal/models"

func (ccdr *ContentDistributorRepository) CreateContentDistributorRepository(distributorContent model.ContentDistributor) (model.ContentDistributor, error) {
	insertDistributor, err := ccdr.connection.Prepare(`
		INSERT INTO tb_imagens (nome, descricao, url, tipo, status) VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return model.ContentDistributor{}, err
	}
	defer insertDistributor.Close()

	_, err = insertDistributor.Exec(
		distributorContent.Name,
		distributorContent.Description,
		distributorContent.DistributorType,
		distributorContent.Status,
	)
	if err != nil {
		return model.ContentDistributor{}, err
	}

	insertDistributor.Close()
	return distributorContent, nil
}
