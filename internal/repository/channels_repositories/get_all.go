package channels_repositories

import (
	"contents-api/internal/config"
	model "contents-api/internal/models"
	"fmt"
	"log"
)

func (ga *ChannelsRepository) FindAll() ([]model.Channels, error) {
	oldDB, err := config.ConnectDBNxtv()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco novo:", err)

	}
	defer oldDB.Close()
	query := "SELECT id, titulo, descricao, imagem, imagem_banner, imagem_bannersinopse, tempoVideo, `status`, url_video, url_hls, url_dash, epg_programacao_idxml FROM tbConteudoCanais"

	rows, err := oldDB.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Channels{}, nil
	}
	defer rows.Close()

	var channelsList []model.Channels

	for rows.Next() {
		var channelsOBJ model.Channels

		err = rows.Scan(
			&channelsOBJ.ID,
			&channelsOBJ.Title,
			&channelsOBJ.Description,
			&channelsOBJ.Image,
			&channelsOBJ.Banner,
			&channelsOBJ.Poster,
			&channelsOBJ.VideoTime,
			&channelsOBJ.Status,
			&channelsOBJ.UrlVideo,
			&channelsOBJ.UrlHls,
			&channelsOBJ.UrlDash,
			&channelsOBJ.Epg_Programacao_IDxml,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Channels{}, nil
		}
		defer rows.Close()

		channelsList = append(channelsList, channelsOBJ)
	}
	fmt.Println("total:", len(channelsList))

	return channelsList, nil
}
