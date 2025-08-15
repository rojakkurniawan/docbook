package services

import (
	"docbook/config"
	"docbook/entity"
	"docbook/repository"
	"docbook/utils"
	errormodel "docbook/utils/errorModel"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserService interface {
	Register(user *entity.User) (*entity.TokenResponse, error)
	Login(loginRequest *entity.LoginRequest) (*entity.TokenResponse, error)
	RefreshToken(refreshToken string) (*entity.TokenResponse, error)
	UpdateUser(id uint, user *entity.User) error
	ChangePassword(id uint, changePasswordRequest *entity.UserChangePasswordRequest) error
	DeleteUser(id uint) error
	GetUserByID(id uint) (*entity.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (as *userService) Register(user *entity.User) (*entity.TokenResponse, error) {
	existingUser, err := as.userRepo.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingUser != nil {
		return nil, errormodel.ErrUserAlreadyExists
	}

	user.Role = "user"
	if err := as.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return as.generateTokenPair(user.ID, user.Role)
}

func (as *userService) Login(loginRequest *entity.LoginRequest) (*entity.TokenResponse, error) {
	existingUser, err := as.userRepo.GetUserByEmail(loginRequest.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingUser == nil {
		return nil, errormodel.ErrInvalidCredentials
	}

	if err := existingUser.CheckPassword(loginRequest.Password); err != nil {
		return nil, errormodel.ErrInvalidCredentials
	}

	return as.generateTokenPair(existingUser.ID, existingUser.Role)
}

func (as *userService) UpdateUser(id uint, user *entity.User) error {
	existingUser, err := as.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errormodel.ErrUserNotFound
		}
		return err
	}

	user.Password = existingUser.Password

	return as.userRepo.UpdateUser(id, user)
}

func (as *userService) ChangePassword(id uint, changePasswordRequest *entity.UserChangePasswordRequest) error {
	existingUser, err := as.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errormodel.ErrUserNotFound
		}
		return err
	}

	if err := existingUser.CheckPassword(changePasswordRequest.OldPassword); err != nil {
		return errormodel.ErrInvalidOldPassword
	}

	return as.userRepo.ChangePassword(id, changePasswordRequest.NewPassword)
}

func (as *userService) DeleteUser(id uint) error {
	existingUser, err := as.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errormodel.ErrUserNotFound
		}
		return err
	}

	if existingUser.Role == "admin" {
		return errormodel.ErrUserNotAuthorized
	}

	return as.userRepo.DeleteUser(id)
}

func (as *userService) GetUserByID(id uint) (*entity.User, error) {
	existingUser, err := as.userRepo.GetByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errormodel.ErrUserNotFound
		}
		return nil, err
	}

	return existingUser, nil
}

func (as *userService) RefreshToken(refreshToken string) (*entity.TokenResponse, error) {
	claims, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		// Handle different types of JWT errors specifically
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errormodel.ErrRefreshTokenExpired
		}
		if errors.Is(err, jwt.ErrTokenInvalidClaims) {
			return nil, errormodel.ErrInvalidRefreshToken
		}
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errormodel.ErrInvalidRefreshToken
		}
		// For any other JWT errors, return invalid token
		return nil, errormodel.ErrInvalidRefreshToken
	}

	user, err := as.userRepo.GetByID(claims.UserID)
	if err != nil || user == nil {
		return nil, errormodel.ErrUserNotFound
	}

	return as.generateTokenPair(user.ID, user.Role)
}

func (as *userService) generateTokenPair(userId uint, role string) (*entity.TokenResponse, error) {
	accessToken, err := utils.GenerateToken(userId, role)
	if err != nil {
		return nil, errormodel.ErrTokenGenerationFailed
	}

	refreshToken, err := utils.GenerateRefreshToken(userId, role)
	if err != nil {
		return nil, errormodel.ErrRefreshTokenGenerationFailed
	}

	expirationTime := time.Now().Add(config.GetJWTExpirationDuration()).Unix()

	return &entity.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expirationTime,
	}, nil
}
