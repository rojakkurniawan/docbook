package config

import (
	"docbook/entity"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

const (
	dbHost     = "DB_HOST"
	dbPort     = "DB_PORT"
	dbUser     = "DB_USER"
	dbPassword = "DB_PASSWORD"
	dbName     = "DB_NAME"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func GetConfig() Config {
	return Config{
		Port: os.Getenv(dbPort),
		DatabaseURL: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv(dbUser),
			os.Getenv(dbPassword),
			os.Getenv(dbHost),
			os.Getenv(dbPort),
			os.Getenv(dbName),
		),
	}
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Specialization{},
		&entity.Doctor{},
		&entity.DoctorSchedule{},
		&entity.Booking{},
		&entity.BookingPatient{},
		&entity.TimeSlot{},
		&entity.MedicalHistory{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Successfully migrated database")
}
