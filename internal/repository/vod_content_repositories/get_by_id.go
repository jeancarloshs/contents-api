package vod_content

import (
	model "contents-api/internal/models"
	"fmt"
)

func (fbidr *VodContentRepository) FindByID(vodID int) (model.VodContent, error) {
	query := `
	SELECT
		vod.id, vod.titulo, vod.subtitulo, vod.descricao, img_imagem.url AS imagem, img_banner.url AS banner, img_poster.url AS poster, img_thumbnail.url AS thumbnail, vod.trailer, vod.url_video, vod.url_hls, vod.url_dash, vod.legenda, cat.nome AS categoria, vod.tempo_video, vod.id_tmdb, vod.classificacao, dist.nome AS distribuidora, vod.status
	FROM tb_conteudos_vod vod
	LEFT JOIN tb_imagens img_imagem ON vod.imagem_id = img_imagem.id AND img_imagem.tipo = 'imagem'
	LEFT JOIN tb_imagens img_banner ON vod.banner_id = img_banner.id AND img_banner.tipo = 'banner'
	LEFT JOIN tb_imagens img_poster ON vod.poster_id = img_poster.id AND img_poster.tipo = 'poster'
	LEFT JOIN tb_imagens img_thumbnail ON vod.thumbnail_id = img_thumbnail.id AND img_thumbnail.tipo = 'thumbnail'
	LEFT JOIN tb_categorias cat ON vod.categoria_id = cat.id
	LEFT JOIN tb_conteudo_distribuidora dist ON vod.distribuidora_id = dist.id
	WHERE vod.id = ?;
	`
	rows, err := fbidr.connection.Query(query, vodID)

	if err != nil {
		fmt.Println(err)
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
