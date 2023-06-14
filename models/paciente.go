package models

import (
	"time"
)

type Paciente struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Prioridade   string    `json:"prioridade"`
	Paciente     string    `json:"paciente"`
	Sexo         string    `json:"sexo"`
	Id_Paciente  string    `json:"id_paciente"`
	Procedimento string    `json:"procedimento"`
	Modalidade   string    `json:"modalidade"`
	Data_Exame   time.Time `json:"data_exame"`
	Data_Laudo   time.Time `json:"data_laudo"`
}

/*



 */
