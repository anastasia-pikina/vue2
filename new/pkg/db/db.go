package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/new/pkg/models"
)

func Init() *gorm.DB {
	dbURL := "postgres://user:password@localhost:5432/dbname"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.Phone{})
	db.AutoMigrate(&models.New{})

	return db
}
