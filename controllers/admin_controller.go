package controllers

import (
	"docbook/entity"
	"docbook/services"
	errormodel "docbook/utils/errorModel"
	"docbook/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService services.AdminService
}

func NewAdminController(adminService services.AdminService) *AdminController {
	return &AdminController{adminService: adminService}
}

func (ac *AdminController) CreateDoctor(c *gin.Context) {
	var createDoctorRequest entity.CreateDoctorRequest

	if err := c.ShouldBindJSON(&createDoctorRequest); err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	var (
		detailError = make(map[string]any)
	)

	if createDoctorRequest.User.FirstName == "" {
		detailError["user_first_name"] = "User first name is required"
	}

	if createDoctorRequest.User.LastName == "" {
		detailError["user_last_name"] = "User last name is required"
	}

	if createDoctorRequest.User.Email == "" {
		detailError["user_email"] = "User email is required"
	}

	if createDoctorRequest.User.Password == "" {
		detailError["user_password"] = "User password is required"
	}

	if createDoctorRequest.Doctor.LicenseNumber == "" {
		detailError["doctor_license_number"] = "Doctor license number is required"
	}

	if createDoctorRequest.Doctor.ConsultationFee == 0 {
		detailError["doctor_consultation_fee"] = "Doctor consultation fee is required"
	}

	if createDoctorRequest.Doctor.YearsOfExperience == 0 {
		detailError["doctor_years_of_experience"] = "Doctor years of experience is required"
	}

	if createDoctorRequest.Doctor.Biography == "" {
		detailError["doctor_biography"] = "Doctor biography is required"
	}

	if createDoctorRequest.Specialization.Name == "" {
		detailError["specialization_name"] = "Specialization name is required"
	}

	if createDoctorRequest.Specialization.Description == "" {
		detailError["specialization_description"] = "Specialization description is required"
	}

	if len(detailError) > 0 {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput.AttachDetail(detailError))
		return
	}

	err := ac.adminService.CreateDoctor(&createDoctorRequest)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusCreated, "Doctor created successfully", nil, nil)
}
