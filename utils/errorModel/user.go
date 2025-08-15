package errormodel

import (
	"docbook/utils/response"
	"net/http"
)

var (
	ErrUserAlreadyExists = response.ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "USER_ALREADY_EXISTS",
		Message:   "User already exists",
	}
	ErrUserNotFound = response.ErrorResponse{
		Status:    http.StatusNotFound,
		ErrorCode: "USER_NOT_FOUND",
		Message:   "User not found",
	}
	ErrInvalidCredentials = response.ErrorResponse{
		Status:    http.StatusUnauthorized,
		ErrorCode: "INVALID_CREDENTIALS",
		Message:   "Email or password is incorrect",
	}
	ErrUserNotAuthorized = response.ErrorResponse{
		Status:    http.StatusUnauthorized,
		ErrorCode: "USER_NOT_AUTHORIZED",
		Message:   "User not authorized",
	}
	ErrInvalidUserInput = response.ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_USER_INPUT",
		Message:   "Some fields are required or invalid",
	}
	ErrInvalidOldPassword = response.ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_OLD_PASSWORD",
		Message:   "Old password is incorrect",
	}
)
