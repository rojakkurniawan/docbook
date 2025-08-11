package services

import (
	"docbook/consts"
	"docbook/entity"
	"docbook/repository"
	"docbook/utils"
	"errors"

	"gorm.io/gorm"
)

type AuthService interface {
	Register(user *entity.User) (string, error)
	Login(loginRequest *entity.LoginRequest) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (as *authService) Register(user *entity.User) (string, error) {
	existingUser, err := as.userRepo.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	if existingUser != nil {
		return "", consts.ErrUserExists
	}

	user.Role = "user"
	if err := as.userRepo.CreateUser(user); err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", consts.ErrTokenGenerationFailed
	}

	return token, nil
}

func (as *authService) Login(loginRequest *entity.LoginRequest) (string, error) {

	existingUser, err := as.userRepo.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return "", err
	}

	if existingUser == nil {
		return "", consts.ErrUserNotFound
	}

	if err := existingUser.CheckPassword(loginRequest.Password); err != nil {
		return "", consts.ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(existingUser.ID)
	if err != nil {
		return "", consts.ErrTokenGenerationFailed
	}

	return token, nil
}
