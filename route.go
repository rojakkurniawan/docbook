package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")
	{
		SetupAuthRoutes(api, db)

	}
}

func SetupAuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	authController := controllers.NewAuthController(db)

	protected := router.Group("/")
	{
		protected.POST("/register", authController.Register)
		protected.POST("/login", authController.Login)
	}
}
