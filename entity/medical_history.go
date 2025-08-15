package entity

import (
	"time"

	"gorm.io/gorm"
)

type MedicalHistory struct {
	gorm.Model
	BookingID         uint    `json:"booking_id" gorm:"not null"`
	UserID            uint    `json:"user_id" gorm:"not null"`
	Allergies         string  `json:"allergies" gorm:"type:text"`
	ChronicDiseases   string  `json:"chronic_diseases" gorm:"type:text"`
	CurrentMedication string  `json:"current_medication" gorm:"type:text"`
	BloodType         string  `json:"blood_type" gorm:"type:ENUM('A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-')"`
	Height            float64 `json:"height" gorm:"type:decimal(5,2)"`
	Weight            float64 `json:"weight" gorm:"type:decimal(5,2)"`
	Notes             string  `json:"notes" gorm:"type:text"`
	Booking           Booking `json:"booking" gorm:"foreignKey:BookingID"`
	User              User    `json:"user" gorm:"foreignKey:UserID"`
}

// CreateMedicalHistoryRequest untuk request dari user
type CreateMedicalHistoryRequest struct {
	BookingID         uint    `json:"booking_id" binding:"required"`
	Allergies         string  `json:"allergies"`
	ChronicDiseases   string  `json:"chronic_diseases"`
	CurrentMedication string  `json:"current_medication"`
	BloodType         string  `json:"blood_type" binding:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O-"`
	Height            float64 `json:"height" binding:"omitempty,min=0,max=300"`
	Weight            float64 `json:"weight" binding:"omitempty,min=0,max=500"`
	Notes             string  `json:"notes"`
}

// UpdateMedicalHistoryRequest untuk update medical history
type UpdateMedicalHistoryRequest struct {
	Allergies         string  `json:"allergies"`
	ChronicDiseases   string  `json:"chronic_diseases"`
	CurrentMedication string  `json:"current_medication"`
	BloodType         string  `json:"blood_type" binding:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O-"`
	Height            float64 `json:"height" binding:"omitempty,min=0,max=300"`
	Weight            float64 `json:"weight" binding:"omitempty,min=0,max=500"`
	Notes             string  `json:"notes"`
}

// MedicalHistoryResponse untuk response yang lebih bersih
type MedicalHistoryResponse struct {
	ID                uint      `json:"id"`
	BookingID         uint      `json:"booking_id"`
	UserID            uint      `json:"user_id"`
	Allergies         string    `json:"allergies"`
	ChronicDiseases   string    `json:"chronic_diseases"`
	CurrentMedication string    `json:"current_medication"`
	BloodType         string    `json:"blood_type"`
	Height            float64   `json:"height"`
	Weight            float64   `json:"weight"`
	Notes             string    `json:"notes"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
