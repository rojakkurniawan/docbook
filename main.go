package main

import (
	"docbook/config"
	"docbook/entity"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitializeApp() *gin.Engine {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading ENV")
	}

	r := gin.Default()

	db := config.ConnectDatabase()

	db.AutoMigrate(&entity.User{}, &entity.Specialization{}, &entity.Doctor{}, &entity.DoctorSchedule{}, &entity.Booking{}, &entity.BookingPatient{}, &entity.TimeSlot{}, &entity.Hospital{}, &entity.MedicalHistory{})

	SetupRoutes(r, db)

	return r
}

func main() {
	app := InitializeApp()
	app.Run(":8080")
}
