package errormodel

import (
	"docbook/utils/response"
	"net/http"
)

var (
	ErrDoctorAlreadyExists = response.ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "DOCTOR_ALREADY_EXISTS",
		Message:   "Doctor already exists",
	}
	ErrDoctorLicenseNumberAlreadyExists = response.ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "DOCTOR_LICENSE_NUMBER_ALREADY_EXISTS",
		Message:   "Doctor license number already exists",
	}
	ErrDoctorNotFound = response.ErrorResponse{
		Status:    http.StatusNotFound,
		ErrorCode: "DOCTOR_NOT_FOUND",
		Message:   "Doctor not found",
	}
	ErrUserNotDoctor = response.ErrorResponse{
		Status:    http.StatusForbidden,
		ErrorCode: "USER_NOT_DOCTOR",
		Message:   "User is not a doctor",
	}
)
