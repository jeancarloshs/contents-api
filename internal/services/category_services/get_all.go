package category_services

import model "contents-api/internal/models"

func (fas *CategoryServices) FindAll() ([]model.Category, error) {
	return fas.repository.FindAll()
}
