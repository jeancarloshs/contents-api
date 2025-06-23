package vod_content

import (
	model "contents-api/internal/models"
	"fmt"
)

func (favr *VodContentRepository) FindAll() ([]model.VodContent, error) {
	query := "CALL SP_get_all_contents();"

	rows, err := favr.connection.Query(query)
	if err != nil {
		fmt.Println("Erro ao executar a procedure:", err)
		return []model.VodContent{}, err
	}
	defer rows.Close()

	var contentList []model.VodContent

	for rows.Next() {
		var contentObj model.VodContent

		err = rows.Scan(
			&contentObj.ID,
			&contentObj.Title,
			&contentObj.Subtitle,
			&contentObj.Description,
			&contentObj.Image,
			&contentObj.Banner,
			&contentObj.Poster,
			&contentObj.Thumbnail,
			&contentObj.Trailer,
			&contentObj.UrlVideo,
			&contentObj.UrlHls,
			&contentObj.UrlDash,
			&contentObj.Caption,
			&contentObj.Category,
			&contentObj.VideoTime,
			&contentObj.IdTmdb,
			&contentObj.Classification,
			&contentObj.Distribution,
			&contentObj.Status,
		)

		if err != nil {
			fmt.Println(err)
			return []model.VodContent{}, err
		}

		contentList = append(contentList, contentObj)
	}

	// rows.Close()

	return contentList, nil
}
