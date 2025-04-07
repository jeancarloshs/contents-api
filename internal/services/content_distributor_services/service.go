package content_distributor_services

import (
	model "contents-api/internal/models"
	"contents-api/internal/repository/content_distributor_repository"
)

type Service interface {
	FindAll() ([]model.ContentDistributor, error)
	FindByID(cdID int) ([]model.ContentDistributor, error)
	Create(distributorContent model.ContentDistributor) (model.ContentDistributor, error)
}

type ContentDistributorService struct {
	repository content_distributor_repository.Repository
}

func NewContentVodService(repo content_distributor_repository.Repository) Service {
	return &ContentDistributorService{
		repository: repo,
	}
}
