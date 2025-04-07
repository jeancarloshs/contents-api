package category_services

import model "contents-api/internal/models"

func (fios *CategoryServices) FindByID(catID int) (model.Category, error) {
	return fios.repository.FindByID(catID)
}
