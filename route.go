package main

import (
	"docbook/controllers"
	"docbook/middleware"
	"docbook/repository"
	"docbook/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// =============================== DEPENDENCY INJECTION ===============================
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	adminRepo := repository.NewAdminRepository(db)
	adminService := services.NewAdminService(adminRepo)
	adminController := controllers.NewAdminController(adminService)

	doctorRepo := repository.NewDoctorRepository(db)
	doctorService := services.NewDoctorService(doctorRepo)
	doctorController := controllers.NewDoctorController(doctorService)

	scheduleRepo := repository.NewScheduleRepository(db)
	scheduleService := services.NewScheduleService(scheduleRepo)
	scheduleController := controllers.NewScheduleController(scheduleService)

	timeslotRepo := repository.NewTimeslotRepository(db)
	timeslotService := services.NewTimeslotService(timeslotRepo)
	timeslotController := controllers.NewTimeslotController(timeslotService)

	bookingRepo := repository.NewBookingRepository(db)
	bookingService := services.NewBookingService(bookingRepo)
	bookingController := controllers.NewBookingController(bookingService)

	medicalHistoryRepo := repository.NewMedicalHistoryRepository(db)
	medicalHistoryService := services.NewMedicalHistoryService(medicalHistoryRepo)
	medicalHistoryController := controllers.NewMedicalHistoryController(medicalHistoryService)

	// =============================== API ROUTES ===============================
	api := router.Group("/api")

	// Authentication routes
	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("/register", userController.Register)
		authRoutes.POST("/login", userController.Login)
		authRoutes.POST("/refresh-token", userController.RefreshToken)
	}

	// User routes (requires authentication)
	userRoutes := api.Group("/user")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.GET("/profile", userController.GetUserByID)
		userRoutes.PUT("/profile", userController.UpdateUser)
		userRoutes.PUT("/change-password", userController.ChangePassword)
		userRoutes.DELETE("/account", userController.DeleteUser)
	}

	// Admin routes (requires authentication and admin role)
	adminRoutes := api.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware(), middleware.RBACAdminMiddleware())
	{
		adminRoutes.POST("/doctors", adminController.CreateDoctor)
	}

	// Doctor routes (requires authentication and doctor role)
	doctorRoutes := api.Group("/doctor")
	doctorRoutes.Use(middleware.AuthMiddleware(), middleware.RBACDoctorMiddleware())
	{
		doctorRoutes.GET("/profile", doctorController.GetDoctorProfileByUserID)

		// Doctor schedule management
		doctorScheduleRoutes := doctorRoutes.Group("/schedules")
		{
			doctorScheduleRoutes.GET("", scheduleController.GetScheduleByUserID)
			doctorScheduleRoutes.POST("", scheduleController.CreateSchedule)
			doctorScheduleRoutes.PUT("/:id", scheduleController.UpdateSchedule)
			doctorScheduleRoutes.DELETE("/:id", scheduleController.DeleteSchedule)
		}

		// Timeslot management
		timeslotRoutes := doctorRoutes.Group("/timeslots")
		{
			timeslotRoutes.POST("", timeslotController.CreateTimeslot)
			timeslotRoutes.GET("", timeslotController.GetAllTimeslots)
			timeslotRoutes.GET("/:id", timeslotController.GetTimeslotByID)
			timeslotRoutes.PUT("/:id", timeslotController.UpdateTimeslot)
			timeslotRoutes.DELETE("/:id", timeslotController.DeleteTimeslot)
		}
	}

	// Public doctor routes (requires authentication)
	doctorPublicRoutes := api.Group("/doctors")
	doctorPublicRoutes.Use(middleware.AuthMiddleware())
	{
		doctorPublicRoutes.GET("", doctorController.GetAllDoctors)
	}

	// Public schedule routes (requires authentication)
	scheduleRoutes := api.Group("/schedules")
	scheduleRoutes.Use(middleware.AuthMiddleware())
	{
		scheduleRoutes.GET("/:id", scheduleController.GetScheduleByID)
	}

	// Public timeslot routes (requires authentication)
	timeslotRoutes := api.Group("/timeslots")
	timeslotRoutes.Use(middleware.AuthMiddleware())
	{
		timeslotRoutes.GET("", timeslotController.GetAllTimeslots)
		timeslotRoutes.GET("/:id", timeslotController.GetTimeslotByID)
	}

	// Booking routes (requires authentication)
	bookingRoutes := api.Group("/bookings")
	bookingRoutes.Use(middleware.AuthMiddleware())
	{
		bookingRoutes.POST("", bookingController.CreateBooking)
		bookingRoutes.GET("/:id", bookingController.GetBookingByID)
		bookingRoutes.PUT("/:id", bookingController.UpdateBooking)
	}

	// Medical History routes (requires authentication)
	medicalHistoryRoutes := api.Group("/medical-history")
	medicalHistoryRoutes.Use(middleware.AuthMiddleware())
	{
		medicalHistoryRoutes.POST("", medicalHistoryController.CreateMedicalHistory)
		medicalHistoryRoutes.GET("/:id", medicalHistoryController.GetMedicalHistoryByID)
		medicalHistoryRoutes.GET("/user", medicalHistoryController.GetMedicalHistoryByUserID)
		medicalHistoryRoutes.PUT("/:id", medicalHistoryController.UpdateMedicalHistory)
		medicalHistoryRoutes.DELETE("/:id", medicalHistoryController.DeleteMedicalHistory)
	}

}
