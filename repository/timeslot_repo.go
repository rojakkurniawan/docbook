package repository

import (
	"docbook/entity"

	"gorm.io/gorm"
)

type TimeslotRepository interface {
	CreateMultipleTimeslots(timeslots []entity.TimeSlot) error
	GetTimeslotByID(id uint) (*entity.TimeSlot, error)
	GetAllTimeslots() ([]entity.TimeSlot, error)
	GetTimeslotsWithFilter(filter *entity.TimeslotFilter) ([]entity.TimeSlot, error)
	GetDoctorScheduleByID(id uint) (*entity.DoctorSchedule, error)
	UpdateTimeslot(id uint, timeslot *entity.TimeSlot) error
	DeleteTimeslot(id uint) error
}

type timeslotRepository struct {
	db *gorm.DB
}

func NewTimeslotRepository(db *gorm.DB) TimeslotRepository {
	return &timeslotRepository{db: db}
}

func (r *timeslotRepository) CreateMultipleTimeslots(timeslots []entity.TimeSlot) error {
	return r.db.Create(&timeslots).Error
}

func (r *timeslotRepository) GetTimeslotByID(id uint) (*entity.TimeSlot, error) {
	var timeslot entity.TimeSlot
	if err := r.db.Where("id = ?", id).First(&timeslot).Error; err != nil {
		return nil, err
	}
	return &timeslot, nil
}

func (r *timeslotRepository) GetAllTimeslots() ([]entity.TimeSlot, error) {
	var timeslots []entity.TimeSlot
	if err := r.db.Find(&timeslots).Error; err != nil {
		return nil, err
	}
	return timeslots, nil
}

func (r *timeslotRepository) GetDoctorScheduleByID(id uint) (*entity.DoctorSchedule, error) {
	var schedule entity.DoctorSchedule
	if err := r.db.Where("id = ?", id).First(&schedule).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *timeslotRepository) UpdateTimeslot(id uint, timeslot *entity.TimeSlot) error {
	return r.db.Model(&entity.TimeSlot{}).Where("id = ?", id).Updates(timeslot).Error
}

func (r *timeslotRepository) DeleteTimeslot(id uint) error {
	return r.db.Delete(&entity.TimeSlot{}, id).Error
}

func (r *timeslotRepository) GetTimeslotsWithFilter(filter *entity.TimeslotFilter) ([]entity.TimeSlot, error) {
	var timeslots []entity.TimeSlot
	query := r.db

	// Apply filters
	if filter.DoctorScheduleID != nil {
		query = query.Where("doctor_schedule_id = ?", *filter.DoctorScheduleID)
	}

	if filter.Date != nil {
		query = query.Where("date = ?", *filter.Date)
	}

	if filter.StartDate != nil {
		query = query.Where("date >= ?", *filter.StartDate)
	}

	if filter.EndDate != nil {
		query = query.Where("date <= ?", *filter.EndDate)
	}

	if filter.IsAvailable != nil {
		query = query.Where("is_available = ?", *filter.IsAvailable)
	}

	if filter.IsBlocked != nil {
		query = query.Where("is_blocked = ?", *filter.IsBlocked)
	}

	if filter.StartTime != nil {
		query = query.Where("start_time >= ?", *filter.StartTime)
	}

	if filter.EndTime != nil {
		query = query.Where("end_time <= ?", *filter.EndTime)
	}

	if err := query.Find(&timeslots).Error; err != nil {
		return nil, err
	}

	return timeslots, nil
}
