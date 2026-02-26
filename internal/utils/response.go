package utils

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`  // Data returned on success, omitted if empty
	Error   interface{} `json:"error,omitempty"` // Specific error info on failure, omitted if empty
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.Set("response", data)
}

func ErrorResponse(c *gin.Context, code int, err error) {
	message := "ERROR"
	switch code {
	case 400:
		message = "BAD REQUEST"
	case 401:
		message = "UNAUTHORIZED"
	case 403:
		message = "FORBIDDEN"
	case 404:
		message = "NOT FOUND"
	case 500:
		message = "INTERNAL SERVER ERROR"
	}

	c.Error(err)
	c.AbortWithStatusJSON(code, APIResponse{
		Code:    code,
		Message: message,
		Error:   err.Error(),
	})
}
