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

type MedicalHistoryController struct {
	medicalHistoryService services.MedicalHistoryService
}

func NewMedicalHistoryController(medicalHistoryService services.MedicalHistoryService) *MedicalHistoryController {
	return &MedicalHistoryController{medicalHistoryService: medicalHistoryService}
}

func (mc *MedicalHistoryController) CreateMedicalHistory(c *gin.Context) {
	var request entity.CreateMedicalHistoryRequest
	userID := c.GetUint(utils.DOCBOOK_USERID)

	if err := c.ShouldBindJSON(&request); err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	// Convert request to medical history entity
	medicalHistory := &entity.MedicalHistory{
		BookingID:         request.BookingID,
		UserID:            userID,
		Allergies:         request.Allergies,
		ChronicDiseases:   request.ChronicDiseases,
		CurrentMedication: request.CurrentMedication,
		BloodType:         request.BloodType,
		Height:            request.Height,
		Weight:            request.Weight,
		Notes:             request.Notes,
	}

	if err := mc.medicalHistoryService.CreateMedicalHistory(medicalHistory); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusCreated, "Medical history created successfully", nil, nil)
}

func (mc *MedicalHistoryController) GetMedicalHistoryByID(c *gin.Context) {
	id := c.Param("id")

	medicalHistoryID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	medicalHistory, err := mc.medicalHistoryService.GetMedicalHistoryByID(uint(medicalHistoryID))
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Medical history retrieved successfully", medicalHistory, nil)
}

func (mc *MedicalHistoryController) GetMedicalHistoryByUserID(c *gin.Context) {
	userID := c.GetUint(utils.DOCBOOK_USERID)

	medicalHistories, err := mc.medicalHistoryService.GetMedicalHistoryByUserID(userID)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Medical histories retrieved successfully", medicalHistories, nil)
}

func (mc *MedicalHistoryController) UpdateMedicalHistory(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint(utils.DOCBOOK_USERID)

	medicalHistoryID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	var request entity.UpdateMedicalHistoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	if err := mc.medicalHistoryService.UpdateMedicalHistory(uint(medicalHistoryID), userID, &request); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Medical history updated successfully", nil, nil)
}

func (mc *MedicalHistoryController) DeleteMedicalHistory(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint(utils.DOCBOOK_USERID)

	medicalHistoryID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	if err := mc.medicalHistoryService.DeleteMedicalHistory(uint(medicalHistoryID), userID); err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Medical history deleted successfully", nil, nil)
}
