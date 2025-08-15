package repository

import (
	"docbook/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
	GetByID(id uint) (*entity.User, error)
	UpdateUser(id uint, user *entity.User) error
	ChangePassword(id uint, newPassword string) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	if err := user.HashPassword(user.Password); err != nil {
		return err
	}
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(id uint, user *entity.User) error {
	return r.db.Model(&entity.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *userRepository) ChangePassword(id uint, newPassword string) error {
	user := entity.User{
		Password: newPassword,
	}

	if err := user.HashPassword(user.Password); err != nil {
		return err
	}

	return r.db.Model(&entity.User{}).Where("id = ?", id).Update("password", user.Password).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}
