package repository

import (
	"docbook/entity"

	"gorm.io/gorm"
)

type AdminRepository interface {
	GetUserByID(id uint) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	CreateDoctor(createDoctorRequest *entity.CreateDoctorRequest) error
	GetDoctorByLicenseNumber(licenseNumber string) (*entity.Doctor, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *adminRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *adminRepository) GetDoctorByLicenseNumber(licenseNumber string) (*entity.Doctor, error) {
	var doctor entity.Doctor
	if err := r.db.Where("license_number = ?", licenseNumber).First(&doctor).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *adminRepository) CreateDoctor(createDoctorRequest *entity.CreateDoctorRequest) error {
	user := &createDoctorRequest.User
	doctor := &createDoctorRequest.Doctor
	specialization := &createDoctorRequest.Specialization

	if err := user.HashPassword(user.Password); err != nil {
		return err
	}

	user.Role = "doctor"
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	if err := r.db.Create(specialization).Error; err != nil {
		return err
	}

	doctor.UserID = user.ID
	doctor.IsActive = true
	doctor.TotalPatients = 0
	doctor.SpecializationID = specialization.ID

	return r.db.Create(doctor).Error
}
