package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"CarSharing/models"
)

var DB *gorm.DB

func Connect() {
	var err error

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	//Logger
	//fmt.Println("DB_HOST:", dbHost)
	//fmt.Println("DB_PORT:", dbPort)
	//fmt.Println("DB_USER:", dbUser)
	//fmt.Println("DB_PASSWORD:", dbPassword)
	//fmt.Println("DB_NAME:", dbName)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	// Миграция схемы
	err = DB.AutoMigrate(&models.User{}, &models.Car{}, &models.Booking{})
	if err != nil {
		log.Fatal("Ошибка миграции схемы:", err)
	}

	fmt.Println("Успешное подключение к базе данных!")
}
