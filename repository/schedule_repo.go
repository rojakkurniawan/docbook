package repository

import (
	"docbook/entity"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	GetUserByID(id uint) (*entity.User, error)
	CreateSchedule(schedule *entity.DoctorSchedule) error
	CreateMultipleSchedules(schedules []entity.DoctorSchedule) error
	GetScheduleByDoctorID(doctorID uint) ([]entity.DoctorSchedule, error)
	GetDoctorByUserID(userID uint) (*entity.Doctor, error)
	GetScheduleByID(id uint) (*entity.DoctorSchedule, error)
	UpdateSchedule(id uint, schedule *entity.DoctorSchedule) error
	DeleteSchedule(id uint) error
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *scheduleRepository) CreateSchedule(schedule *entity.DoctorSchedule) error {
	return r.db.Create(schedule).Error
}

func (r *scheduleRepository) GetScheduleByDoctorID(doctorID uint) ([]entity.DoctorSchedule, error) {
	var schedules []entity.DoctorSchedule
	if err := r.db.Where("doctor_id = ?", doctorID).Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *scheduleRepository) GetDoctorByUserID(userID uint) (*entity.Doctor, error) {
	var doctor entity.Doctor
	if err := r.db.Where("user_id = ?", userID).First(&doctor).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *scheduleRepository) GetScheduleByID(id uint) (*entity.DoctorSchedule, error) {
	var schedule entity.DoctorSchedule
	if err := r.db.Where("id = ?", id).First(&schedule).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *scheduleRepository) UpdateSchedule(id uint, schedule *entity.DoctorSchedule) error {
	return r.db.Model(&entity.DoctorSchedule{}).Where("id = ?", id).Updates(schedule).Error
}

func (r *scheduleRepository) DeleteSchedule(id uint) error {
	return r.db.Delete(&entity.DoctorSchedule{}, id).Error
}

func (r *scheduleRepository) CreateMultipleSchedules(schedules []entity.DoctorSchedule) error {
	return r.db.Create(&schedules).Error
}
