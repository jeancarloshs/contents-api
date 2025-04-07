package category_services

import model "contents-api/internal/models"

func (fios *CategoryServices) FindOne(catID int) (model.Category, error) {
	return fios.repository.FindOne(catID)
}
