package entity

import (
	"time"

	"gorm.io/gorm"
)

type Specialization struct {
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"type:text"`
	gorm.Model
}

type Doctor struct {
	UserID            uint           `json:"user_id" gorm:"not null"`
	SpecializationID  uint           `json:"specialization_id" gorm:"not null"`
	LicenseNumber     string         `json:"license_number" gorm:"unique;not null"`
	ConsultationFee   float64        `json:"consultation_fee" gorm:"default:0"`
	YearsOfExperience int            `json:"years_of_experience" gorm:"default:0"`
	Biography         string         `json:"biography" gorm:"type:text"`
	IsActive          bool           `json:"is_active" gorm:"default:true"`
	TotalPatients     int            `json:"total_patients" gorm:"default:0"`
	User              User           `json:"user" gorm:"foreignKey:UserID"`
	Specialization    Specialization `json:"specialization" gorm:"foreignKey:SpecializationID"`
	gorm.Model
}

type DoctorSchedule struct {
	DoctorID    uint   `json:"doctor_id" gorm:"not null"`
	DayOfWeek   int    `json:"day_of_week" gorm:"not null"`
	StartTime   string `json:"start_time" gorm:"not null"`
	EndTime     string `json:"end_time" gorm:"not null"`
	IsAvailable bool   `json:"is_available" gorm:"default:true"`
	MaxPatients int    `json:"max_patients" gorm:"default:10"`
	Duration    int    `json:"duration" gorm:"default:30"`
	Doctor      Doctor `json:"doctor" gorm:"foreignKey:DoctorID"`
	gorm.Model
}

type CreateScheduleRequest struct {
	DoctorID  uint                    `json:"doctor_id" binding:"required"`
	Schedules []ScheduleDetailRequest `json:"schedules" binding:"required"`
}

type ScheduleDetailRequest struct {
	DayOfWeek   int    `json:"day_of_week" binding:"required,min=1,max=7"`
	StartTime   string `json:"start_time" binding:"required"`
	EndTime     string `json:"end_time" binding:"required"`
	IsAvailable bool   `json:"is_available"`
	MaxPatients int    `json:"max_patients" binding:"required,min=1"`
	Duration    int    `json:"duration" binding:"required,min=1"`
}

type ScheduleResponse struct {
	ID          uint      `json:"id"`
	DoctorID    uint      `json:"doctor_id"`
	DayOfWeek   int       `json:"day_of_week"`
	StartTime   string    `json:"start_time"`
	EndTime     string    `json:"end_time"`
	IsAvailable bool      `json:"is_available"`
	MaxPatients int       `json:"max_patients"`
	Duration    int       `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DoctorProfileResponse struct {
	User           UserProfileResponse    `json:"user"`
	Doctor         DoctorResponse         `json:"doctor"`
	Specialization SpecializationResponse `json:"specialization"`
}

type DoctorResponse struct {
	DoctorID          uint    `json:"doctor_id"`
	LicenseNumber     string  `json:"license_number"`
	ConsultationFee   float64 `json:"consultation_fee"`
	YearsOfExperience int     `json:"years_of_experience"`
	TotalPatients     int     `json:"total_patients"`
	Biography         string  `json:"biography"`
}

type UserProfileResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type SpecializationResponse struct {
	SpecializationID uint   `json:"specialization_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
}
