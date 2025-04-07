package model

type VodContent struct {
	ID             *int    `json:"id"`
	Title          *string `json:"titulo"`
	Subtitle       *string `json:"subtitulo"`
	Description    *string `json:"descricao"`
	Image          *string `json:"imagem"`
	Banner         *string `json:"banner"`
	Poster         *string `json:"poster"`
	Thumbnail      *string `json:"thumbnail"`
	Trailer        *string `json:"trailer"`
	UrlVideo       *string `json:"url_video"`
	UrlHls         *string `json:"url_hls"`
	UrlDash        *string `json:"url_dash"`
	Caption        *string `json:"legenda"`
	Category       *string `json:"categoria"`
	CategoryID     *int    `json:"categoria_id"`
	VideoTime      *string `json:"tempo_video"`
	IdTmdb         *string `json:"id_tmdb"`
	Classification *string `json:"classificacao"`
	Distribution   *string `json:"distribuidora"`
	DistributionID *int    `json:"distribuidora_id"`
	Status         *string `json:"status"`
}
