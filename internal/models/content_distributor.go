package model

type ContentDistributor struct {
	ID              int    `json:"id"`
	Name            string `json:"nome"`
	Description     string `json:"descricao"`
	Status          string `json:"status"`
	DistributorType string `json:"tipo_distribuidora"`
}
