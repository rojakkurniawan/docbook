package services

import (
	"docbook/entity"
	"docbook/repository"
)

type BookingService interface {
	CreateBooking(booking *entity.Booking, patient *entity.PatientRequest) error
	GetBookingByID(id uint) (*entity.BookingResponse, error)
	UpdateBooking(id uint, booking *entity.UpdateBookingRequest) error
	UpdateBookingAdmin(id uint, booking *entity.UpdateBookingAdminRequest) error
}

type bookingService struct {
	bookingRepo repository.BookingRepository
}

func NewBookingService(bookingRepo repository.BookingRepository) BookingService {
	return &bookingService{bookingRepo: bookingRepo}
}

func (s *bookingService) CreateBooking(booking *entity.Booking, patient *entity.PatientRequest) error {
	return s.bookingRepo.CreateBooking(booking, patient)
}

func (s *bookingService) GetBookingByID(id uint) (*entity.BookingResponse, error) {
	return s.bookingRepo.GetBookingByID(id)
}

func (s *bookingService) UpdateBooking(id uint, booking *entity.UpdateBookingRequest) error {
	return s.bookingRepo.UpdateBooking(id, booking)
}

func (s *bookingService) UpdateBookingAdmin(id uint, booking *entity.UpdateBookingAdminRequest) error {
	return s.bookingRepo.UpdateBookingAdmin(id, booking)
}
