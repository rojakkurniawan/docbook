package config

import (
	"docbook/consts"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dbHost := os.Getenv(consts.DBHostEnv)
	dbPort := os.Getenv(consts.DBPortEnv)
	dbUser := os.Getenv(consts.DBUserEnv)
	dbPass := os.Getenv(consts.DBPasswordEnv)
	dbName := os.Getenv(consts.DBNameEnv)

	if dbHost == "" || dbPort == "" || dbUser == "" || dbName == "" {
		log.Fatal("Missing required database environment variables")
	}

	dsn := fmt.Sprintf(consts.MySQLDSNFormat, dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("%s: %v", consts.ErrDatabaseConnectionError, err)
	}

	log.Println("Successfully connected to database")
	return db
}
