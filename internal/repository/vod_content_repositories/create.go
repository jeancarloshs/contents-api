package vod_content

import (
	model "contents-api/internal/models"
	"fmt"
)

func (cvcr *VodContentRepository) CreateVodContentRepository(vodContent model.VodContent) (model.VodContent, error) {
	imgID := 1
	bannerID := 1
	posterID := 1
	// tipo := "banner"

	// insertImg, err := cvcr.connection.Prepare(`
	// 	INSERT INTO tb_imagens (descricao, url, tipo, status) VALUES (?, ?, ?, ?)
	// `)

	// if err != nil {
	// 	fmt.Println("Erro ao preparar query de imagem:", err)
	// 	return err
	// }
	// defer insertImg.Close()

	// err = insertImg.QueryRow(
	// 	vodContent.Title,
	// 	vodContent.Image,
	// 	tipo,
	// 	"ativo",
	// ).Scan(imgID)

	// if err != nil {
	// 	fmt.Println("Erro ao inserir imagem:", err)
	// 	return err
	// }

	// if tipo != "imagem" && tipo != "banner" && tipo != "poster" {
	// 	fmt.Println("Tipo inválido para o ENUM")
	// 	return fmt.Errorf("tipo inválido")
	// }

	insertVod, err := cvcr.connection.Prepare(`
		INSERT INTO 
			tb_conteudos_vod (titulo, subtitulo, descricao, imagem_id, banner_id, poster_id, trailer, url_video,
			url_hls, url_dash, legenda, categoria_id, tempo_video, id_tmdb, classificacao, distribuidora_id, status) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`)

	if err != nil {
		fmt.Println(err)
		return model.VodContent{}, err
	}
	defer insertVod.Close()

	_, err = insertVod.Exec(
		vodContent.Title,
		vodContent.Subtitle,
		vodContent.Description,
		imgID,
		bannerID,
		posterID,
		vodContent.Trailer,
		vodContent.UrlVideo,
		vodContent.UrlHls,
		vodContent.UrlDash,
		vodContent.Caption,
		vodContent.CategoryID,
		vodContent.VideoTime,
		vodContent.IdTmdb,
		vodContent.Classification,
		vodContent.DistributionID,
		vodContent.Status,
	)

	if err != nil {
		fmt.Println("Erro ao inserir conteúdo VOD:", err)
		return model.VodContent{}, err
	}

	insertVod.Close()
	return vodContent, nil
}
