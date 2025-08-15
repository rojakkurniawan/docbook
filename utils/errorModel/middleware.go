package errormodel

import (
	"docbook/utils/response"
	"net/http"
)

var (
	// Authorization header errors
	ErrInvalidAuthorizationHeader = response.ErrorResponse{
		Status:    http.StatusUnauthorized,
		ErrorCode: "INVALID_AUTHORIZATION_HEADER",
		Message:   "The authorization header is invalid",
	}
	ErrTokenRequired = response.ErrorResponse{
		Status:    http.StatusUnauthorized,
		ErrorCode: "TOKEN_REQUIRED",
		Message:   "The token is required",
	}

	// Token validation errors
	ErrInvalidToken = response.ErrorResponse{
		Status:    http.StatusUnauthorized,
		ErrorCode: "INVALID_TOKEN",
		Message:   "The token provided is invalid",
	}
	ErrInvalidTokenType = response.ErrorResponse{
		Status:    http.StatusUnauthorized,
		ErrorCode: "INVALID_TOKEN_TYPE",
		Message:   "Invalid token type provided",
	}
	ErrTokenExpired = response.ErrorResponse{
		Status:    http.StatusUnauthorized,
		ErrorCode: "TOKEN_EXPIRED",
		Message:   "The token provided has expired",
	}

	// Refresh token errors
	ErrInvalidRefreshToken = response.ErrorResponse{
		Status:    http.StatusUnauthorized,
		ErrorCode: "INVALID_REFRESH_TOKEN",
		Message:   "The refresh token provided is invalid",
	}
	ErrRefreshTokenExpired = response.ErrorResponse{
		Status:    http.StatusUnauthorized,
		ErrorCode: "REFRESH_TOKEN_EXPIRED",
		Message:   "The refresh token provided has expired",
	}

	// Token generation errors
	ErrTokenGenerationFailed = response.ErrorResponse{
		Status:    http.StatusInternalServerError,
		ErrorCode: "TOKEN_GENERATION_FAILED",
		Message:   "Failed to generate token",
	}
	ErrRefreshTokenGenerationFailed = response.ErrorResponse{
		Status:    http.StatusInternalServerError,
		ErrorCode: "REFRESH_TOKEN_GENERATION_FAILED",
		Message:   "Failed to generate refresh token",
	}

	// Permission errors
	ErrNotAllowed = response.ErrorResponse{
		Status:    http.StatusForbidden,
		ErrorCode: "NOT_ALLOWED",
		Message:   "You are not allowed to access this resource",
	}
)
