package content_distributor_services

import model "contents-api/internal/models"

func (ccds *ContentDistributorService) Create(distributorContent model.ContentDistributor) (model.ContentDistributor, error) {
	distributor, err := ccds.repository.Create(distributorContent)
	if err != nil {
		return model.ContentDistributor{}, err
	}
	return distributor, nil
}
