package content_distributor_repository

import (
	model "contents-api/internal/models"
)

func (ccdr *ContentDistributorRepository) Create(distributorContent model.ContentDistributor) (model.ContentDistributor, error) {
	insertDistributor, err := ccdr.connection.Prepare(`
		INSERT INTO tb_conteudo_distribuidora (nome, descricao, status, tipo_distribuidora) VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return model.ContentDistributor{}, err
	}
	defer insertDistributor.Close()

	result, err := insertDistributor.Exec(
		distributorContent.Name,
		distributorContent.Description,
		distributorContent.Status,
		distributorContent.DistributorType,
	)
	if err != nil {
		return model.ContentDistributor{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return model.ContentDistributor{}, err
	}

	distributorContent.ID = int(lastID)
	return distributorContent, nil
}
