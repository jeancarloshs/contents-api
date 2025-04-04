package content_distributor_services

import (
	model "contents-api/internal/models"
	"contents-api/internal/repository/content_distributor_repository"
)

type Service interface {
	GetAllContentDistributorService() ([]model.ContentDistributor, error)
	CreateContentDistributorService(distributorContent model.ContentDistributor) (model.ContentDistributor, error)
}

type ContentDistributorService struct {
	repository content_distributor_repository.Repository
}

func NewContentVodService(repo content_distributor_repository.Repository) Service {
	return &ContentDistributorService{
		repository: repo,
	}
}
