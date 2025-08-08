package entity

import "gorm.io/gorm"

type Specialization struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"type:text"`
}

type Doctor struct {
	gorm.Model
	UserID            uint           `json:"user_id" gorm:"not null"`
	SpecializationID  uint           `json:"specialization_id" gorm:"not null"`
	LicenseNumber     string         `json:"license_number" gorm:"unique;not null"`
	ConsultationFee   float64        `json:"consultation_fee" gorm:"default:0"`
	YearsOfExperience int            `json:"years_of_experience" gorm:"default:0"`
	Biography         string         `json:"biography" gorm:"type:text"`
	Specialization    Specialization `json:"specialization" gorm:"foreignKey:SpecializationID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type DoctorSchedule struct {
	gorm.Model
	DoctorID  uint   `json:"doctor_id" gorm:"not null"`
	DayOfWeek int    `json:"day_of_week" gorm:"not null"`
	StartTime string `json:"start_time" gorm:"not null"`
	EndTime   string `json:"end_time" gorm:"not null"`
	Doctor    Doctor `json:"doctor" gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
