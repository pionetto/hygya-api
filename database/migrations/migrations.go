package migrations

import (
	"hygya-api/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.Debug().AutoMigrate(models.Paciente{})
	db.Debug().AutoMigrate(models.User{})
	db.Debug().AutoMigrate(models.Especialidade{})
	db.Debug().AutoMigrate(models.Medico{})
	// db.AutoMigrate(models.Login{})
}
