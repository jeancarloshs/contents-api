package category_repository

import model "contents-api/internal/models"

func (gacr *CategoryRepository) GetAllCategoryRepository() ([]model.Category, error) {
	query := `SELECT * FROM tb_categorias`

	rows, err := gacr.connection.Query(query)
	if err != nil {
		return []model.Category{}, err
	}

	var categoryList []model.Category
	var categoryObj model.Category

	for rows.Next() {
		err = rows.Scan(
			&categoryObj.ID,
			&categoryObj.CategoryName,
			&categoryObj.Description,
			&categoryObj.Status,
			&categoryObj.ShowMenu,
			&categoryObj.ContentType,
			&categoryObj.ParentalControl,
		)
		if err != nil {
			return []model.Category{}, err
		}

		categoryList = append(categoryList, categoryObj)
	}

	return categoryList, nil
}
