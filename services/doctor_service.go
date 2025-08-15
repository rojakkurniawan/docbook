package services

import (
	"docbook/entity"
	"docbook/repository"
	errormodel "docbook/utils/errorModel"
)

type DoctorService interface {
	GetDoctorProfileByUserID(userID uint) (*entity.Doctor, error)
	GetAllDoctors() ([]entity.Doctor, error)
}

type doctorService struct {
	doctorRepo repository.DoctorRepository
}

func NewDoctorService(doctorRepo repository.DoctorRepository) DoctorService {
	return &doctorService{doctorRepo: doctorRepo}
}

func (s *doctorService) GetDoctorProfileByUserID(userID uint) (*entity.Doctor, error) {
	user, err := s.doctorRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if user.Role != "doctor" {
		return nil, errormodel.ErrDoctorNotFound
	}

	// Get doctor profile
	doctor, err := s.doctorRepo.GetDoctorByUserID(userID)
	if err != nil {
		return nil, err
	}

	return doctor, nil
}

func (s *doctorService) GetAllDoctors() ([]entity.Doctor, error) {
	doctor, err := s.doctorRepo.GetAllDoctors()
	if err != nil {
		return nil, err
	}

	return doctor, nil
}
