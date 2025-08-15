package services

import (
	"docbook/entity"
	"docbook/repository"
)

type MedicalHistoryService interface {
	CreateMedicalHistory(medicalHistory *entity.MedicalHistory) error
	GetMedicalHistoryByID(id uint) (*entity.MedicalHistoryResponse, error)
	GetMedicalHistoryByUserID(userID uint) ([]entity.MedicalHistoryResponse, error)
	UpdateMedicalHistory(id uint, userID uint, request *entity.UpdateMedicalHistoryRequest) error
	DeleteMedicalHistory(id uint, userID uint) error
}

type medicalHistoryService struct {
	medicalHistoryRepo repository.MedicalHistoryRepository
}

func NewMedicalHistoryService(medicalHistoryRepo repository.MedicalHistoryRepository) MedicalHistoryService {
	return &medicalHistoryService{medicalHistoryRepo: medicalHistoryRepo}
}

func (s *medicalHistoryService) CreateMedicalHistory(medicalHistory *entity.MedicalHistory) error {
	return s.medicalHistoryRepo.CreateMedicalHistory(medicalHistory)
}

func (s *medicalHistoryService) GetMedicalHistoryByID(id uint) (*entity.MedicalHistoryResponse, error) {
	return s.medicalHistoryRepo.GetMedicalHistoryByID(id)
}

func (s *medicalHistoryService) GetMedicalHistoryByUserID(userID uint) ([]entity.MedicalHistoryResponse, error) {
	return s.medicalHistoryRepo.GetMedicalHistoryByUserID(userID)
}

func (s *medicalHistoryService) UpdateMedicalHistory(id uint, userID uint, request *entity.UpdateMedicalHistoryRequest) error {
	return s.medicalHistoryRepo.UpdateMedicalHistory(id, userID, request)
}

func (s *medicalHistoryService) DeleteMedicalHistory(id uint, userID uint) error {
	return s.medicalHistoryRepo.DeleteMedicalHistory(id, userID)
}
