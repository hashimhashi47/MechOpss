package db

import (
	"MechOpss/internal/src/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	root := os.Getenv("DB_ROOT")

	DB, err = gorm.Open(mysql.Open(root), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect Database:", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Booking{},
		&models.Admin{},
		&models.Staff{},
		&models.Booked{},
		&models.Slot{},
	)

	if err != nil {
		log.Fatal("Failed to AutoMigrate:", err)
	}

	log.Println("Database Connected Successfully")
	return DB
}
