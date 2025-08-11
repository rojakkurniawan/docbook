package main

import (
	"docbook/controllers"
	"docbook/repository"
	"docbook/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	authService := services.NewAuthService(repository.NewUserRepository(db))
	authController := controllers.NewAuthController(authService)

	api := router.Group("/api")
	{
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)
	}
}
