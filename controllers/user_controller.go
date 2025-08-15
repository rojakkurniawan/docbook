package controllers

import (
	"docbook/entity"
	"docbook/services"
	errormodel "docbook/utils/errorModel"
	"docbook/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	authService services.UserService
}

func NewUserController(authService services.UserService) *UserController {
	return &UserController{authService: authService}
}

func (ac *UserController) Register(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	var (
		detailError = make(map[string]any)
	)

	if user.FirstName == "" {
		detailError["first_name"] = "First name is required"
	}

	if user.LastName == "" {
		detailError["last_name"] = "Last name is required"
	}

	if user.Email == "" {
		detailError["email"] = "Email is required"
	}

	if user.Password == "" {
		detailError["password"] = "Password is required"
	}

	if len(detailError) > 0 {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput.AttachDetail(detailError))
		return
	}

	token, err := ac.authService.Register(&user)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	data := response.LoginInfo{
		Email:        user.Email,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
	}

	response.BuildSuccessResponse(c, http.StatusCreated, "User registered successfully", data, nil)
}

func (ac *UserController) Login(c *gin.Context) {
	var loginRequest entity.LoginRequest

	var (
		detailError = make(map[string]any)
	)

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		if loginRequest.Email == "" {
			detailError["email"] = "Email is required"
		}

		if loginRequest.Password == "" {
			detailError["password"] = "Password is required"
		}

		if len(detailError) > 0 {
			response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput.AttachDetail(detailError))
			return
		}

		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	token, err := ac.authService.Login(&loginRequest)

	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	data := response.LoginInfo{
		Email:        loginRequest.Email,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
	}

	response.BuildSuccessResponse(c, http.StatusOK, "User logged in successfully", data, nil)
}

func (ac *UserController) UpdateUser(c *gin.Context) {
	var user entity.User

	userID := c.GetUint("user_id")

	if err := c.ShouldBindJSON(&user); err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	var (
		detailError = make(map[string]any)
	)

	if user.FirstName == "" {
		detailError["first_name"] = "First name is required"
	}

	if user.LastName == "" {
		detailError["last_name"] = "Last name is required"
	}

	if user.Email == "" {
		detailError["email"] = "Email is required"
	}

	if user.Phone == "" {
		detailError["phone"] = "Phone is required"
	}

	if len(detailError) > 0 {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput.AttachDetail(detailError))
		return
	}

	err := ac.authService.UpdateUser(userID, &user)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	data := map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"phone":      user.Phone,
	}

	response.BuildSuccessResponse(c, http.StatusOK, "User updated successfully", data, nil)
}

func (ac *UserController) ChangePassword(c *gin.Context) {
	var changePasswordRequest entity.UserChangePasswordRequest

	userID := c.GetUint("user_id")

	if err := c.ShouldBindJSON(&changePasswordRequest); err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	var (
		detailError = make(map[string]any)
	)

	if changePasswordRequest.OldPassword == "" {
		detailError["old_password"] = "Old password is required"
	}

	if changePasswordRequest.NewPassword == "" {
		detailError["new_password"] = "New password is required"
	}

	if len(detailError) > 0 {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput.AttachDetail(detailError))
		return
	}

	err := ac.authService.ChangePassword(userID, &changePasswordRequest)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Password changed successfully", nil, nil)
}

func (ac *UserController) DeleteUser(c *gin.Context) {
	userID := c.GetUint("user_id")

	err := ac.authService.DeleteUser(userID)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "User deleted successfully", nil, nil)
}

func (ac *UserController) GetUserByID(c *gin.Context) {
	userID := c.GetUint("user_id")

	user, err := ac.authService.GetUserByID(userID)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	data := map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"phone":      user.Phone,
	}

	response.BuildSuccessResponse(c, http.StatusOK, "User retrieved successfully", data, nil)
}

func (ac *UserController) RefreshToken(c *gin.Context) {
	var refreshRequest entity.RefreshTokenRequest
	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		response.BuildErrorResponse(c, errormodel.ErrInvalidUserInput)
		return
	}

	tokenResponse, err := ac.authService.RefreshToken(refreshRequest.RefreshToken)
	if err != nil {
		response.BuildErrorResponse(c, err)
		return
	}

	response.BuildSuccessResponse(c, http.StatusOK, "Token refreshed successfully", tokenResponse, nil)
}
