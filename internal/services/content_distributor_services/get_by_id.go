package content_distributor_services

import model "contents-api/internal/models"

func (gbis *ContentDistributorService) GetByIDContentDistributorService(cdID int) ([]model.ContentDistributor, error) {
	return gbis.repository.GetByIDContentDistributorRepository(cdID)
}
