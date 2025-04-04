package content_distributor_repository

import model "contents-api/internal/models"

func (gbicdr *ContentDistributorRepository) GetByIDContentDistributorRepository(cdID int) ([]model.ContentDistributor, error) {
	query := `SELECT * FROM tb_conteudo_distribuidora WHERE ID = ?`

	row, err := gbicdr.connection.Query(query, cdID)
	if err != nil {
		return []model.ContentDistributor{}, err
	}
	defer row.Close()

	var disContentList []model.ContentDistributor
	var disContentObj model.ContentDistributor

	for row.Next() {
		err := row.Scan(
			&disContentObj.ID,
			&disContentObj.Name,
			&disContentObj.Description,
			&disContentObj.Status,
			&disContentObj.DistributorType,
		)
		if err != nil {
			return []model.ContentDistributor{}, err
		}

		disContentList = append(disContentList, disContentObj)
	}

	return disContentList, nil
}
