package vod_content

import (
	model "contents-api/internal/models"
	"fmt"
)

func (fbidr *VodContentRepository) FindByID(vodID int) (model.VodContent, error) {
	query := fmt.Sprintf("CALL SP_get_content_by_id(%d);", vodID)

	rows, err := fbidr.connection.Query(query)
	if err != nil {
		fmt.Println("Erro ao executar a procedure:", err)
		return model.VodContent{}, err
	}

	// var contentList []model.VodContent
	var contentObj model.VodContent

	for rows.Next() {
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
			return model.VodContent{}, err
		}

		rows.Close()
	}

	return contentObj, nil
}
