package services

import (
	"docbook/entity"
	"docbook/repository"
	errormodel "docbook/utils/errorModel"
)

type ScheduleService interface {
	CreateMultipleSchedules(request *entity.CreateScheduleRequest) error
	GetScheduleByUserID(userID uint) ([]entity.DoctorSchedule, error)
	GetScheduleByID(id uint) (*entity.DoctorSchedule, error)
	UpdateSchedule(id uint, schedule *entity.DoctorSchedule) error
	DeleteSchedule(id uint) error
}

type scheduleService struct {
	scheduleRepo repository.ScheduleRepository
}

func NewScheduleService(scheduleRepo repository.ScheduleRepository) ScheduleService {
	return &scheduleService{scheduleRepo: scheduleRepo}
}

func (s *scheduleService) CreateMultipleSchedules(request *entity.CreateScheduleRequest) error {
	// Buat multiple schedules
	var schedules []entity.DoctorSchedule
	for _, scheduleReq := range request.Schedules {
		schedule := entity.DoctorSchedule{
			DoctorID:    request.DoctorID,
			DayOfWeek:   scheduleReq.DayOfWeek,
			StartTime:   scheduleReq.StartTime,
			EndTime:     scheduleReq.EndTime,
			IsAvailable: scheduleReq.IsAvailable,
			MaxPatients: scheduleReq.MaxPatients,
			Duration:    scheduleReq.Duration,
		}
		schedules = append(schedules, schedule)
	}

	return s.scheduleRepo.CreateMultipleSchedules(schedules)
}

func (s *scheduleService) GetScheduleByUserID(userID uint) ([]entity.DoctorSchedule, error) {
	user, err := s.scheduleRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if user.Role != "doctor" {
		return nil, errormodel.ErrUserNotDoctor // Buat error baru
	}

	// Cari doctor berdasarkan userID
	doctor, err := s.scheduleRepo.GetDoctorByUserID(user.ID)
	if err != nil {
		return nil, errormodel.ErrDoctorNotFound
	}

	schedule, err := s.scheduleRepo.GetScheduleByDoctorID(doctor.ID) // Gunakan doctor.ID, bukan userID
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (s *scheduleService) GetScheduleByID(id uint) (*entity.DoctorSchedule, error) {
	return s.scheduleRepo.GetScheduleByID(id)
}

func (s *scheduleService) UpdateSchedule(id uint, schedule *entity.DoctorSchedule) error {
	return s.scheduleRepo.UpdateSchedule(id, schedule)
}

func (s *scheduleService) DeleteSchedule(id uint) error {
	return s.scheduleRepo.DeleteSchedule(id)
}
