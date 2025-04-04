package content_distributor_repository

import model "contents-api/internal/models"

func (gacdr *ContentDistributorRepository) GetAllContentDistributorRepository() ([]model.ContentDistributor, error) {
	query := `SELECT * FROM tb_conteudo_distribuidora`

	rows, err := gacdr.connection.Query(query)
	if err != nil {
		return []model.ContentDistributor{}, err
	}
	defer rows.Close()

	var distributorList []model.ContentDistributor

	for rows.Next() {
		var distributorObj model.ContentDistributor

		err = rows.Scan(
			&distributorObj.ID,
			&distributorObj.Name,
			&distributorObj.Description,
			&distributorObj.DistributorType,
			&distributorObj.Status,
		)
		if err != nil {
			return []model.ContentDistributor{}, err
		}

		distributorList = append(distributorList, distributorObj)
	}

	return distributorList, nil
}
