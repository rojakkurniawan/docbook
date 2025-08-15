package main

import (
	"docbook/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitializeApp() *gin.Engine {
	godotenv.Load()

	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	r := gin.Default()

	conf := config.GetConfig()
	db := config.InitDatabaseMySQL(&conf)
	config.AutoMigrate(db)

	SetupRoutes(r, db)

	return r
}

func main() {
	app := InitializeApp()
	app.Run(":8080")
}
