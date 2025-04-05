package category_repository

import (
	model "contents-api/internal/models"
	"database/sql"
)

type CategoryRepository struct {
	connection *sql.DB
}

func NewCategoryRepository(connection *sql.DB) CategoryRepository {
	return CategoryRepository{
		connection: connection,
	}
}

type Repository interface {
	GetAllCategoryRepository() ([]model.Category, error)
}
