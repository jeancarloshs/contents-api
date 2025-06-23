package model

type Category struct {
	ID              int    `json:"id"`
	CategoryName    string `json:"nome"`
	Description     string `json:"descricao"`
	Status          string `json:"status"`
	ShowMenu        int    `json:"mostrar_menu"`
	ContentType     string `json:"tipo_conteudo"`
	ParentalControl bool   `json:"controle_parental"`
}
