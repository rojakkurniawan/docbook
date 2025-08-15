package entity

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	BookingCode string         `json:"booking_code" gorm:"unique;not null"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	DoctorID    uint           `json:"doctor_id" gorm:"not null"`
	TimeSlotID  uint           `json:"time_slot_id" gorm:"not null"`
	BookingDate string         `json:"booking_date" gorm:"not null"`
	BookingTime string         `json:"booking_time" gorm:"not null"`
	Status      string         `json:"status" gorm:"type:ENUM('pending', 'confirmed', 'cancelled', 'completed', 'no_show');default:'pending'"`
	Notes       string         `json:"notes" gorm:"type:text"`
	Symptoms    string         `json:"symptoms" gorm:"type:text"`
	User        User           `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Doctor      Doctor         `json:"doctor" gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TimeSlot    TimeSlot       `json:"time_slot" gorm:"foreignKey:TimeSlotID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Patient     BookingPatient `json:"patient" gorm:"foreignKey:BookingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	gorm.Model
}

// BookingResponse untuk response yang lebih bersih
type BookingResponse struct {
	ID          uint            `json:"id"`
	BookingCode string          `json:"booking_code"`
	UserID      uint            `json:"user_id"`
	DoctorID    uint            `json:"doctor_id"`
	TimeSlotID  uint            `json:"time_slot_id"`
	BookingDate string          `json:"booking_date"`
	BookingTime string          `json:"booking_time"`
	Status      string          `json:"status"`
	Notes       string          `json:"notes"`
	Symptoms    string          `json:"symptoms"`
	Patient     PatientResponse `json:"patient"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// PatientResponse untuk response patient
type PatientResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	BirthDate string    `json:"birth_date"`
	Sex       string    `json:"sex"`
	NIK       string    `json:"nik"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateBookingRequest untuk request dari user
type CreateBookingRequest struct {
	UserID      uint           `json:"user_id" binding:"required"`
	DoctorID    uint           `json:"doctor_id" binding:"required"`
	TimeSlotID  uint           `json:"time_slot_id" binding:"required"`
	BookingDate string         `json:"booking_date" binding:"required"`
	BookingTime string         `json:"booking_time" binding:"required"`
	Notes       string         `json:"notes"`
	Symptoms    string         `json:"symptoms"`
	Patient     PatientRequest `json:"patient" binding:"required"`
}

// PatientRequest untuk data patient dalam request
type PatientRequest struct {
	Name      string `json:"name" binding:"required"`
	BirthDate string `json:"birth_date" binding:"required"`
	Sex       string `json:"sex" binding:"required,oneof=male female"`
	NIK       string `json:"nik" binding:"required,len=16"`
	Address   string `json:"address" binding:"required"`
}

// UpdateBookingRequest untuk request dari user (hanya bisa cancel)
type UpdateBookingRequest struct {
	Status   string `json:"status" binding:"required,oneof=cancelled"`
	Notes    string `json:"notes"`
	Symptoms string `json:"symptoms"`
}

// UpdateBookingAdminRequest untuk request dari admin/doctor
type UpdateBookingAdminRequest struct {
	Status   string `json:"status" binding:"required,oneof=pending confirmed cancelled completed no_show"`
	Notes    string `json:"notes"`
	Symptoms string `json:"symptoms"`
}

type BookingPatient struct {
	BookingID uint   `json:"booking_id" gorm:"not null;uniqueIndex"`
	Name      string `json:"name" gorm:"not null"`
	BirthDate string `json:"birth_date" gorm:"not null"`
	Sex       string `json:"sex" gorm:"type:ENUM('male', 'female');not null"`
	NIK       string `json:"nik" gorm:"not null;size:16"`
	Address   string `json:"address" gorm:"type:text;not null"`
	gorm.Model
}
