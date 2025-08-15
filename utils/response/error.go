package response

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrInternalServer = ErrorResponse{
		Status:    http.StatusInternalServerError,
		ErrorCode: "INTERNAL_SERVER_ERROR",
		Message:   "Please contact support",
	}

	ErrBadRequest = ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorCode: "BAD_REQUEST",
		Message:   "Please check your request",
	}
)

type ErrorResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message,omitempty"`
	Details map[string]any `json:"detail_error,omitempty"`

	Status    int    `json:"-"`
	ErrorCode string `json:"-"`
}

func (e ErrorResponse) Error() string {
	if e.ErrorCode == "" {
		return e.Message
	}

	if len(e.Details) > 0 {
		return fmt.Sprintf("%s: %v", e.ErrorCode, e.Details)
	}

	return fmt.Sprintf("%s: %v", e.ErrorCode, e.Message)
}

func (e ErrorResponse) Return() error {
	if e.ErrorCode == "" {
		return e
	}

	return nil
}

func (e *ErrorResponse) AttachDetail(details map[string]any) *ErrorResponse {
	e.Details = details
	return e
}

func BuildError(err error) ErrorResponse {
	var response ErrorResponse

	// Coba type assertion untuk ErrorResponse
	if checkErr, ok := err.(ErrorResponse); ok {
		response.Success = false
		response.Message = checkErr.Message
		response.Details = checkErr.Details
		response.Status = checkErr.Status
		response.ErrorCode = checkErr.ErrorCode
	} else {
		// Coba type assertion untuk *ErrorResponse (pointer)
		if checkErr, ok := err.(*ErrorResponse); ok {
			response.Success = false
			response.Message = checkErr.Message
			response.Details = checkErr.Details
			response.Status = checkErr.Status
			response.ErrorCode = checkErr.ErrorCode
		} else {
			response = ErrInternalServer
			log.Printf("Internal Server Error : %v", err)
		}
	}

	return response
}

func BuildErrorResponse(c *gin.Context, err error) {
	response := BuildError(err)
	log.Printf("error code: %s | error message: %s", response.ErrorCode, response.Message)
	c.JSON(response.Status, response)
}
