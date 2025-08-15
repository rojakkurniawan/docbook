package repository

import (
	"docbook/entity"
	"time"

	"gorm.io/gorm"
)

type MedicalHistoryRepository interface {
	CreateMedicalHistory(medicalHistory *entity.MedicalHistory) error
	GetMedicalHistoryByID(id uint) (*entity.MedicalHistoryResponse, error)
	GetMedicalHistoryByUserID(userID uint) ([]entity.MedicalHistoryResponse, error)
	UpdateMedicalHistory(id uint, userID uint, request *entity.UpdateMedicalHistoryRequest) error
	DeleteMedicalHistory(id uint, userID uint) error
}

type medicalHistoryRepository struct {
	db *gorm.DB
}

func NewMedicalHistoryRepository(db *gorm.DB) MedicalHistoryRepository {
	return &medicalHistoryRepository{db: db}
}

func (r *medicalHistoryRepository) CreateMedicalHistory(medicalHistory *entity.MedicalHistory) error {
	return r.db.Create(medicalHistory).Error
}

func (r *medicalHistoryRepository) GetMedicalHistoryByID(id uint) (*entity.MedicalHistoryResponse, error) {
	var medicalHistory entity.MedicalHistory
	if err := r.db.Where("id = ?", id).First(&medicalHistory).Error; err != nil {
		return nil, err
	}

	return &entity.MedicalHistoryResponse{
		ID:                medicalHistory.ID,
		BookingID:         medicalHistory.BookingID,
		UserID:            medicalHistory.UserID,
		Allergies:         medicalHistory.Allergies,
		ChronicDiseases:   medicalHistory.ChronicDiseases,
		CurrentMedication: medicalHistory.CurrentMedication,
		BloodType:         medicalHistory.BloodType,
		Height:            medicalHistory.Height,
		Weight:            medicalHistory.Weight,
		Notes:             medicalHistory.Notes,
		CreatedAt:         medicalHistory.CreatedAt,
		UpdatedAt:         medicalHistory.UpdatedAt,
	}, nil
}

func (r *medicalHistoryRepository) GetMedicalHistoryByUserID(userID uint) ([]entity.MedicalHistoryResponse, error) {
	var medicalHistories []entity.MedicalHistory
	if err := r.db.Where("user_id = ?", userID).Find(&medicalHistories).Error; err != nil {
		return nil, err
	}

	var responses []entity.MedicalHistoryResponse
	for _, mh := range medicalHistories {
		responses = append(responses, entity.MedicalHistoryResponse{
			ID:                mh.ID,
			BookingID:         mh.BookingID,
			UserID:            mh.UserID,
			Allergies:         mh.Allergies,
			ChronicDiseases:   mh.ChronicDiseases,
			CurrentMedication: mh.CurrentMedication,
			BloodType:         mh.BloodType,
			Height:            mh.Height,
			Weight:            mh.Weight,
			Notes:             mh.Notes,
			CreatedAt:         mh.CreatedAt,
			UpdatedAt:         mh.UpdatedAt,
		})
	}

	return responses, nil
}

func (r *medicalHistoryRepository) UpdateMedicalHistory(id uint, userID uint, request *entity.UpdateMedicalHistoryRequest) error {
	// Check if medical history exists and belongs to user
	var medicalHistory entity.MedicalHistory
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&medicalHistory).Error; err != nil {
		return err
	}

	updates := map[string]interface{}{
		"updated_at": time.Now(),
	}

	if request.Allergies != "" {
		updates["allergies"] = request.Allergies
	}
	if request.ChronicDiseases != "" {
		updates["chronic_diseases"] = request.ChronicDiseases
	}
	if request.CurrentMedication != "" {
		updates["current_medication"] = request.CurrentMedication
	}
	if request.BloodType != "" {
		updates["blood_type"] = request.BloodType
	}
	if request.Height > 0 {
		updates["height"] = request.Height
	}
	if request.Weight > 0 {
		updates["weight"] = request.Weight
	}
	if request.Notes != "" {
		updates["notes"] = request.Notes
	}

	return r.db.Model(&entity.MedicalHistory{}).Where("id = ?", id).Updates(updates).Error
}

func (r *medicalHistoryRepository) DeleteMedicalHistory(id uint, userID uint) error {
	// Check if medical history exists and belongs to user
	var medicalHistory entity.MedicalHistory
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&medicalHistory).Error; err != nil {
		return err
	}

	return r.db.Delete(&medicalHistory).Error
}
