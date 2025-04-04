package content_distributor_services

import model "contents-api/internal/models"

func (gacds *ContentDistributorService) GetAllContentDistributorService() ([]model.ContentDistributor, error) {
	return gacds.repository.GetAllContentDistributorRepository()
}
