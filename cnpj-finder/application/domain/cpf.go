package domain

import "time"

type CPF struct {
	CPF                string    `json:"cpf"`
	Nome               string    `json:"nome"`
	Situacao           string    `json:"situacao"`
	DataNascimento     string    `json:"data_nascimento"`
	Sexo               string    `json:"sexo"`
	TituloEleitor      string    `json:"titulo_eleitor"`
	UltimaAtualizacao  time.Time `json:"ultima_atualizacao"`
}
