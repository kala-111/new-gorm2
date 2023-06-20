package config

import (
	"github.com/kala-111/gin-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres.postgres@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Mobile{})
	DB = db

}
