package migrations

import (
	"hygya-api/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Paciente{})
}
