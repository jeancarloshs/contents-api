package content_distributor_services

import model "contents-api/internal/models"

func (fbis *ContentDistributorService) FindByID(cdID int) ([]model.ContentDistributor, error) {
	return fbis.repository.FindByID(cdID)
}
