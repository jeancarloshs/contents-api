package category_services

import model "contents-api/internal/models"

func (gcbis *CategoryServices) GetCategoryByIDService(catID int) ([]model.Category, error) {
	return gcbis.repository.GetCategoryByIDRepository(catID)
}
