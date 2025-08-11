package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	Phone     string `json:"phone"`
	Role      string `json:"role" gorm:"type:ENUM('user', 'doctor', 'admin');default:'user'"`
	gorm.Model
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
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
