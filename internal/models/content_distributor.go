package model

type ContentDistributor struct {
	ID              int    `json:"id"`
	Name            string `json:"nome"`
	Description     string `json:"descricao"`
	Status          Status `json:"status"`
	DistributorType string `json:"tipo_distribuidora"`
}
