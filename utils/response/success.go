package response

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success  bool        `json:"success"`
	Status   int         `json:"status_code,omitempty"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}

func BuildSuccessResponse(c *gin.Context, status int, message string, data, metadata interface{}) {
	payload := buildSuccess(data, metadata, message)
	c.JSON(status, payload)
}

func buildSuccess(data, metadata interface{}, message string) SuccessResponse {
	var response SuccessResponse

	response.Success = true
	response.Message = message
	response.Data = data
	response.Metadata = metadata

	return response
}
