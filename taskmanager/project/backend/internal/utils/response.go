package utils

import "github.com/gin-gonic/gin"

// Response is the consistent JSON envelope returned by every endpoint.
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success sends a 2xx JSON response with the standard envelope.
func Success(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error sends an error JSON response with the standard envelope.
func Error(c *gin.Context, status int, message string) {
	c.JSON(status, Response{
		Success: false,
		Message: message,
		Data:    nil,
	})
}
