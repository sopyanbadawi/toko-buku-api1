package config

import (
	"log"
	"toko-buku-api1/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	dsn := "host=localhost 	user=postgres password=123 dbname=toko-buku port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database Postgre man!", err)
	}

	database.AutoMigrate(&models.Books{})

	DB = database
}
