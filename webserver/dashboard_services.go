package webserver

import (
	"errors"
	"net/http"
	"workflows/internal/dashboard"

	"github.com/gin-gonic/gin"
)

func registerDashboardServices(rg *gin.RouterGroup) {
	workflows := rg.Group("/dashboard")
	{
		workflows.GET("/:id", getDashboardConfig)
	}
}

func getDashboardConfig(c *gin.Context) {
	var workflowId string = c.Param("id")

	workflow, exists := WFHelper.WorkflowById(workflowId)
	if !exists {
		c.String(http.StatusBadRequest, errors.New("workflow does not exist").Error())
		return
	}

	config, exists := dashboard.ConfigFromWorkflow(*workflow)
	if !exists {
		c.String(http.StatusOK, "workflow has no canvas node")
		return
	}

	// return c.JSON(http.StatusOK, workflow)
	c.JSON(http.StatusOK, gin.H{"canvas": config})
}
