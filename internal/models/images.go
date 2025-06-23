package model

type Images struct {
	ID          *int    `json:"id"`
	Name        *string `json:"nome"`
	Description *string `json:"descricao"`
	Url         *string `json:"url"`
	Type        *string `json:"tipo"`
	Status      *Status `json:"status"`
}
