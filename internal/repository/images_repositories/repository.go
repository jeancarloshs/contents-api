package images_repositories

import (
	model "contents-api/internal/models"
	"database/sql"
)

type ImageRepository struct {
	connection *sql.DB
}

func NewImageRepository(connection *sql.DB) ImageRepository {
	return ImageRepository{
		connection: connection,
	}
}

type Repository interface {
	FindAll() ([]model.Images, error)
	FindByID(imgID int) (model.Images, error)
	Create(imgContent model.Images) (model.Images, error)
}
