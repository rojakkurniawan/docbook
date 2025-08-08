package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"full_name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Phone    string `json:"phone"`
	Role     string `json:"role" gorm:"type:ENUM('patient', 'doctor', 'admin');default:'patient'"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
