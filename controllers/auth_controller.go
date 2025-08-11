package controllers

import (
	"docbook/consts"
	"docbook/entity"
	"docbook/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Register(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, entity.APIResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	token, err := ac.authService.Register(&user)
	if err != nil {
		log.Println("error registering user", err)
		switch err {
		case consts.ErrUserExists:
			c.JSON(http.StatusBadRequest, entity.APIResponse{
				Success: false,
				Message: err.Error(),
			})
		case consts.ErrTokenGenerationFailed:
			c.JSON(http.StatusInternalServerError, entity.APIResponse{
				Success: false,
				Message: err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, entity.APIResponse{
				Success: false,
				Message: err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, entity.APIResponse{
		Success: true,
		Message: "User registered successfully",
		Data:    token,
	})
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginRequest entity.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, entity.APIResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	token, err := ac.authService.Login(&loginRequest)

	if err != nil {
		log.Println("error logging in user", err)
		switch err {
		case consts.ErrUserNotFound:
			c.JSON(http.StatusBadRequest, entity.APIResponse{
				Success: false,
				Message: err.Error(),
			})
		case consts.ErrInvalidCredentials:
			c.JSON(http.StatusBadRequest, entity.APIResponse{
				Success: false,
				Message: err.Error(),
			})
		case consts.ErrTokenGenerationFailed:
			c.JSON(http.StatusInternalServerError, entity.APIResponse{
				Success: false,
				Message: err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, entity.APIResponse{
				Success: false,
				Message: err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, entity.APIResponse{
		Success: true,
		Message: "User logged in successfully",
		Data:    token,
	})
}
