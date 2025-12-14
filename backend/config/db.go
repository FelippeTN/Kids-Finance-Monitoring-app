package config

import (
	"log"

	"github.com/FelippeTN/Kids-Finance-Monitoring-app/backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	//TODO: Passar isso para variáveis de ambiente
	dsn := "host=localhost user=postgres password=1234 dbname=kids-finance-monitoring-app port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	database.AutoMigrate(&models.User{})

	DB = database
}