package models

import (
	"time"

	"gorm.io/gorm"
)

type Medico struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	Nome              string         `json:"nome"`
	Sexo              string         `json:"sexo"` //relacionamento
	Data_nasc         time.Time      `json:"data_nasc"`
	Cpf               string         `json:"cpf"`
	EspecialidadeID   uint           `json:"especialidadeID"` //relacionamento
	Especialidade     Especialidade  `json:"especialidade"`
	Rqe               string         `json:"rqe"`
	Crm_estado        string         `json:"crm_estado"` //relacionamento
	Crm_numero        string         `json:"crm_numero"`
	End_logradouro    string         `json:"end_logradouro"`
	End_numero        string         `json:"end_numero"`
	End_complemento   string         `json:"end_complemento"`
	End_bairro        string         `json:"end_bairro"`
	End_cidade        string         `json:"end_cidade"`
	End_uf            string         `json:"end_uf"` //relacionamento
	End_cep           string         `json:"end_cep"`
	Contato_email     string         `json:"contato_email"`
	Contato_telefone1 string         `json:"contato_telefone1"`
	Contato_telefone2 string         `json:"contato_telefone2"`
	Contato_telefone3 string         `json:"contato_telefone3"`
	Contato_obs       string         `json:"contato_obs"`
	CreatedAt         time.Time      `json:"created"`
	UpdatedAt         time.Time      `json:"updated"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted"`
}
