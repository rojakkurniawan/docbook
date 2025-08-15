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

type ScheduleController struct {
	scheduleService services.ScheduleService
}

func NewScheduleController(scheduleService services.ScheduleService) *ScheduleController {
	return &ScheduleController{scheduleService: scheduleService}
}

func (sc *ScheduleController) CreateSchedule(c *gin.Context) {
	var request entity.CreateScheduleRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	if err := sc.scheduleService.CreateMultipleSchedules(&request); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusCreated, "Schedule created successfully", nil, nil)
}

func (sc *ScheduleController) GetScheduleByUserID(c *gin.Context) {
	userID := c.GetUint("user_id")

	schedule, err := sc.scheduleService.GetScheduleByUserID(userID)
	if err != nil {
		response.BuildErrorResponse(c, errormodel.ErrDoctorNotFound)
		return
	}

	var scheduleResponse []entity.ScheduleResponse

	for _, schedule := range schedule {
		scheduleResponse = append(scheduleResponse, entity.ScheduleResponse{
			ID:          schedule.ID,
			DoctorID:    schedule.DoctorID,
			DayOfWeek:   schedule.DayOfWeek,
			StartTime:   schedule.StartTime,
			EndTime:     schedule.EndTime,
			IsAvailable: schedule.IsAvailable,
			MaxPatients: schedule.MaxPatients,
			Duration:    schedule.Duration,
			CreatedAt:   schedule.CreatedAt,
			UpdatedAt:   schedule.UpdatedAt,
		})
	}
	response.BuildSuccessResponse(c, http.StatusOK, "Schedule fetched successfully", scheduleResponse, nil)
}

func (sc *ScheduleController) GetScheduleByID(c *gin.Context) {
	id := c.Param("id")

	scheduleID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	schedule, err := sc.scheduleService.GetScheduleByID(uint(scheduleID))
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	scheduleResponse := entity.ScheduleResponse{
		ID:          schedule.ID,
		DoctorID:    schedule.DoctorID,
		DayOfWeek:   schedule.DayOfWeek,
		StartTime:   schedule.StartTime,
		EndTime:     schedule.EndTime,
		IsAvailable: schedule.IsAvailable,
		MaxPatients: schedule.MaxPatients,
		Duration:    schedule.Duration,
		CreatedAt:   schedule.CreatedAt,
		UpdatedAt:   schedule.UpdatedAt,
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Schedule fetched successfully", scheduleResponse, nil)
}

func (sc *ScheduleController) UpdateSchedule(c *gin.Context) {
	id := c.Param("id")

	scheduleID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	var schedule entity.DoctorSchedule

	if err := c.ShouldBindJSON(&schedule); err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	schedule.ID = uint(scheduleID)

	if err := sc.scheduleService.UpdateSchedule(uint(scheduleID), &schedule); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Schedule updated successfully", nil, nil)
}

func (sc *ScheduleController) DeleteSchedule(c *gin.Context) {
	id := c.Param("id")

	scheduleID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	if err := sc.scheduleService.DeleteSchedule(uint(scheduleID)); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Schedule deleted successfully", nil, nil)
}
