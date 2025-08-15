package repository

import (
	"docbook/entity"

	"gorm.io/gorm"
)

type DoctorRepository interface {
	GetUserByID(id uint) (*entity.User, error)
	GetDoctorByUserID(userID uint) (*entity.Doctor, error)
	GetAllDoctors() ([]entity.Doctor, error)
}

type doctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) DoctorRepository {
	return &doctorRepository{db: db}
}

func (r *doctorRepository) GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *doctorRepository) GetDoctorByUserID(userID uint) (*entity.Doctor, error) {
	var doctor entity.Doctor
	if err := r.db.Preload("User").Preload("Specialization").Where("user_id = ?", userID).First(&doctor).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *doctorRepository) GetAllDoctors() ([]entity.Doctor, error) {
	var doctors []entity.Doctor
	if err := r.db.Preload("User").Preload("Specialization").Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}
