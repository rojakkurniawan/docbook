package services

import (
	"docbook/entity"
	"docbook/repository"
	errormodel "docbook/utils/errorModel"
	"errors"

	"gorm.io/gorm"
)

type AdminService interface {
	CreateDoctor(createDoctorRequest *entity.CreateDoctorRequest) error
}

type adminService struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(adminRepo repository.AdminRepository) AdminService {
	return &adminService{adminRepo: adminRepo}
}

func (s *adminService) CreateDoctor(createDoctorRequest *entity.CreateDoctorRequest) error {
	user, err := s.adminRepo.GetUserByEmail(createDoctorRequest.User.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if user != nil {
		return errormodel.ErrDoctorAlreadyExists
	}

	doctor, err := s.adminRepo.GetDoctorByLicenseNumber(createDoctorRequest.Doctor.LicenseNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if doctor != nil {
		return errormodel.ErrDoctorLicenseNumberAlreadyExists
	}

	return s.adminRepo.CreateDoctor(createDoctorRequest)
}
