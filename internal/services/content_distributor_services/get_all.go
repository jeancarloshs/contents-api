package content_distributor_services

import model "contents-api/internal/models"

func (facds *ContentDistributorService) FindAll() ([]model.ContentDistributor, error) {
	return facds.repository.FindAll()
}
