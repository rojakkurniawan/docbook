package entity

import "gorm.io/gorm"

type TimeSlot struct {
	DoctorScheduleID uint           `json:"doctor_schedule_id" gorm:"not null"`
	Date             string         `json:"date" gorm:"not null"`
	StartTime        string         `json:"start_time" gorm:"not null"`
	EndTime          string         `json:"end_time" gorm:"not null"`
	IsAvailable      bool           `json:"is_available" gorm:"default:true"`
	IsBlocked        bool           `json:"is_blocked" gorm:"default:false"`
	BlockReason      string         `json:"block_reason"`
	DoctorSchedule   DoctorSchedule `json:"doctor_schedule" gorm:"foreignKey:DoctorScheduleID"`
	gorm.Model
}

type CreateTimeslotRequest struct {
	DoctorScheduleID uint             `json:"doctor_schedule_id" binding:"required"`
	Date             string           `json:"date" binding:"required"`
	TimeSlots        []TimeslotDetail `json:"time_slots" binding:"required"`
}

type TimeslotDetail struct {
	ID          uint   `json:"id"`
	StartTime   string `json:"start_time" binding:"required"`
	EndTime     string `json:"end_time" binding:"required"`
	IsAvailable bool   `json:"is_available"`
	IsBlocked   bool   `json:"is_blocked"`
}

type TimeslotResponse struct {
	DoctorScheduleID uint             `json:"doctor_schedule_id"`
	Date             string           `json:"date"`
	Timeslots        []TimeslotDetail `json:"timeslots"`
}

type TimeslotFilter struct {
	DoctorScheduleID *uint   `json:"doctor_schedule_id" form:"doctor_schedule_id"`
	Date             *string `json:"date" form:"date"`
	StartDate        *string `json:"start_date" form:"start_date"`
	EndDate          *string `json:"end_date" form:"end_date"`
	IsAvailable      *bool   `json:"is_available" form:"is_available"`
	IsBlocked        *bool   `json:"is_blocked" form:"is_blocked"`
	StartTime        *string `json:"start_time" form:"start_time"`
	EndTime          *string `json:"end_time" form:"end_time"`
}
