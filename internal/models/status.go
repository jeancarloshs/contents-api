package model

type Status struct {
	Ativo                 *string `json:"ativo"`
	Inativo               *string `json:"inativo"`
	Sem_Conteudo          *string `json:"sem_conteudo"`
	Conteudo_Com_Problema *string `json:"conteudo_com_problema"`
}
