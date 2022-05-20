package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping
// @Summary health ping
// @Schemes
// @Description do ping
// @Produce json
// @Success 200 {string} Pong
// @Router /health/ping [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}
