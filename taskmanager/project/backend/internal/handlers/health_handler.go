package handlers

import (
	"net/http"

	"taskmanager/internal/utils"

	"github.com/gin-gonic/gin"
)

// HealthCheck handles GET /api/health
func HealthCheck(c *gin.Context) {
	utils.Success(c, http.StatusOK, "Service is healthy", gin.H{
		"status": "ok",
	})
}
