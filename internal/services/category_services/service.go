package category_services

import (
	model "contents-api/internal/models"
	category_repository "contents-api/internal/repository/category_repositories"
)

type Service interface {
	GetAllCategoryService() ([]model.Category, error)
	GetCategoryByIDService(catID int) ([]model.Category, error)
}

type CategoryServices struct {
	repository category_repository.Repository
}

func NewCategoryService(repo category_repository.Repository) Service {
	return &CategoryServices{
		repository: repo,
	}
}
