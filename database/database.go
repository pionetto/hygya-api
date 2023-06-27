package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {

	str := "host=localhost port=5432 user=pacs dbname=pacsdb sslmode=disable password=pacs"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("error: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetConnMaxIdleTime(10)
	config.SetMaxIdleConns(10)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDataBase() *gorm.DB {
	return db
}
