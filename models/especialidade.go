package models

import (
	"time"

	"gorm.io/gorm"
)

type Especialidade struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Especialidade string         `json:"especialidade"`
	CreatedAt     time.Time      `json:"created"`
	UpdatedAt     time.Time      `json:"updated"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted"`
	Medico        []Medico       `gorm:"ForeignKey:EspecialidadeID" json:"medico"`
}

/*



 */
