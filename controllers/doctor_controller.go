package controllers

import (
	"docbook/entity"
	"docbook/services"
	"docbook/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DoctorController struct {
	doctorService services.DoctorService
}

func NewDoctorController(doctorService services.DoctorService) *DoctorController {
	return &DoctorController{doctorService: doctorService}
}

func (dc *DoctorController) GetDoctorProfileByUserID(c *gin.Context) {
	userID := c.GetUint("user_id")

	doctor, err := dc.doctorService.GetDoctorProfileByUserID(userID)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	doctorResponse := entity.DoctorProfileResponse{
		User: entity.UserProfileResponse{
			FirstName: doctor.User.FirstName,
			LastName:  doctor.User.LastName,
			Email:     doctor.User.Email,
			Phone:     doctor.User.Phone,
		},
		Doctor: entity.DoctorResponse{
			DoctorID:          doctor.ID,
			LicenseNumber:     doctor.LicenseNumber,
			ConsultationFee:   doctor.ConsultationFee,
			YearsOfExperience: doctor.YearsOfExperience,
			TotalPatients:     doctor.TotalPatients,
			Biography:         doctor.Biography,
		},
		Specialization: entity.SpecializationResponse{
			SpecializationID: doctor.Specialization.ID,
			Name:             doctor.Specialization.Name,
			Description:      doctor.Specialization.Description,
		},
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Doctor profile retrieved successfully", doctorResponse, nil)
}

func (dc *DoctorController) GetAllDoctors(c *gin.Context) {
	doctor, err := dc.doctorService.GetAllDoctors()
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	var doctorResponse []entity.DoctorProfileResponse

	for _, doctor := range doctor {
		doctorResponse = append(doctorResponse, entity.DoctorProfileResponse{
			User: entity.UserProfileResponse{
				FirstName: doctor.User.FirstName,
				LastName:  doctor.User.LastName,
				Email:     doctor.User.Email,
				Phone:     doctor.User.Phone,
			},
			Doctor: entity.DoctorResponse{
				DoctorID:          doctor.ID,
				LicenseNumber:     doctor.LicenseNumber,
				ConsultationFee:   doctor.ConsultationFee,
				YearsOfExperience: doctor.YearsOfExperience,
				TotalPatients:     doctor.TotalPatients,
				Biography:         doctor.Biography,
			},
			Specialization: entity.SpecializationResponse{
				SpecializationID: doctor.Specialization.ID,
				Name:             doctor.Specialization.Name,
				Description:      doctor.Specialization.Description,
			},
		})
	}

	response.BuildSuccessResponse(c, http.StatusOK, "All doctors retrieved successfully", doctorResponse, nil)
}
