package services

import (
	"docbook/entity"
	"docbook/utils"
	"errors"

	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (as *AuthService) Register(user *entity.User) (string, error) {
	if err := user.HashPassword(user.Password); err != nil {
		return "", errors.New("error hashing password")
	}

	if err := as.DB.Create(user).Error; err != nil {
		return "", errors.New("error creating user")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("error generating token")
	}

	return token, nil
}

func (as *AuthService) Login(loginReq *entity.LoginRequest) (string, error) {
	var user entity.User

	if err := as.DB.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	if err := user.CheckPassword(loginReq.Password); err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("error generating token")
	}

	return token, nil
}
