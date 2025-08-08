package entity

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	BookingCode string `json:"booking_code" gorm:"unique;not null"`
	UserID      uint   `json:"user_id" gorm:"not null"`
	DoctorID    uint   `json:"doctor_id" gorm:"not null"`
	BookingDate string `json:"booking_date" gorm:"not null"`
	BookingTime string `json:"booking_time" gorm:"not null"`
	Status      string `json:"status" gorm:"type:ENUM('pending', 'confirmed', 'cancelled', 'completed');default:'pending'"`
	Notes       string `json:"notes" gorm:"type:text"`
	User        User   `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
