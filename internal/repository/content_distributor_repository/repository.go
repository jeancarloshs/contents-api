package content_distributor_repository

import (
	model "contents-api/internal/models"
	"database/sql"
)

type ContentDistributorRepository struct {
	connection *sql.DB
}

func NewContentDistributorRepository(connection *sql.DB) ContentDistributorRepository {
	return ContentDistributorRepository{
		connection: connection,
	}
}

type Repository interface {
	GetAllContentDistributorRepository() ([]model.ContentDistributor, error)
	GetByIDContentDistributorRepository(cdID int) ([]model.ContentDistributor, error)
	CreateContentDistributorRepository(contentDistributor model.ContentDistributor) (model.ContentDistributor, error)
}
