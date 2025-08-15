package controllers

import (
	"docbook/entity"
	"docbook/services"
	"docbook/utils"
	errormodel "docbook/utils/errorModel"
	"docbook/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	bookingService services.BookingService
}

func NewBookingController(bookingService services.BookingService) *BookingController {
	return &BookingController{bookingService: bookingService}
}

func (bc *BookingController) CreateBooking(c *gin.Context) {
	var request entity.CreateBookingRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	// Convert request to booking entity
	booking := &entity.Booking{
		UserID:      request.UserID,
		DoctorID:    request.DoctorID,
		TimeSlotID:  request.TimeSlotID,
		BookingDate: request.BookingDate,
		BookingTime: request.BookingTime,
		Status:      "pending", // Default status
		Notes:       request.Notes,
		Symptoms:    request.Symptoms,
	}

	if err := bc.bookingService.CreateBooking(booking, &request.Patient); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	// Get the created booking with patient data for response
	createdBooking, err := bc.bookingService.GetBookingByID(booking.ID)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusCreated, "Booking created successfully", createdBooking, nil)
}

func (bc *BookingController) GetBookingByID(c *gin.Context) {
	id := c.Param("id")

	bookingID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	booking, err := bc.bookingService.GetBookingByID(uint(bookingID))
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Booking retrieved successfully", booking, nil)
}

func (bc *BookingController) UpdateBooking(c *gin.Context) {
	id := c.Param("id")
	userRole := c.GetString(utils.DOCBOOK_ROLE)
	userID := c.GetUint(utils.DOCBOOK_USERID)

	bookingID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	// Get current booking to check ownership and current status
	currentBooking, err := bc.bookingService.GetBookingByID(uint(bookingID))
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	// Check if user is admin or doctor
	if userRole == "admin" || userRole == "doctor" {
		var request entity.UpdateBookingAdminRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
			return
		}

		if err := bc.bookingService.UpdateBookingAdmin(uint(bookingID), &request); err != nil {
			response.BuildErrorResponse(c, err)
			return
		}
	} else {
		// Regular user - can only cancel their own booking
		if currentBooking.UserID != userID {
			response.BuildErrorResponse(c, errormodel.ErrNotAllowed)
			return
		}

		var request entity.UpdateBookingRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
			return
		}

		// User can only cancel booking
		if request.Status != "cancelled" {
			response.BuildErrorResponse(c, errormodel.ErrNotAllowed)
			return
		}

		if err := bc.bookingService.UpdateBooking(uint(bookingID), &request); err != nil {
			response.BuildErrorResponse(c, err)
			return
		}
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Booking updated successfully", nil, nil)
}
