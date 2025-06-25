package model

type Channels struct {
	ID                    *int    `json:"id"`
	Title                 *string `json:"titulo"`
	Description           *string `json:"descricao"`
	Image                 *string `json:"imagem"`
	Banner                *string `json:"imagem_banner"`
	Poster                *string `json:"imagem_bannersinopse"`
	UrlVideo              *string `json:"url_video"`
	UrlHls                *string `json:"url_hls"`
	UrlDash               *string `json:"url_dash"`
	VideoTime             *string `json:"tempoVideo"`
	Epg_Programacao_IDxml *string `json:"epg_programacao_idxml"`
	Status                *int    `json:"status"`
}
