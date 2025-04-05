package category_repository

import model "contents-api/internal/models"

func (gcbir *CategoryRepository) GetCategoryByIDRepository(catID int) ([]model.Category, error) {
	query := `SELECT * FROM tb_categorias WHERE ID = ?`

	row, err := gcbir.connection.Query(query, catID)
	if err != nil {
		return []model.Category{}, err
	}

	var catList []model.Category
	var catObj model.Category

	for row.Next() {
		err = row.Scan(
			&catObj.ID,
			&catObj.CategoryName,
			&catObj.Description,
			&catObj.Status,
			&catObj.ShowMenu,
			&catObj.ContentType,
			&catObj.ParentalControl,
		)
		if err != nil {
			return []model.Category{}, err
		}

		catList = append(catList, catObj)
	}

	return catList, nil
}
