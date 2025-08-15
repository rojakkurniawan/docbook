package repository

import (
	"docbook/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking *entity.Booking, patient *entity.PatientRequest) error
	GetBookingByID(id uint) (*entity.BookingResponse, error)
	UpdateBooking(id uint, booking *entity.UpdateBookingRequest) error
	UpdateBookingAdmin(id uint, booking *entity.UpdateBookingAdminRequest) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) CreateBooking(booking *entity.Booking, patient *entity.PatientRequest) error {
	// Generate booking code
	booking.BookingCode = r.generateBookingCode()

	// Start transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Create booking first
	if err := tx.Create(booking).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Create patient data
	bookingPatient := &entity.BookingPatient{
		BookingID: booking.ID,
		Name:      patient.Name,
		BirthDate: patient.BirthDate,
		Sex:       patient.Sex,
		NIK:       patient.NIK,
		Address:   patient.Address,
	}

	if err := tx.Create(bookingPatient).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	return tx.Commit().Error
}

func (r *bookingRepository) GetBookingByID(id uint) (*entity.BookingResponse, error) {
	var booking entity.Booking
	var patient entity.BookingPatient

	// Get booking data
	if err := r.db.Where("id = ?", id).First(&booking).Error; err != nil {
		return nil, err
	}

	// Get patient data
	if err := r.db.Where("booking_id = ?", id).First(&patient).Error; err != nil {
		return nil, err
	}

	// Convert to response format
	response := &entity.BookingResponse{
		ID:          booking.ID,
		BookingCode: booking.BookingCode,
		UserID:      booking.UserID,
		DoctorID:    booking.DoctorID,
		TimeSlotID:  booking.TimeSlotID,
		BookingDate: booking.BookingDate,
		BookingTime: booking.BookingTime,
		Status:      booking.Status,
		Notes:       booking.Notes,
		Symptoms:    booking.Symptoms,
		CreatedAt:   booking.CreatedAt,
		UpdatedAt:   booking.UpdatedAt,
		Patient: entity.PatientResponse{
			ID:        patient.ID,
			Name:      patient.Name,
			BirthDate: patient.BirthDate,
			Sex:       patient.Sex,
			NIK:       patient.NIK,
			Address:   patient.Address,
			CreatedAt: patient.CreatedAt,
			UpdatedAt: patient.UpdatedAt,
		},
	}

	return response, nil
}

func (r *bookingRepository) UpdateBooking(id uint, booking *entity.UpdateBookingRequest) error {
	// Update only allowed fields for user
	updates := map[string]interface{}{
		"status":     booking.Status,
		"updated_at": time.Now(),
	}

	// Only update notes and symptoms if provided
	if booking.Notes != "" {
		updates["notes"] = booking.Notes
	}
	if booking.Symptoms != "" {
		updates["symptoms"] = booking.Symptoms
	}

	return r.db.Model(&entity.Booking{}).Where("id = ?", id).Updates(updates).Error
}

func (r *bookingRepository) UpdateBookingAdmin(id uint, booking *entity.UpdateBookingAdminRequest) error {
	// Admin/doctor can update all fields
	updates := map[string]interface{}{
		"status":     booking.Status,
		"updated_at": time.Now(),
	}

	if booking.Notes != "" {
		updates["notes"] = booking.Notes
	}
	if booking.Symptoms != "" {
		updates["symptoms"] = booking.Symptoms
	}

	return r.db.Model(&entity.Booking{}).Where("id = ?", id).Updates(updates).Error
}

func (r *bookingRepository) generateBookingCode() string {
	// Generate unique booking code with timestamp
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("BK%s", timestamp)
}
