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
	FindAll() ([]model.ContentDistributor, error)
	FindByID(cdID int) ([]model.ContentDistributor, error)
	Create(contentDistributor model.ContentDistributor) (model.ContentDistributor, error)
}
