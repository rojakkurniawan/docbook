package controllers

import (
	"docbook/entity"
	"docbook/services"
	errormodel "docbook/utils/errorModel"
	"docbook/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TimeslotController struct {
	timeslotService services.TimeslotService
}

func NewTimeslotController(timeslotService services.TimeslotService) *TimeslotController {
	return &TimeslotController{timeslotService: timeslotService}
}

func (tc *TimeslotController) CreateTimeslot(c *gin.Context) {
	var request entity.CreateTimeslotRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	if err := tc.timeslotService.CreateMultipleTimeslots(&request); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusCreated, "Timeslots created successfully", nil, nil)
}

func (tc *TimeslotController) GetAllTimeslots(c *gin.Context) {
	// Parse filter parameters from query string
	var filter entity.TimeslotFilter

	// Parse doctor_schedule_id
	if doctorScheduleIDStr := c.Query("doctor_schedule_id"); doctorScheduleIDStr != "" {
		if doctorScheduleID, err := strconv.ParseUint(doctorScheduleIDStr, 10, 32); err == nil {
			doctorScheduleIDUint := uint(doctorScheduleID)
			filter.DoctorScheduleID = &doctorScheduleIDUint
		}
	}

	// Parse date
	if date := c.Query("date"); date != "" {
		filter.Date = &date
	}

	// Parse start_date
	if startDate := c.Query("start_date"); startDate != "" {
		filter.StartDate = &startDate
	}

	// Parse end_date
	if endDate := c.Query("end_date"); endDate != "" {
		filter.EndDate = &endDate
	}

	// Parse is_available
	if isAvailableStr := c.Query("is_available"); isAvailableStr != "" {
		if isAvailable, err := strconv.ParseBool(isAvailableStr); err == nil {
			filter.IsAvailable = &isAvailable
		}
	}

	// Parse is_blocked
	if isBlockedStr := c.Query("is_blocked"); isBlockedStr != "" {
		if isBlocked, err := strconv.ParseBool(isBlockedStr); err == nil {
			filter.IsBlocked = &isBlocked
		}
	}

	// Parse start_time
	if startTime := c.Query("start_time"); startTime != "" {
		filter.StartTime = &startTime
	}

	// Parse end_time
	if endTime := c.Query("end_time"); endTime != "" {
		filter.EndTime = &endTime
	}

	// Check if any filter is applied
	hasFilter := filter.DoctorScheduleID != nil || filter.Date != nil ||
		filter.StartDate != nil || filter.EndDate != nil ||
		filter.IsAvailable != nil || filter.IsBlocked != nil ||
		filter.StartTime != nil || filter.EndTime != nil

	var timeslots []entity.TimeslotResponse
	var err error

	if hasFilter {
		timeslots, err = tc.timeslotService.GetTimeslotsWithFilter(&filter)
	} else {
		timeslots, err = tc.timeslotService.GetAllTimeslots()
	}

	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Timeslots retrieved successfully", timeslots, nil)
}

func (tc *TimeslotController) GetTimeslotByID(c *gin.Context) {
	id := c.Param("id")

	timeslotID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	timeslot, err := tc.timeslotService.GetTimeslotByID(uint(timeslotID))
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Timeslot retrieved successfully", timeslot, nil)
}

func (tc *TimeslotController) UpdateTimeslot(c *gin.Context) {
	id := c.Param("id")

	timeslotID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	var timeslot entity.TimeSlot

	if err := c.ShouldBindJSON(&timeslot); err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	timeslot.ID = uint(timeslotID)

	if err := tc.timeslotService.UpdateTimeslot(uint(timeslotID), &timeslot); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Timeslot updated successfully", nil, nil)
}

func (tc *TimeslotController) DeleteTimeslot(c *gin.Context) {
	id := c.Param("id")

	timeslotID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	if err := tc.timeslotService.DeleteTimeslot(uint(timeslotID)); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Timeslot deleted successfully", nil, nil)
}
