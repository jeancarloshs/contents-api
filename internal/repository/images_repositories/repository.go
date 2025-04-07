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
	GetAllImageRepository() ([]model.Images, error)
	GetImageByIDRepository(imgID int) (model.Images, error)
	CreateImageRepository(imgContent model.Images) (model.Images, error)
}
