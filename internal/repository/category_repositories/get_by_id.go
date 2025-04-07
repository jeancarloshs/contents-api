package category_repository

import (
	model "contents-api/internal/models"
	"fmt"
)

func (fior *CategoryRepository) FindByID(catID int) (model.Category, error) {
	query := `SELECT * FROM tb_categorias WHERE ID = ?`

	row, err := fior.connection.Query(query, catID)
	if err != nil {
		fmt.Println(err)
		return model.Category{}, err
	}

	var catObj model.Category

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
		fmt.Println(err)
		return model.Category{}, err
	}

	row.Close()

	return catObj, nil
}
