package category_services

import model "contents-api/internal/models"

func (gacs *CategoryServices) GetAllCategoryService() ([]model.Category, error) {
	return gacs.repository.GetAllCategoryRepository()
}
