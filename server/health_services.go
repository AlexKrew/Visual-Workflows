package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerHealthServices(rg *gin.RouterGroup) {
	rg.GET("/ping", healthCheck)
}

// Ping
// @Summary health ping
// @Schemes
// @Description do ping
// @Produce json
// @Success 200 {string} Pong
// @Router /ping [get]
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}
